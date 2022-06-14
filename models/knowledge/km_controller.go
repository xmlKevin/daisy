package knowledge

import (
	"daisy/models/base"
)

/*
  @Author : xmlKevin
*/

// 流程分类
type KM_Controller struct {
	base.Model
	Name string `gorm:"column:name; type: varchar(128)" json:"name" form:"name"` // 分类名称
	//userId,targetId类型改变
	userId   int `gorm:"column:userId; type: int(11)" json:"userId" form:"userId"`       // 创建者
	targetId int `gorm:"column:targetId; type: int(11)" json:"targetId" form:"targetId"` // 收藏目标 Id
	//collectType名称类型都改变
	collectType int `gorm:"column:collectType; type: int(11)" json:"collectType" form:"collectType"` // 收藏目标类型
}

func (KM_Controller) TableName() string {
	return "km_controller"
}
