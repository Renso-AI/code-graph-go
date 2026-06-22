// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.4.1"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "6cb6424bf2f0cde654309ed791ed28e4f9799501d61df71be1ac246fc01c6aa5",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "70d559b3e91efeee93ef334a64162cde483149381902cbd3500cf3d8aeafbfa2",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "ba5ad39d4259352fb8c7940dc2932739f2f43a872ae842700c9bf0b397f82bfb",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "d32039023f4035ed8ef69312123db0dc9074bb41e1854e3ae071c15375cfed34",
	"code_graph-x86_64-apple-darwin.tar.gz":           "9613a21f2097e72b87c9c2792f51f60ce4c89d2e12d3962518d71d4facdb3e63",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "0fa0efa6f07ea068dc512122dede5bfb6174b21a2c5588f2549c857f99d6f5f1",
	"code_graph-x86_64-pc-windows-msvc.zip":           "255005933449868732db24fdcf82f725fe5e468978e247bd4a8dd02029e98729",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "3337fa07c6dfbe77042745b9b0afbe4669fd7e2e07636d9ba40eaa2e0683c5a3",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "af7fda9cb0a016555acdee3d56b23bda5e49a33083d2531a146fbb480d8623e2",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "c8c933c2adca72c41e86eae23381df7507af0919b6e98a73ecae974cc6be07e9",
}
