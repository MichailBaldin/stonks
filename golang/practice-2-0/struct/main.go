package main

// ======================================
// Тренажёр: Структуры, встраивание, коллекции (v1)
// ======================================
// Формат как раньше: короткие условия, базовое API, без изысков.

// --- Блок A. База структур (литералы, zero value, сравнение) ---

// 1. Опиши структуру User{ID int; Name string}. Создай zero value и литерал, выведи поля.
// 2. Добавь метод String() для User, который печатает "ID:Name". Вызови через fmt.Println(u).
// 3. Сделай тип Address{City string; Street string}. Вложи Address в User как поле (композиция), выведи City.
// 4. Сравни два значения Address оператором ==. Объясни, почему User с полем-срезом так сравнить нельзя (закомментируй попытку).
// 5. Создай слайс []User из 3 элементов и измени Name у второго — проверь выводом.

// --- Блок B. Методы: значение vs указатель ---

// 6. Добавь к User метод Rename(new string) с value receiver. Вызови и проверь, изменился ли Name.
// 7. Переделай Rename на pointer receiver и повтори проверку.
// 8. Реализуй метод Clone() User, который возвращает копию со сменой Name (оригинал не меняется).
// 9. Сделай функцию NewUser(id int, name string) *User (конструктор) — верни указатель, проверь, что поля выставлены.
// 10. Для Address сделай метод Full() string = "City, Street". Вызови через u.Address.Full().

// --- Блок C. Встраивание (embedding) полей и метод-promotion ---

// 11. Создай тип Timestamps{CreatedAt int64; UpdatedAt int64}. Встрой его анонимно в User: Timestamps.
//     Доступ к полям через u.CreatedAt.
// 12. Добавь Timestamps метод Touch() — обновляет UpdatedAt. Вызови u.Touch() (через promotion).
// 13. Создай тип Named{Name string} и встрой в User: Named. Покажи конфликт имён полей (если уже есть Name).
// 14. Разреши конфликт: убери поле Name в User, оставь его во встраиваемом Named. Доступ к имени — u.Name.
// 15. Встрой *Address (указатель) в User: *Address. Проверь панику при обращении u.City без инициализации. Почини инициализацией.
// 16. Добавь к Address метод Move(city, street). Вызови u.Move(... ) через promotion.
// 17. Создай админов: тип Admin с встраиваемым User. Добавь метод IsAdmin() bool. Проверь доступ к полям User из Admin.
// 18. Сделай два уровня встраивания: SuperAdmin{Admin}, вызови метод Touch() (из Timestamps через User) на SuperAdmin.
// 19. Встрой sync.Mutex в тип SafeCounter{sync.Mutex; m map[string]int}. Покажи Lock/Unlock и инкремент по ключу.
// 20. Удали embedding Mutex и сделай именованное поле. Покажи, что теперь нужен s.Mutex.Lock() вместо promotion.

// --- Блок D. Коллекции структур (слайсы) ---

// 21. Создай []User и реализуй функции: AppendUser([]User, User) []User; IndexByID([]User, id) int.
// 22. Отфильтруй []User по предикату (Name != ""), используя "in-place filter": dst := users[:0].
// 23. Отсортируй []User по Name с помощью sort.Slice.
// 24. Реализуй Paginate(users []User, page, per int) []User — верни подотрезок с границами.
// 25. Сделай глубокое копирование []User (копии значений), измени копию и покажи независимость от оригинала.

// --- Блок E. Коллекции структур (мапы) ---

// 26. Построй индекс: map[int]*User по слайсу пользователей. Проверь, что изменения через указатель видны в исходном слайсе.
// 27. Реализуй MergeUsersByID(dst map[int]*User, src []User): добавляй отсутствующих, обновляй Name по совпавшим ID.
// 28. Инвертируй индекс имён: map[string][]int где ключ — Name, значение — список ID.
// 29. Реализуй DiffUsersByID(a, b map[int]*User]): верни списки added/removed/updated ID.
// 30. Реализуй безопасное чтение: GetUser(m map[int]*User, id int) (*User, bool), не создавая лишних записей.

// --- Блок F. Вложенные слайсы и мапы внутри структур ---

