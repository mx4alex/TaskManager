# ConsoleTaskManager

## Основная идея
Программа, которая позволяет удобно управлять задачами в консоли, классический CRUD.

## Реализованные функции
- Create - создать задачу
	- если задача уже существует, сообщить об этом
- Read - показать список задач
- Update - обновить задачу
- Mark - пометить задачу выполненной
- Delete - удалить задачу
	- если задача не удалена, то вернуть ошибку

## Вариации хранилища задач
- memory
- SQLite

## Статус
- [x] unit tests
- [x] clean architecture
- [x] storage variations (memory, sqlite)
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
- [ ] GRPC