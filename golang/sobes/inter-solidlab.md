**Ключевые акценты для подготовки:**

## 🎯 Обязательно знать:

**OWASP Top 10 с примерами на Go:**

- SQL Injection + параметризованные запросы
- XSS + HTML escaping
- CSRF + SameSite cookies
- Broken Authentication + bcrypt hashing

**DAST специфика:**

- Активное vs пассивное сканирование
- Web crawling и fuzzing
- Payload management
- False positive анализ

## 🚀 Что покажет экспертизу:

**Практические навыки:**

- Умение найти уязвимость в Go коде за 30 секунд
- Знание secure coding practices
- Понимание архитектуры веб-сканеров

**Системное мышление:**

- Как интегрировать DAST в CI/CD
- Безопасность самого сканера
- Мониторинг и alerting

## 💡 Лайфхак для собеседования:

Подготовьте **конкретный пример** уязвимости, которую вы находили или исправляли. Это покажет практический опыт и выделит вас среди других кандидатов.

**Пример структуры ответа:**

1. "Какую уязвимость нашли"
2. "Как её эксплуатировали"
3. "Как исправили в Go коде"
4. "Как предотвратили в будущем"

# Безопасность для DAST Go разработчика

## 🎯 OWASP Top 10 - основа DAST сканирования

### 1. **Injection (Инъекции)**

**SQL Injection:**

```sql
-- Уязвимый запрос
SELECT * FROM users WHERE id = '$user_input'

-- Эксплойт
' OR 1=1--

-- Безопасный подход в Go
db.Query("SELECT * FROM users WHERE id = $1", userID)
```

**Command Injection:**

```go
// Уязвимо
cmd := exec.Command("ping", userInput)

// Безопасно
if !regexp.MustMatch(`^[a-zA-Z0-9.-]+$`, userInput) {
    return errors.New("invalid input")
}
```

**Как DAST это находит:**

- Отправка специальных payload'ов
- Анализ ответов на наличие SQL ошибок
- Timing-based атаки

### 2. **Broken Authentication (Нарушенная аутентификация)**

```go
// Проблемы:
- Слабые пароли
- Session hijacking
- Brute force атаки
- Неправильное хранение сессий

// Безопасная реализация JWT в Go
func generateJWT(userID string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })
    return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
```

### 3. **Sensitive Data Exposure (Утечка чувствительных данных)**

```go
// Проблемы:
- Передача паролей в открытом виде
- Логирование sensitive данных
- Отсутствие HTTPS

// Безопасное хеширование паролей
import "golang.org/x/crypto/bcrypt"

func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}
```

### 4. **XML External Entities (XXE)**

```go
// Уязвимый XML парсинг
decoder := xml.NewDecoder(req.Body)

// Безопасный подход
decoder := xml.NewDecoder(req.Body)
decoder.Strict = false
decoder.Entity = make(map[string]string) // Отключение внешних сущностей
```

### 5. **Broken Access Control**

```go
// Проблема: отсутствие проверки прав доступа
func getUser(w http.ResponseWriter, r *http.Request) {
    userID := r.URL.Query().Get("id")
    // Любой может получить данные любого пользователя
    user := getUserFromDB(userID)
    json.NewEncoder(w).Encode(user)
}

// Правильно: проверка авторизации
func getUser(w http.ResponseWriter, r *http.Request) {
    currentUser := getCurrentUser(r)
    requestedUserID := r.URL.Query().Get("id")

    if currentUser.ID != requestedUserID && !currentUser.IsAdmin {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }
    // ...
}
```

---

## 🔍 Как работает DAST сканирование

### Активное сканирование

```go
// Пример простого SQL injection теста
func testSQLInjection(url string, param string) {
    payloads := []string{
        "' OR 1=1--",
        "'; DROP TABLE users--",
        "' UNION SELECT null,null,null--",
    }

    for _, payload := range payloads {
        resp := sendRequest(url + "?" + param + "=" + payload)
        if containsSQLError(resp.Body) {
            reportVulnerability("SQL Injection", url, param)
        }
    }
}
```

### Пассивное сканирование

```go
// Анализ ответов без отправки атакующих запросов
func analyzeResponse(resp *http.Response) {
    body := readBody(resp)

    // Поиск sensitive информации
    if containsCreditCard(body) {
        reportIssue("Credit card exposed")
    }

    // Проверка заголовков безопасности
    if resp.Header.Get("X-Frame-Options") == "" {
        reportIssue("Missing X-Frame-Options header")
    }
}
```

### Crawling и fuzzing

```go
// Поиск скрытых эндпоинтов
func fuzzDirectories(baseURL string) {
    commonDirs := []string{"admin", "api", "backup", "config", "test"}

    for _, dir := range commonDirs {
        url := baseURL + "/" + dir
        resp := sendRequest(url)
        if resp.StatusCode != 404 {
            addToScanQueue(url)
        }
    }
}
```

---

