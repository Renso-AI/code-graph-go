// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.0.0"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code_graph/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz": "11966fa1d8589ed50b249060fb160b57bd487c48514e6a206b706bb65770fb80",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz": "61dfa8b25d731be17fe46566c4827c30910e013d5e5762d051da2e6401144902",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz": "f23715755e4151558b51abf18485e3b4f764e9c1d0c90cfafd32c4a1f6f8e76e",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "d7ce4dbeea1b80360eb26ba0242aa7d5abbc7266c069fe56e3d70d8660fd5946",
	"code_graph-x86_64-apple-darwin.tar.gz": "502d508ad7194bbd480a1639668ea004b24beca1f91dd04f4727294ab6fc2dc4",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz": "01597cdaf616d83a4640533bff1c19bc66f8be03aa39fc898bd8ac2d6eb55fd5",
	"code_graph-x86_64-pc-windows-msvc.zip": "8f78f6ffe771010f1130655c9a667807c4a22e047c2fd160599b1c25dc6e78f4",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip": "9d216db9ad50ed531fea98528296824b74345adbe24fdf0de87deb76a2ed29d9",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz": "590469a8a1b1b672b5028104294970ac9e8e3cbf4f87123e983d23268f65580e",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz": "e72d8fc7055c35f9fe167a0bff0a915a02b533f4a9acee361fdc98119c2fd510",
}
