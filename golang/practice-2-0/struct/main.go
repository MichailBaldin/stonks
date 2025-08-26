package main

// 1. Храни температуру по дням недели. Выбери структуру данных, чтобы:
//   - быстро получить температуру по названию дня ("Mon", "Tue"...),
//   - уметь перебрать дни по порядку.
//     Добавь 3 дня и выведи "Wed".
// type DayTemp struct {
// 	Date  time.Time
// 	Value float64
// }

// type Forecast struct {
// 	order []time.Weekday
// 	temps map[time.Weekday]DayTemp
// }

// func NewForecast() *Forecast {
// 	return &Forecast{
// 		order: []time.Weekday{},
// 		temps: make(map[time.Weekday]DayTemp),
// 	}
// }

// func (f *Forecast) Add(temp DayTemp) {
// 	wd := temp.Date.Weekday()
// 	f.temps[wd] = temp
// 	for _, d := range f.order {
// 		if d == wd {
// 			return
// 		}
// 	}
// 	f.order = append(f.order, wd)
// }

// func main() {
// 	f := NewForecast()

// 	f.Add(DayTemp{Date: time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC), Value: 26.1})
// 	f.Add(DayTemp{Date: time.Date(2024, 8, 2, 0, 0, 0, 0, time.UTC), Value: 27.1})
// 	f.Add(DayTemp{Date: time.Date(2024, 8, 3, 0, 0, 0, 0, time.UTC), Value: 21.5})

// 	if t, ok := f.temps[time.Saturday]; ok {
// 		fmt.Printf("Quick access: %s -> %.1f°C\n", t.Date.Weekday(), t.Value)
// 	}

// 	for _, d := range f.order {
// 		t := f.temps[d]
// 		fmt.Printf("Ordered: %s -> %.1f°C\n", d, t.Value)
// 	}
// }

//  2. Модель "Заказ" в магазине: id, список товаров (название + количество), статус (string).
//     Создай заказ с двумя позициями и измени статус на "shipped".
// type Status string

// const (
// 	StatusCreated Status = "created"
// 	StatusShipped Status = "shipped"
// )

// type Order struct {
// 	ID       int
// 	Products []*Product
// 	Status   Status
// }

// func NewOrder(id int, status Status, products ...*Product) *Order {
// 	return &Order{
// 		ID:       id,
// 		Status:   StatusCreated,
// 		Products: products,
// 	}
// }

// type Product struct {
// 	Name  string
// 	Count int
// }

// func NewProduct(name string, count int) *Product {
// 	return &Product{
// 		Name:  name,
// 		Count: count,
// 	}
// }

// func main() {
// 	apples := NewProduct("apple", 10)
// 	bananas := NewProduct("banana", 50)

// 	order1 := NewOrder(1, "created", apples, bananas)

// 	order1.Status = StatusShipped
// }

// 3. Частотный словарь: по слайсу слов подсчитай, сколько раз встречается каждое слово.
//    Выведи топ-1 по частоте (без сортировки — просто найди максимум).

// type Dictionary struct {
// 	freq map[string]int
// }

// func NewDictionary() *Dictionary {
// 	return &Dictionary{
// 		make(map[string]int),
// 	}
// }

// func (d *Dictionary) AddSlice(sl []string) {
// 	for _, v := range sl {
// 		d.freq[v]++
// 	}
// }

// func (d *Dictionary) FindMax() string {

// 	keyMax, valMax := "", 0
// 	for k, v := range d.freq {
// 		if v > valMax {
// 			keyMax = k
// 			valMax = v
// 		}
// 	}
// 	return keyMax
// }

// func main() {
// 	d := NewDictionary()
// 	sl := []string{"программа", "программа", "код", "код", "код", "данные", "данные",
// 		"система", "файл", "файл", "функция", "функция", "переменная",
// 		"массив", "слайс", "слайс", "структура", "интерфейс", "память",
// 		"сервер", "клиент", "база", "данных", "запрос", "ответ", "сеть"}

// 	d.AddSlice(sl)
// 	fmt.Println(d.FindMax()) // код
// }

//  4. Контакты: у пользователя может быть несколько телефонов разных типов ("home", "work").
//     Спроектируй структуры, добавь 2 типа телефонов и выведи "work".
// type PhoneType string

// const (
// 	Home PhoneType = "home"
// 	Work PhoneType = "work"
// )

// type User struct {
// 	Name    string
// 	Numbers map[PhoneType][]string
// }

// func NewUser(name string) *User {
// 	return &User{
// 		Name:    name,
// 		Numbers: make(map[PhoneType][]string),
// 	}
// }

// func (u *User) AddNumber(kind PhoneType, num string) {
// 	u.Numbers[kind] = append(u.Numbers[kind], num)
// }

// type Contacts struct {
// 	Users map[string]*User
// }

// func NewContacts() *Contacts {
// 	return &Contacts{
// 		Users: make(map[string]*User),
// 	}
// }

// func (c *Contacts) AddUser(u *User) {
// 	c.Users[u.Name] = u
// }

