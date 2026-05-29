// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.0.13"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "f6b951888be35f37c2cd130f3b6fe2470fd5ecefd1b78f5668243b1d5a191099",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "d0bda880a8e856a77176114d756e55f2fc277d2676793660bc36657bd9130c52",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "7ca95f27a232efcb2ab372fd8dbc0348f2d762ec3756f28e0fdb1352a4272c7c",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "98d1a14e790fc6d65506fccb2b60d86aca710baef0158eeccba894d836da5ec6",
	"code_graph-x86_64-apple-darwin.tar.gz":           "61222b3738d5555381319c657bcb740b7d4f759fa62ebb5d0a681362b83ecb3b",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "8ccd00fec3fb799d109c456bdc1fbb647d2b879af488e1bdae5212066e937ad6",
	"code_graph-x86_64-pc-windows-msvc.zip":           "be368cd1d1807e5891a8a73cb7631bc2ba24a765991e8c2a57f402aa0222e2c1",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "c9d0e78dfe12487630f5600c791ba0fb3aa75120270e61c6a053336d4b422de2",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "a0a3e79d41ed46478050567e4cffba6a786cefbc857ed6a65bb79d2bdb695e4d",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "ac55551546c343f7887daef5946833205bd538d470668e04b3b4d8ca7de9cff9",
}
