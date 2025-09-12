> **сделай HTTP-сервер на net/http, который на `/version` отвечает строкой `"v1.0.0"`**

---

## Уровень L0 — просто работает

```go
mux := http.NewServeMux()
mux.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("v1.0.0"))
})
log.Fatal(http.ListenAndServe(":8080", mux))
```

☑️ ты видишь «v1.0.0» в браузере или через curl.

---

## Уровень L1 — корректность (метод + заголовки)

```go
mux.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("v1.0.0"))
})
```

✅ теперь ты явно разрешаешь только GET
✅ клиенту всегда ясно, что это текст в UTF-8

---

## Уровень L2 — защита и устойчивость

для этой задачи тело не читаем, но стоит добавить **границу по времени** на весь сервер:

```go
srv := &http.Server{
	Addr:         ":8080",
	Handler:      mux,
	ReadTimeout:  5 * time.Second,
	WriteTimeout: 5 * time.Second,
	IdleTimeout:  60 * time.Second,
}
log.Fatal(srv.ListenAndServe())
```

✅ если клиент зависнет и не дочитает/не допишет — сервер не залипнет навсегда.

---

## Уровень L3 — наблюдаемость

оборачиваем хендлер в простую «прослойку» с логированием:

```go
func WithLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

mux := http.NewServeMux()
mux.Handle("/version", WithLog(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// сам обработчик, как в L1
})))
```

✅ теперь в логах видно метод, путь и задержку.

---

## Чек-лист для задачи №1

- [ ] GET /version возвращает `v1.0.0`
- [ ] Другие методы → 405
- [ ] Content-Type: text/plain; charset=utf-8
- [ ] Сервер падает по таймауту, если клиент «подвис»
- [ ] Логирует метод, путь, время

## Немного улучшим L3 - добавим логирование статуса и добавим обработку метода HEAD

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// 1. Сделай HTTP-сервер на net/http,
// который на /version отвечает строкой "v1.0.0".

// 2. На запрос GET /user?id=123 сервер должен вернуть "User ID: 123".
type sr struct {
	http.ResponseWriter // встраивание (embedding)
	code int // сюда запишем статус-код
}

func (s *sr) WriteHeader(c int) {
	s.code = c // 1) запомнили код
	s.ResponseWriter.WriteHeader(c) // 2) пробросили «вниз» (реально отправили клиенту)
}

func WithLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := &sr{
			ResponseWriter: w,
			code:           http.StatusOK, // ← по умолчанию считаем 200
		}
		start := time.Now()
		next.ServeHTTP(rec, r) // ← передаём ВМЕСТО w наш rec
		log.Printf("%s %s %d %v\n", r.Method, r.URL.Path, rec.code, time.Since(start))
	})
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/user", WithLog(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Разрешаем GET и HEAD
		if r.Method != http.MethodGet && r.Method != http.MethodHead {
			w.Header().Set("Allow", "GET, HEAD")
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Missing 'id' parametr", http.StatusBadRequest)
			return
		}

		// Для HEAD — только заголовки, без тела
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		if r.Method == http.MethodHead {
			w.WriteHeader(http.StatusOK)
			return
		}
		fmt.Fprintf(w, "User ID: %s", id)
	})))

	srv := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
