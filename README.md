## CSV Searcher

Данное приложение работает с файлами формата CSV: ищет заданные подстроки в полях файлов.

---
##### Как работает:
1. Читает все файлы формата CSV, которые находятся в директории `files`. В том числе может находить файлы, вложенные в поддиректории.
2. Находит все записи, содержащие любую из подстрок, заданных в массиве строк `queryWords` в файле `main.go`
3. Все найденные записи из всех файлов записывает в отдельный файл в `result/output.txt`

---

Для того, чтобы запустить приложение, необходимо:
1. Установить версию [Go](https://go.dev/dl/) не ниже **1.21**.
2. Запустить программу командой `go run main.go`
