// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.4.4"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "5c98bd5b478114e48e581c33f2a5680792639bc602f6b9feb6745f48bd8d5815",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "b09d3f0fb16b242798480ff7f313be31ea0ceb63731d0101fefe03effced38a1",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "f4b8a65419daef6bc188d7f3f52e2f7592a17c61c4ecc6f4db447c2133d826dd",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "bbf45ff3b8be93aee3da1766d401ef512f9c9b80e2eaaf9e3dfe63fd26e7e4b2",
	"code_graph-x86_64-apple-darwin.tar.gz":           "5c0b79e91f529b4589f25ae744a262e1bb16480e07111d3164c165f9a641796a",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "49c2248c7c9e472991923dc27b07a72e94ba57cf33339bfaabd09c3060d2e7cd",
	"code_graph-x86_64-pc-windows-msvc.zip":           "42557e8e94d90c2f57116d9d2e7ffb3b60c32f7bb5ff9039d5193836edb996f7",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "c581b46a5d0d0bfa821e280ef0262557ebde551d0227f20867834b6d3d728a2d",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "4d4bef43af2c63979242c3153637c343117d1abb20a1afe6c86dd12e6df12d94",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "ab5d3dddea172427d7002a58925ef0527fb3a32ed33554ce9e9880ac2424466e",
}
