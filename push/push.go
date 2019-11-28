package push

import (
	"encoding/json"
	"net/http"
	"socket/conn"
)

type Push struct {
	hub *conn.Hub

}
func NewPush(Hub *conn.Hub) *Push{
	p:=&Push{hub:Hub}
	return p
}
func (this *Push) TestPush(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	var jsonResult []byte

	defer func() {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(jsonResult)
	}()
	r.ParseForm()
	postData := r.PostForm
	result := make(map[string]interface{})
	result["code"] = 0
	result["msg"] = ""
	result["data"] = ""
	if msg, ok := postData["msg"]; ok {
		result["data"] = msg[0]
		this.hub.BroadPush(msg[0]) //推送数据
	} else {
		result["code"] = 1001
		result["msg"] = "msg 参数缺失"
	}
	jsonResult, _ = json.Marshal(result)
}
