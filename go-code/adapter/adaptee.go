package adapter

// Adapterによって包み込まれる構造体
// Adapteeと呼ばれる
type GetDataClient struct {
	db map[string]string
}

func NewGetDataClient() *GetDataClient {
	gdc := new(GetDataClient)
	gdc.db = map[string]string{
		"キー": "バリュー",
	}
	return gdc
}

func (gdc *GetDataClient) GetDataFromDB(key string) string {
	return gdc.db[key]
}
