// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.2.1"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "81ded91ccfd9628ef76568d014f3867ccdcffad90e4baf1b12dc487a2fef1190",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "d26cae749cf2bf9a63e1c47e4c8c1c0fda5cda6614f2d308291b9a43904c7ac8",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "d24aacd8fa6fd057d2d7511f39ef6cab771c12757e704347501ebd09d4865b78",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "a01c32da0e7dc3266fd2c0e36bd1822f80479bb3ddc53f1c8a236400b05f9772",
	"code_graph-x86_64-apple-darwin.tar.gz":           "b6c5b6ecf29ffc874c14ea340d16ef133b6bebd6ae0a32f3b4368475fd3fda56",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "785adfd17639f975fc0652a5c1dc00e4e7190befdf736f4aa56e7d9788693849",
	"code_graph-x86_64-pc-windows-msvc.zip":           "2b079da5d1941e757a410ea3db3e48fbc4b53bcccbbb8e1a3a10550a172f732f",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "40950157442256614e66abbd31cb8bb233641125eb34edd8092552d16293ce09",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "55cc5b8a5ddd41f405e1fd3f69ec5c54266324537240c9ea86c5cb431306bef3",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "b8155b0a9c43cf106a26ee76beca2d78ced0f06202a5eddbc4a38f8bb5c70d0d",
}
