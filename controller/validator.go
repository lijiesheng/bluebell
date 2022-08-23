package controller

import ut "github.com/go-playground/universal-translator"

// 定义一个全局翻译器T
var trans ut.Translator

//// 去除提示信息中的结构体名称  todo 不懂
//func removeTopStruct(fields map[string]string) map[string]string {
//	res := map[string]string{}
//
//}