package main

// ======================================
// Тренажёр: Дженерики (v1)
// ======================================
// Делай по 5–10 заданий за подход. Компилируй, запускай, проверяй вывод/ошибки.
// Подсказка-мнемо: [Параметры типа] + (Ограничение) + {Тело} + Инстанцирование.

// --- Блок A. База синтаксиса ---

// 1. Напиши Identity[T any](v T) T. Вызови с int и string.
// 2. Напиши Pair[A any, B any] { First A; Second B }. Создай Pair[int,string].
// 3. Напиши Swap[T any](a, b T) (T, T). Проверь на int и bool.
// 4. Напиши Equal[T comparable](a, b T) bool. Проверь на string и int.
// 5. Покажи некомпилирующийся пример: Equal на []int (закомментируй и подпиши, почему).

// --- Блок B. Ограничения (constraints) и наборы типов ---

// 6. Создай Ordered интерфейс: объединение ~int | ~int64 | ~float64 | ~string.
// 7. Напиши Min[T Ordered](a, b T) T и Max[T Ordered](a, b T) T.
// 8. Напиши Clamp[T Ordered](x, lo, hi T) T.
// 9. Сделай Number интерфейс: ~int | ~int32 | ~int64 | ~float32 | ~float64. Напиши Sum[T Number]([]T) T.
// 10. Создай Stringer интерфейс с методом String() string. Напиши PrintAll[T Stringer]([]T).

// --- Блок C. Вывод типов (type inference) ---

// 11. Вызови Identity без явного [T] — пусть компилятор выведет тип по аргументу.
// 12. Сделай MakeZero[T any]() T и покажи, что без T вызвать нельзя — инстанцируй явным образом.
// 13. Для Pair[A,B] создай конструктор NewPair[A,B](a A, b B) Pair[A,B]; вызови без явного указания [A,B].
// 14. Покажи кейс, когда inference не срабатывает (например, пустой срез) — исправь явной аннотацией [T].

// --- Блок D. Слайсы + дженерики ---

// 15. MapSlice[T any, R any](in []T, f func(T) R) []R. Прогони на []int -> []string.
// 16. Filter[T any](in []T, pred func(T) bool) []T. Отфильтруй чётные.
// 17. Reduce[T any, R any](in []T, init R, f func(R, T) R) R. Посчитай сумму/конкатенацию.
// 18. Contains[T comparable](in []T, x T) bool. Проверь для строк.
// 19. Unique[T comparable]([]T) []T. Удали дубликаты, сохрани порядок.
// 20. IndexOf[T comparable]([]T, x T) int. Верни -1, если нет.
// 21. Chunk[T any]([]T, size int) [][]T. Разбей на чанки.
// 22. ReverseInPlace[T any]([]T). Проверь на []rune.

// --- Блок E. Мапы + дженерики ---

// 23. Keys[K comparable, V any](m map[K]V) []K — ключи (без упорядочивания).
// 24. Values[K comparable, V any](m map[K]V) []V.
// 25. Invert[K comparable, V comparable](m map[K]V) map[V]K (при коллизиях — последняя запись побеждает).
// 26. GroupBy[T any, K comparable](in []T, key func(T) K) map[K][]T.
// 27. CountBy[T any, K comparable](in []T, key func(T) K) map[K]int.
// 28. MergeMaps[K comparable, V any](a, b map[K]V) map[K]V (b перекрывает a).
// 29. CloneMap[K comparable, V any](m map[K]V) map[K]V (поверхностная копия).

// --- Блок F. Обобщённые структуры и методы ---

// 30. Стек: type Stack[T any] struct { data []T }. Методы Push, Pop (ok bool), Len.
// 31. Очередь: type Queue[T any] с методами Enqueue, Dequeue (ok bool).
// 32. Множество: type Set[T comparable] map[T]struct{}. Методы Add, Has, Delete, List.
// 33. Пара: type KV[K comparable, V any] struct { K K; V V }. Сконвертируй []KV в map[K]V.
// 34. Кольцевой буфер: type Ring[T any] { buf []T; head, size int }. Методы Put, Get (ok bool).
// 35. Пул: type Pool[T any] { new func() T; free []T }. Методы Get() T, Put(T).

// --- Блок G. Подстановки с методами (method sets) ---

// 36. Интерфейс Adder[T any]: Add(T) T. Напиши ApplyAdd[T any](x T, y T, a interface{ Add(T) T }) T (или через дженерик-параметр).
// 37. Ограничение JSONable: interface { MarshalJSON() ([]byte, error) }. Напиши MarshalAll[T JSONable]([]T) [][]byte.
// 38. Ограничение WriterLike: interface { Write([]byte) (int, error) }. Напиши WriteAll[T WriterLike](w T, parts ...[]byte) (int, error).

// --- Блок H. Подстановки по производным типам (tilde ~) ---

// 39. Создай constraint IntLike: ~int | ~int32 | ~int64. Напиши Abs[T IntLike](x T) T.
// 40. Создай свой тип MyInt с базовым типом int. Проверь, что Abs работает с MyInt благодаря ~.
// 41. Создай constraint StringLike: ~string. Напиши Join[T StringLike]([]T, sep string) string.

// --- Блок I. Ошибки компиляции как часть тренажа ---

// 42. Покажи ошибку: использование non-comparable в Set[T comparable]; затем закомментируй тест с []int.
// 43. Покажи ошибку: попытка Ordered без ~ для пользовательского типа; исправь, добавив ~ к базовому типу.
// 44. Покажи ошибку вывода типа при вызове Generic(nil) — исправь явной аннотацией [T].

// --- Блок J. Алгоритмы на дженериках ---

// 45. BinarySearch[T Ordered](in []T, x T) int (индекс или -1).
// 46. MergeSorted[T Ordered](a, b []T) []T (слияние двух отсортированных).
// 47. QuickSortInPlace[T Ordered]([]T). Прогони на []int и []string.
// 48. TopN[T Ordered]([]T, n int) []T (без полной сортировки, можно частично).
// 49. DedupSorted[T comparable]([]T) []T (для заранее отсортированного списка).

// --- Блок K. Практика с JSON/IO и дженериками ---

// 50. DecodeJSONFromFile[T any](path string, out *T) error и EncodeJSONToFile[T any](path string, v T) error — проверь на T=map[string]int и T=[]Pair[string,int].
