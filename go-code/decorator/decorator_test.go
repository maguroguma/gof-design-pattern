package decorator

import "testing"

func TestDecorator(t *testing.T) {
	vanilla := NewVanillaIcecream()
	green := NewGreenTeaIcecream()
	decoratedVanilla := NewCashewNutsToppingIcecream(vanilla)
	decoratedGreen := NewCashewNutsToppingIcecream(green)

	testcases := []struct {
		ice   Icecream
		name  string
		taste string
	}{
		{
			ice:   vanilla,
			name:  "バニラアイスクリーム",
			taste: "バニラ味",
		},
		{
			ice:   green,
			name:  "抹茶アイスクリーム",
			taste: "抹茶味",
		},
		{
			ice:   decoratedVanilla,
			name:  "カシューナッツバニラアイスクリーム",
			taste: "バニラ味",
		},
		{
			ice:   decoratedGreen,
			name:  "カシューナッツ抹茶アイスクリーム",
			taste: "抹茶味",
		},
	}

	for _, tc := range testcases {
		actualName := tc.ice.getName()
		actualTaste := tc.ice.howSweet()
		if actualName != tc.name {
			t.Errorf("want %v, got %v", actualName, tc.name)
		}
		if actualTaste != tc.taste {
			t.Errorf("want %v, got %v", actualTaste, tc.taste)
		}
	}
}
