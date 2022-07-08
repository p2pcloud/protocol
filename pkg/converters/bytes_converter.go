package converters

import (
	"fmt"
	"strings"
)

func UrlToBytes32(url string) ([32]byte, error) {
	var result [32]byte
	bytes := []byte(url)

	isHttp := strings.HasPrefix(url, "http://")
	isHttps := strings.HasPrefix(url, "https://")

	if !isHttp && !isHttps {
		return result, fmt.Errorf("string %s is not url", url)
	}

	if isHttp {
		bytes = bytes[len("http://"):]
		result[0] = 0
	} else {
		bytes = bytes[len("https://"):]
		result[0] = 1
	}

	if len(bytes) > 31 { // 1 byte for http/https
		return result, fmt.Errorf("invalid url length %v. at most 32 expected", len(bytes))
	}

	copy(result[1:], bytes[:])

	return result, nil
}

func Bytes32ToUrl(bytes32 [32]byte) (string, error) {
	var prefix = ""
	if bytes32[0] == 0 {
		prefix = "http://"
	} else if bytes32[0] == 1 {
		prefix = "https://"
	} else {
		return "", fmt.Errorf("invalid first byte of URL. 0 or 1 expected")
	}
	decoded := string(bytes32[1:])
	decoded = strings.TrimRight(decoded, string([]byte{0}))
	return prefix + decoded, nil
}
