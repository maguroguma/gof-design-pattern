package factory_method

import "strconv"

type stationaryGameFactory struct {
	nextProdNumber int
	record         []product
}

func NewStationaryGameFactory() *stationaryGameFactory {
	sgf := new(stationaryGameFactory)
	sgf.nextProdNumber = 1
	sgf.record = []product{}
	return sgf
}

func (sgf *stationaryGameFactory) createProduct() product {
	p := newStationaryGame(sgf.nextProdNumber)
	sgf.nextProdNumber++
	return p
}

func (sgf *stationaryGameFactory) registerProduct(p product) {
	sgf.record = append(sgf.record, p)
}

type stationaryGame struct {
	prodNumber int
}

func newStationaryGame(prodNumber int) *stationaryGame {
	sg := new(stationaryGame)
	sg.prodNumber = prodNumber
	return sg
}

func (sg *stationaryGame) use() string {
	return "[Start] Stationary game No." + strconv.Itoa(sg.prodNumber)
}
