package discovery

import (
	"strings"
)

func StackoverflowURLToDiscoveryID(url string) string {
	urlPieces := strings.SplitAfter(url, "stackoverflow.com/")
	url = urlPieces[1]
	urlPieces = strings.Split(url, "/")
	urlPieces = urlPieces[:3]
	url = strings.Join(urlPieces, "_")
	url = strings.Replace(url, "-", "_", -1)
	return url
}
