# HTTP - что это и зачем нужно?

## Основная задача HTTP
HTTP решает простую задачу: **как программы на разных компьютерах могут обмениваться данными через интернет**.

Представь: у тебя есть программа на компьютере А, которая хочет получить данные от программы на компьютере Б. HTTP - это набор правил, как они должны "разговаривать".

## Основные концепции

**Клиент-сервер модель:**
- **Клиент** - тот кто просит (браузер, мобильное приложение, твоя Go программа)
- **Сервер** - тот кто отвечает (веб-сервер, API)
- Клиент всегда инициирует общение

**Запрос-ответ:**
- Клиент отправляет **запрос** (request) 
- Сервер отправляет **ответ** (response)
- Каждый обмен независим (stateless)

## Простейший пример в жизни

Ты заходишь в кафе (клиент), говоришь официанту "Дайте кофе" (запрос), получаешь кофе (ответ).

HTTP работает так же:
```bash
Клиент: "GET /coffee HTTP/1.1"
Сервер: "200 OK, вот ваш кофе"
```

---

# Структура HTTP сообщений

## Как выглядит HTTP запрос

HTTP запрос состоит из трех частей:

```
GET /users/123 HTTP/1.1        ← Стартовая строка
Host: api.example.com          ← Заголовки  
User-Agent: MyApp/1.0          ← Заголовки
                               ← Пустая строка (обязательно!)
                               ← Тело запроса (если есть)
```

**Стартовая строка содержит:**
- **Метод** (GET, POST, PUT, DELETE) - что хотим сделать
- **Путь** (/users/123) - с чем хотим работать  
- **Версия** (HTTP/1.1) - какую версию протокола используем

**Заголовки** - дополнительная информация:
- `Host` - на какой сервер обращаемся
- `Content-Type` - какие данные отправляем
- `Authorization` - токен авторизации

## Как выглядит HTTP ответ

```
HTTP/1.1 200 OK                ← Статус строка
Content-Type: application/json ← Заголовки
Content-Length: 45             ← Заголовки
                               ← Пустая строка
{"id": 123, "name": "Иван"}    ← Тело ответа
```

**Статус строка:**
- **Версия** (HTTP/1.1)
- **Код статуса** (200, 404, 500) - что получилось
- **Фраза** (OK, Not Found) - человеко-читаемое описание

## Простой пример в Go

```go
// Отправляем GET запрос
resp, err := http.Get("https://api.github.com/users/octocat")

// Получаем:
// GET /users/octocat HTTP/1.1
// Host: api.github.com
// User-Agent: Go-http-client/1.1
```

---

# HTTP методы - зачем нужны разные способы обращения?

## Проблема, которую решают методы

Представь, что у нас есть только один способ общения - просто "дай данные". Но в реальности нам нужно:
- Получить информацию 
- Создать что-то новое
- Изменить существующее
- Удалить ненужное

HTTP методы решают эту проблему - они говорят серверу **что именно** мы хотим сделать.

## Основные методы

**GET** - "дай мне данные"
```
GET /users/123  → получить пользователя с ID 123
GET /posts      → получить список постов
```
- Безопасный (не изменяет данные на сервере)
- Можно кэшировать
- Параметры в URL

**POST** - "создай что-то новое"
```
POST /users     → создать нового пользователя
POST /posts     → написать новый пост
```
- Данные в теле запроса
- Изменяет состояние сервера
- Не безопасный для повтора

**PUT** - "обнови/замени полностью"
```
PUT /users/123  → заменить пользователя 123 новыми данными
```

**DELETE** - "удали"
```
DELETE /users/123  → удалить пользователя 123
```

## Простой пример в Go

```go
// Получить данные
resp, err := http.Get("https://api.example.com/users")

// Создать данные  
http.Post("https://api.example.com/users", "application/json", body)
```

## Почему важно использовать правильные методы?

