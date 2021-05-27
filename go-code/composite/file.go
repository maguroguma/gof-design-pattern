package composite

import (
	"errors"
	"strconv"
)

type File struct {
	name string
	size int
}

func NewFile(name string, size int) *File {
	f := new(File)
	f.name, f.size = name, size
	return f
}

func (f *File) GetName() string {
	return f.name
}
func (f *File) GetSize() int {
	return f.size
}
func (f *File) Add(e Entry) (error, Entry) {
	return errors.New("This is File, not Directory"), nil
}
func (f *File) PrintListDefault() string {
	return f.PrintList("")
}
func (f *File) PrintList(prefix string) string {
	return prefix + "/" + f.String()
}
func (f *File) String() string {
	return f.GetName() + " (" + strconv.Itoa(f.GetSize()) + ")\n"
}