// func main() {
// 	mick := NewUser("Mick")
// 	mick.AddNumber(Work, "+7-925-359-04-46")
// 	mick.AddNumber(Home, "+7-916-123-66-54")

// 	tanya := NewUser("Tanya")
// 	tanya.AddNumber(Work, "+7-986-732-55-11")
// 	tanya.AddNumber(Home, "+7-988-567-99-00")

// 	book := NewContacts()
// 	book.AddUser(mick)
// 	book.AddUser(tanya)

// 	for name, u := range book.Users {
// 		for _, num := range u.Numbers[Work] {
// 			fmt.Printf("Name: %s, Status: %s, Number: %s\n", name, Work, num)
// 		}
// 	}
// }

//  5. Кэш конфигов по имени сервиса: быстрый доступ по ключу, плюс счетчик попаданий.
//     Инициализируй кэш для "auth", увеличь счетчик 3 раза, выведи значение.

// 6. Список задач (todo) с приоритетом (int) и метками ([]string).
//    Добавь 3 задачи, отфильтруй по метке "urgent".

// 7. Каталог книг: хранить книги по жанру. У книги: title, author.
//    Добавь 2 жанра и по 2 книги, выведи все книги жанра "Sci-Fi".

// 8. Учет посещений URL: для каждого домена — список путей, отсортированный по времени добавления.
//    Добавь 3 записи для "example.com" и выведи последний путь.

// 9. Гео-координаты сенсоров: id -> (lat, lon). Нужен быстрый доступ по id и обновление.
//    Добавь 2 сенсора, обнови координаты первого, выведи их.

// 10. Модель чата: каналы, в каждом сообщения (author, text, timestamp int64).
//     Создай 2 канала, добавь сообщения, выведи все авторы канала "#general".

// 11. Рейтинг игроков: хранить очки и уметь быстро повышать/снижать.
//     Добавь 3 игрока, повысь очки двоих, выведи лидера (по максимальным очкам).

// 12. Множество уникальных тегов (set строк). Добавь теги, проверь принадлежность "go".
//     Объясни в комментарии выбор структуры для множества.

// 13. Инвентарь склада: позиция имеет SKU (string), количество (int) и историю событий ([]string).
//     Добавь позицию, проведи "in +10", "out -3", выведи текущее кол-во и историю.

// 14. Индекс по двум ключам: (год, месяц) -> список сумм продаж (int).
//     Добавь данные за 2024/06 и 2024/07, посчитай сумму за 2024/06.

// 15. Настройки приложения с fallback по окружению: env->key->value.
//     Добавь "prod.DB_URL", "dev.DB_URL", выведи значение для "prod" и "DB_URL".

// 16. Очередь задач на обработку (FIFO). Тип задачи: name, payload(map[string]string).
//     Добавь 3 задачи, снимай по одной и печатай name.

// 17. Таблица смежности графа дорог: город -> соседи ([]городов) с расстояниями.
//     Добавь 3 города и дороги, выведи всех соседей города "A".

// 18. Расписание занятий: день недели -> слоты времени -> предмет.
//     Заполни 2 дня по 2 слота, выведи предмет в "Tue 10:00".

// 19. Калькулятор корзины: позиции с ценой (float64) и скидкой (опционально).
//     Добавь 3 позиции, посчитай итог с учетом скидок.

// 20. Профили пользователей: обязательные поля (id, name), опциональные (bio, avatarURL).
//     Смоделируй опциональные значения, выведи "bio" если задано.

// 21. Логи по уровням: level -> []Log{msg, ts}. Добавь INFO и ERROR записи.
//     Выведи все ERROR-сообщения по порядку добавления.

// 22. Результаты экзамена: studentID -> {сумма баллов, ответы map[qID]int}.
//     Добавь 2 студентов, у одного обнови ответ на qID=3, пересчитай сумму.

// 23. Дедупликация событий по id с истечением (expiresAt int64).
//     Добавь 3 события, удалите все «просроченные» (ts < now), выведи оставшиеся id.

// 24. Каталог медиа: по типу ("audio","video") — слайс элементов {title,durationSec}.
//     Добавь по 2 элемента каждого типа, выведи суммарную длительность video.

// 25. Конфигурация фич-флагов: флаг -> включен bool, целевая аудитория ([]userID).
//     Включи флаг "NewUI" для трех пользователей, проверь, включен ли для userID=42.

// 26. Счет в банке: операции (time, kind: "deposit"/"withdraw", amount).
//     Добавь 4 операции и вычисли текущий баланс.

// 27. Телеметрия: метрика -> кольцевой буфер последних N значений (int).
//     Реализуй вставку с переполнением, добавь 7 значений при N=5, выведи содержимое.

// 28. Индекс обратный для поиска по словам: слово -> ids документов ([]int).
//     Добавь связи для 3 слов и 3 документов, по слову "golang" выведи документы без дублей.

// 29. Таблица хеш-сумм файлов: путь -> {size, sha256}.
//     Добавь 2 записи, обнови sha256 одного файла, выведи все пути и их хеши.

// 30. Прайс-лист с историей цен: product -> []PricePoint{ts, price}.
//     Добавь 3 точки для одного товара и найди последнюю цену по максимальному ts.
