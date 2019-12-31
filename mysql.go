package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/guregu/null.v3"
	"time"
)

type CategoryLimit struct {
	CategoryID           int        `gorm:"column:category_id;primary_key" json:"category_id"`
	CreationDate         time.Time  `gorm:"column:creation_date" json:"-"`
	FreeShippingFee      int        `gorm:"column:free_shipping_fee" json:"free_shipping_fee"`
	IsCoverAll           int        `gorm:"column:is_cover_all" json:"is_cover_all"`
	IsSupportInstallment int        `gorm:"column:is_support_installment" json:"is_support_installment"`
	LastChangedDate      time.Time  `gorm:"column:last_changed_date" json:"-"`
	Recyclable           null.Int   `gorm:"column:recyclable" json:"recyclable"`
	SupportSevenDays     null.Int   `gorm:"column:support_seven_days" json:"support_seven_days"`
	TaxRate              null.Float `gorm:"column:tax_rate" json:"tax_rate"`
}

func main() {
	cate := new(CategoryLimit)
	query(cate, "category_id = 1")
	fmt.Println(cate)
}

func query(obj interface{}, sql string) {
	conn, err := gorm.Open("mysql", "writeuser:ddbackend@tcp(10.255.255.22)/ProductDB?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	data := obj.(*CategoryLimit)
	conn.Table("category_limit").Where(sql).Find(&data)
	fmt.Println(data)
}
