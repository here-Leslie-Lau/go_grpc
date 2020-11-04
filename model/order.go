package model

import (
	"github.com/jinzhu/gorm"
	"strconv"
)

type Order struct {
	gorm.Model
	OrderNo string `gorm:"not null;unique" form:"orderNo"`
	UserName string `gorm:"not null;unique" form:"userName"`
	Amount float64 `gorm:"type:DECIMAL(10,2);not null" form:"amount"`
	Status string `gorm:"not null" form:"status"`
	FileUrl string `gorm:"size:255" form:"fileUrl"`
}

//重写String方法
func (o Order) String() string {
	return "(orderNo:" + o.OrderNo + ", userName:" + o.UserName +
		", amount:" + strconv.FormatFloat(o.Amount, 'f', -1, 64) + ", status:" + o.Status +
		", fileUrl:" + o.FileUrl + ")"
}