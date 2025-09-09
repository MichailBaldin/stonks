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
