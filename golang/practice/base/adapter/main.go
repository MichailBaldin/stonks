package main

// 1. Есть `type LegacyPrinter struct{}` с методом `Print(msg string)`.
//    Напиши адаптер под интерфейс `type Printer interface { PrintText(s string) }` и вызови `PrintText("Hello")`.

// 2. У тебя есть `AudioPlayer.PlayMP3(file string)` и интерфейс `MediaPlayer.Play(file string)`.
//    Напиши адаптер, который вызывает MP3-плеер через `MediaPlayer`.

// 3. Интерфейс `type Logger interface { Log(s string) }` и тип `Console struct{}` с методом `WriteLine(text string)`.
//    Создай адаптер и выведи `Log("test")`.

// 4. Адаптируй тип `File struct{}` с методом `SaveToDisk(name string)` к интерфейсу `type Saver interface { Save(name string) }`.

// 5. Интерфейс `Notificator.Notify(msg string)` и структура `SlackAPI.Send(message string)`.
//    Напиши адаптер `SlackAdapter`.

// 6. Интерфейс `type Counter interface { Inc() int }` и структура `OldCounter struct { Count int }` с методом `AddOne() int`.
//    Сделай адаптер.

// 7. Интерфейс `Database.Get(id int) string` и структура `OldStorage.FetchByID(i int) string`.
//    Адаптируй `OldStorage`.

// 8. Интерфейс `Shape.Draw()` и структура `Square.Render()`.
//    Создай адаптер и выведи `"Square drawn"`.

// 9. Интерфейс `Greeter.Greet(name string)` и структура `HelloService.Hello(n string)` — адаптируй.

// 10. Интерфейс `type BoolParser interface { ParseBool(input string) bool }` и структура `LegacyBool.Decide(str string) bool`.

// 11. Интерфейс `Alarm.Trigger(level int)` и структура `LegacyAlert.Warn(severity int)` — вызов с выводом `Triggered`.

// 12. Интерфейс `type Formatter interface { FormatText(t string) string }` и структура `TextBox.RenderText(s string) string`.

// 13. Интерфейс `type Reader interface { ReadLine() string }` и структура `OldReader.Read() string`.

// 14. Интерфейс `Encoder.Encode(data []byte) string`, структура `OldBase64.DoEncode(input []byte) string`.

// 15. Интерфейс `TimeProvider.Now() time.Time`, структура `LegacyClock.GetCurrentTime() time.Time`.

// 16. Интерфейс `type Validator interface { IsValid(s string) bool }`, структура `OldValidator.Validate(v string) bool`.

// 17. Интерфейс `Calculator.Add(x, y int) int`, структура `LegacyCalc.Sum(a, b int) int`.

// 18. Интерфейс `type Sender interface { Send(to string, msg string) }`, структура `LegacySMTP.SendMail(addr, content string)`.

// 19. Интерфейс `Progressor.Step()`, структура `LegacyStep.Increment()` — адаптер с выводом `Step done`.

// 20. Интерфейс `Storage.Put(key string, value string)`, структура `DiskStore.Save(k, v string)`.
