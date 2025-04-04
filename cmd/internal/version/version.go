package version

import (
	"fmt"
	"net/http"
)

var (
	Version   = "dev"
	GitCommit = "none"
	BuildTime = "unknown"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Version: %s\nGit Commit: %s\nBuild Time: %s\n", Version, GitCommit, BuildTime)
}