package docker

import (
	"context"

	"github.com/docker/docker/client"
)

func DockerEngineVersion() string {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	versionInfo, err := cli.ServerVersion(ctx)
	if err != nil {
		panic(err)
	}
	cli.Close()

	return versionInfo.Version
}
