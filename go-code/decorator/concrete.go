package decorator

// ConcreteComponent

type VanillaIcecream struct{}

func NewVanillaIcecream() *VanillaIcecream {
	return new(VanillaIcecream)
}

func (vi *VanillaIcecream) getName() string {
	return "バニラアイスクリーム"
}

func (vi *VanillaIcecream) howSweet() string {
	return "バニラ味"
}

type GreenTeaIcecream struct{}

func NewGreenTeaIcecream() *GreenTeaIcecream {
	return new(GreenTeaIcecream)
}

func (gi *GreenTeaIcecream) getName() string {
	return "抹茶アイスクリーム"
}

func (gi *GreenTeaIcecream) howSweet() string {
	return "抹茶味"
}

// Decorater, ConcreteDecorator

type CashewNutsToppingIcecream struct {
	ice Icecream
}

func NewCashewNutsToppingIcecream(ice Icecream) *CashewNutsToppingIcecream {
	i := new(CashewNutsToppingIcecream)
	i.ice = ice
	return i
}

func (ci *CashewNutsToppingIcecream) getName() string {
	return "カシューナッツ" + ci.ice.getName()
}

func (ci *CashewNutsToppingIcecream) howSweet() string {
	return ci.ice.howSweet()
}
