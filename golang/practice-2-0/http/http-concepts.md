отлично, понял запрос: давай «ду́пло»-уровень — самые частые микрошаблоны, без роутеров и «архитектур». Сначала клиент (делаем запросы), потом сервер (принимаем), и отдельно — когда работать с `[]byte`, а когда с JSON/структурами.

---

# Клиент: 10 ежедневных микрошаблонов

1. **Быстрый GET (и не забыть закрыть Body)**

```go
resp, err := http.Get("https://example.com/data")
if err != nil { /* handle */ }
defer resp.Body.Close()
b, err := io.ReadAll(resp.Body) // осторожно: без лимита
_ = b
```

2. **GET с таймаутом (рекомендуемый базовый)**

```go
cli := &http.Client{Timeout: 5 * time.Second}
resp, err := cli.Get("https://example.com/data")
defer resp.Body.Close()
```

3. **POST JSON (минимум)**

```go
payload := bytes.NewBufferString(`{"a":1,"b":2}`)
req, _ := http.NewRequest(http.MethodPost, "https://api.example.com/sum", payload)
req.Header.Set("Content-Type", "application/json")
resp, err := (&http.Client{Timeout: 5*time.Second}).Do(req)
defer resp.Body.Close()
```

4. **GET с query-параметрами (корректная сборка URL)**

```go
u, _ := url.Parse("https://api.example.com/search")
q := u.Query()
q.Set("q", "golang")
q.Set("limit", "10")
u.RawQuery = q.Encode()
resp, _ := http.Get(u.String())
defer resp.Body.Close()
```

5. **Чтение ответа безопасно (лимитируем размер)**

```go
const max = 1 << 20 // 1MB
limited := io.LimitReader(resp.Body, max)
data, err := io.ReadAll(limited)
```

6. **Стриминг в файл (не держим всё в памяти)**

```go
f, _ := os.Create("out.bin")
defer f.Close()
_, err := io.Copy(f, resp.Body) // потоково, без буфера в память
```

7. **Кастомный Transport (пул соединений по-умолчанию)**

```go
tr := &http.Transport{
	MaxIdleConns:        100,
	MaxIdleConnsPerHost: 10,
	IdleConnTimeout:     90 * time.Second,
}
cli := &http.Client{Transport: tr, Timeout: 5 * time.Second}
```

8. **Запрос с контекстом и дедлайном**

```go
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel()
req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://example.com", nil)
resp, err := (&http.Client{}).Do(req)
defer resp.Body.Close()
```

9. **Форма (application/x-www-form-urlencoded)**

```go
form := url.Values{"a":{"1"}, "b":{"2"}}
resp, _ := http.PostForm("https://api.example.com/sum", form)
defer resp.Body.Close()
```

10. **Multipart upload (файл)**

```go
buf := &bytes.Buffer{}
mw := multipart.NewWriter(buf)
fw, _ := mw.CreateFormFile("file", "photo.jpg")
io.Copy(fw, bytes.NewReader(fileBytes)) // или os.Open(...)
mw.Close()

req, _ := http.NewRequest(http.MethodPost, "https://api.example.com/upload", buf)
req.Header.Set("Content-Type", mw.FormDataContentType())
resp, _ := (&http.Client{Timeout: 10*time.Second}).Do(req)
defer resp.Body.Close()
```

---

# Сервер: 10 ежедневных микрошаблонов

1. **Мини-сервер + health**

```go
mux := http.NewServeMux()
mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK); w.Write([]byte("ok"))
})
log.Fatal(http.ListenAndServe(":8080", mux))
```

2. **Гард по методу**

```go
mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet { http.Error(w, "method not allowed", 405); return }
	w.Write([]byte("pong"))
})
```

3. **Читаем тело как байты (с лимитом)**

```go
mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil { http.Error(w, "too large or read err", 400); return }
	w.Write(b) // отдали как есть
})
```

4. **JSON-вход → JSON-выход**

