package student

import (
	"fmt"
	"os"
	"path/filepath"
)

// 添加姓名学号
func AddS(F, C string) error {
	const space = " "
	var (
		fpath string
		name  string
		id    string
		err   error
	)

	// 创建院系/班级文件
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fpath = filepath.Join(wd, F)
	fmt.Println("fpath: ", fpath)
	err = os.MkdirAll(fpath, 0666)
	if err != nil {
		fmt.Println("- make directory failed, err: ", err)
		return nil
	}
	os.Chdir(fpath)
	os.Create(C)
	fmt.Printf("- File ./ %s / %s have been made successfully.\n",
		F, C)

	stuFile, err := os.OpenFile(C, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(C, "open failed, please check the file.")
		fmt.Println(err)
	}
	defer stuFile.Close()

	// 学号 姓名
	var wContent []byte
	for {
		fmt.Println("- Enter <name> <id>:")
		fmt.Scan(&name, &id)

		// 输入q退出
		if name == "q" || id == "q" {
			return err
		}

		wContent = []byte(id + space + name)
		fmt.Println("write content:", string(wContent))
		_, err = stuFile.Write(wContent)
		if err != nil {
			fmt.Println("write failed, err: ", err)
			return err
		}
		stuFile.Write([]byte("\n"))
		fmt.Println("ok")

	}
}
