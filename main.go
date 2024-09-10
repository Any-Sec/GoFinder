package main

import (
	"enum/pkgs"
	"flag"

	
)


func main (){
	urlAddress := flag.String("url","","The target url address.")
	FilePath := flag.String("file","","wordlist for enumeration")
	flag.Parse()
	if *urlAddress == "" || *FilePath == "" {
		flag.Usage()
	}else {
		
		NewUrl ,_ :=  pkgs.CheckUrlAddress(*urlAddress)
		pkgs.InfosMenu(NewUrl,*FilePath)
	}


}