// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.4.5"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "2c3926ad8fd65b4d6420d3c1d46018a7bfb8688864737ad1d1e0f47a32d998c7",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "8e52350b981fa8f616d2c3f86ed175f4b0a00a023ae0c9953bd503a15d13059d",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "2d1daf3a206f3a387c9e98a4af310a4177a7940a7a4e693d97c0c0d488358eb7",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "5a3c2b009bcaeb9682bdf4a58a0ebebf7157183d7a3aafaa5651ab376ae1e961",
	"code_graph-x86_64-apple-darwin.tar.gz":           "e1103fe7e8f7a296acc5b6f3a035625a93693e5fc52535820fcb0c25a9dbc624",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "97ffdabeaf0f9adbf3593d276f1182db5b2b6af1a9310b20a9a1a79a422e3cbf",
	"code_graph-x86_64-pc-windows-msvc.zip":           "cd0664974fc02ea8f7ac6ca46f78cc223d3fea79facc1bf7a21373f1784f01fa",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "0b1a57e29dbe3254b340f9bded4e3edb8f50fd30cee4db7fe032dbf5cd516f95",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "c88a68f4a9915c78ff98e79f9e19fe6ffc74dbfb4e818c6be55ff31813e8f9f4",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "0c73a9d1adeb5e34debd8e1b6de1c0414920beba305b7b1a692ec6cf5c722a10",
}
