Примеры решений
Пример для задачи 1:

```go
func createUsersTable(db *sql.DB) error {
    query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100),
        email VARCHAR(100)
    )`
    _, err := db.Exec(query)
    return err
}
```

Пример для задачи 6:

```go
func createUser(db *sql.DB, name, email string) (int, error) {
    query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
    var id int
    err := db.QueryRow(query, name, email).Scan(&id)
    return id, err
}
```

Пример для задачи 11:

```go
func getAllUsers(db *sql.DB) ([]User, error) {
    query := `SELECT id, name, email FROM users`
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var u User
        err := rows.Scan(&u.ID, &u.Name, &u.Email)
        if err != nil {
            return nil, err
        }
        users = append(users, u)
    }
    return users, rows.Err()
}
```

Структуры для работы

```go
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

type Product struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Price     float64   `json:"price"`
    CreatedAt time.Time `json:"created_at"`
}

type Order struct {
    ID     int     `json:"id"`
    UserID int     `json:"user_id"`
    Total  float64 `json:"total"`
    Status string  `json:"status"`
}
```

Полезные советы

Всегда проверяй ошибки при работе с базой данных
Используй подготовленные запросы ($1, $2) для защиты от SQL-инъекций
Закрывай rows с помощью defer rows.Close()
Используй транзакции для критических операций
Не забывай про контекст - используй context.Context в продакшн коде

# Настройка PostgreSQL для тренажера

## Способ 1: Через командную строку (psql)

### Шаг 1: Подключись к PostgreSQL

```bash
# Подключение к PostgreSQL (обычно под пользователем postgres)
psql -U postgres

# Или если у тебя другой пользователь
psql -U your_username
```

### Шаг 2: Создай базу данных

```sql
-- Создание базы данных
CREATE DATABASE traindb;

-- Переключение на созданную базу
\c traindb

-- Проверка подключения
\conninfo
```

### Шаг 3: Создай пользователя (опционально)

```sql
-- Создание пользователя для приложения
CREATE USER trainuser WITH PASSWORD 'trainpass';

-- Предоставление прав на базу данных
GRANT ALL PRIVILEGES ON DATABASE traindb TO trainuser;

-- Выход из psql
\q
```

---

## Способ 2: Через Docker (рекомендуемый)

### Создай docker-compose.yml

```yaml
version: "3.8"
services:
  postgres:
    image: postgres:15
    environment:
      POSTGRES_DB: traindb
      POSTGRES_USER: trainuser
      POSTGRES_PASSWORD: trainpass
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data

volumes:
  postgres_data:
```

### Запуск

```bash
# Запуск контейнера
docker-compose up -d

# Проверка, что контейнер запущен
docker-compose ps

# Подключение к базе для проверки
docker-compose exec postgres psql -U trainuser -d traindb
```

---

## Способ 3: Через VSCode расширения

### Установи расширение PostgreSQL

1. Открой VSCode
2. Перейди в Extensions (Ctrl+Shift+X)
3. Найди и установи "PostgreSQL" от ckolkman

### Подключение к базе

1. Нажми F1 → "PostgreSQL: New Connection"
2. Введи данные подключения:
   - Host: localhost
   - User: trainuser (или postgres)
   - Password: trainpass (или твой пароль)
   - Port: 5432
   - Database: traindb

### Создание базы через расширение

1. Подключись к серверу PostgreSQL
2. Правый клик → "New Query"
3. Выполни SQL:

```sql
CREATE DATABASE traindb;
```

---

## Способ 4: Готовый скрипт для инициализации

### Создай файл init_db.go

```go
package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

func main() {
    // Подключение к PostgreSQL (без указания базы данных)
    db, err := sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable")
    if err != nil {
        log.Fatal("Ошибка подключения к PostgreSQL:", err)
    }
    defer db.Close()

    // Создание базы данных
    _, err = db.Exec("CREATE DATABASE traindb")
    if err != nil {
        fmt.Println("База данных уже существует или ошибка:", err)
    } else {
        fmt.Println("База данных traindb создана успешно!")
    }

    // Подключение к созданной базе
    trainDB, err := sql.Open("postgres", "user=postgres dbname=traindb sslmode=disable")
    if err != nil {
        log.Fatal("Ошибка подключения к traindb:", err)
    }
    defer trainDB.Close()

    // Тестовое подключение
    err = trainDB.Ping()
    if err != nil {
        log.Fatal("Не удается подключиться к базе:", err)
    }

    fmt.Println("Успешное подключение к базе traindb!")
    fmt.Println("Можно начинать решать задачи!")
}
```

### Запуск скрипта

```bash
go mod init postgres-trainer
go get github.com/lib/pq
go run init_db.go
```

---

## Строки подключения для тренажера

### Для локального PostgreSQL

```go
db, err := sql.Open("postgres", "user=postgres dbname=traindb sslmode=disable")
```

### Для Docker

```go
db, err := sql.Open("postgres", "user=trainuser password=trainpass dbname=traindb host=localhost sslmode=disable")
```

### Для продакшена (с SSL)

```go
db, err := sql.Open("postgres", "user=trainuser password=trainpass dbname=traindb host=localhost sslmode=require")
```

---

## Полезные команды psql

```sql
-- Показать все базы данных
\l

-- Подключиться к базе
\c traindb

-- Показать все таблицы
\dt

-- Показать структуру таблицы
\d table_name

-- Выполнить SQL файл
\i /path/to/file.sql

-- Выход
\q
```

---

## Проверка настройки

### Создай файл test_connection.go

```go
package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

func main() {
    // Строка подключения - измени под свои настройки
    connStr := "user=postgres dbname=traindb sslmode=disable"

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Проверка подключения
    if err = db.Ping(); err != nil {
        log.Fatal("Не удается подключиться к базе:", err)
    }

    fmt.Println("✅ Подключение к PostgreSQL успешно!")
    fmt.Println("🚀 Можно начинать тренировки!")

    // Проверим версию PostgreSQL
    var version string
    err = db.QueryRow("SELECT version()").Scan(&version)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("📊 Версия PostgreSQL: %s\n", version)
}
```

---

## Рекомендации

1. **Для начинающих**: используй Docker - проще и чище
2. **Для продвинутых**: локальная установка PostgreSQL
3. **Для работы в команде**: Docker Compose с общими настройками
4. **Обязательно**: установи расширение PostgreSQL в VSCode для удобства

После настройки можешь сразу переходить к решению задач из тренажера!
