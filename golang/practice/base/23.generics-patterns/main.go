package main

// Задание 1: Map-утилита
// Реализуй `Map[T, R any](in []T, f func(T) R) []R`
// Проверь на `[]int → []string`


// Задание 2: Filter-утилита
// Реализуй `Filter[T any](in []T, f func(T) bool) []T`
// Фильтруй только чётные числа.


// Задание 3: Reduce-утилита
// Реализуй `Reduce[T any](in []T, zero T, f func(T, T) T) T`
// Подсчитай сумму всех элементов.


// Задание 4: Обёртка `Optional[T]`
// Создай `Optional[T]` с методами `IsPresent()`, `Get()`, `OrElse(default T) T`


// Задание 5: Дженерик `Stack[T]`
// Реализуй стек с методами `Push(T)`, `Pop() (T, bool)`, `Len() int`


// Задание 6: Дженерик `Queue[T]`
// То же самое, но `Enqueue(T)` / `Dequeue() (T, bool)`


// Задание 7: Маппинг DTO → Entity
// Создай:
// type DTO struct { Name string }
// type Entity struct { Name string }
// func mapDTO(d DTO) Entity
// И используй:
// MapStructs[DTO, Entity]([]DTO, mapDTO)


// Задание 8: Генератор `Repeat[T]`
// Реализуй `Repeat[T any](val T, n int) []T`
// Верни срез из `n` копий `val`.


// Задание 9: Валидатор списка
// Реализуй `ValidateAll[T any](in []T, rule func(T) error) error`
// Верни первую ошибку, если есть.


// Задание 10: Тестовый ассерт
// Создай `AssertEqual[T comparable](t *testing.T, expected, actual T)`


// Задание 11: Кэш `Cache[K, V]`
// Создай:
// type Cache[K comparable, V any] struct {
//     data map[K]V
// }

// func (c *Cache[K, V]) Get(k K) (V, bool)
// func (c *Cache[K, V]) Set(k K, v V)


// Задание 12: Типобезопасный флаг
// Создай `type Flag[T any] struct { val T; set bool }`
// Добавь методы `Set()`, `Get()`, `IsSet()`


// Задание 13: Структура `Repository[T]`
// Определи интерфейс с методами `GetByID(id int) (T, error)`, `Save(T) error`


// Задание 14: ReaderWrapper
// Создай:
// type ReaderWrapper[T any] struct {
//     R io.Reader
//     Parse func([]byte) (T, error)
// }
// Метод `ReadParsed() (T, error)`


// Задание 15: Обёртка над map
// Создай `MapWrapper[K comparable, V any]`, добавь `Keys() []K`, `Values() []V`


// Задание 16: Merge\[T] из двух срезов
// Реализуй `func Merge[T any](a, b []T) []T`


// Задание 17: Обёртка над ошибкой
// Создай `Result[T any] struct { Value T; Err error }`


// Задание 18: Convert[T any, U any](T) U
// Реализуй через `f func(T) U`


// Задание 19: Генерик-обёртка над слайсом
// Создай `List[T any]` с методами `Append(T)`, `Map`, `Filter`, `Len()`


// Задание 20: Дженерик `GroupBy`
// func GroupBy[T any, K comparable](items []T, keyFunc func(T) K) map[K][]T
// Сгруппируй строки по длине.

