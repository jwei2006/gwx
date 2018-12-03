package gongzh

import (
	"fmt"
	"github.com/jwei2006/gwx/context"
	"github.com/jwei2006/gwx/util"
)

func CheckSignature(wechatCtx context.Context, signature, timestamp, nonce, echostr string) (bool, string) {
	s := util.Signature(timestamp, nonce, wechatCtx.Token)
	right := signature == s
	fmt.Println(right)
	return right, echostr
}