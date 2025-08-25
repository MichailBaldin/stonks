package main

// 1. Структура `Server{Host string, Port int}`.
//    `NewServer()` → host: "localhost", port: 80.
//    Опции: `WithHost(string)`, `WithPort(int)`.
//    Вывод: `"localhost 80"`, `"127.0.0.1 8080"`.

// 2. Структура `User{Name string, Active bool}`.
//    `WithName(name string)`, `WithActive()`.
//    Вывод: `"Bob true"`.

// 3. Структура `Job{Title string, Remote bool}`.
//    `WithRemote()` — делает `Remote=true`.
//    Вывод: `"Dev true"`.

// 4. Структура `Circle{Radius float64}`.
//    По умолчанию 1.0.
//    `WithRadius(r float64)` → `"2.5"`.

// 5. Структура `Book{Title string, Pages int}`.
//    `WithTitle(string)`, `WithPages(int)`.
//    Вывод: `"Go Basics 300"`.

// 6. Структура `Mail{From, To string}`.
//    Опции: `WithFrom(...)`, `WithTo(...)`.
//    Вывод: `"a@b.com → x@y.com"`.

// 7. Структура `Product{Name string, Price float64}`.
//    По умолчанию "Unknown", 0.0.
//    Вывод: `"Pen 9.99"`.

// 8. Структура `Logger{Level string}`.
//    `WithLevel("DEBUG")`. По умолчанию: `"INFO"`.
//    Вывод: `"DEBUG"`.

// 9. Структура `Account{ID int, Frozen bool}`.
//    `WithID(5)`, `WithFrozen()` — true.
//    Вывод: `"5 true"`.

// 10. Структура `Window{Width, Height int}`.
//     `WithSize(w,h int)` устанавливает оба.
//     Вывод: `"800x600"`.

// 11. Структура `Note{Title, Content string}`.
//     `WithTitle`, `WithContent`.
//     Вывод: `"Memo: ..."`.

// 12. Структура `Timer{Duration int}`.
//     По умолчанию 10.
//     `WithDuration(30)` → `"30"`.

// 13. Структура `Service{Retries int, Timeout int}`.
//     `WithRetries(3)`, `WithTimeout(1000)`
//     Вывод: `"3 retries, 1000ms"`.

// 14. Структура `Image{Format string}`.
//     `WithFormat("png")`. По умолчанию: `"jpg"`.
//     Вывод: `"Format: png"`.

// 15. Структура `Message{To, Text string}`.
//     `WithTo`, `WithText`.
//     Вывод: `"To: A, Msg: Hi"`.

// 16. Структура `Car{Model string, Electric bool}`.
//     `WithModel("Tesla")`, `WithElectric()` → true.
//     Вывод: `"Tesla EV"`.

// 17. Структура `Cache{Enabled bool}`.
//     По умолчанию: false.
//     `WithEnabled()` → `"Cache on"`.

// 18. Структура `Queue{Name string, Size int}`.
//     `WithName("Q1")`, `WithSize(100)`
//     Вывод: `"Q1[100]"`.

// 19. Структура `Session{Token string, Expired bool}`.
//     `WithToken("abc123")`, `WithExpired()`
//     Вывод: `"Token: abc123 (expired)"`.

// 20. Структура `Config{Debug bool, Mode string}`.
//     `WithDebug()`, `WithMode("safe")`
//     Вывод: `"Debug ON, Mode: safe"`.
