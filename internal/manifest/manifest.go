// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.0.14"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "2f07a6c2f09dddde7edfa1d328d30ffa7f0fe409374def0bc9bc3a503c865790",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "00053150e613e5ac2d18a1ef9bb2c74b226839180a001b33b57bd6b8eb4ae393",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "983075b633e41a88a133fabfad045eb5c6ce7a7eb9c0f72da2dbf2fe3549185d",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "52311f96afaac2dc1a0ad5428c8f7018d643f59c067adb553d55d30782ea45c5",
	"code_graph-x86_64-apple-darwin.tar.gz":           "4a8f2f8ae14e4ac60b3d81a2cb6be5aa257369a16f1e244908330b576ec6b481",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "500f5d6c6115b6e0e7152de1998be44feb9cd8403d323816eb69dd5ba32e5024",
	"code_graph-x86_64-pc-windows-msvc.zip":           "d0ed90c3eadf6c0392fa1d67e4d081df6967a7f55d434d3ac00ce8138a09d192",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "5ced589133f08e959af50aed6e60145dd05b1b2f0454c376af5faf8fdb5d78f3",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "106c2ec32ef6f40274a239891ad594b669fd0b4ef8269a852aeb7c4ce7b58fcb",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "e0322658a619e61eba3ec79ed5c3ee1651f17a11cd27692f1d92a72c339e4202",
}