```go
mux.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
	type In struct{ A, B int }
	type Out struct{ Sum int }
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	defer r.Body.Close()
	var in In
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil { http.Error(w, "bad json", 400); return }
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Out{Sum: in.A + in.B})
})
```

5. **Query-параметры с дефолтами**

```go
mux.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" { name = "world" }
	fmt.Fprintf(w, "hi %s", name)
})
```

6. **Ограничение времени обработки (по контексту запроса)**

```go
mux.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
	select {
	case <-time.After(2 * time.Second):
		w.Write([]byte("done"))
	case <-r.Context().Done():
		http.Error(w, "canceled", 499) // пользователь ушёл
	}
})
```

7. **Выдача файла (скачивание)**

```go
mux.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Disposition", `attachment; filename="report.txt"`)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, "hello\n")
})
```

8. **Приём multipart файла (потоково)**

```go
mux.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil { http.Error(w, "bad multipart", 400); return }
	f, _, err := r.FormFile("file")
	if err != nil { http.Error(w, "no file", 400); return }
	defer f.Close()
	n, _ := io.Copy(io.Discard, f) // здесь могла быть запись в диск/обработчик
	fmt.Fprintf(w, "got %d bytes", n)
})
```

9. **Стриминговый ответ (chunked)**

```go
mux.HandleFunc("/stream", func(w http.ResponseWriter, r *http.Request) {
	fl, ok := w.(http.Flusher); if !ok { http.Error(w, "no flusher", 500); return }
	for i := 1; i <= 5; i++ { fmt.Fprintf(w, "chunk %d\n", i); fl.Flush(); time.Sleep(200*time.Millisecond) }
})
```

10. **Быстрый CORS (минимализм для dev)**

```go
func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == http.MethodOptions { w.WriteHeader(204); return }
		next.ServeHTTP(w, r)
	})
}
```

---

# Когда `[]byte`, а когда JSON/структуры?

## Держать как `[]byte` (или поток `io.Reader`) имеет смысл, когда:

* **Не знаешь** схему заранее (проксируешь/логируешь/пробрасываешь «как есть»).
* **Большие данные** или бинарь (изображения, архивы): читаем/пишем **потоком** (через `io.Copy`), чтобы не грузить память.
* **Минимальная латентность** и «zero-copy»-подход: меньше трансформаций — быстрее.
* **Content-Type ≠ JSON** (PDF, CSV, Protobuf и т.п.) — часто передаём «как есть».

## Преобразовывать в структуру (JSON) лучше, когда:

* Ты **контролируешь контракт** и нужна **валидация**/осмысленные ошибки.
* Нужны **семантические поля**, типобезопасность, автодокументация.
* Размер умеренный (десятки/сотни КБ), и важнее удобство, чем микро-перф.

## Частые решения-паттерны:

* **Вход неизвестный/большой** → `io.Copy` в файл/обработчик (никакого `ReadAll`).
* **Вход JSON, контролируешь схему** → `json.Decoder` в struct + лимит через `MaxBytesReader`.
* **Выход большой/стрим** → writer + `Flusher` или `io.Copy` из источника (файл/pipe) в `w`.
* **Выход JSON** → заголовок `Content-Type: application/json` + `json.Encoder`.

---

# Маленький чек-лист «каждый день»

* Клиент всегда с **Timeout**.
* Всегда `defer resp.Body.Close()`.
* Чтение тела: **лимитируй** (`LimitReader`/`MaxBytesReader`).
* Для больших тел: **streaming** (`io.Copy`), не `ReadAll`.
* Уважай `Content-Type` и код ответа (`StatusCode`).
* Не забывай про `Context` на долгих операциях.

---

хочешь, закрепим? выбери **одну** микросцену на практику прямо сейчас:

A) клиент: скачиваем **большой файл** и пишем на диск без загрузки в память;
B) сервер: принимаем **большой upload** и потоково кладём в файл;
C) клиент: корректный **GET с query**, таймаутом и «безопасным» чтением с лимитом;
D) сервер: **JSON-вход → JSON-выход** с валидацией диапазона (например, `0 <= a,b <= 1e6`).

