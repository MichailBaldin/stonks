# Сервер
## Простой сервер
```go
func main() {
    // запуск сервера, который слушает 8080 порт
    http.ListenAndServe(":8080", nil)
}
```

## Добавляем ручку серверу
```go
func main() {
    // регистрируем ручку /hello
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        // ответ сервера на запрос к /hello
		w.Write([]byte("Hi!\n"))
	})
}
```

## Добавляем параментры к запросу ручки сервера
### Один параметр
```go
func main() {
    // регистрируем ручку /hello
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        // читаем параметр name из запроса
        name := r.URL.Query().Get("name")
        // формируем ответ
        response := fmt.Sprintf("Hello, %s", name)
        // передаем ответ в поток стандартного вывода
        fmt.Fprint(w, response)
	})
}
```

Запрос
необходимо экранировать `?` в zsh-оболочке
```bash
curl http://localhost:8080/hello\?name=Mick
```

### Несколько параметров в одном запросе
```go
func main() {
    // регистрируем ручку /hello
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        // читаем параметр name из запроса
        name := r.URL.Query().Get("name")
        // формируем ответ
        response := fmt.Sprintf("Hello, %s", name)
        // передаем ответ в поток стандартного вывода
        fmt.Fprint(w, response)
	})
}
```

Запрос
необходимо экранировать `?` и `&` в zsh-оболочке
```bash
curl http://localhost:8080/hello\?name=Mick\&age=34
```

### Передача парметров через несколько последовательных запросов
Для реализации подобной логики придется использовать
1. Cookies (данные хранятся в браузере)
2. Sessions (данные хранятся на сервере)

## Post-запросы
Помимо запросов `GET` существуют запросы `POST`.
```go
func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		defer r.Body.Close()

		w.Write(body)
	})
}
```
В `POST` запросе мы передаем информацию в теле запроса. Поэтому его необходимо распарсить.
НЕЗАБЫВАЕМ закрыть тело после использования.

## Возвращение JSON в ответе на запрос
```go
func main() {
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
        // заводим простую мапу, чтобы потом ее преобразовать в JSON
		response := map[string]string{"status": "ok"}
        // Создаем новый Encoder, который будет писать в `w` и преобразует мапу в JSON
		json.NewEncoder(w).Encode(response)
	})

	http.ListenAndServe(":8080", nil)
}
```

## Работа с заголовками
```go
func main() {
	http.HandleFunc("/headers", func(w http.ResponseWriter, r *http.Request) {
		var headersBuilder strings.Builder

        // доступ к заголовкам получаем через r.Header
        // тип r.Header map[string][]string
        // значит нужный заголовок можем достать с помощью `r.Header["Key"]`
        // Ниже просто проходим по всем заголовкам и соединяем в одну строку
		for name, values := range r.Header {
			headersBuilder.WriteString(fmt.Sprintf("%s: %s\n", name, strings.Join(values, ", ")))
		}

        // результат отправляем в `w`
		fmt.Fprint(w, headersBuilder.String())
	})

	http.ListenAndServe(":8080", nil)
}
```

Для тестирования заголовки придется отправить в запросе явно
```bash
curl -H "X-Custom-Header: Value1" -H "X-Another-Header: Value2" http://localhost:8080/headers
```

## Обработка JSON сервером
```go
// заводим структуры для запроса и ответа
type Request struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Response struct {
	Sum int `json:"sum"`
}

func main() {
	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		// создаем переменную, в которую будем парсить запрос
        var req Request
		_ = json.NewDecoder(r.Body).Decode(&req)
		defer r.Body.Close()

        // делаем необходимые вычисления
		sum := req.X + req.Y

        // формируем ответ и заполняем структуру для ответа
		resp := Response{Sum: sum}
        
        // преобразуем ее в JSON и отправляем в `w`
		_ = json.NewEncoder(w).Encode(resp)
	})

	http.ListenAndServe(":8080", nil)
}
```

## Middleware
```go
func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	http.Handle("/", LogMiddleware(handler))

	http.ListenAndServe(":8080", nil)
}
```

## Парсинг набора URL
```go
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/user/") {
			pathSuffix := r.URL.Path[len("/user/"):]

			id := strings.TrimSuffix(pathSuffix, "/")

			if id != "" {
				fmt.Fprintf(w, "User ID: %s\n", id)
				return
			}
		}

		http.NotFound(w, r)
	})

	http.ListenAndServe(":8080", nil)
}
```

# Клиент