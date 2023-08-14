package main

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"log/slog"
)

var PORT = 8080

func main() {

	commit, dirty := GitCommit()

	slog.Info("Starting", "commit", commit, "dirty", dirty, "port", PORT)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Version", commit)
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}

// Stolen from https://github.com/acorn-io/baaah/blob/main/pkg/version/version.go#L33
func GitCommit() (commit string, dirty bool) {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return "", false
	}
	for _, setting := range bi.Settings {
		switch setting.Key {
		case "vcs.modified":
			dirty = setting.Value == "true"
		case "vcs.revision":
			commit = setting.Value
		}
	}
	return
}