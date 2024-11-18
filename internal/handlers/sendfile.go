package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/KarmaBeLike/doodocs_days/internal/errors"
)

func (h *ArchiveHandler) SendFileHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to send file")
	if r.Method != http.MethodPost {
		log.Println("Method not allowed")
		http.Error(w, errors.ErrMethodNotAllowed.Message, errors.ErrMethodNotAllowed.Code)
		return
	}

	// Парсим данные формы
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, errors.ErrFormParsingFailed.Message, errors.ErrFormParsingFailed.Code)
		return
	}

	// Получаем файлы
	files := r.MultipartForm.File["file"]
	if len(files) != 1 {
		http.Error(w, errors.ErrMultipleFiles.Message, errors.ErrMultipleFiles.Code)
		return
	}

	// Получаем первый файл
	file, header, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error reading file from form: %s\n", err)
		http.Error(w, errors.ErrNoFileProvided.Message, errors.ErrNoFileProvided.Code)
		return
	}
	defer file.Close()

	// Проверка email
	emailsStr := r.FormValue("emails")
	emails := strings.Split(emailsStr, ",")
	if len(emails) == 0 || emails[0] == "" {
		log.Println("No recipient emails provided")
		http.Error(w, errors.ErrNoEmailsProvided.Message, errors.ErrNoEmailsProvided.Code)
		return
	}

	// Отправка файла
	log.Println("Sending file to provided email addresses:", emails)
	err = h.archiveService.SendFile(file, header, emails)
	if err != nil {
		log.Printf("Error sending file: %s\n", err)
		http.Error(w, errors.ErrFileSendFailed.Message, errors.ErrFileSendFailed.Code)
		return
	}
	log.Println("File sent successfully")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Файл успешно отправлен! 🎉📧")
}
