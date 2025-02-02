package utils

import "net/url"

func NormalizeUrl(url, baseURL string) string {

	if url == "/" || url == "" {
		return baseURL
	}

	url = TrimSlash(url)

	if url[0] == '/' {
		return baseURL + url
	}

	return url
}

func TrimSlash(url string) string {
	if url == "" || url[len(url)-1] != '/' {
		return url
	}

	return url[:len(url)-1]
}

func HaveSameHost(u1, u2 string) bool {
	url1, _ := url.Parse(u1)
	url2, _ := url.Parse(u2)

	return url1.Host == url2.Host
}
