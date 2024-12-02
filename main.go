package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	customFormat := "2006-01-02_150405"
	root := "/home/enma/code/gitHub/dnkFuns/dnkBook/2024"

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 只管理文件，忽略目录
		if !info.IsDir() {
			createdTime := info.ModTime()

			fmt.Printf("File name: %s\n", info.Name())
			fmt.Printf("File path：%s\n", path)
			fmt.Printf("File created time：%s\n", createdTime.Format(customFormat))

			parts := strings.Split(info.Name(), ".")
			if len(parts) != 2 {
				panic("no file name extension")
			}
			err := os.Rename(path, root+"/"+createdTime.Format(customFormat)+"."+parts[1])
			if err != nil {
				panic(err)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}
