// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.1.3"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "516e87f83f7c3f0a9084412b38130b6a1f59f9d50c3a62058e17a043148038fc",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "edc16cfe6a051051407746210ee4c1aa5f29dd28f30d96a0aa40f31be507a8d2",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "902233aff6d6cd9f88019fd6a93d066f4fc7265d68b735a9f50d5e16b44c84de",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "f651130499fb9d47b6a29843fe3143ff5849949adab4fd891d1c7e42630aaf76",
	"code_graph-x86_64-apple-darwin.tar.gz":           "19917315c0a97f6c7c037be8b70a2c98b68af5480d10f7d07fb79cc41866ba85",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "e3507f1b46862e897d3d0fd6d2bc5ca55e670e77cd01f9e1809d261640b2797c",
	"code_graph-x86_64-pc-windows-msvc.zip":           "1380b0e1dddfacec86ce37edfbf342fe05f76510596ce648ee0cdf5bda83dcec",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "03ba8ebfb1b7d9a072c437c393a8b8318707426c94923f2b4e4586930fee431c",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "65e5b8da9f650cfdc74820cf92da2132ee46ab1f52cc6df15884c5e63e09a5c5",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "fe48e7f4c23dbdf68cea188a7956614d22f21d2a9c1b7d4bbd2912efab804966",
}
