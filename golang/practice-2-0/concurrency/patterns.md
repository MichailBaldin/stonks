отлично — фокус узкий и практичный. возьмём 2 недели и прокачаем «моторную память» до автоматизма, а затем — применение. я дам крошечные, запоминающиеся шаблоны (минимум кода), а потом — график повторений и прогрессию задач.

---

# мини-шаблоны (перепечатать до автоматизма)

## **1) Worker Pool (N воркеров, одна очередь)**

```go
// создать каналы с тасками и результатами
// объявить кол-во воркеров и вейт группу
// запустить N воркеров
// создать структуру воркера
// создать нужное кол-во тасок
// создать отдельную горутину чтобы дождаться записи всех результатов в канал и закрыть его
// вычитать все результаты из канала

package main

import (
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		// do work(j)
		results <- j
	}
}

func main() {
	const N = 4
	jobs := make(chan int)
	results := make(chan int)

	var wg sync.WaitGroup
	for i := 0; i < N; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	go func() {
		for i := 0; i < 10; i++ { jobs <- i }
		close(jobs)
	}()

	go func() { wg.Wait(); close(results) }()

	for r := range results {
		_ = r // use result
	}
}
```

## **2) Rate Limiter (тикер «по одному каждые T»)**

```go
// задаем частоту через тикер, не забываем закрыть тикер
// пускай в цикле мы что то делаем, но тормозим следующее выполнение через чтение тикера

package main

import (
	"time"
)

func main() {
	lim := time.NewTicker(100 * time.Millisecond)
	defer lim.Stop()

	for i := 0; i < 10; i++ {
		<-lim.C        // gate
		// do request(i)
	}
}
```

(вариант «токены» — просто `tokens := make(chan struct{}, B)` и периодически добавлять в него токены через тикер.)

## **3) Semaphore (ограничить параллелизм до K)**

```go
// задаем кол-во параллельных воркеров
// создаем семафор (буф канал на нужное нам кол-во)
// не забываем waitgroup
// допустим в цикле запускаем необходимое кол-во тасок
// занимаем семафор пустой структорой
// запускам горутину-таску
// после завершения читаем из семафора структуру
// ждем завершения работы всех тасок

package main

import "sync"

func main() {
	const K = 3
	sem := make(chan struct{}, K)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		sem <- struct{}{}        // acquire
		go func(i int) {
			defer wg.Done()
			// do work(i)
			<-sem                 // release
		}(i)
	}
	wg.Wait()
}
```

\*\*4) generator (порождает поток значений в канал)

```go
package main

func generator(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}
```

## **4) Fan-out (одна очередь → много воркеров, без результатов)**
```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan int)
	const N = 4
	wg := &sync.WaitGroup{}

	for i := 0; i < N; i++ {
		wg.Add(1)
		go worker(jobs, wg)
	}

	go func() {
		for j := range 10 {
			jobs <- j
		}
		close(jobs)
	}()

	wg.Wait()
}

func worker(jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Println(job * 2)
	}
}
```
Немного пошалим и добавим еще контекст для жиру
```go
package main

import (
	"context"
	"sync"
	"time"
)

func worker(ctx context.Context, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case v, ok := <-jobs:
			if !ok { return }        // канал закрыт — выходим
			_ = v                    // process(v)
		}
	}
}

func main() {
	const N = 4
	jobs := make(chan int)
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// воркеры
	wg.Add(N)
	for i := 0; i < N; i++ {
		go worker(ctx, jobs, &wg)
	}

	// продюсер
	go func() {
		for j := 0; j < 10; j++ { jobs <- j }
		close(jobs)                 // мы открыли — мы и закрываем
	}()

	wg.Wait()
}
```

## **5) Fan-in (несколько источников → один канал)**

```go
package main

import (
	"fmt"
	"sync"
)

// forward читает из одного входного канала и пишет всё в out.
func forward(c <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range c {
		out <- v
	}
}

// fanIn объединяет несколько входных каналов в один.
func fanIn(chs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(chs))
	for _, c := range chs {
		go forward(c, out, &wg)
	}

	// Закроем out, когда все forward завершат работу
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	a := make(chan int)
	b := make(chan int)

	// два источника
	go func() {
		defer close(a)
		for i := 0; i < 3; i++ {
			a <- i
		}
	}()
	go func() {
		defer close(b)
		for i := 100; i < 103; i++ {
			b <- i
		}
	}()

	out := fanIn(a, b)
	for v := range out {
		fmt.Println(v)
	}
}

```
fan-in с контекстом

