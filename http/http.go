/*
 * @Author: dowell87
 * @Date: 2021-10-27 16:13:58
 * @Descripttion:
 * @LastEditTime: 2021-10-27 16:25:01
 */
package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
 * @description: http Get获取信息
 * @param {string} uri
 * @return {*}
 */

func HTTPGet(uri string) ([]byte, error) {
	response, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}

/**
 * @description: http Post获取信息
 * @param {string} uri
 * @param {string} data
 * @return {*}
 */

func HTTPPost(uri string, data string) ([]byte, error) {
	body := bytes.NewBuffer([]byte(data))
	response, err := http.Post(uri, "", body)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}
