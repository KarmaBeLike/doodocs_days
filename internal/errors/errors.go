package errors

import "fmt"

type ArchiveError struct {
	Code    int
	Message string
}

var (
	ErrFileRead    = &ArchiveError{Code: 400, Message: "Мы старались, но файл прочесть не удалось 📂🥲"}
	ErrNotArchive  = &ArchiveError{Code: 400, Message: "Мы ожидали .zip файл, но это что-то другое 📦🥲"}
	ErrInvalidMime = &ArchiveError{Code: 415, Message: "Неверный тип файла 📄🥲"}
	ErrInternal    = &ArchiveError{Code: 500, Message: "Неизвестная ошибка 😞🥲"}
)

func (e *ArchiveError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func NewArchiveError(code int, message string) error {
	return &ArchiveError{
		Code:    code,
		Message: message,
	}
}
