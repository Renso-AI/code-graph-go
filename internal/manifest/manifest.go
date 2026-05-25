// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.0.7"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "3f9a3bfae359a803bf69ebe5e411c1280ed6fa70c020eb2b18a8dbff703063a8",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "263f83382a861ddfd70f8ad45b783a84ed31e431cac429dde7d0a76226974590",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "9521d27410b3f93a1c2cccdea8669c6c42e8e57ab37df6ccc72190b39addba32",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "8bca56f076f243d034b7e712a6b530c00f4f4fbc445ce4c62119499311f77a5c",
	"code_graph-x86_64-apple-darwin.tar.gz":           "eeb74f3a728a94e083ece1fd2f7f5b1a9870fc5b94ba3adad9f05d77f2aae45f",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "366f0bf29f084d91c1d15fd8ebef94977fdf0f5140c8132fedc3069adb52a368",
	"code_graph-x86_64-pc-windows-msvc.zip":           "90178d2e9740ccbf859521cc0415b45ab8ced342b41a6a3c53d6596522ff31d8",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "89731507faac3053c5e97604278e4a71b23300ca0cf6dbbf6f2e3725c1a9fcd3",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "9b75e3ffbabfbac4c3e80d114e586424e2bd42935c05ad789fb482b4de0f4b89",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "607ea67f093908690f66a78ced345f05df1e41d3aaff2748d9cf237b0ebd3dcf",
}
