// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.0.15"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "d7af2676f9f8384f6509e610f889f3fd6adc36f488e2b80d6c013c687b5fb8b6",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "dbbf3b71fdcebd558576f71c00f5614517e0d6002bd5d57e534ea8e96b543bbb",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "e1aaf9afb2c1a09f011329380fb504e7982677cbc9fc39eaeef82d2c830a8c17",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "338de37a313a2731e03eafa3d8782fe84ef96fb8ab97ceec8af53a2e588ffb66",
	"code_graph-x86_64-apple-darwin.tar.gz":           "40c63b8cc0e124893a11908c6bee676250481e0dafdc6771c229780a8b01702c",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "7af715feca5a79ea4ae5d1879d9a9af9e098c5ffd6703fc949c9b4a10fb74ecf",
	"code_graph-x86_64-pc-windows-msvc.zip":           "b62429433a51851e0ff1fb01a1779c4bc1cdc28199e1c5ff9bbd37cf89efde67",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "f8485c92b08613172990c0d7270cc3b693496711406522fd6cb84e8bbee9e316",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "93f66714cf3012ce51666d22a62dd1bc9b6d0bfddfcd4af7942414b87590434e",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "5102bb6ffc2a60d68716a4ef37af993841027cdd3eed5f804a16797904799964",
}
