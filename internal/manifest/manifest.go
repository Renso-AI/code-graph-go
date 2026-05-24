// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.0.6"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "0be2aa8da8b9bb9b0979ff9e07d04f222e844ff05e51cbd21951cbac4ef5ec92",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "f0aa72b0935df71d24237c60b1fc9122e2088c00ad3d3fe4b54b49694d9ddf27",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "c5af0487e271922a1ff1184059c0c30493f2b520bf1aa6647d318a5ed7080e4d",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "a34424826b8a367403513e06d3cd6259f632b633697724591239eff89b5ae751",
	"code_graph-x86_64-apple-darwin.tar.gz":           "b8e8d816870089c0c8c35a4fefe4218c91cb20e4f0917f4ffdbc0f2f1463a6b0",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "285671482d0541e69076df27e9f7bdba4093473e970386ff2c9ac348d5939552",
	"code_graph-x86_64-pc-windows-msvc.zip":           "8085eee9cfa0937880c2a2fa9b91ab5d3a395c95caee353605f9bd292004a2b2",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "fa21b657629c641600496bc1531167f922adecbc7dc2faff4fec040147298ccc",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "2640163070bc5fa5f248f0d597c03413cb70a1d8827bf3299c8c1f1f51a407ac",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "30eaa1237e696549614639409a0926a4d0e2aefb6dc92cda8ae336e72d6adf4a",
}
