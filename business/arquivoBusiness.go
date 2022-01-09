package business

import (
	"os"

	"labs.com/ftp/service"
)

type ArquivoBusiness struct {
	FtpService service.IFtpService
}

type IArquivoBusiness interface {
	LerArquivo() ([]byte, error)
	CriarArquivo() error
}

func NewArquivoBusiness(ftpService service.IFtpService) IArquivoBusiness {
	return &ArquivoBusiness{ftpService}
}

func (b *ArquivoBusiness) LerArquivo() ([]byte, error) {
	client, error := b.FtpService.ConnectFTP("localhost", "21")

	if error != nil {
		return nil, error
	}

	if error := b.FtpService.LoginFTP(*client, "alexandre", "21012001"); error != nil {
		return nil, error
	}

	buf, error := b.FtpService.ReadFileFTP(*client, "arquivos/", "FechamentoCoopluiza_202201.txt")

	if error != nil {
		return nil, error
	}

	return buf, nil
}

func (b *ArquivoBusiness) CriarArquivo() error {
	file, error := os.Create("./sample.txt")

	if error != nil {
		return error
	}
	defer file.Close()

	_, error = file.WriteString("teste")

	if error != nil {
		return error
	}

	return nil
}
