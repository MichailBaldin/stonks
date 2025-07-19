package main

// 1. Интерфейс `Sorter.Sort([]int) []int`, стратегии — `AscSort`, `DescSort`.
//    Выведи отсортированный массив `[3, 1, 2]`.

// 2. Интерфейс `Formatter.Format(text string) string`.
//    Стратегия `UpperCaseFormatter` → `"hello"` → `"HELLO"`.

// 3. Интерфейс `Discount.Calc(price float64) float64`.
//    Стратегия `TenPercentOff` возвращает `price * 0.9`.

// 4. Интерфейс `Filter.Apply([]int) []int`.
//    Стратегия: оставить только чётные.
//    Ввод: `[1,2,3,4]` → Вывод: `[2 4]`.

// 5. Интерфейс `Validator.Validate(s string) bool`.
//    Стратегия `NotEmptyValidator`.
//    Проверь строку `""` → false, `"ok"` → true.

// 6. Интерфейс `Comparator.Compare(a, b int) bool`.
//    Стратегия: `GreaterThan`, проверка `3,2` → `true`.

// 7. Интерфейс `Encoder.Encode(text string) string`.
//    Стратегия `Reverser`, строка `"abc"` → `"cba"`.

// 8. Стратегия как `func(int, int) int`.
//    Внедри функцию `add`, передай `2, 3` → `5`.

// 9. Интерфейс `Renderer.Render(content string) string`.
//    Стратегия `HtmlRenderer` → `"Go"` → `"<p>Go</p>"`.

// 10. Стратегия `CSVExporter.Export([]string) string`.
//     Массив `["a", "b", "c"]` → `"a,b,c"`.

// 11. Интерфейс `Logger.Log(text string)`, стратегии: `ConsoleLogger`, `SilentLogger`.
//     `"msg"` печатается или игнорируется в зависимости от стратегии.

// 12. Интерфейс `Divider.Divide(a, b int) int`.
//     Стратегия `SafeDivide` → `a / b`, но если `b == 0`, возвращает `0`.

// 13. Интерфейс `Printer.Print(s string) string`, стратегия добавляет префикс `"LOG: "` к тексту.

// 14. Интерфейс `Hasher.Hash(data string) string`.
//     Стратегия: вернуть `"hashed:" + data`.

// 15. Стратегия как `func([]int) int` — вернуть максимум.
//     `[1, 5, 3]` → `5`.

// 16. Интерфейс `PayStrategy.Pay(amount float64) string`.
//     Стратегия: `"Paid %.2f with card"`.

// 17. Интерфейс `Joiner.Join(a, b string) string`.
//     Стратегия добавляет `"::"` между строками.

// 18. Интерфейс `Alarm.Trigger(code int) string`, стратегия возвращает `"Code: %d"`.

// 19. Интерфейс `Colorizer.Color(name string) string`, стратегия возвращает `"🔵 name"`.

// 20. Стратегия `func(string) bool`, имя → true, если длиннее 5 символов.
//     `"Ivan"` → `false`, `"Vladimir"` → `true`.
