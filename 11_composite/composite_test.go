package composite

import "fmt"

// ExampleComposite 使用者(client)
func ExampleComposite() {
	fmt.Println("Making root entries...")
	rootDir := NewDirectory("root")
	binDir := NewDirectory("bin")
	tmpDir := NewDirectory("tmp")
	usrDir := NewDirectory("usr")
	rootDir.Add(binDir)
	rootDir.Add(tmpDir)
	rootDir.Add(usrDir)
	binDir.Add(NewFile("vi", 10000))
	binDir.Add(NewFile("latex", 20000))
	rootDir.PrintList()

	fmt.Println()
	fmt.Println("Making user entries...")
	yuki := NewDirectory("yuki")
	hanako := NewDirectory("hanako")
	tomura := NewDirectory("tomura")
	usrDir.Add(yuki)
	usrDir.Add(hanako)
	usrDir.Add(tomura)
	yuki.Add(NewFile("diary.html", 100))
	yuki.Add(NewFile("Composite.go", 200))
	hanako.Add(NewFile("memo.tex", 300))
	tomura.Add(NewFile("gome.doc", 400))
	tomura.Add(NewFile("junk.mail", 500))
	rootDir.PrintList()

	// Output:
	// Making root entries...
	// /root (30000)
	// /root/bin (30000)
	// /root/bin/vi (10000)
	// /root/bin/latex (20000)
	// /root/tmp (0)
	// /root/usr (0)
	//
	// Making user entries...
	// /root (31500)
	// /root/bin (30000)
	// /root/bin/vi (10000)
	// /root/bin/latex (20000)
	// /root/tmp (0)
	// /root/usr (1500)
	// /root/usr/yuki (300)
	// /root/usr/yuki/diary.html (100)
	// /root/usr/yuki/Composite.go (200)
	// /root/usr/hanako (300)
	// /root/usr/hanako/memo.tex (300)
	// /root/usr/tomura (900)
	// /root/usr/tomura/gome.doc (400)
	// /root/usr/tomura/junk.mail (500)
}
