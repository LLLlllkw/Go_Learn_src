package biz

import "courseWork/pkg/data"
import "courseWork/pkg/CRDerror"

type ItemDetail struct {
	Id    string
	Price float32
	Date  *PriceDate
}
type PriceDate struct {
	Year  int
	Month int
	Day   int
}

type ItemPrices []map[*PriceDate]*ItemDetail

func GetItemPricesById(Id string, Date *PriceDate) []ItemDetail,error {
	if Date.Year >= 2022 || Date.Year <= 2018 {
		return nil,CRDerror.errWrongDate
	}
	handler := data.BUFFItemDataHandler{Id: Id, Date: Date}

}
