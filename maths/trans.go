/*
 * @Author: dowell87
 * @Date: 2021-11-21 18:23:34
 * @Descripttion: 进制转换
 * @LastEditTime: 2021-11-21 18:29:30
 */
package maths

import (
	"fmt"
	"strconv"
)

/**
 * @description: 十进制转换为2进制
 * @param {int} num
 * @return {int}
 */

func DecimalToBinary(num int) int {
	str := ""
	for ; num > 0; num /= 2 {
		lsb := num % 2
		fmt.Println(lsb)
		str = strconv.Itoa(lsb) + str
	}
	n, _ := strconv.Atoi(str)
	return n
}
