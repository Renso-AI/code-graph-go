// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.1.0"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "69c33cc5ee104ccdb271a741674e033de8dafe1d6ca42aaea35195859b980d9a",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "c2e304598e74f41e1255a3741e823a5f4083b342c1ac9165220504c5eb729412",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "a4e9543f590825b1a8b1fa940d9aba8098d1f46af0785edd0f073b067f409bda",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "3ec2a6b1868ef51c241588ff7724e48b2748848cca7837aeac78e8c4a8fcc485",
	"code_graph-x86_64-apple-darwin.tar.gz":           "4058a1138a2a5f57a5b0a09add6f0b0238eb2a24ec0068c9ab39c23ff52edd9c",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "45dcad3234dc04329948fbedb524267970f55b0f6cc7c329ee0aaee43fcf7e60",
	"code_graph-x86_64-pc-windows-msvc.zip":           "a83e335c66baca01c693837e5b0fe559261796d4ac0f484ff1b21bc11644b4b2",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "2112d2eaede366cbf133f7400e10717b6c36a59901ca14ab4a323bc4c433388b",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "9809c5fbe031c31dd70c2505bf0d299f388f180d4ac5a7cc29156c7d95711dfd",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "93bb59b641a6a6aa01be408ae231058c83a381fb7749272e4ab07c4c2eb3e860",
}
