/*
 * @Author: dowell87
 * @Date: 2021-10-27 16:29:04
 * @Descripttion:
 * @LastEditTime: 2021-10-27 16:30:16
 */
package pagination

/**
 * @description: 获取页数
 * @param {*} page
 * @param {int} pageSize
 * @return {*}
 */
func GetPage(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}

/**
 * @description: 获取分页的总数
 * @param {*} total
 * @param {int} pagesize
 * @return {*}
 */
func GetPageSum(total, pagesize int) int {
	res := total % pagesize
	if res == 0 {
		return total / pagesize
	} else {
		return (total / pagesize) + 1
	}
}
