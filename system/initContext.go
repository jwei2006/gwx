package system

func InitContext() error {
	params := new(InitParams)
	params.ContextType = 1
	params.DB.Init.DbType = "mysql"
	params.DB.Init.Url = "root:qweqwe@tcp(10.1.16.10:3306)/jw_test2?charset=utf8&parseTime=True&loc=Local"
	params.DB.Init.GetSql = "select * from t_wx where status = 11"
	params.DbAttend = false
	e := params.InitContext()
	return e
}