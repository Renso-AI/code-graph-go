// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.0.20"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "22595e2090a75437ec30d24ce4ab097b6656186292491abce25ebf82bf98b899",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "3ac2c6a26dc6f41c8e2d2e6aa9292f9266d66eb0b0b948b09202d9a8ffd64b48",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "cc96ddfe7a2a178e2bdcad8ce59c1a0c3464df7ca12221d45fb03f03ab6a913a",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "fcd453d96a0906977376e620ecdc159d3abad23d08aa72a795491ffd0322049b",
	"code_graph-x86_64-apple-darwin.tar.gz":           "24f26050e9d88280b923a9ba840af682c24a9400ebcca793965f73f8469d1ad8",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "5ef53d4ef9ed636abdb33dd2b67ce66636394ee6a3b916f553271f951886d7d5",
	"code_graph-x86_64-pc-windows-msvc.zip":           "151a0b22278bdf65f66d170b5b9381596bb1c7995ab02a0ed3a02c65b8af219b",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "b4a7486cf7822209829a7ee455a094b564e7641f1208d2d5ccd2343da0f12c24",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "8cdfe06415e1d313b867a50a9b164700cdb35df495d6c70c044d6dca9abe92ea",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "644a3fc48404d735be34ca7d8ca40e2e0940745a4f39debaaab9d340bd9122e7",
}
