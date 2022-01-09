package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ConnectFTP(t *testing.T) {
	mock := NewFtpService()
	t.Run("Sucess", func(t *testing.T) {
		_, error := mock.ConnectFTP("localhost", "21")
		assert.NoError(t, error)
	})
	t.Run("Error", func(t *testing.T) {
		_, error := mock.ConnectFTP("", "")
		assert.Error(t, error)
	})
}

func Test_LoginFTP(t *testing.T) {
	mock := NewFtpService()
	t.Run("Sucess", func(t *testing.T) {
		client, _ := mock.ConnectFTP("localhost", "21")
		error := mock.LoginFTP(*client, "anonymous", "anonymous")
		assert.NoError(t, error)
	})
	t.Run("Error", func(t *testing.T) {
		client, _ := mock.ConnectFTP("localhost", "21")
		error := mock.LoginFTP(*client, "", "")
		assert.Error(t, error)
	})
}

func Test_ReadFileFTP(t *testing.T) {
	mock := NewFtpService()
	t.Run("Sucess", func(t *testing.T) {
		client, _ := mock.ConnectFTP("localhost", "21")
		error := mock.LoginFTP(*client, "anonymous", "anonymous")
		assert.NoError(t, error)
		_, error = mock.ReadFileFTP(*client, "arquivos/", "sample.txt")
		assert.NoError(t, error)
	})
	t.Run("Error - ChangeDir", func(t *testing.T) {
		client, _ := mock.ConnectFTP("localhost", "21")
		error := mock.LoginFTP(*client, "anonymous", "anonymous")
		assert.NoError(t, error)
		_, error = mock.ReadFileFTP(*client, "naoexiste/", "")
		assert.Error(t, error)
	})
	t.Run("Error - Retr", func(t *testing.T) {
		client, _ := mock.ConnectFTP("localhost", "21")
		error := mock.LoginFTP(*client, "anonymous", "anonymous")
		assert.NoError(t, error)
		_, error = mock.ReadFileFTP(*client, "arquivos/", "teste.txt")
		assert.Error(t, error)
	})
}
