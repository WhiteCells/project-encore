package rawendpoints

import (
	"encoding/json"
	"io"
	"net/http"
)

//encore:api public raw
func Webhook(w http.ResponseWriter, req *http.Request) {
	// w 直接想客户端写出 HTTP 响应
	// req 客户端发送的 HTTP 请求
	// 1. 接受非标准格式的请求
	// 2. 控制 HTTP 响应
	// 3. 处理未预定义的路径或动态路由

	// if req.Method != http.MethodPost {
	// 	http.Error(w, "Method not allowd", http.StatusMethodNotAllowed)
	// 	return
	// }

	var data map[string]interface{}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "failed to read request body", http.StatusInternalServerError)
		return
	}
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"success"}`))
}
