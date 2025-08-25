package main

// Задание 1: Сериализация структуры с простыми типами
// Условие:
// Реализуй функцию `ToJSON(x any) string`, которая:
// * принимает структуру
// * сериализует только `string`, `int`, `bool`
// * читает `json`-теги (если есть), иначе — имя поля
// * игнорирует неэкспортируемые поля
// * оборачивает всё в JSON-объект
// Пример:
// type User struct {
//     Name string `json:"name"`
//     Age  int    `json:"age"`
//     VIP  bool   `json:"vip"`
// }
// User{"Alice", 30, true} → {"name":"Alice","age":30,"vip":true}


// Задание 2: Добавь поддержку `float64`
// Условие:
// Теперь сериализуй также поля с типом `float32` и `float64`.
// Пример:
// type Product struct {
//     Title string
//     Price float64
// }
// → {"Title":"...", "Price":123.45}


// Задание 3: Обработка срезов простых типов
// Условие:
// Если в поле `[]string`, `[]int`, `[]float64`, сериализуй как JSON-массив:
// type Group struct {
//     Members []string `json:"members"`
// }
// → {"members":["Alice","Bob","Eve"]}


// Задание 4: Рекурсивная сериализация вложенных структур
// Условие:
// Если поле — вложенная структура, сериализуй её с помощью 
// **рекурсивного вызова** `ToJSON`.
// Пример:
// type Address struct {
//     City string
// }
// type User struct {
//     Name    string
//     Address Address
// }
// → {"Name":"Alice","Address":{"City":"Moscow"}}


// Задание 5: Обработка тегов `omitempty`
// Условие:
// Если у поля тег `json:"name,omitempty"` и 
// значение по нулю (`""`, `0`, `false`), то пропусти 
// это поле при сериализации.
// Пример:
// type User struct {
//     Name string `json:"name"`
//     Note string `json:"note,omitempty"`
// }
// User{Name: "A", Note: ""} → {"name":"A"}
