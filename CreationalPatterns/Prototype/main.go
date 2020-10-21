package main

import "fmt"

type inode interface {
	print(string)
	clone() inode
}

type file struct {
	name string
}

func (f *file) print(indentation string)  {
	fmt.Println(indentation + f.name + "_clone")
}

func (f *file) clone() inode {
	return &file{name: f.name}
}

type folder struct {
	childrens []inode
	name      string
}

func (f *folder) print(indentation string)  {
	fmt.Println(indentation + f.name + "_clone")
	for _, i := range f.childrens{
		i.print(indentation + indentation)
	}
}

func (f *folder) clone() inode {
	cloneFolder := &folder{name: f.name}
	var tempChildrens []inode
	for _, i := range f.childrens {
		copy := i.clone()
		tempChildrens = append(tempChildrens, copy)
	}
	cloneFolder.childrens = tempChildrens
	return cloneFolder
}

func main()  {
	file1 := &file{name: "File1"}
	file2 := &file{name: "File2"}
	file3 := &file{name: "File3"}
	folder1 := &folder{
		childrens: []inode{file1},
		name:      "Folder1",
	}

	folder2 := &folder{
		childrens: []inode{folder1, file2, file3},
		name:      "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("  ")
	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("  ")
}