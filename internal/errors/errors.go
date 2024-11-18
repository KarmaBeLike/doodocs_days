package errors

import (
	"fmt"
	"net/http"
)

// ErrorResponse представляет структуру для ошибок архивации.
type ErrorResponse struct {
	Code    int    // HTTP-код ошибки
	Message string // Сообщение об ошибке
}

// Предопределенные ошибки.
var (
	ErrFileRead          = New(http.StatusBadRequest, "Мы старались, но прочесть файл не удалось 📂🥲")
	ErrNotArchive        = New(http.StatusBadRequest, "Мы ожидали .zip файл, но это что-то другое 📦🥲")
	ErrInvalidMime       = New(http.StatusUnsupportedMediaType, "Неверный тип файла 📄❌")
	ErrInternal          = New(http.StatusInternalServerError, "Неизвестная ошибка 😞🥲")
	ErrInvalidFile       = New(http.StatusBadRequest, "Файл не загружен или поврежден 📂😭")
	ErrFileOpenFailed    = New(http.StatusBadRequest, "Не удалось открыть файл 📂😞")
	ErrZipCreation       = New(http.StatusInternalServerError, "Ошибка создания архива 📦❌")
	ErrZipWriteFailed    = New(http.StatusInternalServerError, "Ошибка записи в архив 📦❌")
	ErrZipCloseFailed    = New(http.StatusInternalServerError, "Ошибка закрытия архива 📦😞")
	ErrFileSaveFailed    = New(http.StatusInternalServerError, "Ошибка сохранения файла 📂❌")
	ErrSMTPConfigMissing = New(http.StatusInternalServerError, "Отсутствует SMTP-конфигурация 📧⚙️")
	ErrEmailSendFailed   = New(http.StatusInternalServerError, "Ошибка отправки электронной почты 📤❌")
)

var (
	ErrNoFileProvided    = New(http.StatusBadRequest, "Файл не предоставлен 📂🤔")
	ErrMultipleFiles     = New(http.StatusBadRequest, "Можно прикрепить только один файл 📂❌")
	ErrNoEmailsProvided  = New(http.StatusBadRequest, "Не указаны адреса получателей 📬❌")
	ErrFileSendFailed    = New(http.StatusInternalServerError, "Не удалось отправить файл 📂❌")
	ErrFormParsingFailed = New(http.StatusBadRequest, "Ошибка при разборе формы 🤯")
	ErrMethodNotAllowed  = New(http.StatusMethodNotAllowed, "Метод не разрешен 🙅‍♂️")
)

// Error реализует интерфейс error для ErrorResponse.
func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("Код: %d, Сообщение: %s", e.Code, e.Message)
}

// New создает новую ошибку ErrorResponse с указанным кодом и сообщением.
func New(code int, message string) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: message,
	}
}
