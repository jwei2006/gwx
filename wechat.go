package main
//package gwx
//
//import (
//	"github.com/jwei2006/gwx/context"
//	"github.com/jwei2006/gwx/system"
//	"github.com/jwei2006/gwx/gongzh"
//)
//
//func InitWechat(params *system.InitParams) error {
//	return params.InitContext()
//}
//func Close() {
//	system.FinallyClose()
//}
//func GetContext(tag string) (bool, *context.Context) {
//	return system.GetContext(tag)
//}
//
//func CheckSignature(wechatCtx context.Context, signature, timestamp, nonce, echostr string) (bool, string){
//	return gongzh.CheckSignature(wechatCtx, signature, timestamp, nonce, echostr)
//}