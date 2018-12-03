package gwx

import (
	"github.com/jwei2006/gwx/context"
	"github.com/jwei2006/gwx/init"
	"github.com/jwei2006/gwx/gongzh"
)

func InitWechat(params *init.InitParams) error {
	return params.InitContext()
}
func Close() {
	init.FinallyClose()
}
func GetContext(tag string) (bool, *context.Context) {
	return init.GetContext(tag)
}

func CheckSignature(wechatCtx context.Context, signature, timestamp, nonce, echostr string) (bool, string){
	return gongzh.CheckSignature(wechatCtx, signature, timestamp, nonce, echostr)
}