```go
package main
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// forwardCtx читает из одного канала и пишет в out,
// завершается по закрытию входа или по ctx.Done().
func forwardCtx(ctx context.Context, c <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case v, ok := <-c:
			if !ok {
				return
			}
			select {
			case <-ctx.Done():
				return
			case out <- v:
			}
		}
	}
}

// fanInCtx объединяет несколько каналов в один с поддержкой context.
func fanInCtx(ctx context.Context, chs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(chs))
	for _, c := range chs {
		go forwardCtx(ctx, c, out, &wg)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	a := make(chan int)
	b := make(chan int)

	// источники
	go func() {
		defer close(a)
		for i := 0; i < 5; i++ {
			a <- i
			time.Sleep(50 * time.Millisecond)
		}
	}()
	go func() {
		defer close(b)
		for i := 100; i < 105; i++ {
			b <- i
			time.Sleep(70 * time.Millisecond)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	out := fanInCtx(ctx, a, b)
	for v := range out {
		fmt.Println(v)
	}
}

```

## **6) Цепочка обработчиков / конвейер (stage-by-stage)**

```go
package main

func stage1(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in { out <- v /* transform1 */ }
	}()
	return out
}

func stage2(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in { out <- v /* transform2 */ }
	}()
	return out
}

func main() {
	src := make(chan int)
	go func() { for i := 0; i < 10; i++ { src <- i }; close(src) }()

	p1 := stage1(src)
	p2 := stage2(p1)

	for v := range p2 { _ = v }
}
```

## **7) fan-in → fan-out**
слить несколько источников в один поток, затем распараллелить обработку на N воркеров и собрать результаты обратно в один канал.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

////////////////////////////////////////////////////////////
// 1) FAN-IN: несколько входных каналов int -> один общий //
////////////////////////////////////////////////////////////

