package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/labstack/echo"
)

func TestOutputSistersVersion(t *testing.T) {
	apath, _ := filepath.Abs("../")
	os.Chdir(apath)

	echoServer := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/test/docker/version", strings.NewReader(""))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := echoServer.NewContext(req, rec)

	err := OutputSistersVersion(c)
	if err != nil {
		t.Fatal()
	}
	if rec.Code != http.StatusOK {
		t.Errorf("Status code: %d excepted, but got %d", http.StatusOK, rec.Code)
	}

	var info VersionInformationSisters
	err = json.NewDecoder(rec.Body).Decode(&info)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Client: Sisters")
	fmt.Printf("CodeName: %s\n", info.CodeName)
	fmt.Printf("Version: %s\n", info.Version)
	fmt.Printf("Go version: %s\n", info.GolangVersion)
	fmt.Printf("Docker Engine Version: %s\n", info.DockerEngineVersion)
	fmt.Printf("Git commit: %s\n", info.BuiltGitCommitHash)
	fmt.Printf("Built: %s\n", info.BuiltDate)
	fmt.Printf("OS: %s\n", info.Os)
	fmt.Printf("Arch: %s\n", info.Arch)
}
