/*
How would you implement a program to list go source code files in a given directory including all 
sub-directories recursively?
Please use filepath.WalkDir() function (https://pkg.go.dev/path/filepath#WalkDir) for this purpose.
*/
package main

import (
	"flag"
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
    flag.Parse()
    root := flag.Arg(0)
    err := filepath.WalkDir(root, visit)
    fmt.Printf("filepath.WalkDir() returned %v\n", err)
}
func visit(path string, di fs.DirEntry, err error) error { 
    fmt.Printf("Visited: %s\n", path)
    return nil
}

/*

go run WalkDir.go ex    

Visited: ex
Visited: ex/.DS_Store
Visited: ex/ex2
Visited: ex/ex2/.DS_Store
Visited: ex/ex2/example2.txt
Visited: ex/example.txt
filepath.WalkDir() returned <nil>

*/