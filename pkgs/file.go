package pkgs

import (
    "bufio"
    "fmt"
    "os"
)

func FileReader(file_path string) []string  {
	readFile,err  := os.Open(file_path)
	if err != nil {
		fmt.Println("Cannot find the file path , check your flags !!!")
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var FileLines []string
	for  fileScanner.Scan() {
		FileLines = append(FileLines, fileScanner.Text())
	}
	return FileLines

}