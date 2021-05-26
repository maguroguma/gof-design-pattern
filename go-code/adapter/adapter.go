package adapter

// Adapteeを包む構造体
// Wrapperとも呼ばれる
type GetClient struct {
	gdc *GetDataClient
}

func NewGetClient() *GetClient {
	gc := new(GetClient)
	gc.gdc = NewGetDataClient()
	return gc
}

func (gc *GetClient) Get(key string) string {
	return gc.gdc.GetDataFromDB(key)
}