## 🛡️ Безопасность Go приложений

### 1. **Input Validation**

```go
import (
    "html"
    "net/url"
    "regexp"
)

func sanitizeInput(input string) string {
    // HTML escaping
    escaped := html.EscapeString(input)

    // URL encoding
    encoded := url.QueryEscape(escaped)

    return encoded
}

func validateEmail(email string) bool {
    pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
    matched, _ := regexp.MatchString(pattern, email)
    return matched
}
```

### 2. **Secure HTTP Headers**

```go
func securityMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // CSRF protection
        w.Header().Set("X-Frame-Options", "DENY")
        w.Header().Set("X-Content-Type-Options", "nosniff")
        w.Header().Set("X-XSS-Protection", "1; mode=block")
        w.Header().Set("Strict-Transport-Security", "max-age=31536000")
        w.Header().Set("Content-Security-Policy", "default-src 'self'")

        next.ServeHTTP(w, r)
    })
}
```

### 3. **Rate Limiting**

```go
import "golang.org/x/time/rate"

type RateLimiter struct {
    limiter *rate.Limiter
}

func (rl *RateLimiter) Allow() bool {
    return rl.limiter.Allow()
}

func rateLimitMiddleware(next http.Handler) http.Handler {
    limiter := rate.NewLimiter(10, 100) // 10 requests per second, burst of 100

    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if !limiter.Allow() {
            http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
            return
        }
        next.ServeHTTP(w, r)
    })
}
```

### 4. **Secure Session Management**

```go
import (
    "crypto/rand"
    "encoding/base64"
    "net/http"
    "time"
)

func generateSessionID() string {
    bytes := make([]byte, 32)
    rand.Read(bytes)
    return base64.URLEncoding.EncodeToString(bytes)
}

func setSecureCookie(w http.ResponseWriter, name, value string) {
    cookie := &http.Cookie{
        Name:     name,
        Value:    value,
        HttpOnly: true,  // Защита от XSS
        Secure:   true,  // Только HTTPS
        SameSite: http.SameSiteStrictMode, // CSRF protection
        MaxAge:   3600,  // 1 час
    }
    http.SetCookie(w, cookie)
}
```

---

## 🌐 Сетевая безопасность

### TLS/HTTPS в Go

```go
// Настройка безопасного TLS
func createSecureTLSConfig() *tls.Config {
    return &tls.Config{
        MinVersion:   tls.VersionTLS12,
        CipherSuites: []uint16{
            tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
            tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
        },
        PreferServerCipherSuites: true,
    }
}

// HTTP клиент с проверкой сертификатов
func createSecureHTTPClient() *http.Client {
    return &http.Client{
        Timeout: 30 * time.Second,
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{
                InsecureSkipVerify: false, // Всегда проверяем сертификаты
            },
        },
    }
}
```

### Защита от SSRF

```go
import "net"

func isPrivateIP(ip string) bool {
    private := []string{
        "10.0.0.0/8",
        "172.16.0.0/12",
        "192.168.0.0/16",
        "127.0.0.0/8",
    }

    addr := net.ParseIP(ip)
    for _, block := range private {
        _, subnet, _ := net.ParseCIDR(block)
        if subnet.Contains(addr) {
            return true
        }
    }
    return false
}

func validateURL(targetURL string) error {
    u, err := url.Parse(targetURL)
    if err != nil {
        return err
    }

    ips, err := net.LookupIP(u.Hostname())
    if err != nil {
        return err
    }

    for _, ip := range ips {
        if isPrivateIP(ip.String()) {
            return errors.New("private IP not allowed")
        }
    }
    return nil
}
```

---

## 🔐 Криптография в Go

### Хеширование и подписи

```go
import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
)

func calculateHMAC(message, secret string) string {
    key := []byte(secret)
    h := hmac.New(sha256.New, key)
    h.Write([]byte(message))
    return hex.EncodeToString(h.Sum(nil))
}

func verifyHMAC(message, messageMAC, secret string) bool {
    expectedMAC := calculateHMAC(message, secret)
    return hmac.Equal([]byte(messageMAC), []byte(expectedMAC))
}
```

### Шифрование данных

```go
import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "io"
)

func encrypt(plaintext []byte, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }

    nonce := make([]byte, gcm.NonceSize())
    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, err
    }

    ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
    return ciphertext, nil
}
```

---

## 📝 Возможные вопросы на собеседовании

### Скрининг (быстрые вопросы)

**1. OWASP Top 10:**

- "Назовите топ-3 веб-уязвимости из OWASP Top 10"
- "Чем отличается SQL injection от XSS?"
- "Что такое CSRF и как его предотвратить?"

**2. DAST vs другие виды тестирования:**

- "В чем разница между DAST и SAST?"
- "Какие уязвимости может найти DAST, а какие нет?"
- "Когда лучше использовать DAST, а когда penetration testing?"

**3. Безопасность Go:**

