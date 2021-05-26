package prototype

import "strconv"

type AddString struct {
	base string
}

func NewAddString(base string) *AddString {
	as := new(AddString)
	as.base = base
	return as
}

func (as *AddString) Use(num int) string {
	return as.base + " + " + strconv.Itoa(num)
}

func (as *AddString) CreateClone() ClonableObject {
	newAs := new(AddString)
	newAs.base = as.base
	return newAs
}