- **Семантика** - сразу понятно что делает запрос
- **Кэширование** - браузеры кэшируют только GET
- **Безопасность** - GET не должен изменять данные
- **REST API** - стандартный подход к построению API

---

Ты прав, извини. Вот важные технические детали HTTP:

# Детали передачи данных в HTTP

## GET vs POST - где передаются данные

**GET - данные в URL:**
```
GET /search?username=admin&password=123456 HTTP/1.1
Host: example.com
```

**POST - данные в теле:**
```
POST /login HTTP/1.1
Host: example.com
Content-Type: application/x-www-form-urlencoded

username=admin&password=123456
```

## Почему логины/пароли НЕЛЬЗЯ через GET

**1. Логи сервера**
```
[access.log] GET /login?user=admin&pass=secret123 - вся авторизация в логах!
```

**2. История браузера**
```
Браузер сохраняет: example.com/login?password=secret123
```

**3. Referrer заголовки**
```
User-Agent переходит с example.com/login?pass=123 на google.com
Google получает: Referer: example.com/login?pass=123
```

**4. Кэширование**
```
Браузер/прокси кэшируют GET запросы → пароль остается в кэше
```

## Ограничения длины URL

**Браузеры:**
- Chrome: ~8000 символов
- IE: 2048 символов  
- Firefox: ~65000 символов

**Серверы:**
- Apache: 8000 по умолчанию
- Nginx: 4096-8192
- IIS: 16384

**Практический пример:**
```go
// ЭТО ПЛОХО - может не работать в IE
longURL := "http://example.com/search?data=" + strings.Repeat("x", 3000)
http.Get(longURL) // Может вернуть 414 Request-URI Too Large
```

## Кэширование GET vs POST

**GET кэшируется:**
```go
// Этот запрос может вернуть старые данные из кэша браузера
resp, _ := http.Get("http://api.com/users")
```

**POST НЕ кэшируется:**
```go
// Этот запрос всегда идет на сервер
http.Post("http://api.com/users", "application/json", body)
```

## Content-Length в POST

**Обязательно для POST с телом:**
```
POST /api HTTP/1.1
Content-Length: 25    ← Без этого сервер не поймет где заканчивается тело

{"name": "test user"}
```

**В Go это делается автоматически:**
```go
body := strings.NewReader(`{"name": "test"}`)
http.Post(url, "application/json", body) // Content-Length добавится сам
```

---

# Заголовки HTTP - технические детали

## Структура и ограничения заголовков

**Формат заголовка:**
```
Имя-Заголовка: значение\r\n
```

**Технические ограничения:**
- Имя заголовка: только ASCII, без пробелов, case-insensitive
- Значение: может содержать UTF-8, но лучше кодировать
- Общий размер всех заголовков: обычно 8KB лимит

**Пример проблемы:**
```go
// ЭТО СЛОМАЕТСЯ
req.Header.Set("My Header", "value") // Пробел в имени = ошибка

// ПРАВИЛЬНО  
req.Header.Set("My-Header", "value")
```

## Важные заголовки и их подводные камни

**Host - обязательный в HTTP/1.1:**
```
GET /api HTTP/1.1
Host: example.com    ← БЕЗ ЭТОГО сервер вернет 400 Bad Request
```

**User-Agent - идентификация клиента:**
```
User-Agent: Mozilla/5.0 (compatible; MyBot/1.0)
```
Многие API блокируют запросы без User-Agent или с подозрительным.

**Content-Type - критично для POST/PUT:**
```
// JSON данные
Content-Type: application/json; charset=utf-8

// Форма
Content-Type: application/x-www-form-urlencoded

// Файлы
Content-Type: multipart/form-data; boundary=something123
```

**Accept - что клиент понимает:**
```
Accept: application/json, text/plain, */*
```

## Кодировки и Content-Encoding

**Проблема:** HTTP изначально для текста, бинарные данные нужно кодировать.

