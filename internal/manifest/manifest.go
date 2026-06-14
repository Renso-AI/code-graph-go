// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.2.2"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "8ef7592285299491ebccea75e6f6ce8c9cc24be224a18bdf4c211828211e89db",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "a78bf019589962701374da84f5c56e2b2b85ad95ad381831ab299e7be032dd28",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "f744cc2932c418a67d17cba20f70ce17db64cf8d93ac20bd187423cd04d3bc26",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "47a4f9f79bb3ff0d4292d9a3e99e47992c57f0919b0e64ca63b5ea3817cbf4ec",
	"code_graph-x86_64-apple-darwin.tar.gz":           "05b2843620edb14da685c6c750fc5335991da281c1d1c90e0b15a52cbe748f97",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "208034fa7bafc57078f997eaab7dd6dd7d535b828d60030f9cef69eba742cf5c",
	"code_graph-x86_64-pc-windows-msvc.zip":           "7f2d25590acfeebda67da3fc3b6d560d17d825d24d2db56b684daed5205762e9",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "8d84fcb54ae27584bff5dfb65819e1038eff338b01075c99d86783c609ceddb8",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "8fc3df3fadea4643f60ad0294013ff2710570450f40c00ba787721c0c58ac040",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "7d0e93a54ff2894032785232337244f85ef366e1f71703880ec286326bbabdac",
}
