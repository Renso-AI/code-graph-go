// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.4.3"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "e9e7f50198f2a1523708728e779560a9067b691da17b7d19083005a6ced3d15e",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "4f78ff2f9c766537065247dee278b02a8cf02f110a71bbed6c853cf29ae1d01d",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "b733964aa41ff81c416b4f0e9e5ea7ba67ea3ec4acea5d14ca7b3d4f8fdb6d72",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "64897cd8b23a7f9aad3efd4967e4c8872c306bd2462f2f6c1fc226c560b8f186",
	"code_graph-x86_64-apple-darwin.tar.gz":           "1ccd5d585c4b21b352553591ce91b050390f7b3c00ab44f6b61a0219c0eac42a",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "68152092f3c374ea52df4857fa7058b5e02e2142263b5a61fc5d71c859e11cd5",
	"code_graph-x86_64-pc-windows-msvc.zip":           "3c28148ce75d168b69eb1b281ba190ec8ffc4382c32789313a3f3745186e6870",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "7699d8d10d05badbbe4e5c2553d88779258a68852b5e8cc6b3628bed3a15d4bc",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "d052f88f4fddbba9efdaab55a45076f711970036314ce1fa0b6fff09ddd65bf9",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "48ddfe0e3567f9c9fd48e1f20a9a03a5be35b4642f27b9aa1cbc54ec7575a478",
}
