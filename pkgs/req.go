package pkgs

import (
	"fmt"
	"net/http"
	"os"
)

func GetRequest(url_addres string) InformationWebsite {
	res, err := http.Get(url_addres)
	if err != nil {
		fmt.Printf("error making http request: %s", err)
		os.Exit(1)
	}
	if res.StatusCode != 404{
		return InformationWebsite{url: url_addres, StatusCode: res.StatusCode,Status: res.Status , Proto: res.Proto }	
	}else{
		return InformationWebsite{url: "", StatusCode: res.StatusCode, Status: "" , Proto: ""}
	}
	

}

type InformationWebsite struct {
	url string
	StatusCode int
	Status string
	Proto string
	
}
