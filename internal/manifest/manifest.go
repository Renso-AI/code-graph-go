// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.0.11"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "77006fd43dd64f1a32f4f57e768ef1aca7100ee0f479eae14d90d216d914b2b4",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "9ab7b78aa9c36168081720ee41413b6c4823ba8497aa868f712b155a6ab488df",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "d94637e068b2173a93789d5fb59597b92760c08264ed9c691c39e9f96c9333cd",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "03a1f752b84c86acd956fb28f6550aec03e3185b743e989d1136b1f5f8cc7d0e",
	"code_graph-x86_64-apple-darwin.tar.gz":           "8194d1151a3167ded0f65aa8fc7a1741e470254bc23c8b083c3eaa11d5cf9df4",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "692480c83511c61c4db7fe4d095f6a7cd6212720843129cd233b6706569ceb04",
	"code_graph-x86_64-pc-windows-msvc.zip":           "2c69016cda18c56c9944a71040999c377daee39dc92dbd28350f5bb8ef79dba0",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "945e880126eb6b1b4d940578a6c092ff94c86aa89d7de7cce58d8fb0ac2d4735",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "c0be822e232e670eb6936f6bca86e350d6c95f9049a15753c1d2ac1e35e45848",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "e6c888304c413eaf1aadb0462acc6fbf74132bf7a31f95689b0c6cc5347b894d",
}