// 31. Тип Group{Name string; Members []User}. Добавь метод Add(u User) и RemoveByID(id int) с сохранением порядка.
// 32. Тип Org{Name string; Teams map[string]*Group}. Функция AddToTeam(org *Org, team string, u User) с инициализацией map и Members.
// 33. Тип Catalog{ByCity map[string][]*User}. Добавь пользователя по городу, инициализируя при первом добавлении.
// 34. Тип KVStore{Buckets map[string]map[string]string}. Реализуй Upsert(bucket, key, val string).
// 35. Покажи aliasing: два Group ссылаются на один и тот же слайс Members. Измени один — увидь изменения в другом. Почини глубокой копией.

// --- Блок G. JSON-теги и сериализация ---

// 36. Добавь json-теги в User: id, name. Сериализуй []User в JSON и обратно.
// 37. Убери поле Address из JSON (omitempty + "-" для внутренних). Проверь результат.
// 38. Тип Profile{User; Skills []string}. Сериализуй и проверь, что поля User промотируются в JSON.
// 39. Создай map[int]User и сериализуй — поймай конвертацию ключей в строки. Прочитай обратно в map[string]User или в []pair.

// --- Блок H. Конструкторы и валидация ---

// 40. Напиши NewAddress(city, street string) (*Address, error) с простой валидацией (не пустые строки).
// 41. Сделай NewUserWithAddress(id int, name string, addr *Address) (*User, error) — собирает валидный объект.
// 42. Добавь метод Validate() error для User (проверка Name != ""). Вызови перед добавлением в коллекцию.
// 43. Тип Registry{byID map[int]*User}. Реализуй Register(u *User) error — ошибка при дубликате ID, иначе сохранить.
// 44. Реализуй FindByCity(reg *Registry, city string) []User — по вложенному адресу (*Address).

// --- Блок I. Копирование, мутабельность, границы ---

// 45. Покажи разницу копии значения и указателя: u2 := u; p := &u; измени p.Name — проверь u2.Name.
// 46. В Group реализуй DeepClone() Group: новая копия слайса и значений.
// 47. Сделай SafeSlice(users []User) []User, который возвращает копию для чтения "без утечки" внутреннего слайса наружу.
// 48. Реализуй метод TrimTo(n int) для Group — усечь Members до n безопасно (границы, не паниковать).
// 49. Реализуй ReplaceMember(g *Group, i int, u User) error — безопасная замена по индексу с проверкой границ.
// 50. Реализуй Snapshot(reg *Registry) map[int]User — вернуть "замороженную" копию (значения, не указатели).

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

// 31. Сессии пользователей в системе: userID -> время последнего действия.
//     Добавь 3 сессии, обнови у одного пользователя и выведи "самый свежий".

// 32. История переводов в банке: перевод {from, to, amount}, хранить список для каждого userID.
//     Добавь 2 пользователя и по 2 перевода, выведи все переводы "user1".

// 33. Библиотека: книга имеет ISBN, название и список авторов.
//     Храни книги в map по ISBN, добавь 2 книги и выведи авторов второй.

// 34. Планировщик задач: очередь задач по времени (timestamp, text).
//     Добавь 3 задачи, выведи ближайшую по времени.

// 35. Календарь встреч: день -> список встреч (start, end, title).
//     Добавь встречи на Mon и Tue, выведи все на Tue.

// 36. Каталог фильмов: жанр -> []Movie{title, year, rating}.
//     Добавь 2 жанра, по 2 фильма, выведи лучший по рейтингу из "Drama".

// 37. Интернет-магазин: корзина пользователя хранит товары (sku -> count).
//     Добавь 2 товара в корзину, увеличь количество одного, выведи итог.

// 38. Соцсеть: посты. У поста: id, текст, список лайков (userIDs).
//     Добавь 2 поста, по 2 лайка, выведи всех, кто лайкнул первый пост.

// 39. Игровой матч: у команды список игроков, у игрока статистика (kills, deaths).
//     Создай 2 команды, по 2 игрока, выведи лучшего игрока матча по kills.

// 40. Лента событий: массив структур {id, type, payload}.
//     Добавь 5 событий разных типов, выведи только события типа "error".
