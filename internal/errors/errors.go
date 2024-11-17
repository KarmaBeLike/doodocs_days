package errors

import "fmt"

// ArchiveError представляет структуру для ошибок архивации.
type ArchiveError struct {
	Code    int    // HTTP-код ошибки
	Message string // Сообщение об ошибке
}

// Предопределенные ошибки.
var (
	ErrFileRead        = &ArchiveError{Code: 400, Message: "Мы старались, но прочесть файл не удалось 📂🥲"}
	ErrNotArchive      = &ArchiveError{Code: 400, Message: "Мы ожидали .zip файл, но это что-то другое 📦🥲"}
	ErrInvalidMime     = &ArchiveError{Code: 415, Message: "Неверный тип файла 📄🥲"}
	ErrInternal        = &ArchiveError{Code: 500, Message: "Неизвестная ошибка 😞🥲"}
	ErrInvalidMimeType = &ArchiveError{Code: 415, Message: "Неверный формат файла 📂😢"}
	ErrInvalidFile     = &ArchiveError{Code: 400, Message: "Файл не загружен или поврежден 📂😭"}
)

// Error реализует интерфейс error для ArchiveError.
func (e *ArchiveError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

// NewArchiveError создает новую ошибку ArchiveError с указанными кодом и сообщением.
func NewArchiveError(code int, message string) error {
	return &ArchiveError{
		Code:    code,
		Message: message,
	}
}
