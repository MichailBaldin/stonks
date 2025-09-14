обожаю такой формат. Ни решений — только чёткие, маленькие, «законченные» задачи (по 15–20 минут каждая). Каждая — мини-сервис/утилита с ясной целью, акцентом на структуры, интерфейсы и чистый код. Держим слои (handler → service → storage), ошибки — через `wrap %w`, тестируем curl’ом.

---

### 1) env-echo (config + HTTP + чистые слои)

**Цель:** развернуть минимальный сервис с конфигом и хендлерами.

* Структура: `Config{Addr string, AppName string}`, `Server{svc Service}`, `Service{}`.
* Endpoints:

  * `GET /health` → `{"ok":true,"app":"<AppName>"}`.
  * `GET /echo?msg=hi` → `{"msg":"hi","len":2}` (len по рунам).
* Требования: `Config` из `env` (дефолты), валидируй. Без глобалов. Логи — через интерфейс `Logger` (минимум `Println`).
* Done: `curl :8080/health` отдаёт JSON; `:8080/echo?msg=привет` корректно считает руны.

---

### 2) key-counter (sync.Map + сортировка)

**Цель:** потокобезопасные счётчики по ключам.

* API: `POST /inc?k=a` инкрементит; `GET /metrics` → отсортированные по убыванию значения: `[{k:"a", n:10}]`.
* Структуры: `CounterStore` с `sync.Map`; DTO `Metric{K string; N int}`.
* Детали: чтения частые, записи редкие — обоснование `sync.Map` в комментарии. Сбор в `[]Metric`, сортировка через `sort.Slice`.
* Done: конкурентные `ab -n 1000 -c 20` на `/inc` не ломают инварианты.

---

### 3) file-sink (bufio.Writer + os.OpenFile)

**Цель:** принять текст и буферизованно писать в файл.

* API: `POST /logs` (text/plain body) → `{"bytes": <n>, "sha256":"..."}`.
* Реализация: `os.OpenFile(..., O_CREATE|O_WRONLY|O_APPEND, 0644)` + `bufio.NewWriter`, `io.TeeReader` для подсчёта SHA-256 на лету.
* Чистый код: интерфейс `Sink{Write(p []byte) (int,error)}`, продакшн-реализация пишет в файл, фейковая — в память (для теста).
* Done: Flush обязательный. Повторные POST накапливают, хэш от присланного куска.

---

### 4) tiny-pipeline (generics + map/filter/reduce)

**Цель:** обобщённые трансформации чисел.

* API: `POST /compute` с JSON `{"nums":[1,2,3,4]}` → `{"even":[2,4],"sum":10}`.
* Реализация: `Filter[T any]`, `Reduce[T any,R any]`; `[]int` → чётные и сумма.
* Архитектура: чистая функция в слое `service`, handler — только (де)сериализация.
* Done: валидные ошибки 400 при битом JSON.

---

### 5) url-short (map + валидация + интерфейсы)

**Цель:** крошечный шортенер без внешней сети.

* API: `POST /shorten` `{"url":"https://example.com"}` → `{"code":"abc123"}`; `GET /r/{code}` → 302 Location.
* Реализация: интерфейс `Store{Save(code,url) error; Find(code)(string,bool)}`; память: `map[string]string` с `RWMutex`.
* Бизнес-логика: `code` = base62(монотонный int) или случайный с проверкой коллизий.
* Done: конкуррентные POST не создают дубликаты.

---

### 6) ttl-cache (map + тикер + очистка)

**Цель:** in-memory кэш с истечением.

* API: `PUT /cache?k=a&ttl=5s` (body — value bytes); `GET /cache?k=a` → value или 404; `GET /stats` → `{size:int}`.
* Реализация: `entry{val []byte; exp time.Time}`, `map[string]entry` + `RWMutex`; janitor-горутинка с `time.Ticker`.
* Чистый код: контекст-управляемая остановка janitor в `Server.Stop`.
* Done: по истечении TTL ключ исчезает; гонок под `-race` нет.

---

### 7) tail-N (io, bufio.Scanner, кольцевой буфер)

**Цель:** хранить последние `N` строк файла.

* CLI/HTTP на выбор: `GET /tail?n=10` → JSON массив последних строк.
* Реализация: открыть файл `os.Open`, читать сканером построчно и класть в `Ring[string]` (свой generic).
* Доп.: в фоне пересчитывать раз в `X` секунд (Ticker) — перезагрузить «хвост».
* Done: большие файлы не грузить целиком; закрытия/Stop корректны.

---

### 8) worker-pool (errgroup + SetLimit + контекст)

**Цель:** запустить `M` задач с параллелизмом `K`.

