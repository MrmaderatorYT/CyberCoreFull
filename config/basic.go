package config

import (
	"encoding/base64"
	"github.com/Tnze/go-mc/chat"
	"io/ioutil"
)

var (
	ProtocolVersion uint16       = 340
	MOTD            chat.Message = chat.Text("Тестовое ядро §aДА")
)

// получить место расположение иконки 64х64 в папке config
func GetFaviconPath() string {
	return "config/icon.png"
}

// GetFaviconBase64 - Возвращает Base64-кодированную строку с иконкой сервера.
func GetFaviconBase64(faviconPath string) (string, error) {
	// Читаем файл с иконкой
	faviconData, err := ioutil.ReadFile(faviconPath)
	if err != nil {
		return "", err
	}
	// Кодируем в Base64
	faviconBase64 := base64.StdEncoding.EncodeToString(faviconData)
	return faviconBase64, nil
}