func fanIn(chs ...<-chan int) <-chan int {
	out := make(chan int) // общий выход
	var wg sync.WaitGroup

	forward := func(c <-chan int) {
		defer wg.Done()
		for v := range c {  // пока вход не закрыт
			out <- v        // перекладываем элемент в общий out
		}
	}

	wg.Add(len(chs))
	for _, c := range chs {
		go forward(c)
	}

	// Закрываем out, когда ВСЕ "писатели" в него завершатся.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

////////////////////////////////////////////////////////////
// 2) FAN-OUT: один вход -> N воркеров, общий выход out  //
////////////////////////////////////////////////////////////

// startWorkers запускает N одинаковых воркеров.
// Каждый читает из общего in (fan-out) и пишет результат в общий out.
// Когда все воркеры завершатся, out закрывается.
func startWorkers(n int, in <-chan int) <-chan int {
	out := make(chan int) // общий канал результатов
	var wg sync.WaitGroup

	worker := func(id int) {
		defer wg.Done()
		for v := range in {         // каждый воркер читает из ОДНОГО in
			// Здесь — "работа". Сейчас просто заглушка.
			// Реальная логика: compute(v), io(v), http(v) и т.п.
			res := process(v)
			out <- res              // отправляем результат
		}
	}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go worker(i)
	}

	// Важно: out закрывает ТОЛЬКО владелец (здесь — мы),
	// после того как все воркеры закончат писать.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

/////////////////////////////////////////////
// 3) Простейшая "работа" — заглушка        //
/////////////////////////////////////////////

func process(v int) int {
	// Имитируем нагрузку (не обязательно для шаблона)
	time.Sleep(10 * time.Millisecond)
	return v * 10 // например, умножаем на 10
}

////////////////////////////////////////////////////////////
// 4) Мини-демо сборки: (srcA, srcB) -> fanIn -> workers //
////////////////////////////////////////////////////////////

func main() {
	// Два источника как пример (в реальном коде это могут быть разные производители)
	srcA := make(chan int)
	srcB := make(chan int)

	// Кто открыл канал — тот и закрывает.
	go func() {
		defer close(srcA)
		for i := 1; i <= 5; i++ { srcA <- i }
	}()
	go func() {
		defer close(srcB)
		for i := 100; i <= 104; i++ { srcB <- i }
	}()

	// 1) Сливаем источники -> общий поток
	in := fanIn(srcA, srcB)

	// 2) Распараллеливаем обработку на N воркеров (fan-out)
	const N = 3
	results := startWorkers(N, in)

	// 3) Читаем единый поток результатов
	for r := range results {
		fmt.Println("result:", r)
	}
}
```

> мнемоники:\
> • **WP** = jobs→workers→results + `wg.Wait()` → `close(results)`\
> • **RL** = «ждём тикер» или «токен из буфера»\
> • **Sem** = `sem <- struct{}{}` (acquire), `<-sem` (release)\
> • **Fan-out** = много `go worker()` читают **один** `in`\
> • **Fan-in** = `merge` с `WaitGroup` и `close(out)` в конце\
> • **Chain/Pipeline** = «каждый stage: читает in → пишет out → закрывает out»

---

# план на 2 недели (короткие ежедневные циклы + повторения)

Каждый день: 30–45 мин. Структура занятия: 5 мин обзор (из головы), 20–30 мин печатаем по памяти, 5–10 мин быстрые проверки.

**Неделя 1 — «механика шаблонов»**

- **День 1 (Пн): Worker Pool.**
  Цель: набрать скелет за ≤5 мин с `jobs/results/wg`.
  Повтор: в конце дня — разочек с нуля.
- **День 2 (Вт): Rate Limiter.**
  Тикер + вариант с токенами (буфер).
  Повтор: WP из Д1 (быстрый).
- **День 3 (Ср): Semaphore.**
  Ограничение параллелизма K.
  Повтор: RL из Д2.
- **День 4 (Чт): Fan-out.**
  Несколько воркеров читают один `in`.
  Повтор: Sem из Д3.
- **День 5 (Пт): Fan-in (merge).**
  Универсальная `merge[T any]`.
  Повтор: Fan-out из Д4.
- **День 6 (Сб): Цепочка обработчиков (pipeline).**
  2–3 стадии, каждая закрывает свой `out`.
  Повтор: Fan-in из Д5.
- **День 7 (Вс): Ревизия недели.**
  6 мини-скелетов — каждый перепечатать за ≤3–5 мин.

**Неделя 2 — «вариации и применение»**

- **День 8:** WP + Sem вместе (пул, но каждая задача может запускать под-горутину с ограничением K).
- **День 9:** RL поверх WP (каждый `send` в `jobs` через тикер или токены).
- **День 10:** Fan-out → Pipeline (каждый воркер — стадия 1, потом стадия 2).
- **День 11:** Fan-in нескольких источников в один pipeline.
- **День 12:** Отмена/таймауты (добавь `context.Context` в один из шаблонов).
- **День 13:** Безопасное завершение: кто закрывает какой канал, где `wg.Wait()`.
- **День 14:** Контрольная: «на чистом листе» 6 шаблонов + один комбинированный.

---

# график повторения (spaced)

- **D1 вечером**: повтор темы D1.
- **D3 утром**: повтор D1+D2 (из головы, по 3–5 мин).
- **D5 утром**: повтор D1–D3.
- **D7**: повтор D1–D6 (уже есть).
- **D10**: быстрый повтор всех 6.
- **D14**: финальный повтор + самотест.

---

# прогрессия задач (после механики)

К каждой теме — по возрастающей сложности:

**Worker Pool**

1. Фиксированный N, 100 задач, суммируй результаты.
2. Задачи разной «тяжести» (sleep), проверь отсутствие «затыков».
3. Динамическая остановка: закрыть `jobs` по сигналу `ctx.Done()`.

**Rate Limiter**

1. 1 запрос/100мс — измерь фактический интервал.
2. «всплеск» из B токенов, затем 1/Δt — токен-бакет.
3. Пер-ключевой лимит (map\[key]лимитер).

**Semaphore**

1. Ограничь одновременные I/O до K.
2. Семофор на ресурсы: K=кол-во подключений к БД.
3. Справедливость (очередность) — обсуди, нужно ли.

**Fan-out**

1. Парсинг строк файла параллельно.
2. Убедись, что `in` закрывается, воркеры завершаются.
3. Замерь ускорение при разных N.

**Fan-in**

1. Слить два тикера разной частоты.
2. Слить результаты от 3 HTTP-источников.
3. Обработать закрытие любого источника корректно.

**Цепочка обработчиков**

1. 3 стадии: parse → transform → encode.
2. Вставь `context` для отмены по таймауту.
3. Backpressure: ограничь буфер между стадиями.

> правило безопасности: «Кто открыл канал — тот и закрывает». `WaitGroup` закрывает **выходной** канал после уходов всех писателей.

---

# как тренироваться эффективно

- **Активное вспоминание:** печатай из головы. Подглядывать — только после компиляции.
- **Ограничения по времени:** ставь таймер на 5 минут на шаблон.
- **Мини-чеклист:** для каждого шаблона проговаривай: «кто закрывает канал, где WaitGroup, кто ждёт кого».
- **Сигналы завершения:** добавь `ctx` минимум в один шаблон на второй неделе.

---

хочешь — начнём прямо сейчас с первого шага. за 5 минут набросай из памяти **worker pool** (jobs/results/wg, как выше), и скажи, на каком месте споткнулся. если удобнее — я дам «скелет без тел», а ты впишешь пустые места.

**вопрос на старт:** сколько минут в день ты реально готов уделять (30, 45 или 60)? от этого подстрою длительность блоков и объём задач.
