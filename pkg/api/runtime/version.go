package runtime

import (
	"context"
	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/docker/docker/client"
)

type versionInformation struct {
	CodeName  string
	Version   string
	BuiltDate string
}

func GetVersion() (string, string, string, string, string, string, string, string) {
	var conf map[string]versionInformation
	conf, err := encodeToml("../../VERSION.toml")
	if err != nil {
		panic(err)
	}

	return conf["version"].CodeName, conf["version"].Version, golangVersion(), dockerEngineVersion(), "", conf["version"].BuiltDate, getOs(), arch()
}

func encodeToml(filePath string) (map[string]versionInformation, error) {
	conf := map[string]versionInformation{}
	_, err := toml.DecodeFile(filePath, &conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func golangVersion() string {
	return runtime.Version()
}

func getOs() string {
	return runtime.GOOS
}

func arch() string {
	return runtime.GOARCH
}

func dockerEngineVersion() string {
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
