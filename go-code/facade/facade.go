package facade

// HTTPリクエストとDBクエリを束ねる窓口
type Facade struct{}

func NewFacade() *Facade {
	return new(Facade)
}

func (f *Facade) Get() string {
	hc := NewHttpClient()
	dc := NewDbClient()

	res := ""
	res += hc.GetData()
	res += " と "
	res += dc.SelectData()
	res += " で作られた何か"

	return res
}