Скажи букву — и я дам короткое задание + эталон (после твоей попытки).

мне нравится твой тренажёр: он ровно про «ду́пло»-навыки, которые реально нужны каждый день. Как улучшить — чтобы он учил не только «как вернуть строку», но и **микропаттерны качества**: таймауты, лимиты, заголовки, статусы, валидация, байты vs JSON, потоковая обработка.

ниже — апгрейд-рамка + мини-шаблоны.

# Как улучшить тренажёр

## 1) Разбей задачи на «уровни качества»

* **L0 — I/O база:** как у тебя сейчас (прочитал → ответил).
* **L1 — корректность:** проверка метода, статусы (`400/405/415/500`), `Content-Type`.
* **L2 — защита и устойчивость:** лимит тела (`MaxBytesReader`), «безопасное чтение» (`LimitReader`), явные таймауты/дедлайны.
* **L3 — производительность:** стриминг (`io.Copy`), избегаем `ReadAll` на больших телах.
* **L4 — наблюдаемость:** логируем метод/путь/латентность, корреляционный ID (из заголовка), базовый метрик-хук.

Каждую твою задачу можно «прокачать» по этим уровням — 5 минут на уровень.

## 2) Добавь чёткие критерии принятия (Definition of Done)

Для каждой задачи короткий чек-лист:

* Разрешённый метод?
* Правильный `Content-Type` **в ответе**?
* Обработка пустых/битых входов?
* Предел размера тела?
* Для чисел: что при `n` не число / отсутствует параметр?

## 3) Три «режима передачи данных»

Тренируй каждый эндпоинт в одной из трёх схем:

* **Bytes-in / bytes-out** (эхо, реверс, copy): потоково, без лишних аллокаций.
* **JSON-in / JSON-out** (health, data): `Decoder`/`Encoder`, теги, нулевые значения.
* **Query-in / text-out** (subtract, multiply, square, greet): валидация query.

---

# Мини-шаблоны (вставляй в апгрейды)

## A) Метод-гард + тип ответа

```go
if r.Method != http.MethodGet {
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	return
}
w.Header().Set("Content-Type", "text/plain; charset=utf-8")
```

## B) Безопасное чтение тела (bytes-in/out)

```go
r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1MB
defer r.Body.Close()
data, err := io.ReadAll(r.Body) // ок для малых тел
if err != nil { http.Error(w, "read error", 400); return }
```

## C) Стриминг без буфера в память (copy/upload)

```go
_, err := io.Copy(w, io.LimitedReader{R: r.Body, N: 1<<20}) // прокидываем ≤1MB
if err != nil { http.Error(w, "stream error", 500) }
```

## D) JSON-in / JSON-out с валидацией

```go
type In struct{ Name string `json:"name"`; Age int `json:"age"` }
type Out struct{ Message string `json:"message"` }

r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
defer r.Body.Close()
dec := json.NewDecoder(r.Body); dec.DisallowUnknownFields()
var in In
if err := dec.Decode(&in); err != nil { http.Error(w, "bad json", 400); return }
if in.Name == "" || in.Age < 0 { http.Error(w, "bad data", 422); return }

w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(Out{Message: fmt.Sprintf("Hello %s, you are %d", in.Name, in.Age)})
```

## E) Query-параметры с дефолтами и проверкой

```go
q := r.URL.Query()
a, errA := strconv.Atoi(q.Get("a"))
b, errB := strconv.Atoi(q.Get("b"))
if errA != nil || errB != nil { http.Error(w, "bad query", 400); return }
fmt.Fprintf(w, "%d", a-b)
```

## F) Ответ времени (формат + зона)

```go
w.Header().Set("Content-Type", "text/plain; charset=utf-8")
fmt.Fprint(w, time.Now().Format("2006-01-02 15:04:05"))
```

## G) Рандом (crypto/rand → int в диапазоне)

