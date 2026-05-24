// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.0.5"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "b5c29e6159a6fc3c0a2b1c571a0ea926cff770e61122d61721f64d7b5d3b734d",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "425dd9b0a00e0908ebc3a556879abd3c47da09f8d7071756c64ff8f502531d24",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "c5706d827deba87948e042cf0dfa06684f696a5d949a142417042918ff663b1d",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "0499afc551b1d5cf9aaaee9434dc0d67cbe1008ee3e46b132548222181a94a5f",
	"code_graph-x86_64-apple-darwin.tar.gz":           "cb0ad97dd71183fc0ca3458c27a3f78222ce0a245b244496f995dc407e514b1a",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "3b6841672527942911735f872bd7dd01dd56742e84b5cda90ddd14f998add63d",
	"code_graph-x86_64-pc-windows-msvc.zip":           "6daeefc69c19abb616b08e9217d3a3c2ffebe6a797f6754f2b894848b17dc49c",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "a1cc51abeb48b49074440b64c4d5f62bfa42890c40e0e0c5e3d9f282f623fdbf",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "c7aa8304dc7834fe44bd406c1e4d5b4e68f5ee59b428c91e8efbcda0bd398f10",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "2db795bf20eac709bd88ec0aa6f1f834189405cfe7f360294b0625878c3164e7",
}
