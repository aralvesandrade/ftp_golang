package business

import (
	"errors"
	"testing"

	"github.com/jlaffaye/ftp"
	"github.com/stretchr/testify/assert"
	mocking "github.com/stretchr/testify/mock"

	"labs.com/ftp/service/mocks"
)

func Test_LerArquivo(t *testing.T) {
	t.Run("Sucess", func(t *testing.T) {
		mockFtpService := new(mocks.IFtpService)
		mock := NewArquivoBusiness(mockFtpService)
		mockFtpService.On("ConnectFTP", mocking.AnythingOfType("string"), mocking.AnythingOfType("string")).Return(&ftp.ServerConn{}, nil)
		mockFtpService.On("LoginFTP", mocking.AnythingOfType("ftp.ServerConn"), mocking.AnythingOfType("string"), mocking.AnythingOfType("string")).Return(nil)
		mockFtpService.On("ReadFileFTP", mocking.AnythingOfType("ftp.ServerConn"), mocking.AnythingOfType("string"), mocking.AnythingOfType("string")).Return([]byte(""), nil)
		_, error := mock.LerArquivo()
		assert.NoError(t, error)
	})
	t.Run("Error - ConnectFTP", func(t *testing.T) {
		mockFtpService := new(mocks.IFtpService)
		mock := NewArquivoBusiness(mockFtpService)
		mockFtpService.On("ConnectFTP", mocking.AnythingOfType("string"), mocking.AnythingOfType("string")).Return(nil, errors.New("Ocorreu um erro"))
		_, error := mock.LerArquivo()
		assert.Error(t, error)
	})
	t.Run("Error - LoginFTP", func(t *testing.T) {
		mockFtpService := new(mocks.IFtpService)
		mock := NewArquivoBusiness(mockFtpService)
		mockFtpService.On("ConnectFTP", mocking.AnythingOfType("string"), mocking.AnythingOfType("string")).Return(&ftp.ServerConn{}, nil)
		mockFtpService.On("LoginFTP", mocking.AnythingOfType("ftp.ServerConn"), mocking.AnythingOfType("string"), mocking.AnythingOfType("string")).Return(errors.New("Ocorreu um erro"))
		_, error := mock.LerArquivo()
		assert.Error(t, error)
	})
	t.Run("Error - ReadFileFTP", func(t *testing.T) {
		mockFtpService := new(mocks.IFtpService)
		mock := NewArquivoBusiness(mockFtpService)
		mockFtpService.On("ConnectFTP", mocking.AnythingOfType("string"), mocking.AnythingOfType("string")).Return(&ftp.ServerConn{}, nil)
		mockFtpService.On("LoginFTP", mocking.AnythingOfType("ftp.ServerConn"), mocking.AnythingOfType("string"), mocking.AnythingOfType("string")).Return(nil)
		mockFtpService.On("ReadFileFTP", mocking.AnythingOfType("ftp.ServerConn"), mocking.AnythingOfType("string"), mocking.AnythingOfType("string")).Return(nil, errors.New("Ocorreu um erro"))
		_, error := mock.LerArquivo()
		assert.Error(t, error)
	})
}
