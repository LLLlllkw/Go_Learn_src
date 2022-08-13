package data

import "courseWork/pkg/CRDerror"

type DataHandler interface {
	InsertNew(*Date, *PriceInfo) error
	Query(*Date) *PriceInfo
	Delete(*Date) error
}

type BUFFItemDataHandler struct {
	Id   string // Item Id
	Date *Date
}

type Date struct {
	Year  int
	Month int
	Day   int
}

type PriceInfo struct {
	Value float32
	Unit  string
}

func (*BUFFItemDataHandler) GetItemPrices(date *Date, price *PriceInfo) [] {
	// GET item prices for biz
	if _, err := CheckDate(date); err != nil {
		return CRDerror.errInsertFailure.Wrap(err)
	}
	return nil
}

func (*BUFFItemDataHandler) InsertNew(date *Date, price *PriceInfo) error {
	// insert a record into database
	if _, err := CheckDate(date); err != nil {
		return CRDerror.errInsertFailure.Wrap(err)
	}
	return nil
}

func (*BUFFItemDataHandler) Query(date *Date) *PriceInfo {
	// query a record into database
	if _, err := CheckDate(date); err != nil {
		return CRDerror.errInsertFailure.Wrap(err)
	}

	// query from sql handler
	return &PriceInfo{
		Value: 5,
		Unit:  "$",
	}
}

func (*BUFFItemDataHandler) Delete(date *Date, price *PriceInfo) error {
	// insert a record into database
	if _, err := CheckDate(date); err != nil {
		return CRDerror.errInsertFailure.Wrap(err)
	}
	// Delete from sql handler
	return nil
}

func CheckDate(date *Date) (bool, error) {
	// check whether date info is valid
	return true, nil
}
