package main

import (
	"enum/pkgs"
	"fmt"
	"os"
)

func main(){
	URLadd:= os.Args[1]
	Lines := pkgs.FileReader("common.txt")
	fmt.Println("[==================================|GoFinder V0.1|============================================]")
	for _,value := range Lines{
		new_url := URLadd+value
		strct := pkgs.GetRequest(new_url)
		if strct.StatusCode != 404 {
			fmt.Printf("[FOUND] Url: %s  [%s] [%s] \n",new_url,strct.Proto,strct.Status)
		}else {
			continue
		} 
	}
	fmt.Println("[==============================================================================]")