**Content-Encoding - сжатие:**
```
Content-Encoding: gzip    ← Данные сжаты
Content-Encoding: deflate
Content-Encoding: br      ← Brotli (новый, лучше gzip)
```

**Transfer-Encoding - как передаются данные:**
```
Transfer-Encoding: chunked    ← Данные идут частями
```

**В Go автоматически:**
```go
// Go автоматически добавляет Accept-Encoding: gzip
resp, _ := http.Get("http://example.com")
// И автоматически распаковывает gzip ответ
```

## Chunked Transfer Encoding - передача по частям

**Проблема:** Сервер не знает размер ответа заранее (например, генерирует данные на лету).

**Решение:** Отправлять частями (chunks).

**Формат chunked:**
```
HTTP/1.1 200 OK
Transfer-Encoding: chunked

5\r\n          ← Размер первого chunk (5 байт в hex)
Hello\r\n      ← Данные первого chunk
6\r\n          ← Размер второго chunk (6 байт в hex)  
 World\r\n     ← Данные второго chunk
0\r\n          ← Размер 0 = конец
\r\n           ← Финальный CRLF
```

**Когда используется:**
- Стриминг больших файлов
- API с генерацией данных в реальном времени
- SSE (Server-Sent Events)

**В Go:**
```go
func streamHandler(w http.ResponseWriter, r *http.Request) {
    // Go автоматически использует chunked если не указан Content-Length
    
    for i := 0; i < 5; i++ {
        fmt.Fprintf(w, "Chunk %d\n", i)
        w.(http.Flusher).Flush() // Принудительно отправляем chunk
        time.Sleep(1 * time.Second)
    }
}
```

## Connection управление

**Connection: close**
```
Connection: close    ← Закрыть соединение после ответа
```

**Connection: keep-alive** (по умолчанию в HTTP/1.1)
```
Connection: keep-alive    ← Оставить соединение открытым
Keep-Alive: timeout=60, max=100
```

**Проблема с keep-alive:**
```go
// БЕЗ defer resp.Body.Close() соединения не переиспользуются!
resp, _ := http.Get("http://example.com")
io.ReadAll(resp.Body)
// resp.Body.Close() ← ОБЯЗАТЕЛЬНО для переиспользования соединения
```

## Подводные камни с заголовками в Go

**1. Case-insensitive, но Go канонизирует:**
```go
req.Header.Set("content-type", "application/json")
fmt.Println(req.Header.Get("Content-Type")) // "application/json" - работает!
```

**2. Множественные значения:**
```go
req.Header.Add("Accept", "application/json")
req.Header.Add("Accept", "text/plain")
// Accept: application/json, text/plain
```

**3. Опасные заголовки (Go не дает их менять):**
```go
// Эти заголовки Go устанавливает сам, изменить нельзя:
// Content-Length, Connection, Transfer-Encoding
req.Header.Set("Content-Length", "100") // Будет проигнорировано!
```

---

# HTTP статус-коды - что сервер говорит клиенту

## Зачем нужны статус-коды

**Проблема:** Клиент отправил запрос, получил ответ. Но как понять - все ли прошло хорошо?

**Решение:** Сервер в первой строке ответа сообщает результат числовым кодом:

```
HTTP/1.1 200 OK        ← Все хорошо
HTTP/1.1 404 Not Found ← Ресурс не найден  
HTTP/1.1 500 Internal Server Error ← Ошибка сервера
```

## Категории статус-кодов

**1xx - Информационные** (редко используются)
- 100 Continue - можно продолжать отправку данных

**2xx - Успех**
- 200 OK - запрос успешно обработан
- 201 Created - ресурс создан
- 204 No Content - успех, но нет данных для возврата

**3xx - Перенаправление**
- 301 Moved Permanently - ресурс навсегда переехал
- 302 Found - временное перенаправление
- 304 Not Modified - данные не изменились (для кэша)

