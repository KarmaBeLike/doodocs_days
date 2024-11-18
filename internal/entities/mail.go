package entities

// Mail содержит данные для отправки email
type Mail struct {
	Sender     string   // Отправитель
	To         []string // Список получателей
	Subject    string   // Тема письма
	Body       string   // Тело письма
	Attachment string   // Имя вложения
	File       []byte   // Данные файла
}
