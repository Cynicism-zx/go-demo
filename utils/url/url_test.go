package url

import (
	"net/url"
	"testing"
)

func TestUrl(t *testing.T) {
	u := "http://www.google.com?target=%E4%B8%AD%E6%96%87"
	str, _ := url.QueryUnescape(u)
	t.Log(str)
	t.Log(1672021223 - 1672021163)
}
