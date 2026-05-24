// Launcher for `code_graph`. Downloads the prebuilt Rust binary from
// GitHub Releases on first run, SHA256-verifies it against the
// manifest baked into this module, caches it under XDG_CACHE_HOME,
// then execs it forwarding os.Args + env.

package main

import (
	"github.com/renso-ai/code-graph-go/internal/launcher"
)

func main() {
	launcher.Run("code_graph")
}
