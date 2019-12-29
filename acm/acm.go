package acm

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	yaml "gopkg.in/yaml.v2"
)

type cacheConf struct {
	dataId  string
	group   string
	content string
	md5     string
}

type Client struct {
	AccessKey  string
	SecretKey  string
	EndPoint   string
	NameSpace  string
	TimeOut    int
	servers    map[int]string
	HttpClient *http.Client
	cache      sync.Map
}

// func NewClient(option func(c *Client)) (*Client, error) {
func NewClient(acckey, secret, nsid string) (*Client, error) {
	client := &Client{
		AccessKey:  acckey,
		SecretKey:  secret,
		NameSpace:  nsid,
		EndPoint:   "acm.aliyun.com",
		TimeOut:    15,
		HttpClient: &http.Client{Timeout: 10 * time.Second},
		servers:    make(map[int]string),
		cache:      sync.Map{},
	}

	err := client.initServer()
	return client, err
}

func (c *Client) initServer() error {
	if c.EndPoint == "" {
		return errors.New("endpoint not empty")
	}

	resp, err := c.HttpClient.Get(fmt.Sprintf("http://%s:8080/diamond-server/diamond", c.EndPoint))

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	body := strings.TrimSpace(string(byt))

	if resp.StatusCode != 200 {
		return errors.New(body)
	}

	servers := strings.Split(body, "\n")

	for k, v := range servers {
		if strings.Index(v, ":") == -1 {
			c.servers[k] = v + ":8080"
		} else {
			c.servers[k] = v
		}
	}
	return nil
}

//rand reutrn server
func (c *Client) getServer() string {
	for _, v := range c.servers {
		return v
	}
	return ""
}

func (c *Client) getSign(params []string) string {
	signStr := strings.Join(params, "+")
	hc := hmac.New(sha1.New, []byte(c.SecretKey))
	hc.Write([]byte(signStr))
	return base64.StdEncoding.EncodeToString(hc.Sum(nil))
}

func (c *Client) callApi(api string, params map[string]string, method string) (string, error) {
	server := c.getServer()

	if server == "" {
		return "", errors.New("get server error")
	}

	timeStamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	timeStamp = timeStamp[:13]

	spec := "?"
	if strings.Index(api, "?") != -1 {
		spec = "&"
	}

	var request *http.Request
	var err error
	query := url.Values{}
	for k, v := range params {
		query.Add(k, v)
	}

	if method == "GET" {
		u := fmt.Sprintf("http://%s/%s%s%s", server, api, spec, query.Encode())
		request, err = http.NewRequest(method, u, nil)
	} else {
		u := fmt.Sprintf("http://%s/%s", server, api)
		request, err = http.NewRequest(method, u, strings.NewReader(query.Encode()))
	}

	if err != nil {
		return "", err
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	request.Header.Add("Spas-AccessKey", c.AccessKey)
	request.Header.Add("timeStamp", timeStamp)

	if probe, has := params["Probe-Modify-Request"]; has {
		request.Header.Add("longPullingTimeout", strconv.Itoa(c.TimeOut*1000))
		request.Header.Add("Spas-Signature", c.getSign([]string{probe}))
		c.HttpClient.Timeout = time.Duration(c.TimeOut+30) * time.Second
	} else {
		request.Header.Add("Spas-Signature", c.getSign([]string{c.NameSpace, params["group"], timeStamp}))
	}

	resp, err := c.HttpClient.Do(request)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	body := string(byt)

	if resp.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("response error:%s %s", body, params))
	}
	return body, nil
}

func (c *Client) GetServers() map[int]string {
	return c.servers
}

func (c *Client) getCacheKey(dataId, group string) string {
	return strings.Join([]string{c.NameSpace, dataId, group}, "-")
}

func (c *Client) GetConfig(dataId, group string) (string, error) {
	key := c.getCacheKey(dataId, group)
	var err error
	var ret string
	v, ok := c.cache.Load(key)

	if !ok {
		ret, err = c.callApi("diamond-server/config.co", map[string]string{
			"tenant": c.NameSpace,
			"dataId": dataId,
			"group":  group,
		}, "GET")

		if err == nil {
			c.cache.Store(key, &cacheConf{
				dataId:  dataId,
				group:   group,
				content: ret,
				md5:     c.getMd5(ret),
			})
		}
	} else {
		tmp := v.(*cacheConf)
		ret = tmp.content
	}

	return ret, err
}

// GetConfigObject ...
func (c *Client) GetConfigObject(dataId, group string, v interface{}) error {
	str, err := c.GetConfig(dataId, group)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal([]byte(str), v); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetAllConfigs(pageNo, pageSize int) (string, error) {
	return c.callApi("diamond-server/basestone.do?method=getAllConfigByTenant", map[string]string{
		"pageNo":   strconv.Itoa(pageNo),
		"pageSize": strconv.Itoa(pageSize),
	}, "GET")
}

func (c *Client) getMd5(ret string) string {
	// text := MustUtf8ToGbk([]byte(ret))
	text := []byte(ret)
	algorithm := md5.New()
	algorithm.Write(text)
	return hex.EncodeToString(algorithm.Sum(nil))
}

func (c *Client) Publish(dataId, group, content string) (string, error) {
	// bt, err := Utf8ToGbk([]byte(content))
	// if err != nil {
	// 	return "", err
	// }
	bt := []byte(content)

	return c.callApi("diamond-server/basestone.do?method=syncUpdateAll", map[string]string{
		"tenant":  c.NameSpace,
		"dataId":  dataId,
		"group":   group,
		"content": string(bt),
	}, "POST")
}

func (c *Client) Subscribe(dataId, group, contentMd5 string) (string, error) {
	key := c.getCacheKey(dataId, group)
	v, ok := c.cache.Load(key)
	if ok && contentMd5 == "" {
		tmp := v.(*cacheConf)
		contentMd5 = tmp.md5
	}

	log.Printf("local md5:%s", contentMd5)

	probe := strings.Join([]string{dataId, group, contentMd5, c.NameSpace}, "\x02") + "\x01"
	ret, err := c.callApi("diamond-server/config.co", map[string]string{
		"Probe-Modify-Request": probe,
	}, "POST")

	if err != nil {
		return "", err
	}

	if strings.TrimSpace(ret) == strings.Join([]string{dataId, group, c.NameSpace}, "%02")+"%01" {
		c.cache.Delete(key)
		return c.GetConfig(dataId, group)
	}
	return "", errors.New("no update")
}

func (c *Client) Delete(dateId, group string) (string, error) {
	return c.callApi("diamond-server/datum.do?method=deleteAllDatums", map[string]string{
		"tenant": c.NameSpace,
		"dataId": dateId,
		"group":  group,
	}, "POST")
}
