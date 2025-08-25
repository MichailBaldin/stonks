task3

```go
type WorkerPool struct {
	numWorkers int
	taskChan   chan string
	wg         sync.WaitGroup
}

func NewWorkerPool(numWorkers int) *WorkerPool {
	return &WorkerPool{
		numWorkers: numWorkers,
		taskChan:   make(chan string, 100),
	}
}

func (wp *WorkerPool) SendTask(task string) {
	wp.taskChan <- task
}

func (wp *WorkerPool) Start() {
	for i := range wp.numWorkers {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()

	for task := range wp.taskChan {
		fmt.Printf("Worker#%d make: %s\n", id, task)
		duration := time.Duration(rand.Intn(400) + 100)
		time.Sleep(duration * time.Millisecond)
		fmt.Printf("Worker#%d finish task: %s. Duration: %d\n", id, task, duration)
	}

	fmt.Printf("Worker#%d finished work")
}

func (wp *WorkerPool) Stop() {
	close(wp.taskChan)
	wp.wg.Wait()
}

func main() {
	pool := NewWorkerPool(3)

	pool.Start()

	tasks := []string{
		"Обработать заказ #1",
		"Отправить email",
		"Сгенерировать отчет",
		"Обновить базу данных",
		"Обработать заказ #2",
		"Создать бэкап",
		"Валидировать данные",
		"Обработать заказ #3",
	}

	for _, task := range tasks {
		pool.SendTask(task)
		time.Sleep(time.Millisecond * 100)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Stopping worker pool...")
	pool.Stop()

	fmt.Println("All tasks done")
}
```
