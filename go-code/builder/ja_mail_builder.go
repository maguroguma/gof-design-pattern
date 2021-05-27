package builder

type JaMailBuilder struct{}

func NewJaMailBuilder() *JaMailBuilder {
	return new(JaMailBuilder)
}

func (jmb *JaMailBuilder) opening() string {
	return "拝啓、"
}

func (jmb *JaMailBuilder) closing() string {
	return "敬具。"
}
