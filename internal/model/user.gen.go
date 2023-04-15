// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "User"

// User mapped from table <User>
type User struct {
	ID           int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                      // 主键
	Name         string    `gorm:"column:name;not null" json:"name"`                                       // 姓名
	StudentID    string    `gorm:"column:student_id;not null" json:"student_id"`                           // 学号
	Gender       bool      `gorm:"column:gender;not null" json:"gender"`                                   // 性别：0表示女，1表示男
	Class        string    `gorm:"column:class;not null" json:"class"`                                     // 班级
	Position     string    `gorm:"column:position" json:"position"`                                        // 学生职务
	Mobile       string    `gorm:"column:mobile;not null" json:"mobile"`                                   // 电话号码
	Group        string    `gorm:"column:group;not null" json:"group"`                                     // 报名组别
	Introduction string    `gorm:"column:introduction" json:"introduction"`                                // 个人简介
	Awards       string    `gorm:"column:awards" json:"awards"`                                            // 所获奖项
	Reason       string    `gorm:"column:reason" json:"reason"`                                            // 加入原因
	CreatedAt    time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"` // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"` // 最后修改时间
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
