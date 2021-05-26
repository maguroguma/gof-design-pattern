package prototype

type MultiplyString struct {
	base string
}

func NewMultiplyString(base string) *MultiplyString {
	ms := new(MultiplyString)
	ms.base = base
	return ms
}

func (ms *MultiplyString) Use(num int) string {
	res := ms.base
	for i := 0; i < num-1; i++ {
		res += " " + ms.base
	}
	return res
}

func (ms *MultiplyString) CreateClone() ClonableObject {
	newMs := new(MultiplyString)
	newMs.base = ms.base
	return newMs
}
