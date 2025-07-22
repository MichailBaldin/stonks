Вот материал по **Linked Lists**!

**🧠 Главная идея:** в отличие от массивов, где элементы лежат рядом в памяти, в списках каждый элемент "знает" только где находится следующий.

**🔑 Ключевые преимущества vs массивы:**

- Динамический размер
- O(1) вставка/удаление (если есть указатель на узел)
- Не нужно заранее знать размер

**🔑 Главные недостатки:**

- Нет произвольного доступа O(1) по индексу
- Больше памяти на хранение указателей
- Хуже cache locality

**💡 Самые важные паттерны для собеседований:**

1. **Slow & Fast pointers** - решает кучу задач (поиск середины, циклов, n-го элемента с конца)
2. **Dummy node** - упрощает 90% кода, особенно когда может измениться голова
3. **Разворот списка** - базовый навык, спрашивают очень часто

**⚠️ Частые ловушки:**

- Забыть проверить `node == nil`
- Потерять ссылку на следующий узел при развороте
- Неправильно обработать случай изменения головы списка

**🔥 Алгоритм Floyd'а** (поиск цикла) - это классика! Понимание того, почему slow и fast указатели встретятся именно в цикле, показывает глубину понимания.

**🎯 Признак задачи на Linked List:** если в условии упоминается "список", "цепочка", "последовательность узлов" - скорее всего это оно.

Попробуй мысленно прогнать алгоритмы, особенно разворот списка - его очень часто просят написать на доске.

# Linked Lists - Связанные списки

## Основные концепции

### Что это такое?

**Linked List** - это структура данных, где элементы (узлы) связаны друг с другом через указатели. Каждый узел содержит данные и ссылку на следующий узел.

```go
type ListNode struct {
    Val  int
    Next *ListNode
}
```

### Зачем нужны, если есть массивы?

| Операция          | Массив | Linked List     |
| ----------------- | ------ | --------------- |
| Доступ по индексу | O(1)   | O(n)            |
| Поиск элемента    | O(n)   | O(n)            |
| Вставка в начало  | O(n)   | **O(1)**        |
| Вставка в конец   | O(1)   | O(n) или O(1)\* |
| Удаление          | O(n)   | **O(1)\***      |

\*если есть указатель на нужный узел

**Главное преимущество:** динамический размер + быстрая вставка/удаление.

## Типы связанных списков

### 1. Односвязный список (Singly Linked List)

```go
type ListNode struct {
    Val  int
    Next *ListNode
}

// Пример: 1 -> 2 -> 3 -> nil
```

### 2. Двусвязный список (Doubly Linked List)

```go
type DoublyNode struct {
    Val  int
    Next *DoublyNode
    Prev *DoublyNode
}

// Пример: nil <- 1 <-> 2 <-> 3 -> nil
```

### 3. Циклический список

Последний узел ссылается на первый (может быть как односвязным, так и двусвязным).

## Базовые операции

### Создание и вставка

```go
// Вставка в начало
func insertAtHead(head *ListNode, val int) *ListNode {
    newNode := &ListNode{Val: val}
    newNode.Next = head
    return newNode
}

// Вставка в конец
func insertAtTail(head *ListNode, val int) *ListNode {
    newNode := &ListNode{Val: val}

    if head == nil {
        return newNode
    }

    current := head
    for current.Next != nil {
        current = current.Next
    }
    current.Next = newNode

    return head
}

// Вставка после определенного узла
func insertAfter(node *ListNode, val int) {
    if node == nil {
        return
    }

    newNode := &ListNode{Val: val}
    newNode.Next = node.Next
    node.Next = newNode
}
```

### Удаление

```go
// Удаление по значению
func deleteValue(head *ListNode, val int) *ListNode {
    // Если удаляем голову
    if head != nil && head.Val == val {
        return head.Next
    }

    current := head
    for current != nil && current.Next != nil {
        if current.Next.Val == val {
            current.Next = current.Next.Next
            return head
        }
        current = current.Next
    }

    return head
}

// Удаление узла, имея только ссылку на него (хитрый трюк!)
func deleteNode(node *ListNode) {
    // Копируем значение следующего узла в текущий
    node.Val = node.Next.Val
    // Удаляем следующий узел
    node.Next = node.Next.Next
}
```

## Ключевые паттерны для собеседований

### 1. Two Pointers в списке (Slow & Fast)

**Поиск середины списка:**

