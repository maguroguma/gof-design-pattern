package adapter

import "testing"

func TestAdapterPattern(t *testing.T) {
	// 想像: 古いコードでは具象クラスが使われていたが、
	// リファクタリング後はinterfaceを用いたい。
	// getter := GetDataClient{}

	var getter Getter
	getter = NewGetClient()

	actual := getter.Get("キー")
	if actual != "バリュー" {
		t.Errorf("Error")
	}
}
