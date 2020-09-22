package conf

type CananConfig struct {
	// 访问数据库地址
	Address string
	// 用户名
	UserName string
	//  密码
	Password string
	// schema - db_table list
	SchemaInfo map[string][]string
	// 延迟写入
	Delay int
	// 写入 方向 暂定redis
}


