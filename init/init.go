package init

import (
	"errors"
	"github.com/go-xorm/xorm"
	"github.com/jwei2006/gwx/context"
)

var contextMap = make(map[string]context.Context, 0)
var engine *xorm.Engine
var isDbAttend = false
var attendTable string

type InitParams struct {
	//1来自于数据库，2直接传入
	ContextType int8
	//关注记录数据库
	DbAttend bool
	//如果来自于数据库，则不传
	Apps []context.Context
	DB   struct {
		Init struct {
			DbType string
			Url    string
			GetSql string
		}
		Attend struct {
			//与init相同
			IsInit bool
			DbType string
			Url    string
			Table  string
		}
	}
}

func (params *InitParams) InitContext() error {
	var initEngine *xorm.Engine
	e := doInitContext(params, initEngine)
	if e != nil {return e}
	e = doInitAttend(params, initEngine)
	return e
}
func FinallyClose() {
	closeDb(engine)
}
func GetContext(tag string) (bool, *context.Context) {
	context, has := contextMap[tag]
	return has, &context
}
func IsDbAttend() bool {
	return isDbAttend
}
func GetAttendTable() string {
	return attendTable
}

func closeDb(_engine *xorm.Engine) {
	if _engine != nil {
		_engine.Close()
	}
}
func initContext(apps []context.Context) error {
	if apps == nil || len(apps) == 0 {
		return errors.New("参数错误：上下文")
	}
	for _, item := range apps {
		contextMap[item.Tag] = item
	}
	return nil
}
func doInitContext(params *InitParams, initEngine *xorm.Engine) error {
	var e error
	if params.ContextType == 1 {
		if params.DB.Init.DbType == "" || params.DB.Init.Url == "" {
			return errors.New("参数错误:数据库")
		}
		initEngine, e = xorm.NewEngine(params.DB.Init.DbType, params.DB.Init.Url)
		if e != nil {
			return e
		}
		var apps = make([]context.Context, 0)
		e := initEngine.SQL(params.DB.Init.GetSql).Find(&apps)
		if e != nil{return e}
		e = initContext(params.Apps)
		if e != nil {
			return e
		}
	} else {
		e := initContext(params.Apps)
		if e != nil {
			return e
		}
	}
	return e
}
func doInitAttend(params *InitParams, initEngine *xorm.Engine) error {
	var e error
	isDbAttend = params.DbAttend
	attendTable = params.DB.Attend.Table
	if !params.DbAttend {
		defer closeDb(initEngine)
	} else {
		if !params.DB.Attend.IsInit {
			defer closeDb(initEngine)
			if params.DB.Attend.Url == "" || params.DB.Attend.DbType == "" {
				return errors.New("参数错误：关注数据库")
			}
			engine, e = xorm.NewEngine(params.DB.Attend.DbType, params.DB.Attend.Url)
			if e != nil {
				return e
			}
		} else {
			engine = initEngine
		}
	}
	return nil
}