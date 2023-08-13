# TaskManager
REST API сервис для работы с задачами, написанный на языке Go

## Реализованные функции
- Create - создать задачу
	- если задача уже существует, сообщить об этом
- Read - показать список задач
- Update - обновить задачу
- Mark - пометить задачу выполненной
- Delete - удалить задачу
	- если задача не удалена, то вернуть ошибку

## Вариации интерфейса
- CLI
- REST API

## Вариации хранилища задач
- memory
- SQLite
- PostgreSQL

## Статус
- [x] unit tests
- [x] clean architecture
- [x] storage variations (memory, sqlite, postgres)
- [x] config
- [x] migrations
- [x] REST
- [x] export/import json
- [x] concurrency
- [x] graceful shutdown
- [ ] integration tests
- [ ] mocks
- [ ] observity
- [ ] image app
