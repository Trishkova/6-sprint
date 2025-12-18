package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(response http.ResponseWriter, request *http.Request) {

	if request.Method != http.MethodGet {
		http.Error(response, "Не получается обработать запрос, ожидается GET-запрос", http.StatusMethodNotAllowed)
		return
	}

	file, err := os.Open("index.html")
	if err != nil {
		http.Error(response, "Не получилось открыть index.html", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	response.Header().Set("Content-Type", "text/html; charset=utf-8")

	_, err = io.Copy(response, file)
	if err != nil {
		http.Error(response, "Ошибка при отправке файла", http.StatusInternalServerError)
		return
	}
}

func UploadHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Не получается обработать запрос", http.StatusMethodNotAllowed)
		return
	}
	if err := request.ParseMultipartForm(10 << 20); err != nil {
		http.Error(response, fmt.Sprintf("Ошибка парсинга формы: %v", err), http.StatusBadRequest)
		return
	}
	file, header, err := request.FormFile("myFile")
	if err != nil {
		http.Error(response, fmt.Sprintf("Не получается открыть файл: %v", err), http.StatusBadRequest)
		return
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(response, fmt.Sprintf("Не получается прочитать содержимое файла: %v", err), http.StatusBadRequest)
		return
	}
	originalText := strings.TrimSpace(string(content))
	result, err := service.Convert(originalText)
	if err != nil {
		http.Error(response, fmt.Sprintf("Ошибка конвертации: %v", err), http.StatusInternalServerError)
		return
	}

	newFileName := fmt.Sprintf("%s_converted%s", time.Now().UTC().Format("20060102_150405"), filepath.Ext(header.Filename))
	if err := os.WriteFile(newFileName, []byte(result), 0644); err != nil {
		http.Error(response, fmt.Sprintf("Не получилось сохранить результат конвертации: %v", err), http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "text/plain; charset=utf-8")
	response.WriteHeader(http.StatusOK)

	summary := fmt.Sprintf("Конвертация завершена.\n\nИсходный текст: %s\n\nРезультат:\n%s", originalText, result)
	response.Write([]byte(summary))
}
