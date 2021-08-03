package program_repo

import "time"

// 计划
//go:generate gormgen -structs Program -input .
type Program struct {
	Id                int32     //
	Name              string    //
	Slogan            string    // 标语
	CoverUrl          string    // 封面图
	PerformerCount    int32     // 参与人数
	EnrollPriceInCoin int32     // 报名费用（早币）
	CreatedAt         time.Time `gorm:"time"` //
	Type              int32     // 计划类型
	Order             int32     // 顺序
	Official          int32     // 是否官网创建
	IsDeleted         int32     //
	Tags              string    //
	Marking           int32     // 默认为1    2 为已经标记的
	UpdatedAt         time.Time `gorm:"time"` //
}
