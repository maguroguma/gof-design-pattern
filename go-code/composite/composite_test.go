package composite

import "testing"

func TestCompositePattern(t *testing.T) {
	println("Making root entries...")
	rootdir := NewDirectory("root")
	bindir := NewDirectory("bin")
	tmpdir := NewDirectory("tmp")
	usrdir := NewDirectory("usr")
	rootdir.Add(bindir)
	rootdir.Add(tmpdir)
	rootdir.Add(usrdir)
	bindir.Add(NewFile("vi", 10000))
	bindir.Add(NewFile("latex", 20000))
	actual := rootdir.PrintListDefault()
	println(actual)

	println("Making user entries...")
	suzukaze := NewDirectory("suzukaze")
	yagami := NewDirectory("yagami")
	takimoto := NewDirectory("takimoto")
	usrdir.Add(suzukaze)
	usrdir.Add(yagami)
	usrdir.Add(takimoto)
	suzukaze.Add(NewFile("diary.html", 100))
	suzukaze.Add(NewFile("Composite.java", 200))
	yagami.Add(NewFile("memo.tex", 300))
	takimoto.Add(NewFile("game.doc", 400))
	takimoto.Add(NewFile("junk.mail", 500))
	actual = rootdir.PrintListDefault()
	println(actual)
}
