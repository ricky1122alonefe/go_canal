package module

// type BaseTable struct {
// 	// 版本号
// 	Version string `gorm:"column:_VERSION" json:"_VERSION"`
// 	// 创建时间
// 	CreatedTime time.Time `gorm:"column:CREATED_TIME" json:"CREATED_TIME"`
// 	// 创建用户
// 	CreatedUser string `gorm:"column:CREATED_USER" json:"CREATED_USER"`
// 	// 创建的用户id
// 	CreatedUserId int `gorm:"column:CREATED_USER_ID" json:"CREATED_USER_ID"`
// 	// 更新时间
// 	UpdatedTime time.Time `gorm:"column:UPDATED_TIME" json:"UPDATED_TIME"`
// 	// 更新用户
// 	UpdatedUser string `gorm:"column:UPDATED_USER" json:"UPDATED_USER"`
// 	// 修改的用户id
// 	UpdatedUserId int `gorm:"column:UPDATED_USER_ID" json:"UPDATED_USER_ID"`
// }

type TBAResource struct {
	// BaseTable

	// 编号
	Id int `gorm:"column:ID" json:"ID"`
	// 资源名称
	ResName string `gorm:"column:NAME" json:"NAME"`
	// 资源类型
	ResType string `gorm:"column:TYPE" json:"TYPE"`
	// 父资源编号
	ResParentId int `gorm:"column:PARENT_ID" json:"PARENT_ID"`
	// 资源路径
	Path string `gorm:"column:PATH" json:"PATH"`
	// 资源图标
	Icon string `gorm:"column:ICON" json:"ICON"`
	// 功能编码
	FuncCode string `gorm:"column:FUNC_CODE" json:"FUNC_CODE"`
	// 排序值
	OrderValue int `gorm:"column:ORDER_VALUE" json:"ORDER_VALUE"`
	// 对应系统id
	ProductionId int `gorm:"column:PRODUCTION_ID" json:"PRODUCTION_ID"`
	// 对应系统code
	ProductionCode string `gorm:"column:PRODUCTION_CODE" json:"PRODUCTION_CODE"`
	// 内部代码
	InnerCode string `gorm:"column:INNER_CODE" json:"INNER_CODE"`
	// 平台id
	PlatformId int `gorm:"column:PLATFORM_ID" json:"PLATFORM_ID"`
}

