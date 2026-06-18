// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.4.0"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "2d195e7af6f1174e6c3df28a1f274b364a8b6a0f60a478ee6470c40058e00f2f",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "0b9b309b7aeeb18298c267a23abdaf4c54c023ef06c973d2cee9583dfedda785",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "d8aa893a40470ef76ace78d8ea282d13d3e20f5d8ef9df4baf93f6c7a3d477c2",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "6829825f89f9dc3d505aa4e20a5bb7176927ebadade00d14964f2626329736e2",
	"code_graph-x86_64-apple-darwin.tar.gz":           "7b53290cbf41690478f5ec4b8f493bb415cac97b21dd7eb6acc9537084d98433",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "17d3799b1bb1c3512fbd94533e5f47403e5cdbc7efd91b69b400ac1fe46b3af6",
	"code_graph-x86_64-pc-windows-msvc.zip":           "dd0ce8b35d3769170b7119f004126a943bdea8d7ff339690a88bd13c71e0c0f5",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "c511e29a955f895932a60ae457b144c50cc6ad3b7ceaecccbf6084914a1cf0a3",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "bc8cd5363722e1b32727363e1697357ca074730954cdf39a1af54c2dda7339d3",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "0716897806706a4ed5aa52a84670a89ce4827600c50da49c366612974e08c194",
}