- "Как безопасно хранить пароли в Go?"
- "Что такое timing attack и как его избежать?"
- "Как защититься от directory traversal в Go?"

### Техническое собеседование

**1. Практические задачи:**

```go
// Задача 1: Найти уязвимость в коде
func login(w http.ResponseWriter, r *http.Request) {
    username := r.FormValue("username")
    password := r.FormValue("password")

    query := "SELECT id FROM users WHERE username='" + username + "' AND password='" + password + "'"
    rows, _ := db.Query(query)
    // ... что здесь не так?
}

// Задача 2: Реализовать безопасную аутентификацию
func secureLogin(w http.ResponseWriter, r *http.Request) {
    // Ваша реализация
}
```

**2. Архитектурные вопросы:**

- "Как бы вы спроектировали DAST сканер?"
- "Какие компоненты нужны для веб-сканера уязвимостей?"
- "Как обеспечить безопасность самого сканера?"

**3. Системные вопросы:**

- "Как мониторить безопасность Go приложения?"
- "Какие метрики безопасности важно отслеживать?"
- "Как интегрировать DAST в CI/CD pipeline?"

---

## 🎯 Специфичные для DAST темы

### 1. **Web Crawling Security**

```go
// Безопасный веб-краулер
type SecureCrawler struct {
    client     *http.Client
    rateLimiter *rate.Limiter
    userAgent   string
}

func (c *SecureCrawler) Crawl(url string) error {
    // Проверка URL
    if err := validateURL(url); err != nil {
        return err
    }

    // Rate limiting
    if !c.rateLimiter.Allow() {
        return errors.New("rate limit exceeded")
    }

    // Безопасный запрос
    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("User-Agent", c.userAgent)

    resp, err := c.client.Do(req)
    // ...
}
```

### 2. **Payload Management**

```go
// Управление тестовыми payload'ами
type PayloadManager struct {
    sqlInjection []string
    xss          []string
    xxe          []string
}

func (pm *PayloadManager) GetPayloads(vulnType string) []string {
    switch vulnType {
    case "sql":
        return pm.sqlInjection
    case "xss":
        return pm.xss
    case "xxe":
        return pm.xxe
    default:
        return nil
    }
}
```

### 3. **Result Analysis**

```go
// Анализ результатов сканирования
type VulnerabilityScanner struct{}

func (vs *VulnerabilityScanner) AnalyzeResponse(resp *http.Response, payload string) *Finding {
    body := readResponseBody(resp)

    // SQL injection detection
    sqlErrors := []string{
        "SQL syntax error",
        "mysql_fetch_array",
        "ORA-01756",
    }

    for _, errorPattern := range sqlErrors {
        if strings.Contains(body, errorPattern) {
            return &Finding{
                Type:     "SQL Injection",
                Severity: "High",
                Payload:  payload,
                Evidence: errorPattern,
            }
        }
    }

    return nil
}
```

---

## 🚨 Red Flags в коде

### Что искать при код-ревью:

```go
// ❌ Плохо
password := r.FormValue("password")
log.Printf("User login attempt: %s with password: %s", username, password)

// ✅ Хорошо
log.Printf("User login attempt: %s", username)

// ❌ Плохо
query := "SELECT * FROM users WHERE id = " + userID

// ✅ Хорошо
query := "SELECT * FROM users WHERE id = $1"
rows, err := db.Query(query, userID)

// ❌ Плохо
if user.Password == providedPassword {
    // timing attack возможен
}

// ✅ Хорошо
if subtle.ConstantTimeCompare([]byte(user.Password), []byte(providedPassword)) == 1 {
    // защита от timing attack
}
```

---

## 📚 Ресурсы для изучения

### Документация и стандарты:

- **OWASP Testing Guide** - методология тестирования
- **NIST Cybersecurity Framework** - стандарты безопасности
- **Go Security Guide** - best practices для Go

### Инструменты для практики:

- **DVWA** (Damn Vulnerable Web App) - для изучения уязвимостей
- **WebGoat** - OWASP проект для обучения
- **Burp Suite** - для анализа веб-приложений

### Полезные Go пакеты:

```go
// Безопасность
import "golang.org/x/crypto/bcrypt"
import "golang.org/x/crypto/scrypt"
import "crypto/subtle"

// Валидация
import "github.com/go-playground/validator/v10"
import "html"
import "net/url"

// Rate limiting
import "golang.org/x/time/rate"
```

---

## 💡 Ключевые takeaways для собеседования

1. **Знайте OWASP Top 10** наизусть с примерами
2. **Понимайте разницу** между DAST, SAST, IAST
3. **Умейте находить уязвимости** в Go коде
4. **Знайте как защищаться** от основных атак
5. **Понимайте архитектуру** веб-сканеров
6. **Практикуйте coding** - безопасные функции на Go

**Главное:** покажите, что вы не просто знаете теорию, но понимаете практическое применение в контексте разработки DAST сканера! 🎯