**4xx - Ошибки клиента**
- 400 Bad Request - плохой запрос
- 401 Unauthorized - нужна авторизация
- 403 Forbidden - доступ запрещен
- 404 Not Found - ресурс не найден

**5xx - Ошибки сервера**
- 500 Internal Server Error - ошибка на сервере
- 502 Bad Gateway - прокси получил плохой ответ
- 503 Service Unavailable - сервис недоступен

## Практический пример - простой сервер

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/ok", okHandler)
    http.HandleFunc("/created", createdHandler)
    http.HandleFunc("/notfound", notFoundHandler)
    http.HandleFunc("/error", errorHandler)
    
    fmt.Println("Сервер на :8080")
    http.ListenAndServe(":8080", nil)
}

func okHandler(w http.ResponseWriter, r *http.Request) {
    // 200 - по умолчанию
    fmt.Fprintf(w, "Все хорошо!")
}

func createdHandler(w http.ResponseWriter, r *http.Request) {
    // 201 - ресурс создан
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "Пользователь создан")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
    // 404 - не найден
    w.WriteHeader(http.StatusNotFound)
    fmt.Fprintf(w, "Страница не найдена")
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
    // 500 - ошибка сервера
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "Что-то пошло не так")
}
```

## Практический пример - проверка статуса в клиенте

```go
package main

import (
    "fmt"
    "io"
    "net/http"
)

func main() {
    urls := []string{
        "http://localhost:8080/ok",
        "http://localhost:8080/notfound", 
        "http://localhost:8080/error",
    }
    
    for _, url := range urls {
        checkStatus(url)
    }
}

func checkStatus(url string) {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Printf("Ошибка запроса: %v\n", err)
        return
    }
    defer resp.Body.Close()
    
    body, _ := io.ReadAll(resp.Body)
    
    // Проверяем статус код
    switch resp.StatusCode {
    case 200:
        fmt.Printf("✅ %s: %s\n", url, string(body))
    case 404:
        fmt.Printf("❌ %s: Не найден - %s\n", url, string(body))
    case 500:
        fmt.Printf("💥 %s: Ошибка сервера - %s\n", url, string(body))
    default:
        fmt.Printf("❓ %s: Статус %d - %s\n", url, resp.StatusCode, string(body))
    }
}
```

## Важные детали про статус-коды

**1. Порядок имеет значение:**
```go
// НЕПРАВИЛЬНО - статус после записи данных не сработает
fmt.Fprintf(w, "Данные")
w.WriteHeader(500) // Будет проигнорировано!

// ПРАВИЛЬНО
w.WriteHeader(500)
fmt.Fprintf(w, "Ошибка")
```

**2. Автоматический 200:**
```go
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello") // Автоматически установится 200 OK
}
```

**3. Перенаправления:**
```go
func redirectHandler(w http.ResponseWriter, r *http.Request) {
    // Автоматически устанавливает 302 и Location заголовок
    http.Redirect(w, r, "/new-page", http.StatusFound)
}
```

**4. Проверка в клиенте - важные паттерны:**
```go
resp, err := http.Get("http://example.com")
if err != nil {
    // Ошибка сети/подключения
    return err
}
defer resp.Body.Close()

if resp.StatusCode >= 400 {
    // Любая ошибка (4xx или 5xx)
    return fmt.Errorf("HTTP ошибка: %d", resp.StatusCode) 
}

// Только для 2xx статусов читаем данные
body, err := io.ReadAll(resp.Body)
```

## Практический пример - REST API с правильными статусами

```go
func userHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        w.WriteHeader(200) // OK
        fmt.Fprintf(w, `{"id": 1, "name": "Иван"}`)
        
    case "POST":
        w.WriteHeader(201) // Created
        fmt.Fprintf(w, `{"id": 2, "name": "Новый пользователь"}`)
        
    case "PUT":
        w.WriteHeader(200) // OK (обновлен)
        fmt.Fprintf(w, `{"id": 1, "name": "Обновленный"}`)
        
    case "DELETE":
        w.WriteHeader(204) // No Content (удален, нет данных)
        
    default:
        w.WriteHeader(405) // Method Not Allowed
        fmt.Fprintf(w, "Метод не поддерживается")
    }
}
```

---

# JSON в HTTP - базовые принципы

## Основная идея

HTTP передает текст. JSON - способ превратить объект в текст и обратно.

```
Объект → JSON строка → HTTP → JSON строка → Объект
```

## Минимальный пример сервера

```go
type User struct {
    Name string `json:"name"`
}

