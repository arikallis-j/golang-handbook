package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Header struct {
		Code    int    `json:"code"`
		Message string `json:"message,omitempty"`
	} `json:"header"`
	Data []struct {
		Type       string `json:"type"`
		ID         int    `json:"id"`
		Attributes struct {
			Email       string `json:"email"`
			Article_ids []int  `json:"article_ids"`
		} `json:"attributes"`
	} `json:"data,omitempty"`
}

func ReadResponse(rawResp string) (Response, error) {
	res := Response{}
	err := json.Unmarshal([]byte(rawResp), &res)
	return res, err

}

const resRaw = `
{
	"header": {
		"code": 0,
		"message": ""
	},
	"data": [{
		"type": "user",
		"id": 100,
		"attributes": {
			"article_ids": [10, 11, 12],
			"email": "bob@yandex.ru"
		}
	}]
}
`

func main() {
	res, _ := ReadResponse(resRaw)
	fmt.Printf("%+v", res)
}
