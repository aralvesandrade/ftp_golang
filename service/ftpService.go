package service

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jlaffaye/ftp"
)

type FtpService struct{}

type IFtpService interface {
	ConnectFTP(host, port string) (*ftp.ServerConn, error)
	LoginFTP(client ftp.ServerConn, user, pass string) error
	ReadFileFTP(client ftp.ServerConn, path, fileName string) ([]byte, error)
	SendFileFTP(client ftp.ServerConn, path, fileName string) error
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

func (f *FtpService) ReadFileFTP(client ftp.ServerConn, path, fileName string) ([]byte, error) {
	if error := client.ChangeDir(path); error != nil {
		return nil, error
	}

	response, error := client.Retr(fileName)

	if error != nil {
		return nil, error
	}

	buf, _ := ioutil.ReadAll(response)

	return buf, nil
}

func (f *FtpService) SendFileFTP(client ftp.ServerConn, path, fileName string) error {
	if error := client.ChangeDir(path); error != nil {
		return error
	}

	file, error := os.Open(fileName)

	if error != nil {
		return error
	}
	defer file.Close()

	if error := client.Stor(path, file); error != nil {
		return error
	}

	if error := os.Remove(fileName); error != nil {
		return error
	}

	return nil
}
