/*
 * @Author: dowell87
 * @Date: 2021-11-15 11:43:48
 * @Descripttion: 一些基础比较大小的工具
 * @LastEditTime: 2021-11-15 11:46:27
 */
package maths

/**
 * @description: 取最大的数
 * @param {int} a
 * @param {int} b
 * @return {int}
 */

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

/**
 * @description: 取最小的数
 * @param {int} a
 * @param {int} b
 * @return {int}
 */

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
