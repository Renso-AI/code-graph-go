// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.0.1"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz": "298e1b8778b60ac1d66623a2360df4ba60ce5855e1030dfb60f2db413d23c6cd",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz": "213f1edfb4f2b4fcf41d38811d2e70f221978ba479180517f2e90d7712460aa7",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz": "e2c9d5f220dabc2bc04d2ab5089be326fe760afe2a010ce796e4086c837d9ab0",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "60f63509b9881f12e84f2bb78e046be0f4bf0042f592b2d0de7993d561b13eb0",
	"code_graph-x86_64-apple-darwin.tar.gz": "7858cdb3f3e72e3b7dd1b823f1545b161fc149b9e70d0d99abb60108587f0d0e",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz": "04441e6d8824c34218487d3f967bcc8c8a65a547ae24ad26ea9c8051e0babe5e",
	"code_graph-x86_64-pc-windows-msvc.zip": "a17515d1ee5c1afd8580a3310922db3692a440eb52e977f6d5a465187860f24a",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip": "81af39cb994c4894c318207ad09c49ec2c1f60e8df283d2da2d28555274c0f12",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz": "f78ccc2778f98ceae24b6073387b806b452f13557875a8e92af3e3a57b3bfa77",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz": "ac67de67407372f843d485a3f7be88619eceec358818165542b90fd57ffd3e17",
}
