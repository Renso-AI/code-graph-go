// Launcher for `code_graph-mcp`. Mirrors cmd/code_graph/main.go but
// downloads and execs the MCP-server binary.

package main

import (
	"github.com/renso-ai/code-graph-go/internal/launcher"
)

func main() {
	launcher.Run("code_graph-mcp")
}