```go
nBig, _ := rand.Int(rand.Reader, big.NewInt(100)) // 0..99
fmt.Fprint(w, nBig.Int64()+1)                     // 1..100
```

## H) Базовый лог-мидлвар

```go
func WithLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}
```

---

# Пример: апгрейд твоего задания №3 (POST /copy)

**L0:** прочитать тело → вернуть `Copy: <тело>`.
**L1:** только `POST`, `Content-Type` ответа `text/plain`.
**L2:** лимит 1MB, корректная ошибка при превышении.
**L3:** для больших тел — потоково (`io.Copy`): сначала пишем префикс, затем копируем тело без `ReadAll`.

шаблон L3:

```go
mux.HandleFunc("/copy", func(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { http.Error(w, "method not allowed", 405); return }
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// Пишем префикс сразу
	if _, err := io.WriteString(w, "Copy: "); err != nil { return }

	// Копируем тело с лимитом 1MB
	_, err := io.Copy(w, io.LimitedReader{R: r.Body, N: 1<<20})
	if err != nil { http.Error(w, "stream error", 500) }
})
```

---

# Тестирование (скелет для любой задачи)

```go
func TestSubtract(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/subtract?a=10&b=3", nil)
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK { t.Fatalf("code=%d", rr.Code) }
	if strings.TrimSpace(rr.Body.String()) != "7" { t.Fatalf("body=%q", rr.Body.String()) }
}
```

---

хочешь, начнём улучшать твой список **по одному**?
Выбери **одну** задачу из твоих 1–16 — и я превращу её в «уровни L0→L4» (с чек-листом и небольшими подсказками). Какую берём первой?

отлично, давай разберём твою задачу №1:

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

* [ ] GET /version возвращает `v1.0.0`
* [ ] Другие методы → 405
* [ ] Content-Type: text/plain; charset=utf-8
* [ ] Сервер падает по таймауту, если клиент «подвис»
* [ ] Логирует метод, путь, время

---

👉 Видишь идею? Мы берём **очень простую задачу**, и вместо «один раз написать», тренируем на ней 3–4 микро-паттерна, которые потом будут встречаться **в каждой** ручке.

Хочешь, мы так же прокачаем твою задачу №2 (`GET /user?id=123`)?

круто, берём твой список и превращаем его в «ду́пло»-тренажёр с осознанными микропаттернами. Я не буду решать за тебя — дам **рамку (L0→L3)**, **универсальные мини-шаблоны**, и **карту всех 50 задач**: что за метод, какой тип ввода/вывода, какие заголовки/статусы и обязательные проверки. Потом пойдём по партиям.

## Рамка качества для любой ручки

* **L0 (работает):** делает то, что заявлено.
* **L1 (корректность):** метод-гард, `Content-Type`, коды ошибок, базовая валидация.
* **L2 (безопасность/устойчивость):** лимиты тела (`MaxBytesReader`), «безопасное чтение», явные таймауты сервера, отмена по `Context`.
* **L3 (эффективность/наблюдаемость):** стриминг (`io.Copy`) где уместно, минимизация аллокаций, простой лог-мидлвар.

## 4 базовых мини-шаблона (постоянно будешь юзать)

**Q→T (query→text):**

```go
// GET /path?a=..&b=..
func handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet { http.Error(w, "method not allowed", 405); return }
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	q := r.URL.Query()
	// parse+validate q.Get("a"), q.Get("b")
	// write text result via fmt.Fprintf / w.Write
}
```

**J→J (json→json):**

```go
// POST /path  Content-Type: application/json
func handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { http.Error(w, "method not allowed", 405); return }
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20); defer r.Body.Close()

	dec := json.NewDecoder(r.Body); dec.DisallowUnknownFields()
	var in In
	if err := dec.Decode(&in); err != nil { http.Error(w, "bad json", 400); return }
	// validate in...
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Out{/* ... */})
}
```

**B→T (bytes→text, мелкие тела):**

