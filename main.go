package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	customFormat = "2006-01-02_150405"
	createYear   = "2006"
	filesPath    = "/home/enma/code/gitHub/dnkFuns/dnkBook/temp"
)

func moveFiles(fcY string, fileName string) error {
	pathDir := "/home/enma/code/gitHub/dnkFuns/dnkBook"
	oldDir := filepath.Join(filesPath + "/" + fileName)
	destDir := filepath.Join(pathDir + "/" + fcY + "/" + fileName)

	//判断目录是否存在，无则创建
	if _, err := os.Stat(pathDir + "/" + fcY); os.IsNotExist(err) {
		err = os.MkdirAll(pathDir+"/"+fcY, os.ModePerm)
		if err != nil {
			return fmt.Errorf("无法创建目录: %v\n", err)
		}
	}

	//移动
	err := os.Rename(oldDir, destDir)
	if err != nil {
		return fmt.Errorf("无法移动文件 %s: %v\n", fileName, err)
	} else {
		fmt.Printf("文件 %s 移动成功\n", fileName)
	}
	return err
}

func fileRenameDate(filesPath string) error {
	// Walk through the directory
	err := filepath.Walk(filesPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("walk file tree error %s: %v\n", filesPath, err)
		}

		//只管理文件，忽略目录
		if !info.IsDir() {
			createdTime := info.ModTime()

			// fmt.Printf("File name: %s\n", info.Name())
			// fmt.Printf("File path：%s\n", path)
			// fmt.Printf("File created time：%s\n", createdTime.Format(customFormat))

			parts := strings.Split(info.Name(), ".")
			if len(parts) != 2 {
				panic("no file name extension")
			}
			err := os.Rename(path, filesPath+"/"+createdTime.Format(customFormat)+"."+parts[1])
			if err != nil {
				return fmt.Errorf("rename error %s: %v", info.Name(), err)
			}

			fileCreatedYear := createdTime.Format(createYear)
			err = moveFiles(fileCreatedYear, createdTime.Format(customFormat)+"."+parts[1])
			if err != nil {
				return err
			}
		}
		return err
	})
	return err
}

func main() {
	err := fileRenameDate(filesPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Done.!")
	}
}
