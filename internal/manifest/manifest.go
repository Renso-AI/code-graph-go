// Package manifest holds release metadata templated by the release
// workflow before the Go module is pushed to renso-ai/code-graph-go.

package manifest

// Version of the upstream code_graph release this module is pinned to.
const Version = "1.0.23"

// ReleaseURLBase is the GitHub Releases prefix; the launcher appends
// "/v<Version>/<binary>-<target>.<ext>" to form the download URL.
const ReleaseURLBase = "https://github.com/Renso-AI/code-graph-dist/releases/download"

// SHA256SUMS maps "<binary>-<target>.<ext>" filenames to their
// hex-encoded SHA256, baked in so the launcher can verify downloads
// without trusting any extra metadata fetched at runtime.
var SHA256SUMS = map[string]string{
	"code_graph-aarch64-apple-darwin.tar.gz":          "54aba93a0f7eab744f7f146229ca2f86d13c2ffcabbffc750a535621edeb4d03",
	"code_graph-mcp-aarch64-apple-darwin.tar.gz":      "a0ea1dc8297c315b9459de772cd594c02f533ee998926fcb60cf42717536f129",
	"code_graph-aarch64-unknown-linux-gnu.tar.gz":     "11d9d0ce348e009c00dc86e7920a95c4dcf34a14c66166c151d059393d780692",
	"code_graph-mcp-aarch64-unknown-linux-gnu.tar.gz": "cd05e72fa76885071948d7773c5552feb79911221db6a37fa381bd4c4b3c3b3a",
	"code_graph-x86_64-apple-darwin.tar.gz":           "b334d4169c9b4b44c5cd11e0fc1bb06dbd4ccd5eb123d9442da473b9dd26a845",
	"code_graph-mcp-x86_64-apple-darwin.tar.gz":       "2997e742fcd71970ebeca0245532aeee6bc1f3dc5863e6f77bdd537e6f175cff",
	"code_graph-x86_64-pc-windows-msvc.zip":           "f70a871202deb79fd818deda8dc0e5d00a0e2de4b11ff662d31200c684871d6d",
	"code_graph-mcp-x86_64-pc-windows-msvc.zip":       "1723f6be54ffbe8a67b4cc41df81afb81d5b03ebe40d56b29edd57a63662b342",
	"code_graph-x86_64-unknown-linux-gnu.tar.gz":      "0cc60351cad071aa7b59d9653acea126ba20c6e12a077be1378f28fe69ef3e35",
	"code_graph-mcp-x86_64-unknown-linux-gnu.tar.gz":  "6a3c0191c6941352552a7611ee88c1adbf3393e4470958468aaa8672e464db5a",
}
