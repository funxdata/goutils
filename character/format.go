/*
 * @Author: dowell87
 * @Date: 2021-11-21 18:35:06
 * @Descripttion: 字符串格式化
 * @LastEditTime: 2021-11-21 18:38:24
 */
package character

import (
	"fmt"
	"strings"
)

const (
	cRedBegin   = "\033[31m"
	cGreenBegin = "\033[32m"
	cBrownBegin = "\033[33m"
	cEnd        = "\033[0m"
)

func ErrorString(format string, args ...interface{}) string {
	return cRedBegin + fmt.Sprintf(format, args...) + cEnd
}

func SuccessString(format string, args ...interface{}) string {
	return cGreenBegin + fmt.Sprintf(format, args...) + cEnd
}

func WarnString(format string, args ...interface{}) string {
	return cBrownBegin + fmt.Sprintf(format, args...) + cEnd
}

func FormatToCamelWithPrefix(pre, src string) string {
	var str string = pre
	src = strings.Replace(src, "_", "__", -1)
	if len(src) <= 1 {
		str += "_" + strings.ToUpper(src)
	} else {
		offset, i := 0, 1
		for ; i < len(src)-1; i++ {
			if src[i] >= 'A' && src[i] <= 'Z' {
				if (src[i-1] >= 'A' && src[i-1] <= 'Z') && (src[i+1] >= 'A' && src[i+1] <= 'Z') {
					continue
				}
				str += "_" + strings.ToUpper(string(src[offset:i]))
				offset = i
			}
		}
		if src[len(src)-1] >= 'A' && src[len(src)-1] <= 'Z' {
			if src[len(src)-2] >= 'A' && src[len(src)-2] <= 'Z' {
				str += "_" + strings.ToUpper(string(src[offset:]))
			} else {
				str += "_" + strings.ToUpper(string(src[offset:i]))
				str += "_" + strings.ToUpper(string(src[i:]))
			}
		} else {
			str += "_" + strings.ToUpper(string(src[offset:]))
		}
	}
	if pre == "" {
		str = string(str[1:])
	}
	return str
}

func FormatToCamel(src string) string {
	return FormatToCamelWithPrefix("", src)
}
