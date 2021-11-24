/*
 * @Author: dowell87
 * @Date: 2021-10-27 16:00:11
 * @Descripttion:
 * @LastEditTime: 2021-11-24 10:30:08
 */
package http

import (
	"net/url"
	"testing"
	"time"
)

// TestProxy ...
func TestProxy(t *testing.T) {
	return
	purl, _ := url.Parse("http://180.167.34.187:8181/")
	cli := NewProxyClient(purl)
	cli.Timeout = time.Second * 5
	_, err := cli.Get("http://example.ckeyer.com/")
	if err != nil {
		t.Error(err)
		return
	}
}
