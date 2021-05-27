package builder

// builderに具体的な指示を出すもの
type Director struct {
	b builder
}

func NewDirector(b builder) *Director {
	d := new(Director)
	d.b = b
	return d
}

// テンプレートメソッドに該当する
func (d *Director) Construct(content string) string {
	res := ""

	res += d.b.opening() + "\n"
	res += content + "\n"
	res += d.b.closing()

	return res
}
