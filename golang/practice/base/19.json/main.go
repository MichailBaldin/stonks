package main

// 1. Создай структуру `User{ID int, Name string}` и сериализуй в JSON.


// 2. Распарси JSON `{"id":1,"name":"Alice"}` в структуру `User`.


// 3. Сериализуй слайс `[]string{"one", "two"}` и выведи строку.


// 4. Распарси массив `["red", "green", "blue"]` в `[]string`.


// 5. Распарси JSON `{"ok":true,"count":3}` в `map[string]interface{}` и выведи типы.


// 6. Используй `omitempty` — поле `*string Bio` не должно попадать, если `nil`.


// 7. Пропусти поле `Secret` при сериализации (тег `json:"-"`).


// 8. Распарси вложенный JSON `{"title":"Go","author":{"name":"Rob"}}` 
// в структуру `Article{Title string, Author User}`.


// 9. Используй `json.NewEncoder(os.Stdout)` для потоковой записи структуры.


// 10. Прочитай JSON из `strings.NewReader` с помощью `json.NewDecoder`.


// 11. Напиши `UnmarshalJSON` для поля `type Date time.Time`, 
// которое принимает формат `"YYYY-MM-DD"`.


// 12. Сериализуй структуру с `bool` полем, которое в JSON должно быть `"yes"` / `"no"` (через `MarshalJSON`).


// 13. Используй `[]byte` как поле и сериализуй его 
// в base64 (по умолчанию Go так делает).


// 14. Прочитай JSON в `interface{}` и выведи `reflect.TypeOf(val["field"])`.


// 15. Сделай слайс `[]User`, сериализуй его в JSON-массив.


// 16. Распарси JSON-массив объектов в `[]User`.


// 17. Обработай ошибку: попытка распарсить `{bad json}` должна дать `json.SyntaxError`.


// 18. Сделай структуру с вложенной структурой-указателем, 
// проверь `omitempty` поведение при `nil`.


// 19. Реализуй REST-хендлер, который принимает POST 
// с `application/json`, распарсивает и возвращает `200 OK`.


// 20. Распарси JSON-массив с разными типами: `[1, "two", true, null]` 
// в `[]interface{}` и выведи значения.
