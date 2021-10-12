# Golang Mastery Test for Godel

Все програамы компилируются и запускаются через стандартный `go run task#.go`

- [X] Задание 1
- [X] Задание 2
- [ ] Задание 3 
- [X] Задание 4
- [X] Задание 5

____

Для пятого задания потребуется установить дополнительные библиотеки и поднять базу данных.

Лично я использовал драйвер mysql. Устанавливается он из [репозитория](https://github.com/go-sql-driver/mysql) командой:
```
go get -u github.com/go-sql-driver/mysql
```

Так же для запуска нужен марштуризатор [gorilla/mux](https://github.com/gorilla/mux):
```
go get github.com/gorilla/mux
```

____

Адрес и порт к базе данных указывается в `144` строке:
```
db, err:= sql.Open("mysql", "<логин>:<пароль>@tcp(<адрес:порт>)/<название бд>")
```
Пример: `db, err:= sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")`

В бд нужно создать таблицу messages. Код sql зароса для создания полей лежит в файле /task5/db.sql

____

Запустив программу, мы сможем обращаться к серверу через 88 порт через GET и POST запросы. Запросы приведены ниже:
| Запрос | Адрес | Описание |
|----------------|---------|----------------:|
| GET | / | Получить все статьи |
| GET | /id/<id статьи> | Получить конкретную статью по её уникальному ID |
