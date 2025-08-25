package main

// 1. Выполни `GET` на `https://httpbin.org/get` и выведи тело ответа.


// 2. Сделай `POST` с `application/json` телом: `{"name": "Go"}` — на 
// `https://httpbin.org/post`.


// 3. Распарси JSON-ответ `httpbin` и выведи поле `"json.name"`.


// 4. Сделай `GET` с заголовком `X-Trace-Id: 12345`.


// 5. Проверь `StatusCode` ответа, если не `200 OK` — логируй ошибку.


// 6. Отправь `POST`-запрос с телом `map[string]int{"a": 1, "b": 2}` 
// и получи обратно JSON с подтверждением.


// 7. Установи `http.Client{Timeout: 1 * time.Second}` 
// и протестируй с `https://httpbin.org/delay/3`.


// 8. Выполни 3 попытки `GET` к `http://localhost:9999`, 
// с паузой между ошибками (retry logic).


// 9. Отправь `GET` с `User-Agent: MyBot/1.0` и проверь, что он отражается в ответе.


// 10. Используй `context.WithTimeout` для `http.NewRequest`,
// чтобы отменить долгий запрос.


// 11. Напиши функцию `DoJSONRequest(method, url string, data any) (*http.Response, error)`, 
// которая маршалит `data` в JSON и делает запрос.


// 12. Напиши обёртку `FetchJSON(url string, out interface{}) error` — получает 
// и распарсивает JSON.


// 13. Напиши клиент `PingServer(url string) bool`, который возвращает `true`, 
// если `GET` даёт `200`.


// 14. Напиши запрос `GET /msg/1`, получи JSON и распарсь в 
// `type Message struct{ ID int; Text string }`.


// 15. Протестируй таймаут соединения с поддельным IP (например, `http://10.255.255.1`) 
// и обработай как `net.Error`.


// 16. Используй `http.DefaultTransport.(*http.Transport).MaxIdleConns = 100` — настрой 
// клиент с пулом соединений.


// 17. Напиши `POST /multiply`, отправь `a=2, b=4`, 
// и проверь ответ `{"result": 8}`.


// 18. Используй `http.NewRequestWithContext` и отключай контекст 
// при превышении лимита времени.


// 19. Добавь поддержку `Bearer Token` в заголовках.


// 20. Организуй несколько параллельных `GET`-запросов в `[]string{...}` 
// и собери все ответы в `[]string` с использованием `sync.WaitGroup`.