```go
func findMiddle(head *ListNode) *ListNode {
    slow := head
    fast := head

    // fast движется в 2 раза быстрее
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    return slow // slow будет в середине
}
```

**Поиск цикла (Floyd's Algorithm):**

```go
func hasCycle(head *ListNode) bool {
    slow := head
    fast := head

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next

        if slow == fast {
            return true // цикл найден
        }
    }

    return false
}

// Найти начало цикла
func detectCycle(head *ListNode) *ListNode {
    slow := head
    fast := head

    // Находим точку встречи
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            break
        }
    }

    if fast == nil || fast.Next == nil {
        return nil // нет цикла
    }

    // Ставим один указатель на начало
    slow = head
    while slow != fast {
        slow = slow.Next
        fast = fast.Next
    }

    return slow // начало цикла
}
```

### 2. Разворот списка

**Итеративный способ:**

```go
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    current := head

    for current != nil {
        next := current.Next  // сохраняем следующий
        current.Next = prev   // разворачиваем связь
        prev = current        // двигаем prev
        current = next        // двигаем current
    }

    return prev // новая голова
}
```

**Рекурсивный способ:**

```go
func reverseListRecursive(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    newHead := reverseListRecursive(head.Next)
    head.Next.Next = head
    head.Next = nil

    return newHead
}
```

### 3. Слияние отсортированных списков

```go
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{} // фиктивный узел для упрощения
    current := dummy

    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            current.Next = l1
            l1 = l1.Next
        } else {
            current.Next = l2
            l2 = l2.Next
        }
        current = current.Next
    }

    // Добавляем оставшиеся элементы
    if l1 != nil {
        current.Next = l1
    }
    if l2 != nil {
        current.Next = l2
    }

    return dummy.Next // возвращаем без фиктивного узла
}
```

### 4. Удаление n-го элемента с конца

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next: head}
    slow := dummy
    fast := dummy

    // Двигаем fast на n+1 шагов вперед
    for i := 0; i <= n; i++ {
        fast = fast.Next
    }

    // Двигаем оба указателя, пока fast не дойдет до конца
    for fast != nil {
        slow = slow.Next
        fast = fast.Next
    }

    // Удаляем узел
    slow.Next = slow.Next.Next

    return dummy.Next
}
```

## Важные детали и ловушки

### 1. Dummy Node (фиктивный узел)

**Зачем:** Упрощает обработку граничных случаев, особенно когда может измениться голова списка.

```go
dummy := &ListNode{}
dummy.Next = head
// работаем с dummy.Next
return dummy.Next
```

### 2. Проверка на nil

**Всегда проверяй:**

- `head == nil`
- `node.Next == nil` перед обращением к `node.Next.Next`

### 3. Обновление головы списка

Если функция может изменить голову списка, она должна возвращать новую голову:

```go
func modifyList(head *ListNode) *ListNode {
    // может измениться голова
    return newHead
}
```

## Классические задачи на собеседованиях

1. **Reverse Linked List** - базовая задача на разворот
2. **Detect Cycle** - поиск цикла алгоритмом Floyd'а
3. **Merge Two Sorted Lists** - слияние отсортированных списков
4. **Remove Nth Node From End** - удаление n-го элемента с конца
5. **Find Middle of List** - поиск середины списка
6. **Palindrome Linked List** - проверка на палиндром
7. **Copy List with Random Pointer** - копирование сложного списка

## Когда использовать Linked List vs Array

**Используй Linked List когда:**

- ✅ Частые вставки/удаления в начале или середине
- ✅ Размер данных заранее неизвестен
- ✅ Не нужен произвольный доступ по индексу

**Используй Array когда:**

- ✅ Нужен быстрый доступ по индексу
- ✅ Важна locality of reference (кэш-эффективность)
- ✅ Много операций поиска

## На собеседовании

**Частые вопросы:**

1. "Чем отличается от массива?" → Динамический размер, O(1) вставка vs O(1) доступ
2. "Как найти цикл?" → Floyd's algorithm с двумя указателями
3. "Как развернуть список?" → Итеративно или рекурсивно
4. "Зачем нужен dummy node?" → Упрощает граничные случаи

**Как решать задачи:**

1. Нарисуй схему списка на бумаге/доске
2. Проговори алгоритм пошагово
3. Обязательно обработай граничные случаи (пустой список, один элемент)
4. Используй dummy node при необходимости
