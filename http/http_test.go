/*
 * @Author: dowell87
 * @Date: 2021-10-27 16:15:55
 * @Descripttion:
 * @LastEditTime: 2021-10-27 16:26:30
 */
package http

import "testing"

func TestHTTPGet(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Log(HTTPGet("https://www.funxdata.com/"))
	}
}

func TestHTTPPost(t *testing.T) {
	data := ""
	for i := 0; i < 10; i++ {
		t.Log(HTTPPost("https://www.funxdata.com/", data))
	}
}