* API: `POST /jobs` `{"n":50,"k":5,"ms":100}` → `{"ok":true,"done":50}` (каждая задача «работает» `ms` миллисекунд).
* Реализация: `errgroup.WithContext`, `g.SetLimit(k)`, уважать `ctx.Done()`.
* Структуры: `JobRunner` интерфейс, реализация «заглушка».
* Done: при `ctx` с таймаутом часть задач отменяется, ответ — 504/408 с оборачиваемой ошибкой.

---

### 9) mini-library (структуры + каналы + инварианты)

**Цель:** доступ к «книгам» через командный канал.

* Домен: `Book{ID, Title}`, `Library{available map[int]Book; loans map[int]Book}`.
* Поток: отдельная горутина `LibraryLoop` принимает команды `Borrow/Return/Add` (через `chan Cmd`), отвечает в `reply chan`.
* HTTP: `POST /borrow?id=1`, `POST /return?id=1`, `GET /stats` → `{"avail":N,"loans":M}`.
* Done: нельзя «взять» уже выданную; инварианты защищены одним «владельцем» — библиотекой.

---

### 10) csv2json (io + encoding/csv + bytes.Buffer)

**Цель:** конвертер CSV → JSON строк объектов.

* API: `POST /convert` (multipart или raw CSV) → `application/json` массив `{header:value}`.
* Реализация: `csv.Reader`, первая строка — заголовки, последующие — данные. Буферизованная запись через `json.Encoder`/`bytes.Buffer`.
* Чистый код: интерфейс `Decoder`/`Encoder` для тестируемости; лимиты по размеру.
* Done: корректно обрабатывает пустые строки и разное число колонок (400 с пояснением).

---

### 11) text-normalize (strings/runes/bytes)

**Цель:** нормализатор текста.

* API: `POST /normalize` `{"s":"  Привет,\tмир!  "}` → `{"s":"Привет, мир!"}` (трим, collapse пробелов, стандартные кавычки).
* Детали: идти по рунам, `unicode`-классы, `strings.Builder` для сборки.
* Архитектура: `Normalizer` интерфейс, реализация Pure (чистая функция).
* Done: не ломает UTF-8; тест на не-ASCII.

---

### 12) safe-fetch (context + retry + backoff + errors)

**Цель:** безопасный вызов «внешнего» источника (замокай).

* API: `GET /price?id=BTC` → проксируй к `Fetcher` (интерфейс), в демо — имитация с вероятной ошибкой.
* Реализация: retry с экспоненциальным backoff (3 попытки), уважаем `ctx`.
* Ошибки: кастомный `Temporary`/`Retryable` тип; `errors.Is/As` в принятии решения ретраить.
* Done: при `ctx` cancel немедленно останавливается; логика — без глобалов, легко тестируется.

---

### 13) audit-middleware (интерфейсы + чистая архитектура)

**Цель:** middleware-обвязка для аудита.

* Реализация: `type Handler interface { ServeHTTP(w http.ResponseWriter,r *http.Request) }`, обёртка `Audit(next Handler, log Logger) Handler`.
* Логируй метод, путь, код ответа (обёртка над `ResponseWriter`), длительность.
* Встрой в задания 1/5/8 — без зависимости на конкретный логгер.
* Done: в тесте проверяй, что код/время пишутся.

---

### 14) snapshot-registry (map + копии + конкурентность)

**Цель:** реестр пользователей и «снимок».

* API: `POST /users` JSON `{"id":1,"name":"A"}`; `GET /users/snapshot` → `map[int]User` по значению (не указатели).
* Реализация: `Registry{mu RWMutex; m map[int]*User}`; метод `Snapshot() map[int]User` (deep copy значений).
* Done: параллельные записи не ломают снимок; тест — `-race` чист.

---

### 15) metrics-pull (ticker + context + errgroup)

**Цель:** периодический сбор метрик с «источников».

* Конфиг: список источников (строки). Каждые `T` секунд опрашивать все с `SetLimit(K)`, агрегировать.
* API: `GET /metrics` → текущие агрегаты; `POST /reload` перечитывает конфиг (в памяти).
* Реализация: `atomic.Value` для безопасной публикации текущего снимка.
* Done: остановка по `ctx`; нет утечек горутин.

---

#### Как работать

* Для каждого задания: распиши небольшие интерфейсы, 2–3 структуры, один сервис-слой, один хендлер, один storage (in-memory).
* Правила чистого кода: короткие функции, явные ошибки `fmt.Errorf("%s: %w", op, err)`, не экспортируй лишнее, зависимости через интерфейсы, без глобалов.
* Мини-чек: `go test -race` (если добавишь тесты), пара `curl` примеров, graceful shutdown по `SIGINT`.
