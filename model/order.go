package model

import (
	"github.com/jinzhu/gorm"
	"strconv"
)

type Order struct {
	gorm.Model
	orderNo string `gorm:"not null;unique" form:"orderNo"`
	userName string `gorm:"not null;unique" form:"userName"`
	amount float64 `gorm:"type:DECIMAL(10,2);not null" form:"amount"`
	status string `gorm:"not null" form:"status"`
	fileUrl string `gorm:"size:255" form:"fileUrl"`
}

//重写String方法
func (o Order) String() string {
	return "(orderNo:" + o.orderNo + ", userName:" + o.userName +
		", amount:" + strconv.FormatFloat(o.amount, 'f', -1, 64) + ", status:" + o.status +
		", fileUrl:" + o.fileUrl + ")"
}