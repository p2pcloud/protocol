package converters

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUrlToBytes32(t *testing.T) {
	bytes, err := UrlToBytes32("https://web.mydomain.com:8080/eb/232323")
	require.NoError(t, err)
	require.Len(t, bytes, 32)

	decodedStr, err := Bytes32ToUrl(bytes)
	require.NoError(t, err)
	require.Equal(t, "https://web.mydomain.com:8080/eb/232323", decodedStr)
}

func TestUrlToBytes32TooLong(t *testing.T) {
	str := "https://web.mydomain.com:8080/eb/232323/too/long"
	_, err := UrlToBytes32(str)
	require.Error(t, err)
}

func TestLength(t *testing.T) {
	for i := 0; i < 40; i++ {
		url := "https://" + strings.Repeat("a", i)
		bytes, err := UrlToBytes32(url)
		if i < 32 {
			require.NoError(t, err)
			require.Len(t, bytes, 32)

			decodedStr, err := Bytes32ToUrl(bytes)
			require.NoError(t, err)
			require.Equal(t, url, decodedStr)
		} else if i >= 32 {
			require.Error(t, err)
		}
	}
}

func TestUrlToBytes32Short(t *testing.T) {
	bytes, err := UrlToBytes32("https://a.com")
	require.NoError(t, err)
	require.Len(t, bytes, 32)

	decodedStr, err := Bytes32ToUrl(bytes)
	require.NoError(t, err)
	require.Equal(t, decodedStr, "https://a.com")
}

func TestHttpHttps(t *testing.T) {
	bytes, err := UrlToBytes32("https://web.mydomain.com:8080/eb/232323")
	require.NoError(t, err)
	require.Len(t, bytes, 32)

	decodedStr, err := Bytes32ToUrl(bytes)
	require.NoError(t, err)
	require.Equal(t, "https://web.mydomain.com:8080/eb/232323", decodedStr)

	bytes, err = UrlToBytes32("http://not.secure/eb/232323")
	require.NoError(t, err)
	require.Len(t, bytes, 32)

	decodedStr, err = Bytes32ToUrl(bytes)
	require.NoError(t, err)
	require.Equal(t, "http://not.secure/eb/232323", decodedStr)
}

func TestNonUrl(t *testing.T) {
	_, err := UrlToBytes32("not a url")
	require.Error(t, err)
}
