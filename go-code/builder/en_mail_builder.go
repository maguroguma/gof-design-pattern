package builder

type EnMailBuilder struct{}

func NewEnMailBuilder() *EnMailBuilder {
	return new(EnMailBuilder)
}

func (emb *EnMailBuilder) opening() string {
	return "Hello,"
}

func (emb *EnMailBuilder) closing() string {
	return "Best Regards."
}