func handler(w http.ResponseWriter, r *http.Request) {
    var user User
    json.NewDecoder(r.Body).Decode(&user)  // JSON → объект
    
    fmt.Printf("Получил: %s\n", user.Name)
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)        // объект → JSON
}
```

## Минимальный пример клиента

```go
user := User{Name: "Иван"}
jsonData, _ := json.Marshal(user)        // объект → JSON

resp, _ := http.Post("http://localhost:8080", 
    "application/json", 
    bytes.NewBuffer(jsonData))
```

## Главное правило

**Всегда указывай Content-Type: application/json** - иначе сервер не поймет что это JSON.

---

# JSON в HTTP - передача структурированных данных

## Проблема, которую решает JSON

**HTTP передает только текст/байты. А как передать сложные данные?**

**Плохо - все в одной строке:**
```
GET /user?id=1&name=Иван&email=ivan@test.com&age=25&city=Москва
```

**Хорошо - структурированные данные:**
```json
{
  "id": 1,
  "name": "Иван", 
  "email": "ivan@test.com",
  "age": 25,
  "city": "Москва"
}
```

## Content-Type для JSON

**Обязательно указывать тип данных:**
```
POST /users HTTP/1.1
Content-Type: application/json

{"name": "Иван", "email": "ivan@test.com"}
```

**Без Content-Type сервер не поймет что это JSON!**

## Простой пример - сервер принимает JSON

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
    // Проверяем Content-Type
    if r.Header.Get("Content-Type") != "application/json" {
        w.WriteHeader(400)
        fmt.Fprintf(w, "Нужен Content-Type: application/json")
        return
    }
    
    // Декодируем JSON из тела запроса
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        w.WriteHeader(400)
        fmt.Fprintf(w, "Ошибка парсинга JSON: %v", err)
        return
    }
    
    // Обрабатываем данные
    user.ID = 123 // Присваиваем ID
    fmt.Printf("Получен пользователь: %+v\n", user)
    
    // Возвращаем JSON ответ
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    json.NewEncoder(w).Encode(user)
}

func main() {
    http.HandleFunc("/users", createUserHandler)
    fmt.Println("Сервер на :8080")
    http.ListenAndServe(":8080", nil)
}
```

## Простой пример - клиент отправляет JSON

```go
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func main() {
    // Создаем пользователя
    user := User{
        Name:  "Петр",
        Email: "petr@test.com",
    }
    
    // Сериализуем в JSON
    jsonData, err := json.Marshal(user)
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Отправляем JSON: %s\n", jsonData)
    
    // Отправляем POST запрос
    resp, err := http.Post(
        "http://localhost:8080/users",
        "application/json",           // Content-Type
        bytes.NewBuffer(jsonData),    // Тело запроса
    )
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    
    // Читаем ответ
    body, _ := io.ReadAll(resp.Body)
    fmt.Printf("Статус: %d\n", resp.StatusCode)
    fmt.Printf("Ответ: %s\n", body)
    
    // Парсим JSON ответ
    var createdUser User
    json.Unmarshal(body, &createdUser)
    fmt.Printf("Создан пользователь с ID: %d\n", createdUser.ID)
}
```

## Важные детали JSON в HTTP

**1. json.NewDecoder vs json.Unmarshal:**
```go
// Для HTTP тела - используйте Decoder (эффективнее)
var user User
json.NewDecoder(r.Body).Decode(&user)

// Для готовых байтов - используйте Unmarshal
var user User  
json.Unmarshal(jsonBytes, &user)
```

