package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func TestCreatePostHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/post"
	r.POST(url, CreatePostHandler)

	//body := `{
	//	"community_id": 1,
	//	"title": "test",
	//	"content": "just a test"
	//}`
	//
	//req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))
	//
	//w := httptest.NewRecorder()
	//r.ServeHTTP(w, req)

	//assert.Equal(t, 200, w.Code)
	//
	//// 判断响应的内容是不是按预期返回了需要登录的错误
	//
	//// 方法1：判断响应内容是不是包含了指定的字符串
	////assert.Contains(t, w.Body.String(), "需要登录")
	//
	//// 方法2：将响应的内容反序列化到ResponseData 然后判断字段与预期是否一致
	//res := new(ResponseData)
	//if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
	//	t.Fatalf("json.Unmarshal w.Body failed,err:%v\n", err)
	//}
	//assert.Equal(t, res.Code, CodeNeedLogin)
	//
	//body1 := `{
	//	"title": "test",
	//	"content": "just a test"
	//}`
	//req1, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body1)))
	//
	//w1 := httptest.NewRecorder()
	//r.ServeHTTP(w1, req1)
	//
	//assert.Equal(t, 200, w1.Code)
	//res1 := new(ResponseData)
	//if err := json.Unmarshal(w1.Body.Bytes(), res1); err != nil {
	//	t.Fatalf("json.Unmarshal w.Body failed,err:%v\n", err)
	//}
	//assert.Equal(t, res1.Code, CodeInvalidParam)

	testCase := []struct {
		name         string
		body         string
		statusCode   int
		expectedCode ResCode
	}{
		{
			name:         "Valid request",
			body:         `{"community_id":123, "title":"test","content":"just a test"}`,
			statusCode:   200,
			expectedCode: CodeNeedLogin,
		},
		{
			name:         "CodeInvalidParam",
			body:         `{"title":"test","content":"just a test"}`,
			statusCode:   200,
			expectedCode: CodeInvalidParam,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(tc.body)))
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tc.statusCode, w.Code)

			res := new(ResponseData)
			if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
				t.Helper() //为了更好地定位错误，可以在需要打印错误消息的地方调用 t.Helper()，以指示测试框架将错误消息与调用测试的代码位置关联起来。
				t.Fatalf("json.Unmarshal w.body failed,err: %v\n", err)

			}
			assert.Equal(t, tc.expectedCode, res.Code)
		})
	}
}
