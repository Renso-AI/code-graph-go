// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.6.2"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "f6d57ecff504914bb75741322444cb48ed0c697f6d012de44d51f23e232eee90",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "9186f19b5a7261303f1357b94c3fce080af4beba8de2d514f4a9f3bbf0d2808a",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "20ca8b26132d65e6e695c5bbf831c513af2c3fcc305d85e5ccd80c759811e7b7",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "e0139f21e087b8b6b0f58ecaf198752f014deafc258c18dcb514d4a7190bfca5",
	"code_graph-x86_64-apple-darwin.tar.gz":           "955ac75dd121a6d7dfaea675dc2bc2818edaafa2c9516a6e5bbe6fcc5674d714",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "404535355388d53147e93d6f826de72ce9a1ab9fd55f40787b6d3abd7b49b1a5",
	"code_graph-x86_64-pc-windows-msvc.zip":           "d12874a6d7ab600dc2cd8a915cfb4566522a931696d4686423a84fb4558a0f64",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "0746cead458f5554f5eeaf147378744323a17e20ccfb8f1007a2ede0c1c42e97",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "073664b538ad97f43235274b7da49ffcada9f715e357c950033beca9e1e08543",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "3ffc4e3fd8f0f0bbf840cab704e55620907c1468b2c55e4a5933b4ba6cb97e3c",
}
