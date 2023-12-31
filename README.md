# Тестовое задание avitoTech

<!-- ToC start -->
# Описание задачи
Требуется реализовать сервис, хранящий пользователя и сегменты, в которых он состоит (создание, изменение, удаление сегментов, а также добавление и удаление пользователей в сегмент).
## Опциональные задания
- Реализовать сохранение истории попадания/выбывания пользователя из сегмента с возможностью получения отчета по пользователю за определенный период. На вход: год-месяц. На выходе ссылка на CSV файл.
- Реализовать возможность задавать TTL (время автоматического удаления пользователя из сегмента)
- В методе создания сегмента, добавить опцию указания процента пользователей, которые будут попадать в сегмент автоматически. В методе получения сегментов пользователя, добавленный сегмент должен отдаваться у заданного процента пользователей.
Полное описание в [TASK](https://github.com/avito-tech/backend-trainee-assignment-2023)

# Реализация
- Следование дизайну REST API.
- Чистая архитектура
- Применение фреймворка [gin-gonic/gin](https://github.com/gin-gonic/gin).
- Применение СУБД Postgres посредствов библиотеки [sqlx](https://github.com/jmoiron/sqlx) и написанием SQL запросов.
- Конфигурация приложения - библиотека [viper](https://github.com/spf13/viper).
- Контейнеризация с помощью Docker и docker-compose

**Структура проекта:**
```
.
├── pkg
│   ├── handler     // слой обработки
│   ├── service     // слой бизнес-логики
│   └── repository  // слой взаимодействия с БД
├── cmd             // точка входа в приложение
├── schema          // SQL файлы миграции
├── configs         // файл конфигурации
```

# Endpoints

- POST /api/segment/ - создание сегмента
    - Тело запроса:
        - name - название сегмента
- POST /api/segment/:per - создание сегмента с указанием процента пользователей, которые будут попадать в сегмент автоматически
    - Тело запроса:
        - name - название сегмента
    - Параметры запроса:
        - :per - процент пользователей
- DELETE /api/segment/ - удаление сегмента
    - Тело запроса:
        - name - название сегмента
- PUT /api/user/ - управление причастности пользователя к сегментам: добавление к ним или удаление из них
    - Тело запроса:
        - user_id - идентификатор пользователя
        - add_to_segments - список сгементов, куда нужно добавить пользователя, элемент списка состоит из названия сегмента и времени автоматического удаления при его наличии  
        - delete_from_segments - список сгементов, откуда нужно удалить пользователя пользователя
- GET /api/user/:user_id - получения активных сегментов пользователя
    - Параметры запроса:
        - :user_id - идентификатор пользователя
- GET /api/user/ - получение отчёта
    - Тело запроса:
        - year - год
        - month - месяц
     
# Запуск
```
make build && make run
```
Если приложение запускается впервые, необходимо применить миграции к базе данных:
```
make migrate-up
```

# Примеры

Запросы сгенерированы командой curl

### 1. POST /api/segment/
**Запрос:**
```
$ curl --location --request POST 'localhost:8000/api/segment/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "myFristSegment"
}'
```
**Тело ответа:**
```
{
    "id": 4
}
```

### 2. POST /api/segment/:per
**Запрос:**
```
$ curl --location --request POST 'localhost:8000/api/segment/30' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "mySecondSegment"
}'
```
**Тело ответа:**
```
{
    "id": 5
}
```

### 3. DELETE /api/segment/
**Запрос:**
```
$ curl --location --request DELETE 'localhost:8000/api/segment/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "mySecondSegment"
}'
```
**Тело ответа:**
```
{
    "status":"ok"
}
```

### 4. PUT /api/user/
**Запрос:**
```
$ curl --location --request PUT 'localhost:8000/api/user/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "user_id": 1004,
    "add_to_segments": [
      {
        "name": "myFristSegment"
      }
    ],
    "delete_from_segments": [
    ]
}'
```
**Тело ответа:**
```
{
    "status":"ok"
}
```

### 5. GET /api/user/:user_id
**Запрос:**
```
$ curl --location --request GET 'localhost:8000/api/user/1004' \
--header 'Content-Type: application/json'
```
**Тело ответа:**
```
[
    {
        "name":"myFristSegment",
    },
    {
        "name":"mySecondSegment",
    }
]
```



### 6. GET /api/user/
**Запрос:**
```
$ curl --location --request GET 'localhost:8000/api/user/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "year": 2023,
    "month": 8
}'
```
**Тело ответа:**
```
{
    "reference": "report.csv"
}
```
Уже видно, что последний пример не демонстрирует корректный возврат того, что требовалось в таске. 
Это из-за того, что у меня не получилось сделать так, чтобы сервис загрузил файл в облачное хранили и вернул ссылку, 
поэтому сервис просто возвращает название файла, созданного в локалке. Вот пример того, как выглядит отчёт: https://disk.yandex.ru/d/1A5lYTtIkmVT5g
