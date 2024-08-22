package common

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

/*
	 {
	  "code": 0,
	  "msg": "OK",
	  "data": {
	    "message": "Hello world!"
	  }
	}
*/
func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	marshal, _ := json.Marshal(resp)
	fmt.Println(string(marshal))
	if err != nil {
		body.Code = -1
		body.Msg = err.Error()
	} else {
		body.Msg = "ok"
		body.Data = resp
	}
	httpx.OkJson(w, body)
}
