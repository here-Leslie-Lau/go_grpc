package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	. "go_grpc/model"
)

type OrderDao struct {
	Db *gorm.DB
}

func (o *OrderDao) Init() {
	var err error

	o.Db, err = gorm.Open("mysql", DatabaseUrl)

	if err != nil {
		panic("连接数据库失败:" + err.Error())
	}

	o.Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Order{})
}

//添加
func (o OrderDao) AddOrder(order Order) error {
	tx := o.Db.Begin()
	defer tx.Close()

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}