```

### Пояснения

отличный вопрос! разберём «по кирпичикам», что делает этот кусок.

## 1) Зачем вообще вся эта конструкция?

Стандартный `net/http` **не говорит тебе**, какой статус-код (200/404/500…) отправил хендлер. Чтобы залогировать статус, нужно **перехватить** момент, когда хендлер вызывает `WriteHeader`. Для этого мы «подменяем» `ResponseWriter` на свой обёрнутый, который запоминает код.

## 2) Что такое `sr` и зачем в нём встраивание

```go
type sr struct {
    http.ResponseWriter // встраивание (embedding)
    code int            // сюда запишем статус-код
}
```

- Мы **встраиваем** (`embedding`) оригинальный `http.ResponseWriter`.
  Благодаря этому `sr` автоматически имеет методы `Header()` и `Write(...)` — как у оригинала.
- Добавляем поле `code` — тут будем хранить отправленный статус.

## 3) Перехватываем статус через `WriteHeader`

```go
func (s *sr) WriteHeader(c int) {
    s.code = c                                // 1) запомнили код
    s.ResponseWriter.WriteHeader(c)           // 2) пробросили «вниз» (реально отправили клиенту)
}
```

- Любой хендлер, который делает `w.WriteHeader(404)`, теперь фактически вызывает **наш** метод, и мы можем записать `404` в `s.code`.

> Почему не перехватываем `Write`? Можно и `Write` перехватывать, но статус можно поменять только через `WriteHeader`. Если хендлер **не** вызывает `WriteHeader`, при первом `Write` Go сам отправит `200 OK` — мы учли это инициализацией (см. следующий пункт).

## 4) Middleware `WithLog`: как работает обёртка

```go
func WithLog(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        rec := &sr{ResponseWriter: w, code: http.StatusOK} // ← по умолчанию считаем 200
        start := time.Now()

        next.ServeHTTP(rec, r) // ← передаём ВМЕСТО w наш rec

        log.Printf("%s %s %d %v", r.Method, r.URL.Path, rec.code, time.Since(start))
    })
}
```

Пошагово:

1. Создаём `rec` — это наш «рекордер» с полем `code`. Сразу ставим `200` на случай, если хендлер **не** вызовет `WriteHeader` (типичный случай для успешного ответа).
2. Вызываем «внутренний» обработчик `next.ServeHTTP`, **подкладывая** ему `rec` вместо оригинального `w`.

   - Если внутри вызовут `w.WriteHeader(404)` — попадём в наш `(*sr).WriteHeader` и зафиксируем `404`.
   - Если внутри **не** вызовут `WriteHeader` и сразу напишут тело (`w.Write(...)`) — Go отправит `200`; у нас уже и так `code=200`.

3. Когда `next` завершится, у нас есть и **длительность**, и **итоговый статус** в `rec.code` — логируем.

## 5) Наглядный «трасс»

Представь хендлер:

```go
func h(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/ok" {
        w.WriteHeader(404)
        w.Write([]byte("not found"))
        return
    }
    w.Write([]byte("ok")) // без WriteHeader → отправится 200 OK
}
```

Путь запроса `"/bad"`:

- WithLog создаёт `rec(code=200)`, запускает `next.ServeHTTP(rec, r)`.
- Внутри `h`: `w.WriteHeader(404)` → это `rec.WriteHeader(404)` → `rec.code=404`.
- Выходим из `next`, логируем `... 404 <latency>`.

Путь запроса `"/ok"`:

- Внутри `h` вызывают сразу `w.Write(...)` → Go сам шлёт `200 OK`.
- `rec.code` остался равен 200 (мы так и инициализировали).
- Логируем `... 200 <latency>`.

## 6) Что ещё можно (но не обязательно)

Если хочешь **считать объём ответа**, можно дополнительно перехватить `Write`:

```go
type sr struct {
    http.ResponseWriter
    code int
    n    int64 // байт записано
}
func (s *sr) Write(p []byte) (int, error) {
    // если кто-то вызвал Write до WriteHeader — статус уже 200 у нас по умолчанию
    n, err := s.ResponseWriter.Write(p)
    s.n += int64(n)
    return n, err
}
```

Тогда в логе можно печатать и размер: `... code bytes duration`.

---

### Короткая памятка (как запомнить)

- `sr` — **s**tatus **r**ecorder: обёртка над `ResponseWriter`, где мы «зеркалим» статус.
- Инициализируй `code=200`, чтобы покрыть путь без `WriteHeader`.
- `WithLog(next)` → создаёт `rec`, вызывает `next(rec, r)`, потом логирует `rec.code`.

Если хочешь, покажу минимальный пример «конца в конец»: хендлер, который иногда даёт 404, и как лог меняется для обоих случаев.
