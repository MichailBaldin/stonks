# Простые задачи для тренажера Go HTTP (20-30 минут каждая)

## Базовые HTTP серверы

### 1. Простой CRUD API для пользователей

Создай HTTP сервер на порту 8080 с эндпоинтами:

- `GET /users` - возвращает список пользователей в JSON
- `POST /users` - создает пользователя из JSON body
- `GET /users/{id}` - возвращает пользователя по ID
- `DELETE /users/{id}` - удаляет пользователя

Структура: `{id: int, name: string, email: string}`. Используй `map[int]User` в памяти. ID генерируй как счетчик. Один тест для GET /users. Dockerfile: `FROM golang:alpine` с простым запуском.

### 2. API со счетчиком запросов

Создай сервер с эндпоинтами:

- `GET /counter` - возвращает текущий счетчик запросов
- `POST /counter/reset` - обнуляет счетчик
- `GET /health` - возвращает `{"status": "ok"}`

Каждый запрос увеличивает глобальный счетчик на 1. Используй `atomic.Int64`. Один тест на увеличение счетчика. Базовый Dockerfile.

### 3. JSON Echo сервер

Создай сервер:

- `POST /echo` - возвращает полученный JSON с добавленным полем `timestamp`
- `GET /time` - возвращает `{"time": "2024-01-01T12:00:00Z"}`

Принимай любой валидный JSON, добавляй поле timestamp с текущим временем UTC. Один тест для каждого эндпоинта. Простой Dockerfile.

### 4. Файловый сервер

Создай сервер:

- `POST /upload` - загружает файл, сохраняет в папку `files/`
- `GET /files` - список файлов с размерами
- `GET /files/{name}` - отдает файл

Принимай только текстовые файлы до 1MB. Имя файла = оригинальное имя. Один тест на загрузку. Dockerfile с COPY папки files.

### 5. Калькулятор API

Создай сервер:

- `POST /calc/add` - принимает `{a: float64, b: float64}`, возвращает сумму
- `POST /calc/multiply` - произведение двух чисел
- `GET /calc/history` - история последних 10 операций

Храни историю в слайсе структур `{operation: string, a: float64, b: float64, result: float64}`. Один тест на сложение. Базовый Dockerfile.

### 6. Простой TODO API

Создай сервер для задач:

- `GET /todos` - список всех задач
- `POST /todos` - создание задачи `{title: string}`
- `PATCH /todos/{id}` - переключение статуса completed

Структура: `{id: int, title: string, completed: bool}`. Используй map в памяти. Один тест на создание задачи. Простой Dockerfile.

### 7. Сервер статистики

Создай сервер:

- `POST /stats/visit` - записывает посещение страницы `{page: string}`
- `GET /stats/pages` - статистика посещений по страницам
- `GET /stats/total` - общее количество посещений

Считай посещения в `map[string]int`. Возвращай `{page: string, visits: int}`. Один тест на запись посещения. Dockerfile стандартный.

### 8. Простой proxy

Создай HTTP proxy:

- `GET /proxy?url=...` - проксирует GET запрос к указанному URL
- `GET /health` - health check

Используй `http.Get()` для проксирования. Возвращай статус и body как есть. Timeout 5 секунд. Один тест с моком. Базовый Dockerfile.

### 9. Key-Value хранилище

Создай in-memory KV сервер:

- `POST /kv/{key}` - сохранить значение (body = value)
- `GET /kv/{key}` - получить значение
- `DELETE /kv/{key}` - удалить ключ
- `GET /kv` - все ключи

Используй `map[string]string`. Body как строку. Один тест на set/get. Простой Dockerfile.

### 10. Random API

Создай сервер случайных данных:

- `GET /random/number` - случайное число 1-100
- `GET /random/string` - случайная строка 10 символов
- `GET /random/uuid` - простой UUID (используй `crypto/rand`)

Используй пакеты `math/rand` и `crypto/rand`. Один тест на формат ответа. Базовый Dockerfile.

### 11. URL shortener

Создай простой сокращатель URL:

- `POST /shorten` - принимает `{url: string}`, возвращает `{short: string}`
- `GET /{short}` - редирект на оригинальный URL
- `GET /stats/{short}` - счетчик переходов

Генерируй короткий код из 6 символов (a-z, 0-9). Храни в map. Один тест на создание. Простой Dockerfile.

### 12. Простой чат API

Создай API для сообщений:

- `POST /messages` - отправить сообщение `{user: string, text: string}`
- `GET /messages` - последние 20 сообщений
- `GET /users` - список уникальных пользователей

Структура: `{id: int, user: string, text: string, timestamp: string}`. Храни в слайсе. Один тест на отправку. Dockerfile стандартный.

### 13. Сервер конвертации

Создай конвертер единиц:

- `GET /convert/temp?c=25` - Цельсий в Фаренгейт
- `GET /convert/length?m=100` - метры в футы
- `GET /convert/weight?kg=70` - килограммы в фунты

