package container

import (
	"archive/tar"
	"bytes"
	"context"
	"io"
	"io/fs"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"

	"github.com/facebookgo/symwalk"
)

var DOCKERFILE_NAME string = "Dockerfile.cuda-11.1.1"

func ContainerCreateAndStart(HostVolumeAbsolutePath string) int {
	buildDockerfile("sample")

	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	cli.NegotiateAPIVersion(ctx)

	imagefilename := "nvcr.io/nvidia/cuda:" + "11.1.1-cudnn8-runtime"

	// If need environmental value, Please write below.
	envvalue1 := "ENV_VALUE=TEST"

	resp, err := cli.ContainerCreate(
		ctx,
		&container.Config{
			Image: imagefilename,
			Cmd:   []string{},
			Env:   []string{envvalue1},
			Tty:   false,
		}, &container.HostConfig{
			Mounts: []mount.Mount{
				{
					Type:   mount.TypeBind,
					Source: HostVolumeAbsolutePath,
					Target: "/workspace",
				},
			},
		}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		panic(err)
	}

	var exitCode container.WaitResponse
	exitCodeCh, errch := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errch:
		if err != nil {
			panic(err)
		}

	case exitCode = <-exitCodeCh:
		// End Process - Remove container
		/*
			err = cli.ContainerRemove(ctx, resp.ID, container.RemoveOptions{})
			if err != nil {
				panic(err)
			}
		*/

		//WA
		if exitCode.StatusCode == 1 {
			slog.Info("Status 1")
		}
		//CE
		if exitCode.StatusCode == 2 {
			slog.Info("Status 2")
		}
		//AC
		if exitCode.StatusCode == 0 {
			slog.Info("Status 0")
		}
	}
	return (int(exitCode.StatusCode))
}

func buildDockerfile(jobname string) {
	imagename := "sisters-client:" + jobname
	// Todo: Automatic create a Dockerfile
	var (
		dockerfileName   string   = DOCKERFILE_NAME
		imageNameAndTags []string = []string{imagename}
	)
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Panic(err)
	}
	cli.NegotiateAPIVersion(ctx)

	buildOptions := types.ImageBuildOptions{
		Dockerfile: dockerfileName,
		Remove:     true,
		Tags:       imageNameAndTags,
	}
	res, err := cli.ImageBuild(
		ctx,
		getArchivedContextdirectory("../test"),
		buildOptions,
	)
	if err != nil {
		log.Panic(err)
	}
	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}

func getArchivedContextdirectory(contextDir string) *bytes.Reader {
	buf := new(bytes.Buffer)

	tw := tar.NewWriter(buf)
	defer tw.Close()

	if err := symwalk.Walk(contextDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		intake_path := strings.Replace(path, contextDir, "", 1)
		tar_mode := info.Mode()
		tar_size := info.Size()

		if info.Mode()&fs.ModeSymlink != 0 {
			path_sym, err := filepath.EvalSymlinks(path)
			if err != nil {
				return err
			}
			path = path_sym
			stat, err := os.Stat(path_sym)
			if err != nil {
				return err
			}
			tar_mode = stat.Mode()
			tar_size = stat.Size()
		}

		if err := tw.WriteHeader(&tar.Header{
			Name:    intake_path,
			Mode:    int64(tar_mode),
			ModTime: info.ModTime(),
			Size:    tar_size,
		}); err != nil {
			return err
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err := io.Copy(tw, f); err != nil {
			return err
		}
		return nil

	}); err != nil {
		panic(err)
	}

	return bytes.NewReader(buf.Bytes())
}
