// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.0.22"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "fc35d2a7115d5eeaa982b9c62a4a8d969ce1b05d3b18a72ce86c56b9606a11cd",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "ad2ff5db3965490a19f31a88f6620c3d42b9ab6ccbe4c63c112468bb2d5b85a8",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "47ebf088b78bf795f452c90fb62dd34a907caa1a159d1750dc9f8bf4752a9487",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "21952e3bc7aed94e04e2e7a0ec849c6d7c8f067288e25ba3a13cd251b80cec98",
	"code_graph-x86_64-apple-darwin.tar.gz":           "81d9b468a4faddcda938c3b442c50d45ef68fcf41c738800530f26406cc240c5",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "d4e491c0c875dc568d8e60d268c26116e6a40544d807745ebc40289dd699d41c",
	"code_graph-x86_64-pc-windows-msvc.zip":           "be3c2bdaf764c049ca9e7bbb428e72c4c1c939a428078cb0d76939953b31bfd8",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "287cadc088ac6be87f93770a9f2a7f628d22c3309f4fb49fe7a038ade77ecd47",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "cbc846fc676f4994f8fed96bf55ada77355fdcaf3ef7e75c56bdeef69eb1e35c",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "c2e278bba9ced2572f46621bbda61f69f6b1bd16341c9fbe111db0cb347fb411",
}
