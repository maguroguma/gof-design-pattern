package factory_method

// Template Method部分を担当する構造体
type factoryManager struct {
	f factory
}

func NewFactoryManager(f factory) *factoryManager {
	fm := new(factoryManager)
	fm.f = f
	return fm
}

// Template Methodによる生成処理（Factory Method）
func (fm *factoryManager) FactoryMethod() product {
	p := fm.f.createProduct()
	fm.f.registerProduct(p)
	return p
}
