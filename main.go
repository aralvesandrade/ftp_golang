package main

import (
	"fmt"

	"labs.com/ftp/business"
	"labs.com/ftp/service"
)

func main() {
	ftpService := service.NewFtpService()
	arquivoBusiness := business.NewArquivoBusiness(ftpService)
	buf, error := arquivoBusiness.LerArquivo()
	if error != nil {
		fmt.Println(error.Error())
	}
	fmt.Println(string(buf))
}
