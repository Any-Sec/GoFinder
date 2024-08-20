package pkgs

import (
	

	"github.com/fatih/color"
)



func InfosMenu(url_addres string , file_path string){
	color.White("===============GoFinder======================\n\n")
	color.White("[#] Url      :   %s \n", url_addres)
	color.White("[#] Wordlist :   %s \n\n", file_path)
	color.White("=============================================\n\n")
	lines := FileReader(file_path)
	for i, value := range lines {
		NewUrl := url_addres+value
		data , _ := GetRequestToWebsite(NewUrl)
		if data.StatusCode != 404 {
			color.Green("[*][FOUND]  %s [%s] [%s] %.2f%% \n",NewUrl,data.Proto,data.Status,CounterLines(float32(i),float32(len(lines))))
		}

		
	}
}

func CounterLines(index float32, lines float32) float32 {

	percent := (index/lines) * 100
	return percent

}