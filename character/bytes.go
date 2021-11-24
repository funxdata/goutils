/*
 * @Author: dowell87
 * @Date: 2021-11-21 18:30:56
 * @Descripttion:
 * @LastEditTime: 2021-11-21 18:38:17
 */
package character

import "unsafe"

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
