package pkgs

import (
	"fmt"
	"net/http"
	"strings"

	
)

//Url check and fix that
func CheckUrlAddress(url_addres string) (string,error)  {

	if strings.HasSuffix(url_addres,"//") {
		RemovedSlash := strings.TrimSuffix(url_addres,"/")
		return RemovedSlash,nil
		
	};if strings.HasSuffix(url_addres, " "){
		RemovedSpace := strings.Split(url_addres, " ")
		return RemovedSpace[0],nil

	};if !strings.HasSuffix(url_addres, "/"){
		AddedSlash := fmt.Sprintf("%s%s",url_addres,"/")
		return AddedSlash,nil
	}
	return url_addres , nil
}



func GetRequestToWebsite(url_addres string) (InformationWebsite,error) {
	
	con , err := http.Get(url_addres)
	if err != nil {
		return InformationWebsite{Status: "Error"} ,err
	}
	ReturnData := InformationWebsite{url: url_addres,StatusCode: con.StatusCode,Status: con.Status,Proto: con.Proto}
	return ReturnData , nil
}






type InformationWebsite struct {
	url string
	StatusCode int
	Status string
	Proto string
	
}
