# TaskManager
REST API сервис для работы с задачами, написанный на языке Go

## Установка и конфигурация
- Склонировать репозиторий:
  ```
  git clone https://github.com/mx4alex/TaskManager.git
  ```
- Настроить конфигурацию в файле `config.yaml`
- Запустить *docker compose*
  ```
  docker compose up --build
  ```

## Использование

### Сервис поддерживает следующие эндпоинты:
- `POST /tasks` создает задачу, которая передана в *body* 
- `GET /tasks` возвращает все задачи
- `PUT /tasks/{id}/mark` помечает задачу с заданным *id* выполненной
- `PUT /tasks/{id}` изменяет задачу с заданным *id* на задачу, которая передана в *body*
- `DELETE /tasks/{id}` удаляет задачу с заданным *id*

Документация находится в папке <a href="https://github.com/mx4alex/TaskManager/tree/main/docs">docs</a>

Визуальная документация Swagger UI доступна по адресу `http://localhost:8080/swagger/index.html#`

## Вариации интерфейса
- CLI
- REST API
- GRPC

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
- [x] dockerized
- [x] GRPC
- [x] mocks
- [ ] observability
