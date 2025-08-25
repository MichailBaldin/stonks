package main

// Задание 1: identity-функция
// Условие:
// Реализуй `func identity[T any](v T) T`, возвращающую переданное значение.


// Задание 2: дженерик `max()`
// Условие:
// Реализуй `func max[T constraints.Ordered](a, b T) T`, возвращающую максимум.


// Задание 3: дженерик `swap()`
// Условие:
// Реализуй `func swap[T any](a, b T) (T, T)`, возвращающую значения в обратном порядке.


// Задание 4: срез и дженерик `printAll()`
// Условие:
// Реализуй `func printAll[T any](list []T)`, которая печатает все элементы среза.


// Задание 5: `sum()` для int и float64
// Условие:
// Реализуй `func sum[T Number](list []T) T`
// Создай собственное ограничение `type Number interface { ~int | ~float64 }`


// Задание 6: структура `Box[T]`
// Условие:
// Создай `type Box[T any] struct { value T }`
// Добавь метод `Get() T`


// Задание 7: структура `Pair[A, B]`
// Условие:
// Создай `type Pair[A, B any] struct { First A; Second B }`
// Создай `Pair[int, string]` и выведи содержимое.


// Задание 8: `mapKeys()` из map\[K]V → \[]K
// Условие:
// Реализуй `func mapKeys[K comparable, V any](m map[K]V) []K`


// Задание 9: `contains()` — проверка в срезе
// Условие:
// Реализуй `func contains[T comparable](list []T, v T) bool`


// Задание 10: `filter()` — отфильтровать срез по условию
// Условие:
// Реализуй `func filter[T any](list []T, f func(T) bool) []T`


// Задание 11: `map()` — применить функцию к каждому элементу
// Условие:
// Реализуй `func mapSlice[T, R any](list []T, f func(T) R) []R`


// Задание 12: `indexOf()` — индекс элемента в срезе
// Условие:
// Реализуй `func indexOf[T comparable](list []T, target T) int`, возвращающую индекс или `-1`


// Задание 13: `reverse()` — перевернуть срез
// Условие:
// Реализуй `func reverse[T any](list []T) []T`


// Задание 14: структура `Stack[T]`
// Условие:
// Создай `type Stack[T any] struct` с методами `Push(T)`, `Pop() T`, `Len() int`


// Задание 15: структура `Queue[T]`
// Условие:
// То же, но для `Queue[T]`: методы `Enqueue(T)`, `Dequeue() T`, `Len() int`


// Задание 16: дженерик `minMax()`
// Условие:
// Реализуй `func minMax[T constraints.Ordered](list []T) (T, T)`
// Возвращает минимум и максимум.


// Задание 17: `zeroValue()` — возвращает значение по умолчанию
// Условие:
// Реализуй `func zeroValue[T any]() T`
// Проверь, что для `int` это 0, для `string` — ""


// Задание 18: сравнение двух значений
// Условие:
// Реализуй `func equal[T comparable](a, b T) bool`


// Задание 19: печать типов во время выполнения
// Условие:
// Реализуй `func debugType[T any](val T)` — печатай `reflect.TypeOf(val)`


// Задание 20: функция `first[T any](values ...T) T`
// Условие:
// Возвращает первый элемент (если есть), иначе — zero value.

