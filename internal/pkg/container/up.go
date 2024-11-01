package container

import (
	"archive/tar"
	"bytes"
	"context"
	"io"
	"log"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
)

var DOCKERFILE_NAME string = "Dockerfile.cuda-11.1.1"

func ContainerCreateAndStart() int {
	buildDockerfile("sample")

	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	cli.NegotiateAPIVersion(ctx)

	imagefilename := "judge-server:" + DOCKERFILE_NAME

	// If need environmental value, Please write below.
	envvalue1 := "ENV_VALUE=TEST"

	absolutePath, err := filepath.Abs("../../docker/images/")
	if err != nil {
		panic(err)
	}
	sourcepath := absolutePath + "/"

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
					Source: sourcepath,
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
		err = cli.ContainerRemove(ctx, resp.ID, container.RemoveOptions{})
		if err != nil {
			panic(err)
		}

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
		getArchivedDockerfile(DOCKERFILE_NAME),
		buildOptions,
	)
	if err != nil {
		log.Panic(err)
	}
	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}

func getArchivedDockerfile(dockerfile string) *bytes.Reader {
	// read the Dockerfile

	filepath := "../../docker/images/" + dockerfile
	f, err := os.Open(filepath)
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Panic(err)
		}
	}()
	b, err := io.ReadAll(f)
	if err != nil {
		log.Panic(err)
	}

	// archive the Dockerfile
	tarHeader := &tar.Header{
		Name: dockerfile,
		Size: int64(len(b)),
	}
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()
	err = tw.WriteHeader(tarHeader)
	if err != nil {
		log.Panic(err)
	}
	_, err = tw.Write(b)
	if err != nil {
		log.Panic(err)
	}

	return bytes.NewReader(buf.Bytes())
}
