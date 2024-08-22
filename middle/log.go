package middle

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"test/common"
)

// 日志处理中间件
type LogMiddleware struct {
}

func NewLogMiddleware() *LogMiddleware {
	return &LogMiddleware{}
}

type logResponseWriter struct {
	writer http.ResponseWriter
	code   int
	// 存响应数据
	buf *bytes.Buffer
}

func newLogResponseWriter(writer http.ResponseWriter, code int) *logResponseWriter {
	var buf bytes.Buffer
	return &logResponseWriter{
		writer: writer,
		code:   code,
		buf:    &buf,
	}
}

func (w *logResponseWriter) Write(bs []byte) (int, error) {
	w.buf.Write(bs)
	return w.writer.Write(bs)
}

func (w *logResponseWriter) Header() http.Header {
	return w.writer.Header()
}

func (w *logResponseWriter) WriteHeader(code int) {
	w.code = code
	w.writer.WriteHeader(code)
}

func (m *LogMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		//get方法不记录数据
		if strings.EqualFold(method, "get") {
			next(w, r)
			return
		}
		var dup io.ReadCloser
		dup, r.Body, _ = drainBody(r.Body)
		lwr := newLogResponseWriter(w, http.StatusOK)
		next(lwr, r)
		r.Body = dup
		logDetailLogic(r, lwr)
	}
}

func logDetailLogic(request *http.Request, response *logResponseWriter) {
	requestPath := request.RequestURI
	requestMethod := request.Method
	bs, _ := io.ReadAll(request.Body)
	str := strings.ReplaceAll(string(bs), "\n", "")
	result := response.buf.String()
	userInfo, err := common.GetJwtUserInfo(request.Context().Value("userInfo"))
	userId := int64(0)
	var userName string
	if userInfo != nil {
		userId = userInfo.UserId
		userName = userInfo.UserName
	}
	if err != nil {
	}
	fmt.Println("用户：", userId, "--", userName, "请求URL：", requestPath, "请求方式：", requestMethod, "请求内容：", str, "响应内容：", result)
	//var resp types.CommonResp
	//err = json.Unmarshal([]byte(result), &resp)
	//if resp.Code == 0 {
	//	//成功处理了数据  插入数据库
	//	operateLog := &model.SysmsOperateLog{
	//		OperateUser: userName,
	//		ReqData:     str,
	//		ReqMethod:   requestMethod,
	//		ReqPath:     requestPath,
	//		OperateType: requestMethod,
	//	}
	//	gplus.Insert(operateLog)
	//} else {
	//	log.Println("请求失败了")
	//}
}

// drainBody from httputil.drainBody
func drainBody(b io.ReadCloser) (r1, r2 io.ReadCloser, err error) {
	if b == nil || b == http.NoBody {
		return http.NoBody, http.NoBody, nil
	}
	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return nil, b, err
	}
	if err = b.Close(); err != nil {
		return nil, b, err
	}
	return io.NopCloser(&buf), io.NopCloser(bytes.NewReader(buf.Bytes())), nil
}
