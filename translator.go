package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func fy(str string) string {
	res := requestTranslate(str)
	if res == nil {
		return ""
	}

	if res.Result.TransResult == nil {
		return ""
	}

	return res.Result.TransResult[0].Dst
}
func requestTranslate(content string) *Res {
	url := "https://aip.baidubce.com/rpc/2.0/mt/texttrans/v1?access_token=" + GetAccessToken()
	reqArgs := Request{
		From: "en",
		To:   "zh",
		Q:    content,
	}
	b, _ := json.Marshal(reqArgs)
	payload := strings.NewReader(string(b))

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var buf bytes.Buffer
	if _, err = io.Copy(&buf, res.Body); err != nil {
		panic(err)
	}

	var resp = &Res{}
	if err = json.Unmarshal(buf.Bytes(), resp); err != nil {
		panic(err)
	}

	return resp
}

/**
 * 使用 AK，SK 生成鉴权签名（Access Token）
 * @return string 鉴权签名信息（Access Token）
 */
func GetAccessToken() string {
	url := "https://aip.baidubce.com/oauth/2.0/token"
	postData := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", API_KEY, SECRET_KEY)
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(postData))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	accessTokenObj := map[string]string{}
	json.Unmarshal(body, &accessTokenObj)
	return accessTokenObj["access_token"]
}

type Request struct {
	From string `json:"from"`
	To   string `json:"to"`
	Q    string `json:"q"`
}

type Res struct {
	Result struct {
		From        string `json:"from"`
		TransResult []struct {
			Dst string `json:"dst"`
			Src string `json:"src"`
		} `json:"trans_result"`
		To string `json:"to"`
	} `json:"result"`
	LogId int64 `json:"log_id"`
}
