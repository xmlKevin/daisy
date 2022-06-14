package knowledge

import (
	"daisy/models/base"
)

/*
  @Author : xmlKevin
*/

// 流程分类
type KM_Comment struct {
	base.Model
	Name string `gorm:"column:name; type: varchar(128)" json:"name" form:"name"` // 分类名称
	//userId,targetId类型改变
	parentCommentId int `gorm:"column:parentCommentId; type: int(11)" json:"parentCommentId" form:"parentCommentId"` // 创建者
	documentId      int `gorm:"column:documentId; type: int(11)" json:"documentId" form:"documentId"`                // 创建者
	createUserId    int `gorm:"column:createUserId; type: int(11)" json:"createUserId" form:"createUserId"`          // 创建者
	replyUserId     int `gorm:"column:replyUserId; type: int(11)" json:"replyUserId" form:"replyUserId"`             // 创建者
	html            int `gorm:"column:html; type: mediumtext" json:"html" form:"html"`                               // 创建者
	userAgent       int `gorm:"column:userAgent; type: mediumtext" json:"userAgent" form:"userAgent"`                // 创建者
	//collectType名称类型都改变
	pass int `gorm:"column:pass; type: boolean" json:"pass" form:"pass"` // 收藏目标类型
}

func (KM_Comment) TableName() string {
	return "km_commnet"
}
