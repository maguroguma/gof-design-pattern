package factory_method

// factory: 何らかの工場、生成処理に加えて他の複雑な処理も行う
type factory interface {
	createProduct() product
	registerProduct(product)
}

// product: 生成されるもの、振る舞いを抽出しておく
type product interface {
	use() string
}