```go
func handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { http.Error(w, "method not allowed", 405); return }
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20); defer r.Body.Close()

	b, err := io.ReadAll(r.Body); if err != nil { http.Error(w, "read error", 400); return }
	// transform b -> text result
}
```

**B→B (bytes→bytes, крупные/неизвестные):**

```go
func handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { http.Error(w, "method not allowed", 405); return }
	// optional: write headers for type
	_, err := io.Copy(w, io.LimitedReader{R: r.Body, N: 1<<20})
	if err != nil { http.Error(w, "stream error", 500) }
}
```

---

## Карта всех 50 задач (что проверять)

**Формат:** `№) PATH — METHOD — тип ввода → тип ответа — обязательные проверки/нюансы`

1. `/version` — GET — ∅ → text — 405 иначе, `text/plain`.
2. `/user` — GET — query `id` → text — id обязателен (не пуст).
3. `/copy` — POST — bytes → text — лимит тела, префикс `"Copy: "`.
4. `/health` — GET — ∅ → JSON — `{"service":"running"}`, `application/json`.
5. `/method` — ANY — ∅ → text — вернуть `r.Method`.
6. `/subtract` — GET — query `a,b` → text — оба ints, 400 при ошибке.
7. `/data` — POST — JSON `{name,age}` → JSON — валидация: name≠"", age≥0.
8. `/time` — GET — ∅ → text — формат `2006-01-02 15:04:05`.
9. `/square` — GET — query `n` → text — int, 400 при ошибке.
10. `/reverse` — POST — text → text — реверс строки, лимит тела.
11. `/random` — GET — ∅ → text — 1..100 (лучше `crypto/rand`).
12. `/greet` — GET — `name,lang` → text — lang∈{en,ru}, дефолты.
13. `/count` — POST — text → text — количество слов (split по пробелам).
14. `/multiply` — GET — `a,b` → text — ints, overflow не критично.
15. `/uppercase` — POST — text → text — `strings.ToUpper`.
16. `/length` — GET — `text` → text — длина в рунах или байтах? (укажи явно).
17. `/contains` — POST — JSON `{text,search}` → JSON `{found}` — substring check.
18. `/divide` — GET — `a,b` → text — b≠0, 400 при делении на ноль.
19. `/split` — POST — JSON `{text,separator}` → JSON `[]string` — `strings.Split`.
20. `/power` — GET — `base,exp` → text — целые, exp≥0.
21. `/max` — POST — JSON `{numbers:[]}` → JSON `{max}` — пустой массив → 400.
22. `/factorial` — GET — `n` → text — n≥0, ограничь n (например ≤20).
23. `/palindrome` — POST — text → JSON `{isPalindrome}` — нормализуй регистр/пробелы по решению.
24. `/fibonacci` — GET — `n` → text — n≥0, лимит n (например ≤92 для int64).
25. `/sort` — POST — JSON `[]number` → JSON отсортированный — стабильность не критична.
26. `/age` — GET — `year` → text — 0\<year≤текущий, иначе 400.
27. `/filter` — POST — JSON `{numbers,min}` → JSON `[]number` — `>= min`.
28. `/even` — GET — `n` → text — "true"/"false".
29. `/sum` — POST — JSON `[]number` → JSON `{sum}` — переполнение не критично.
30. `/prime` — GET — `n` → text — "true"/"false" (простой тест).
31. `/unique` — POST — JSON `[]string` → JSON уникальные — порядок: решить сам.
32. `/gcd` — GET — `a,b` → text — Евклид; a,b≥0.
33. `/anagram` — POST — JSON `{word1,word2}` → JSON `{areAnagram}` — нормализация регистра.
34. `/binary` — GET — `n` → text — `strconv.FormatInt`.
35. `/avg` — POST — JSON `[]number` → JSON `{avg}` — пустой массив → 400.
36. `/abs` — GET — `n` → text — модуль.
37. `/frequency` — POST — text → JSON `{char:count}` — считать по рунам (UTF-8) или байтам — выбери.
38. `/lcm` — GET — `a,b` → text — через gcd; 0 кейсы.
39. `/vowels` — POST — text → JSON `{count}` — набор гласных для выбранного языка(ов).
40. `/celsius` — GET — `f` → text — формула `(f-32)*5/9`, форматировать разумно.
41. `/replace` — POST — JSON `{text,old,new}` → JSON `{text}` — `strings.ReplaceAll`.
42. `/digit-sum` — GET — `n` → text — сумма цифр (учти знак).
43. `/trim` — POST — text → text — `strings.TrimSpace`.
44. `/leap` — GET — `year` → text — "true"/"false" по правилам кратности 4/100/400.
45. `/intersection` — POST — JSON `{arr1,arr2}` → JSON пересечение (без дублей).
46. `/roman` — GET — `n` → text — 1..3999, 400 при выходе.
47. `/median` — POST — JSON `[]number` → JSON `{median}` — чётная длина → среднее двух.
48. `/distance` — GET — `x1,y1,x2,y2` → text — корень из суммы квадратов.
49. `/base64` — POST — text → text — base64.StdEncoding.EncodeToString.
50. `/permutations` — GET — `text` → text — `len!` (учти повторяющиеся символы — оговори правило: считать как все перестановки без учёта повторов или с учётом?).

