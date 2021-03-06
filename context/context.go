package context

import (
	"net/http"
	"sync"
)

// Context struct
type Context struct {
	Tag            string
	AppID          string
	AppSecret      string
	Token          string
	EncodingAESKey string
	PayMchID       string
	PayNotifyURL   string
	PayKey         string

	//Writer  http.ResponseWriter
	//Request *http.Request

	//accessTokenLock 读写锁 同一个AppID一个
	//accessTokenLock *sync.RWMutex

	//jsAPITicket 读写锁 同一个AppID一个
	//jsAPITicketLock *sync.RWMutex
}

// Query returns the keyed url query value if it exists
func (ctx *Context) Query(req *http.Request, writer  http.ResponseWriter, key string) string {
	value, _ := ctx.GetQuery(req, writer, key)
	return value
}

// GetQuery is like Query(), it returns the keyed url query value
func (ctx *Context) GetQuery(req *http.Request, writer  http.ResponseWriter, key string) (string, bool) {
	if values, ok := req.URL.Query()[key]; ok && len(values) > 0 {
		return values[0], true
	}
	return "", false
}

// SetJsAPITicketLock 设置jsAPITicket的lock
func (ctx *Context) SetJsAPITicketLock(lock *sync.RWMutex) {
	//ctx.jsAPITicketLock = lock
}

// GetJsAPITicketLock 获取jsAPITicket 的lock
func (ctx *Context) GetJsAPITicketLock() *sync.RWMutex {
	//return ctx.jsAPITicketLock
	return nil
}
