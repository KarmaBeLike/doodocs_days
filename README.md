# Doodocs Days

## Описание проекта

Doodocs Days — это веб-приложение для работы с архивами, позволяющее пользователям загружать архивы, извлекать информацию о содержимом архивов и создавать новые архивы. Приложение предоставляет API для взаимодействия с архивами, включая создание архивов, извлечение информации о содержимом, и отправку файлов на электронную почту.


### 1. Клонируйте репозиторий

git clone https://github.com/KarmaBeLike/doodocs_days.git
cd doodocs_days

### START PROJECT
- **установка зависимостей:**
```
go mod tidy
```
- **настройка SMTP:**
```
создайте . env с конфигурациями как показано в config.env.example
```
- **запуск:**
```
go run ./cmd
```
- **Получить информацию об архиве:**
   ```http
    POST /archive/info
    ```
    -Формат запроса: multipart/form-data:
        - Поле: file — архивный файл (.ZIP)
    - sample output:
    ```json
    {
  "Filename": "test.zip",
  "ArchiveSize": 9.0,
  "TotalSize": 9.0,
  "TotalFiles": 1,
  "Files": [
    {
      "FilePath": "test.zip",
      "Size": 9.0,
      "MimeType": "application/octet-stream"
    }
  ]
    }
- **Создать архив:**
    ```http
    POST /archive/files
    ```
    -Формат запроса: multipart/form-data:
        - Поле: files[] — массив файлов для добавления в архив
   
   - **Рассылка файла по нескольким email:**
    ```http
    POST /mail/file
    ```
    -Формат запроса: multipart/form-data:
        - Поля: 
        file — файл для отправки
        emails[] — список email-адресов для отправки файла
  
