package logger
//
//import (
//	"fmt"
//	"go.uber.org/zap"
//	"net/http"
//	"testing"
//)
//
//func TestInitLogger(t *testing.T) {
//	simpleHttpGet("www.google.com")
//	simpleHttpGet("http://www.baidu.com")
//}
//
//func simpleHttpGet(uri string) {
//	initLogger()
//	fmt.Println("logger==>", logger)
//	resp, err := http.Get(uri)
//	if err != nil {
//		logger.Error( // 输出 Error 日志
//			"Error fetching url...",
//			zap.String("url : ", "http://www.google.com"),
//			zap.Error(err))
//		// {"level":"error","ts":1660578643.271744,"caller":"logger/logger_test.go:20","msg":"Error fetching url...","url : "
//	} else {
//		logger.Info( // 输出 Info 日志
//			"Success..",
//			zap.String("statusCode", resp.Status),
//			zap.String("url", "www.google.com"))
//		// {"level":"info","ts":1660578643.296084,"caller":"logger/logger_test.go:25","msg":"Success..","statusCode":"200 OK",
//	}
//}
