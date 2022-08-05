package converters

import (
	"fmt"
	"strings"
)

func UrlToBytes32(url string) ([32]byte, error) {
	var result [32]byte
	bytes := []byte(url)

	if !strings.HasPrefix(url, "https://") {
		return result, fmt.Errorf("url '%s' has to start from https", url)
	}

	bytes = bytes[len("https://"):]

	if len(bytes) > 32 {
		return result, fmt.Errorf("invalid url length %v. at most 32 expected", len(bytes))
	}

	copy(result[:], bytes[:])

	return result, nil
}

func Bytes32ToUrl(bytes32 [32]byte) (string, error) {
	decoded := string(bytes32[:])
	decoded = strings.TrimRight(decoded, string([]byte{0}))
	return "https://" + decoded, nil
}