**2. JSON теги в структурах:**
```go
type User struct {
    ID       int    `json:"id"`                    // обычное поле
    Name     string `json:"name"`                  
    Email    string `json:"email"`
    Password string `json:"password,omitempty"`    // не включать если пустое
    Secret   string `json:"-"`                     // никогда не включать в JSON
    Age      int    `json:"age,string"`            // число как строка в JSON
}
```

**3. Обработка ошибок JSON:**
```go
func parseJSON(r *http.Request) (*User, error) {
    var user User
    
    // Ограничиваем размер тела запроса (защита от больших JSON)
    r.Body = http.MaxBytesReader(nil, r.Body, 1048576) // 1MB
    
    decoder := json.NewDecoder(r.Body)
    decoder.DisallowUnknownFields() // Ошибка при неизвестных полях
    
    err := decoder.Decode(&user)
    if err != nil {
        return nil, fmt.Errorf("ошибка JSON: %v", err)
    }
    
    return &user, nil
}
```

**4. Типичные ошибки Content-Type:**
```go
// НЕПРАВИЛЬНО - забыли Content-Type
http.Post("http://example.com/api", "", jsonData)

// ПРАВИЛЬНО
http.Post("http://example.com/api", "application/json", jsonData)

// Или через Request:
req.Header.Set("Content-Type", "application/json; charset=utf-8")
```

## Практический пример - полный CRUD с JSON

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
    "strings"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

var users = make(map[int]User)
var nextID = 1

func usersHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    switch r.Method {
    case "GET":
        // Возвращаем всех пользователей
        userList := make([]User, 0, len(users))
        for _, user := range users {
            userList = append(userList, user)
        }
        json.NewEncoder(w).Encode(userList)
        
    case "POST":
        // Создаем нового пользователя
        var user User
        if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
            w.WriteHeader(400)
            json.NewEncoder(w).Encode(map[string]string{"error": "Неверный JSON"})
            return
        }
        
        user.ID = nextID
        users[nextID] = user
        nextID++
        
        w.WriteHeader(201)
        json.NewEncoder(w).Encode(user)
    }
}

func userHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    // Извлекаем ID из URL
    idStr := strings.TrimPrefix(r.URL.Path, "/users/")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        w.WriteHeader(400)
        json.NewEncoder(w).Encode(map[string]string{"error": "Неверный ID"})
        return
    }
    
    switch r.Method {
    case "GET":
        user, exists := users[id]
        if !exists {
            w.WriteHeader(404)
            json.NewEncoder(w).Encode(map[string]string{"error": "Пользователь не найден"})
            return
        }
        json.NewEncoder(w).Encode(user)
        
    case "DELETE":
        if _, exists := users[id]; !exists {
            w.WriteHeader(404)
            json.NewEncoder(w).Encode(map[string]string{"error": "Пользователь не найден"})
            return
        }
        delete(users, id)
        w.WriteHeader(204) // No Content
    }
}

func main() {
    // Добавляем тестового пользователя
    users[1] = User{ID: 1, Name: "Иван", Email: "ivan@test.com"}
    nextID = 2
    
    http.HandleFunc("/users", usersHandler)
    http.HandleFunc("/users/", userHandler)
    
    fmt.Println("JSON API сервер на :8080")
    fmt.Println("GET  /users     - список пользователей")
    fmt.Println("POST /users     - создать пользователя")  
    fmt.Println("GET  /users/1   - получить пользователя")
    fmt.Println("DELETE /users/1 - удалить пользователя")
    
    http.ListenAndServe(":8080", nil)
}
```

**Тестирование с curl:**
```bash
# Получить всех пользователей
curl http://localhost:8080/users

# Создать пользователя  
curl -X POST -H "Content-Type: application/json" \
     -d '{"name":"Мария","email":"maria@test.com"}' \
     http://localhost:8080/users

