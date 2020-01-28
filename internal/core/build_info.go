package core

import (
	"strings"
)

type BuildInfo struct {
	Version   string
	BuildDate string
	GoVersion string
	GitBranch string
	GitCommit string
	GoArch    string
	GoOS      string
}

// isRelease returns true when the version of the CLI is an official release:
// - version must be non-empty (exclude tests)
// - version must not contain '+dev' label
func (b *BuildInfo) isRelease() bool {
	return b.Version != "" && !strings.Contains(b.Version, "+")
}
