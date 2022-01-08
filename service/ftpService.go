package service

import (
	"fmt"
	"io/ioutil"

	"github.com/jlaffaye/ftp"
)

type FtpService struct{}

type IFtpService interface {
	ConnectFTP(host, port string) (*ftp.ServerConn, error)
	LoginFTP(client ftp.ServerConn, user, pass string) error
	ReadFileFTP(client ftp.ServerConn, path, file string) ([]byte, error)
}

func NewFtpService() IFtpService {
	return &FtpService{}
}

func (f *FtpService) ConnectFTP(host, port string) (*ftp.ServerConn, error) {
	addr := fmt.Sprintf("%s:%s", host, port)

	client, error := ftp.Connect(addr)

	if error != nil {
		return nil, error
	}

	return client, nil
}

func (f *FtpService) LoginFTP(client ftp.ServerConn, user, pass string) error {
	if error := client.Login(user, pass); error != nil {
		return error
	}

	return nil
}

func (f *FtpService) ReadFileFTP(client ftp.ServerConn, path, file string) ([]byte, error) {
	if error := client.ChangeDir(path); error != nil {
		return nil, error
	}

	response, error := client.Retr(file)

	if error != nil {
		return nil, error
	}

	buf, error := ioutil.ReadAll(response)

	if error != nil {
		return nil, error
	}

	return buf, nil
}
