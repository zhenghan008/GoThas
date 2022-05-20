package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// New ArthasRequest 结构体初始化
func (g *ArthasRequest) New() {
	if g.Url == "" {
		g.Url = "http://localhost:8563/api"
	}
	if g.Action == "" {
		g.Action = "exec"
	}
	if g.Command == "" {
		g.Command = "version"
	}
	if g.ExecTimeout == 0 {
		g.ExecTimeout = 1000
	}
}

// GetInfoByCom 通过请求Arthas API 获取ArthasRespond 指针
func (g *ArthasRequest) GetInfoByCom() *ArthasRespond {
	var as = &ArthasRespond{}
	jsonArgs := fmt.Sprintf(`{ "action": "%s","command": "%s", "execTimeout": %d }`, g.Action, g.Command, g.ExecTimeout)
	jsonStr := []byte(jsonArgs)
	req, err := http.NewRequest("POST", g.Url, bytes.NewReader(jsonStr))
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error: ", err)
		}
	}()
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)
	StatusCode := resp.StatusCode
	fmt.Println("StatusCode: ", StatusCode)
	if StatusCode == 200 {
		r, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		err = json.Unmarshal(r, as)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println("response: ", string(r))

	}

	return as

}
