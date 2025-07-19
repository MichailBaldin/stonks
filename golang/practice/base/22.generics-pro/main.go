package main

// Задание 1: Ограничение только `int`
// Условие:
// Создай `type IntOnly interface { ~int }`
// Реализуй `func square[T IntOnly](x T) T`


// Задание 2: Объединение `int` и `float64`
// Условие:
// Создай `type Number interface { ~int | ~float64 }`
// Реализуй `func add[T Number](a, b T) T`


// Задание 3: Использование `constraints.Ordered`
// Условие:
// Скачай `golang.org/x/exp/constraints`
// Создай `func maxOrdered[T constraints.Ordered](a, b T) T`


// Задание 4: Сравнение с `comparable`
// Условие:
// Реализуй `func isEqual[T comparable](a, b T) bool`


// Задание 5: Индекс по значению
// Условие:
// Создай `func indexOf[T comparable](list []T, x T) int`


// Задание 6: Обобщённый `Min()` с ограничением
// Условие:
// Реализуй `func min[T constraints.Ordered](list []T) T`
// Верни минимальный элемент.


// Задание 7: Type union для сложных чисел
// Условие:
// Создай `type Real interface { ~float32 | ~float64 }`
// Реализуй `func half[T Real](x T) T`


// Задание 8: Параметризованный метод
// Условие:
// Создай `type Calculator struct{}`
// И метод `Multiply[T Number](a, b T) T`


// Задание 9: Структура с дженерик-методом
// Условие:
// Создай `type Context struct{}`
// Добавь метод `WithValue[T any](key string, val T)`


// Задание 10: Метод со своим параметром типа
// Условие:
// Создай `func (s *Stack[T]) Map[R any](f func(T) R) []R`
// Применяет функцию к каждому элементу стека.


// Задание 11: Функция `CopyMap()`
// Условие:
// Реализуй `func CopyMap[K comparable, V any](src map[K]V) map[K]V`


// Задание 12: Свой `Filter()`
// Условие:
// Реализуй `func Filter[T any](input []T, f func(T) bool) []T`


// Задание 13: Свой `Reduce()`
// Условие:
// Реализуй `func Reduce[T any](input []T, zero T, f func(T, T) T) T`


// Задание 14: Тип с базой `int`
// Условие:
// Создай `type MyInt int`
// Проверь, входит ли `MyInt` в `~int`
// Проверь, можно ли передать `MyInt` в `Number`.


// Задание 15: Обобщённая `Optional[T]`
// Условие:
// Создай тип `Optional[T any]` со свойствами `value T` и `present bool`
// Добавь методы `Get()`, `IsPresent()`


// Задание 16: Утилита `MapKeys()`
// Условие:
// Реализуй `func MapKeys[K comparable, V any](m map[K]V) []K`


// Задание 17: `TryParse[T Number](s string) (T, error)`
// Условие:
// Реализуй функцию, парсящую строку в `int` или `float64`, в зависимости от `T`


// Задание 18: Типовой алиас и ограничение
// Условие:
// Создай:
// type ID = int
// type IDLike interface { ~int }
// Реализуй функцию `validateID[T IDLike](id T)`


// Задание 19: Проверка типов внутри дженерика (не работает)
// Условие:
// Попробуй сделать:
// if T == int { ... } // ❌
// Почему нельзя? Как сделать иначе?


// Задание 20: Обобщённый обёртчик над io.Reader
// Условие:
// Создай `type ReaderWrapper[T any] struct { R io.Reader; Parser func([]byte) T }`
// Добавь метод `ReadParsed() (T, error)`