# Получить пользователя по ID
curl http://localhost:8080/users/1
```

---

# Формы в HTTP - альтернатива JSON

## Проблема, которую решают формы

JSON - это для API между программами. Но что если данные вводит человек в браузере?

**HTML форма:**
```html
<form method="POST" action="/submit">
    <input type="text" name="username" value="ivan">
    <input type="password" name="password" value="secret">
    <button type="submit">Войти</button>
</form>
```

**Превращается в HTTP запрос:**
```
POST /submit HTTP/1.1
Content-Type: application/x-www-form-urlencoded

username=ivan&password=secret
```

## Ключевые отличия от JSON

**JSON:**
```
Content-Type: application/json
{"username": "ivan", "password": "secret"}
```

**Форма:**
```
Content-Type: application/x-www-form-urlencoded
username=ivan&password=secret
```

## Простой пример - сервер принимает форму

```go
func loginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        // Показываем форму
        fmt.Fprintf(w, `
        <form method="POST">
            <input name="username" placeholder="Логин">
            <input name="password" type="password" placeholder="Пароль">
            <button type="submit">Войти</button>
        </form>`)
        return
    }
    
    // Парсим форму
    r.ParseForm()
    
    username := r.FormValue("username")
    password := r.FormValue("password")
    
    fmt.Fprintf(w, "Логин: %s, Пароль: %s", username, password)
}
```

## Простой пример - клиент отправляет форму

```go
data := url.Values{}
data.Set("username", "ivan")
data.Set("password", "secret")

resp, _ := http.Post("http://localhost:8080/login",
    "application/x-www-form-urlencoded",
    strings.NewReader(data.Encode()))
```

## Важные детали

**1. Обязательно вызывать ParseForm():**
```go
r.ParseForm()                    // Без этого FormValue не работает
username := r.FormValue("username")
```

**2. URL encoding:**
```
Пробелы → +
Специальные символы → %XX
ivan petrov → ivan+petrov
```

---

# Загрузка файлов - multipart/form-data

## Проблема

Обычные формы передают только текст. А как передать файл?

**Не работает:**
```
Content-Type: application/x-www-form-urlencoded
filename=photo.jpg&content=BINARY_DATA_HERE???
```

**Решение - multipart:**
```
Content-Type: multipart/form-data; boundary=----1234

------1234
Content-Disposition: form-data; name="username"

ivan
------1234  
Content-Disposition: form-data; name="file"; filename="photo.jpg"
Content-Type: image/jpeg

[BINARY DATA OF FILE]
------1234--
```

## HTML форма для файлов

```html
<form method="POST" enctype="multipart/form-data" action="/upload">
    <input type="text" name="description" placeholder="Описание">
    <input type="file" name="file">
    <button type="submit">Загрузить</button>
