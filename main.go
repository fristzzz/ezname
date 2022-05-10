// fristzzz 2022.5.10
// version 0.01
// LICENSE MIT
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	stu "./student"
)

func main() {
	var (
		F, C string
		//dataFile *os.File
	)
	flag.Parse()
	if flag.NArg() == 0 {
		help()
		return
	}

	fmt.Printf("Enter %s and %s:\n", stu.F, stu.C)
	fmt.Scan(&F, &C)

	fmt.Println(flag.Arg(0))
	if strings.EqualFold(flag.Arg(0), "add") {
		stu.AddS(F, C)
		fmt.Println("finish adding")
		return
	}

	if flag.NArg() != 2 {
		help()
	}
	path := flag.Arg(0)
	form := flag.Arg(1)

	// get data file
	wd, _ := os.Getwd()
	dataPath := filepath.Join(wd, F, C)
	dataFile, err := os.OpenFile(dataPath, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(`- get data file failed, please try:
$ ./ezname.exe add`)
		panic(err)
	}
	fmt.Println("initializing...")
	students := initStu(dataFile, F, C)
	dataFile.Close()

	// switch work directory to the specified one
	err = os.Chdir(path)
	if err != nil {
		fmt.Println("Open path failed.Please enter correct path")
		panic(err)
	}
	fmt.Println("switch work dir to ", path)

	//raname
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Read path failed, check the path, err: ", err)
		return
	}
	var (
		newName, oldName string
	)
	idReg := regexp.MustCompile(`\d{9}`)
	for i := 0; i < len(files); i++ {
		oldName = files[i].Name()
		fmt.Println("old name found:", oldName)
		ext := filepath.Ext(oldName)

		id := idReg.Find([]byte(oldName))
		newName = GetName(students, string(id), form)
		fmt.Println("new name:", newName)
		os.Rename(oldName, newName+ext)
	}
}

func help() {
	fmt.Println("- usage: $ ./rename.exe <path> <form>")
	fmt.Println(`- or you can input
$ rename.exe add to add student information`)
	fmt.Println(`- at the first time,
you should use "add" function to initialize the data`)
	return
}

func initStu(f *os.File, F, C string) []stu.Student {
	var (
		line     []byte
		err      error
		name, id string
	)
	s := make([]stu.Student, 0)
	reader := bufio.NewReader(f)
	for {
		line, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		id, name = string(line)[:9], string(line)[10:]
		s = append(s, *stu.NewStudent(F, C, name, id))
	}
	fmt.Println("data initialization succeed")
	return s
}
