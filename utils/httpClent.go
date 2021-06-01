package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type ApiResponseData struct {
	Rtn  int                    `json:"rtn"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}

type ApiResponseMsg struct {
	Rtn int    `json:"rtn"`
	Msg string `json:"msg"`
}

// HttpGet 根据struct返回不同形式定义的接口数据
func (s *ApiResponseData) HttpGet(url string) (resp ApiResponseData, err error) {
	apiResponse, err := HttpGetBody(url)
	if err != nil {
		return
	}

	err = json.Unmarshal(apiResponse, &resp)
	if err != nil {
		err = errors.New("Unmarshal failed " + err.Error())
		return
	}

	return resp, err
}

// HttpGet 根据struct返回不同形式定义的接口数据
func (s *ApiResponseMsg) HttpGet(url string) (resp ApiResponseMsg, err error) {
	apiResponse, err := HttpGetBody(url)
	if err != nil {
		return
	}

	err = json.Unmarshal(apiResponse, &resp)
	if err != nil {
		err = errors.New("Unmarshal failed " + err.Error())
		return
	}

	return resp, err
}

// HttpGetBody 通用的获取http的Body
func HttpGetBody(url string) (body []byte, err error) {
	res, _ := http.Get(url)
	defer res.Body.Close()

	if res.StatusCode != 200 {
		err = errors.New("Api接口调用错误:" + res.Status)
		return
	}

	body, err = ioutil.ReadAll(res.Body)

	return body, err
}
