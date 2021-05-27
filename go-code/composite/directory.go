package composite

import (
	"strconv"
)

type Directory struct {
	name      string
	directory []Entry
}

func NewDirectory(name string) *Directory {
	d := new(Directory)
	d.name = name
	d.directory = []Entry{}
	return d
}

func (d *Directory) GetName() string {
	return d.name
}
func (d *Directory) GetSize() int {
	size := 0
	for _, e := range d.directory {
		size += e.GetSize()
	}
	return size
}
func (d *Directory) Add(e Entry) (error, Entry) {
	d.directory = append(d.directory, e)
	return nil, d
}
func (d *Directory) PrintListDefault() string {
	return d.PrintList("")
}
func (d *Directory) PrintList(prefix string) string {
	res := prefix + "/" + d.String()
	for _, e := range d.directory {
		res += e.PrintList(prefix + "/" + d.name)
	}
	return res
}
func (d *Directory) String() string {
	return d.GetName() + " (" + strconv.Itoa(d.GetSize()) + ")\n"
}
