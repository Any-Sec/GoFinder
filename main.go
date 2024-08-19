package main

import (
	"enum/pkgs"
	"os"
	
)


func main (){
	UrlAdd:= os.Args[1]
	FilePath:= os.Args[2]
	NewUrl ,_ :=  pkgs.CheckUrlAddress(UrlAdd)
	pkgs.InfosMenu(NewUrl,FilePath)

}