> Примечание на подумать (и зафиксировать в доке тренажёра):
> — где считать длину/частоты в **рунах**, а где в **байтах**;
> — числа: формат ответа — **plain text** или **JSON**? (я рекомендую в таком тренажёре text для простых вычислений, JSON для структурных ответов).

---

## Детализация L0→L3 для первых 10 (коротко, без кода)

**1 `/version` (Q→T):**
L0: вернуть `"v1.0.0"`. L1: только GET, `text/plain`. L2: таймауты сервера. L3: лог-мидлвар.

**2 `/user?id=` (Q→T):**
L0: `"User ID: X"`. L1: проверить `id` не пуст, 400 иначе. L2: лимита тела нет, но таймауты сервера. L3: лог.

**3 `/copy` (B→T):**
L0: прочитать тело, ответ `"Copy: <body>"`. L1: только POST, `text/plain`. L2: `MaxBytesReader` 1MB. L3: для больших тел — сначала пишем `"Copy: "`, затем `io.Copy` из тела.

**4 `/health` (∅→J):**
L0: JSON `{"service":"running"}`. L1: GET, `application/json`. L2: таймауты. L3: лог.

**5 `/method` (ANY→T):**
L0: вернуть `r.Method`. L1: `text/plain`. L2: —. L3: лог.

**6 `/subtract?a&b` (Q→T):**
L0: `a-b`. L1: 400 при нечислах/отсутствии, `text/plain`. L2: —. L3: лог.

**7 `/data` (J→J):**
L0: принять `{name,age}` → `{"message":"Hello ..."}`. L1: DisallowUnknownFields, валидация (name, age≥0). L2: `MaxBytesReader`. L3: лог.

**8 `/time` (∅→T):**
L0: текущее время по формату. L1: `text/plain`. L2: —. L3: лог.

**9 `/square?n` (Q→T):**
L0: `n*n`. L1: 400 если не число, `text/plain`. L2: —. L3: лог.

**10 `/reverse` (B→T):**
L0: развернуть строку. L1: POST, `text/plain`. L2: `MaxBytesReader`; если хочешь честный реверс по рунам — конвертируй. L3: для крупных тел делай потоковый реверс не так просто — можно оставить как «малые тела» для тренажёра.

---

## Как двигаемся

1. Возьми **1–5** и реализуй по L0→L1 (быстро).
2. Затем для них добавь L2 (лимиты/таймауты) — это 5–7 минут правки.
3. Я проверю чек-листами и дам точечные правки, потом перейдём к 6–10.

единственный вопрос сейчас: хочешь начать с **1–5 (L0→L2)** или с **6–10 (L0→L1)**? (выбери одно)
