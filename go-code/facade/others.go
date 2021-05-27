package facade

type HttpClient struct{}

func NewHttpClient() *HttpClient {
	return new(HttpClient)
}

func (hc *HttpClient) GetData() string {
	return "(HTTP経由で得られた何か)"
}

type DbClient struct{}

func NewDbClient() *DbClient {
	return new(DbClient)
}

func (dc *DbClient) SelectData() string {
	return "(DBから得られた何か)"
}
