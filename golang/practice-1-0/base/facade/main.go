package main

// 1. Создай `VideoPlayer` с методами `Load()`, `Play()`.
//    Фасад `MovieApp.Start()` вызывает оба. Вывод: `"Loading video"`, `"Playing video"`.

// 2. Структуры `Engine.Start()`, `Gear.SetDrive()`, фасад `Car.Drive()`.
//    Вывод: `"Engine started"`, `"Drive mode set"`.

// 3. Структуры `Logger.Init()`, `Database.Connect()`, фасад `App.Start()`.
//    Вывод: `"Logger ready"`, `"Database ready"`.

// 4. Структуры `Auth.Login(user string)`, `Session.Create()`.
//    Фасад `AuthService.SignIn(user)` вызывает оба. Вывод: `"Logged in: Alice"`, `"Session created"`.

// 5. Структуры `File.Open(name string)`, `File.Read()`, фасад `FileReader.ReadFile(name)` вызывает оба.
//    Вывод: `"Opened test.txt"`, `"Reading contents"`.

// 6. Структуры `Image.Load()`, `Image.Render()`, фасад `Viewer.Show()`.
//    Вывод: `"Image loaded"`, `"Rendered"`.

// 7. Структуры `Network.Dial()`, `Network.Handshake()`, фасад `Client.Connect()`.
//    Вывод: `"Dialing..."`, `"Handshake done"`.

// 8. Структуры `CPU.Run()`, `Memory.Allocate()`, фасад `Computer.Boot()`.
//    Вывод: `"CPU running"`, `"Memory allocated"`.

// 9. Структуры `Music.LoadTrack(name string)`, `Music.Play()`, фасад `Player.PlayTrack(name)` вызывает оба.
//    Вывод: `"Track loaded: rock"`, `"Playing music"`.

// 10. Структуры `Cache.Clear()`, `Disk.Clean()`, фасад `Cleaner.FullClean()`.
//     Вывод: `"Cache cleared"`, `"Disk cleaned"`.

// 11. Структуры `Backup.Prepare()`, `Backup.WriteToDisk()`, фасад `BackupService.Backup()`.
//     Вывод: `"Preparing backup"`, `"Backup complete"`.

// 12. Структуры `Compiler.Analyze()`, `Compiler.EmitBinary()`, фасад `Compiler.Compile()`.
//     Вывод: `"Analyzing..."`, `"Binary ready"`.

// 13. Структуры `Account.Load(id int)`, `Account.CheckStatus()`, фасад `AccountFacade.Open(id)`.
//     Вывод: `"Loaded account 5"`, `"Account active"`.

// 14. Структуры `Sensor.CheckTemp()`, `Sensor.CheckPressure()`, фасад `Monitor.AllStatus()`.
//     Вывод: `"Temp OK"`, `"Pressure OK"`.

// 15. Структуры `Mail.Connect()`, `Mail.Send(to string)`, фасад `Mailer.SendTo(to)` вызывает оба.
//     Вывод: `"Connected"`, `"Sent to bob@example.com"`.

// 16. Структуры `Router.LoadRoutes()`, `Server.Listen()`, фасад `WebApp.Launch()`.
//     Вывод: `"Routes loaded"`, `"Listening..."`.

// 17. Структуры `Zip.Pack()`, `Zip.Encrypt()`, фасад `Archive.CreateSecureZip()`.
//     Вывод: `"Packing..."`, `"Encrypted"`.

// 18. Структуры `Chat.Open()`, `Chat.Send(msg string)`, фасад `Messenger.StartChat(msg)`.
//     Вывод: `"Chat opened"`, `"Sent: hello"`.

// 19. Структуры `Camera.Init()`, `Camera.Capture()`, фасад `Photographer.Snap()`.
//     Вывод: `"Camera ready"`, `"Captured"`.

// 20. Структуры `Form.Fill(data string)`, `Form.Submit()`, фасад `UI.SubmitForm(data)`.
//     Вывод: `"Form filled: X"`, `"Form submitted"`.
