package main

// 1. Структура `User{Name string, Age int}`.
//    Создай `NewUser(name string) *User`, где `Age = 18`.
//    Вывод: `"Name: Ivan, Age: 18"`.

// 2. Структура `Point{X, Y int}`.
//    Конструктор `NewPoint(x, y int) *Point`.
//    Ввод: `NewPoint(2, 3)` → Вывод: `"Point: 2, 3"`.

// 3. Структура `Server{Host string, Port int}`.
//    Конструктор `NewServer(host string)` → `Port = 8080`.
//    Вывод: `"Server on localhost:8080"`.

// 4. Структура `Book{Title string, Pages int}`.
//    `NewBook(title string)` → `Pages = 100`.
//    Вывод: `"Book: Go Basics, 100 pages"`.

// 5. Структура `Rectangle{Width, Height int}`.
//    `NewSquare(size int) *Rectangle` — создаёт квадрат.
//    Вывод: `"Square 5x5"`.

// 6. Структура `Account{ID int, Balance float64}`.
//    `NewAccount(id int)` → `Balance = 0.0`.
//    Вывод: `"Account 7 with 0.00"`.

// 7. Структура `Message{To, Body string}`.
//    `NewMessage(to string)` → `Body = "Hello!"`.
//    Вывод: `"To: bob, Body: Hello!"`.

// 8. Структура `Car{Model string, Speed int}`.
//    `NewCar(model string)` → `Speed = 0`.
//    Вывод: `"Model: Tesla, Speed: 0"`.

// 9. Структура `Logger{Level string}`.
//    `NewLogger()` → `Level = "INFO"`.
//    Вывод: `"Logger level: INFO"`.

// 10. Структура `Sensor{ID int, Status string}`.
//     `NewSensor(id int)` → `Status = "idle"`.
//     Вывод: `"Sensor 3: idle"`.

// 11. Структура `Job{Title string, Salary int}`.
//     `NewJob(title string)` → `Salary = 50000`.
//     Вывод: `"Job: Dev, Salary: 50000"`.

// 12. Структура `Timer{Duration int}`.
//     `NewTimer()` → `Duration = 10`.
//     Вывод: `"Timer set for 10s"`.

// 13. Структура `Pair{A, B string}`.
//     `NewPair(a, b string)` возвращает пару.
//     Вывод: `"A: foo, B: bar"`.

// 14. Структура `Email{From, Subject string}`.
//     `NewEmail(from string)` → `Subject = "No subject"`.
//     Вывод: `"From: a@b.com, Subject: No subject"`.

// 15. Структура `Window{Width, Height int}`.
//     `NewWindow()` → 800x600.
//     Вывод: `"Window 800x600"`.

// 16. Структура `Config{Retry int}`.
//     `NewConfig()` → `Retry = 3`.
//     Вывод: `"Retries: 3"`.

// 17. Структура `Item{Name string, Price float64}`.
//     `NewItem(name string)` → `Price = 9.99`.
//     Вывод: `"Item: Pen, $9.99"`.

// 18. Структура `Session{Token string}`.
//     `NewSession()` → токен `"xyz"`.
//     Вывод: `"Token: xyz"`.

// 19. Структура `Circle{Radius float64}`.
//     `NewCircle(r float64)` возвращает круг.
//     Вывод: `"Radius: 2.5"`.

// 20. Структура `Note{Title, Text string}`.
//     `NewNote(title string)` → `Text = "..."`.
//     Вывод: `"Title: Memo, Text: ..."`.
