package handler

import "time"

type VersionInformationSisters struct {
	CodeName            string    `json:"codeName"`
	Version             string    `json:"version"`
	GolangVersion       string    `json:"golangVersion"`
	DockerEngineVersion string    `json:"dockerEngineVersion"`
	BuiltGitCommitHash  string    `json:"builtgitCommitHash"`
	BuiltDate           time.Time `json:"builtDate"`
	Os                  string    `json:"os"`
	Arch                string    `json:"arch"`
}
