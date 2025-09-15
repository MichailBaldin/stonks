### 1. GET-запрос

**Условие:** Сделай GET-запрос на `https://httpbin.org/get` и выведи тело ответа.

```go
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
```

---

### 2. POST-запрос (JSON)

**Условие:** Отправь JSON `{"name":"Bob"}` на `https://httpbin.org/post` и выведи ответ.

```go
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	data := []byte(`{"name":"Bob"}`)

	resp, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
```

---

### 3. Query-параметры

**Условие:** Сделай GET-запрос с параметром `?id=123`.

```go
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	base, _ := url.Parse("https://httpbin.org/get")
	q := base.Query()
	q.Set("id", "123")
	base.RawQuery = q.Encode()

	resp, _ := http.Get(base.String())
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
```

---

### 4. Заголовки

**Условие:** Отправь GET-запрос с заголовком `User-Agent: GoClient`.

```go
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, _ := http.NewRequest("GET", "https://httpbin.org/headers", nil)
	req.Header.Set("User-Agent", "GoClient")

	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
```

---

### 5. Работа с JSON (Unmarshal)

**Условие:** Сделай GET-запрос и распарсь JSON-ответ в структуру.

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Origin string `json:"origin"`
	URL    string `json:"url"`
}

func main() {
	resp, _ := http.Get("https://httpbin.org/get")
	defer resp.Body.Close()

	var r Response
	json.NewDecoder(resp.Body).Decode(&r)

	fmt.Println(r.Origin, r.URL)
}
```

---

### 6. Таймаут и context

**Условие:** Сделай GET-запрос с таймаутом 2 секунды.

```go
package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, "GET", "https://httpbin.org/delay/3", nil)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Ошибка или таймаут:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
```
