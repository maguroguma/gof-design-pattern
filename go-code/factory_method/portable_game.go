package factory_method

import "strconv"

type portableGameFactory struct {
	nextProdNumber int
	record         []product
}

func NewPortableGameFactory() *portableGameFactory {
	pgf := new(portableGameFactory)
	pgf.nextProdNumber = 1
	pgf.record = []product{}
	return pgf
}

func (pgf *portableGameFactory) createProduct() product {
	p := newPortableGame(pgf.nextProdNumber)
	pgf.nextProdNumber++
	return p
}

func (pgf *portableGameFactory) registerProduct(p product) {
	pgf.record = append(pgf.record, p)
}

type portableGame struct {
	prodNumber int
}

func newPortableGame(prodNumber int) *portableGame {
	pg := new(portableGame)
	pg.prodNumber = prodNumber
	return pg
}

func (pg *portableGame) use() string {
	return "[Start] Portable game No." + strconv.Itoa(pg.prodNumber)
}