Формулы: F = C*9/5+32, ft = m*3.28084, lbs = kg\*2.20462. Один тест на температуру. Базовый Dockerfile.

### 14. Middleware logger

Создай сервер с логированием:

- `GET /users` - любой endpoint с данными
- `POST /users` - создание данных
- Middleware логирует: timestamp, method, path, status code

Используй стандартный `log`. Middleware как функция-обертка. Один тест проверяет вызов handler'а. Простой Dockerfile.

### 15. Rate counter API

Создай счетчик запросов по IP:

- `GET /rate/my` - количество запросов с вашего IP
- `POST /rate/reset` - сброс счетчика для IP
- `GET /rate/all` - счетчики всех IP

Используй `map[string]int` где ключ = IP. IP из `r.RemoteAddr`. Один тест со специальным IP. Стандартный Dockerfile.

### 16. JSON валидатор

Создай валидатор JSON:

- `POST /validate` - проверяет валидность JSON в body
- `POST /format` - форматирует JSON (indent)
- `POST /minify` - минифицирует JSON

Используй `json.Valid()`, `json.Indent()`, `json.Compact()`. Возвращай `{valid: bool, error?: string}`. Один тест на валидный JSON. Простой Dockerfile.

### 17. Simple metrics

Создай сервер метрик:

- `POST /metrics/inc/{name}` - увеличить метрику на 1
- `GET /metrics/{name}` - значение метрики
- `GET /metrics` - все метрики

Используй `map[string]int64` с `sync.RWMutex`. Один тест на инкремент. Базовый Dockerfile.

### 18. HTTP client wrapper

Создай обертку над HTTP клиентом:

- `GET /fetch?url=...` - GET запрос к URL, возврат body
- `POST /post` - POST запрос, принимает `{url: string, data: object}`
- `GET /headers?url=...` - только заголовки ответа

Timeout 10 секунд. Возвращай status, headers, body. Один тест с моком. Простой Dockerfile.

### 19. Template API

Создай сервер шаблонизации:

- `POST /template` - принимает `{template: string, data: object}`
- Заменяет `{{.field}}` на значения из data
- `GET /templates` - список сохраненных шаблонов

Простая замена строк `strings.Replace()`. Храни шаблоны в map. Один тест на замену. Стандартный Dockerfile.

### 20. Basic auth server

Создай сервер с простой авторизацией:

- `GET /public` - доступно всем
- `GET /private` - требует заголовок `Authorization: Bearer secret123`
- `POST /login` - принимает `{password: string}`, если "admin" то возвращает токен

Middleware проверяет токен в заголовке. Один тест с правильным токеном. Базовый Dockerfile.

### 21. Data aggregator

Создай агрегатор данных:

- `POST /data` - принимает `{value: float64, category: string}`
- `GET /aggregate/{category}` - сумма, среднее, количество для категории
- `GET /categories` - список всех категорий

Структура для агрегации: `{sum, count, average}`. Считай на лету. Один тест на агрегацию. Простой Dockerfile.

### 22. Simple scheduler

Создай планировщик задач:

- `POST /schedule` - принимает `{task: string, delay: int}` (delay в секундах)
- `GET /tasks` - список задач с статусом
- Через delay секунд меняет статус на "completed"

Используй `time.AfterFunc()`. Статусы: pending, completed. Один тест на создание задачи. Стандартный Dockerfile.

### 23. Text processor

Создай обработчик текста:

- `POST /text/upper` - переводит в верхний регистр
- `POST /text/lower` - в нижний регистр
- `POST /text/count` - считает слова и символы

Body = обычный текст. Возвращай обработанный текст или статистику. Один тест на каждую операцию. Базовый Dockerfile.

### 24. Memory cache

Создай простой кеш:

- `POST /cache/{key}` - сохранить значение (body = value)
- `GET /cache/{key}` - получить значение
- `DELETE /cache/{key}` - удалить
- `GET /cache/stats` - количество ключей

Используй `map[string]string` с `sync.RWMutex`. Статистика: `{keys: int}`. Один тест на set/get. Простой Dockerfile.

### 25. Health checker

Создай проверку состояния:

- `GET /health` - статус сервера `{status: "ok", uptime: "1h2m3s"}`
- `GET /health/deep` - проверяет все компоненты
- `POST /health/toggle` - включает/выключает "maintenance mode"

Компоненты: memory usage, disk space (моки). В maintenance mode все endpoints возвращают 503. Один тест на health check. Стандартный Dockerfile.

## Общие упрощенные требования:

### Минимальная структура:

```
/
├── main.go
├── Dockerfile
├── go.mod
└── README.md
```

### Упрощенные тесты:

- Один unit test на основной функционал
- Только happy path
- Без mocking сложных зависимостей

### Простой Dockerfile:

```dockerfile
FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]
```

### Базовые требования:

- HTTP сервер на порту 8080
- JSON ответы где нужно
- Минимальная обработка ошибок (400, 404, 500)
- Один файл main.go до 200 строк
- Запуск: `go run main.go`
- Docker: `docker build -t task . && docker run -p 8080:8080 task`
