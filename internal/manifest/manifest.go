// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.2.0"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "36fcebc83d346dea2368cac27c9ea53c23647cd7bfd3625225cf0a2a0aef2f60",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "f4a0235ff1270d4aa5755d28f112a8ec62df90196fa47baf19d8b24513e7d99c",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "9c93369f8762b2fe3a3044643428bca8a8e1330fafd24b27e6025026211ea8d8",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "9d4d7058d7417997bc3c42cf5807a3c3eca4019365a27c96e3931ad82a699deb",
	"code_graph-x86_64-apple-darwin.tar.gz":           "e8ca0d2e13221bbf5c51fc4ac1a8fb9f4d0777d5e75504e5caf734808f663db9",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "130e4634195ac8837d9792f549da9677888be3e7e1424553a87048f8a57bf090",
	"code_graph-x86_64-pc-windows-msvc.zip":           "10b39ccea7e807e8026231c6f76fdd377973ab8b42fb2700fecd2b8e08a23c28",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "c28a042f0ab737a301300bb620bbf9d01d0a212d1cbb23ff1ee704bf33d29c8d",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "2feec4352500580a2b6f0a11e78ef78493a4f2a0154ae6b61a82f76f81e500fc",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "e1190cf8eed4524c781dd3ee1fb1220e482333bf310892ed5d2ba0fef16834e7",
}
