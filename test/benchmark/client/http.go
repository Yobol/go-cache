package client

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type HttpClient struct {
	*http.Client
	server string
}

func (c *HttpClient) get(key string) string {
	resp, e := c.Get(c.server + key)
	if e != nil {
		log.Println(key)
		panic(e)
	}
	if resp.StatusCode == http.StatusNotFound {
		return ""
	}
	if resp.StatusCode != http.StatusOK {
		panic(resp.Status)
	}
	b, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		panic(e)
	}
	return string(b)
}

func (c *HttpClient) set(key, value string) {
	req, e := http.NewRequest(http.MethodPut,
		c.server+key, strings.NewReader(value))
	if e != nil {
		log.Println(key)
		panic(e)
	}
	resp, e := c.Do(req)
	if e != nil {
		log.Println(key)
		panic(e)
	}
	if resp.StatusCode != http.StatusOK {
		panic(resp.Status)
	}
}

func (c *HttpClient) Run(cmd *Cmd) {
	if cmd.Name == "get" {
		cmd.Value = c.get(cmd.Key)
		return
	}
	if cmd.Name == "set" {
		c.set(cmd.Key, cmd.Value)
		return
	}
	panic("unknown cmd name " + cmd.Name)
}

func newHTTPClient() *HttpClient {
	client := &http.Client{Transport: &http.Transport{MaxIdleConnsPerHost: 1}}
	return &HttpClient{client, "http://localhost:10615/cache/"}
}

func (c *HttpClient) PipelinedRun([]*Cmd) {
	panic("httpClient pipelined run not implement")
}
