package container

import (
	"bytes"
	"context"
	"io"
	"log"
	"log/slog"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
)

func ContainerCreateAndStart(contextData []byte) int {
	buildDockerfile(contextData)

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
					Source: "/home/admidori/projects/rnp/sisters/tmp/context",
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

func buildDockerfile(contextData []byte) {
	imagename := "sisters-client:" + "samaple"
	// Todo: Automatic create a Dockerfile
	var (
		dockerfileName   string   = "Dockerfile"
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
		bytes.NewReader(contextData),
		buildOptions,
	)
	if err != nil {
		log.Panic(err)
	}
	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}
