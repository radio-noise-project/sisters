package docker

import (
	"fmt"
	"testing"
)

func TestDockerEngineVersion(t *testing.T) {
	infoVersion := DockerEngineVersion()
	fmt.Printf("Docker Engine Version: %s", infoVersion)
}