</form>
```

**Ключевое:** `enctype="multipart/form-data"`

## Простой пример - сервер принимает файл

```go
func uploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        fmt.Fprintf(w, `
        <form method="POST" enctype="multipart/form-data">
            <input name="description" placeholder="Описание">
            <input name="file" type="file">
            <button type="submit">Загрузить</button>
        </form>`)
        return
    }
    
    // Парсим multipart form (лимит 10MB)
    r.ParseMultipartForm(10 << 20)
    
    // Получаем обычное поле
    description := r.FormValue("description")
    
    // Получаем файл
    file, header, err := r.FormFile("file")
    if err != nil {
        fmt.Fprintf(w, "Ошибка: %v", err)
        return
    }
    defer file.Close()
    
    fmt.Fprintf(w, "Файл: %s, Размер: %d, Описание: %s", 
        header.Filename, header.Size, description)
}
```

## Простой пример - клиент отправляет файл

```go
func uploadFile() {
    // Создаем multipart writer
    var b bytes.Buffer
    writer := multipart.NewWriter(&b)
    
    // Добавляем обычное поле
    writer.WriteField("description", "Мое фото")
    
    // Добавляем файл
    fileWriter, _ := writer.CreateFormFile("file", "test.txt")
    fileWriter.Write([]byte("Содержимое файла"))
    
    writer.Close()
    
    // Отправляем
    resp, _ := http.Post("http://localhost:8080/upload",
        writer.FormDataContentType(),  // Content-Type с boundary
        &b)
    
    defer resp.Body.Close()
}
```

## Ключевые детали

**1. Лимиты размера:**
```go
r.ParseMultipartForm(10 << 20)  // 10MB лимит
```

**2. Boundary автоматически:**
```go
writer.FormDataContentType()  // "multipart/form-data; boundary=..."
```

**3. Сохранение файла на диск:**
```go
file, header, _ := r.FormFile("file")
dst, _ := os.Create(header.Filename)
io.Copy(dst, file)  // Копируем содержимое
```

---

# Middleware - общая логика для всех запросов

## Проблема

У тебя есть 10 API endpoints. И для каждого нужно:
- Логировать запрос
- Проверить авторизацию  
- Добавить CORS заголовки

**Плохо - копипаста в каждом handler:**
```go
func handler1(w http.ResponseWriter, r *http.Request) {
    log.Printf("Запрос: %s", r.URL)        // Повторяется
    if !checkAuth(r) { return }            // Повторяется  
    w.Header().Set("Access-Control-Allow-Origin", "*")  // Повторяется
    
    // Реальная логика handler
}

func handler2(w http.ResponseWriter, r *http.Request) {
    log.Printf("Запрос: %s", r.URL)        // Повторяется
    if !checkAuth(r) { return }            // Повторяется
    w.Header().Set("Access-Control-Allow-Origin", "*")  // Повторяется
    
    // Реальная логика handler
}
```

## Решение - Middleware

**Middleware = функция, которая "оборачивает" handler и добавляет общую логику.**

## Простейший пример

```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Запрос: %s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)  // Вызываем следующий handler
    })
}

func myHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Привет!")
}

func main() {
    // Оборачиваем handler в middleware
    http.Handle("/", loggingMiddleware(http.HandlerFunc(myHandler)))
    http.ListenAndServe(":8080", nil)
}
```

**Что происходит:**
1. Запрос приходит в loggingMiddleware
2. Middleware логирует запрос
3. Middleware вызывает myHandler через next.ServeHTTP()
4. myHandler отрабатывает и возвращает ответ

## Цепочка middleware

```go
func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token != "Bearer secret" {
            http.Error(w, "Нет доступа", 401)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        next.ServeHTTP(w, r)
    })
}

func main() {
    handler := http.HandlerFunc(myHandler)
    
    // Цепочка: logging → auth → cors → myHandler
    wrappedHandler := loggingMiddleware(authMiddleware(corsMiddleware(handler)))
    
    http.Handle("/protected", wrappedHandler)
    http.ListenAndServe(":8080", nil)
}
```

## Практический пример - API с middleware

```go
func timeMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("Запрос выполнен за %v", time.Since(start))
    })
}

func jsonMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        next.ServeHTTP(w, r)
    })
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, `{"message": "API работает", "time": "%s"}`, time.Now().Format("15:04:05"))
}

func main() {
    handler := http.HandlerFunc(apiHandler)
    
    // Все API endpoint будут иметь JSON Content-Type и логирование времени
    wrappedHandler := timeMiddleware(jsonMiddleware(handler))
    
    http.Handle("/api", wrappedHandler)
    
    fmt.Println("Сервер на :8080")
    http.ListenAndServe(":8080", nil)
}
```

## Главная идея

**Middleware позволяет добавлять общую логику ко всем handlers без копирования кода.**

Типичное применение:
- Логирование запросов
- Авторизация/аутентификация  
- CORS заголовки
- Измерение времени выполнения
- Ограничение rate limit

---

