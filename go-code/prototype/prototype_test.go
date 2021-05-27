package prototype

import "testing"

func TestPrototypePattern(t *testing.T) {
	m := NewManager()
	AStr := "AAA + 3"
	BStr := "BBB BBB BBB"

	// 登録
	orgA := NewAddString("AAA")
	m.Register("A", orgA)
	orgB := NewMultiplyString("BBB")
	m.Register("B", orgB)

	// クローンの生成
	cloneA := m.Create("A")
	cloneB := m.Create("B")

	// 検証
	if orgA.Use(3) != AStr {
		t.Errorf("Error")
	}
	if cloneA.Use(3) != AStr {
		t.Errorf("Error")
	}
	if orgB.Use(3) != BStr {
		t.Errorf("Error")
	}
	if cloneB.Use(3) != BStr {
		t.Errorf("Error")
	}

	// クローンはオリジナルとは別であることを検証
	orgA.base = "A"
	orgB.base = "B"
	if orgA.Use(3) != "A + 3" {
		t.Errorf("Error")
	}
	if cloneA.Use(3) != AStr {
		t.Errorf("Error")
	}
	if orgB.Use(3) != "B B B" {
		t.Errorf("Error")
	}
	if cloneB.Use(3) != BStr {
		t.Errorf("Error")
	}
}
