package pkgs

import (
	

	"github.com/fatih/color"
)



func InfosMenu(url_addres string , file_path string){
	color.White("===============GoFinder======================")
	color.White("[#] Url      :%s \n", url_addres)
	color.White("[#] Wordlist :%s \n", file_path)
	color.White("=============================================")
	lines := FileReader(file_path)
	for i, value := range lines {
		NewUrl := url_addres+value
		data , _ := GetRequestToWebsite(NewUrl)
		if data.StatusCode != 404 {
			color.Green("[FOUND] url: %s [%s] [%s] \n",NewUrl,data.Proto,data.Status)
		}
		if i%100 == 0 {
			CounterLines(i,len(lines))	
		}
		
	}
}

func CounterLines(index int, lines int)  {
	color.Green("Lines : %d / %d ", index,lines)

}