package model

// HouseInfo 在售房屋数据model，数据库字段格式必须为首字母大写的驼峰格式
// 在结构体内可以继承gorm.Model的模型 ，gorm.Model 会自带ID, CreatedAt, UpdatedAt, DeletedAt这4个字段
type HouseInfo struct {
	Id            string `gorm:"varchar(64); primary_key ;comment: '房子id'"`
	TotalPrice    string `gorm:"varchar(64); comment: '房子总价'"`
	UnitPrice     string `gorm:"varchar(64); comment: '房子单价'"`
	RoomInfo      string `gorm:"varchar(64); comment:'房屋户型信息'"`
	AreaInfo      string `gorm:"varchar(64); comment:'房屋面积信息'"`
	AreaName      string `gorm:"varchar(64); comment:'小区名称'"`
	CommunityInfo string `gorm:"varchar(64); comment:'详细区域'"`
	Transaction   string `gorm:"varchar(64); comment:'交易信息'"`
}
