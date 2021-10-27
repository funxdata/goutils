/*
 * @Author: dowell87
 * @Date: 2021-10-27 16:30:44
 * @Descripttion:
 * @LastEditTime: 2021-10-27 16:32:35
 */
package pagination

import "testing"

func TestGetPage(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Log(GetPage(10, 20))
	}
}

func TestGetPageSum(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Log(GetPageSum(10, 20))
	}
}
