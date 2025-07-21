# 👋 "РАССКАЖИ О СЕБЕ"

## **Краткий ответ (1-2 минуты):**

Привет! Я Go разработчик с 3.5 годами опыта. Всё это время работал в одном большом проекте - системе удаленного обновления прошивок для самосвалов КАМАЗ.

**Начинал джуниором** - писал простые CRUD API для управления заданиями на обновление. Изучал основы: gRPC, PostgreSQL, как правильно структурировать код.

**Через полтора года** перешел на более сложные задачи - обработку высоконагруженного потока событий от устройств через Kafka. Здесь уже разбирался с производительностью, идемпотентностью, отказоустойчивостью. Система обрабатывала более 1000 сообщений в минуту от 500+ самосвалов в карьерах.

**К концу проекта** делал уже самостоятельные задачи - аналитику, отчеты, интеграции с внешними системами. Построил ETL пайплайны для генерации PDF отчетов для менеджмента КАМАЗ.

**Технически выросс от джуна до мида** - изучил event-driven архитектуру, микросервисы, решал сложные задачи с distributed systems. Был несколько серьезных факапов в продакшене, но каждый научил чему-то важному.

**Проект завершился**, команда расформировалась. Ищу новые вызовы где смогу применить опыт с Kafka, PostgreSQL и Go, но уже в другой предметной области. Хочется поработать в команде где есть сениоры, у которых можно учиться дальше.

---

## **Если попросят подробнее про технологии:**

_"Основной стек - Go, gRPC, PostgreSQL, Kafka. Работал с event-driven архитектурой, outbox pattern для гарантированной доставки событий, optimistic locking для concurrent updates. Настраивал Kafka consumer groups для high-throughput обработки. Использовал Redis для кеширования и идемпотентности. Мониторинг через Prometheus и Grafana."_

## **Если спросят про самое сложное:**

_"Наверное, настройка производительной обработки событий в Kafka. Нужно было правильно разбить на партиции, настроить consumer groups, обеспечить идемпотентность и при этом не потерять ни одного сообщения. Плюс отладка deadlock'а в production - там пришлось глубоко разбираться с горутинами и profiling'ом."_

## **Мотивация для новой работы:**

_"Хочется расти дальше, изучать новые технологии и паттерны. В текущем проекте освоил event-driven подход, но интересно попробовать другие архитектурные решения. Плюс хочется поработать в команде где есть опытные разработчики - всегда есть чему учиться."_

# **Команда:**

- **TeamLead & Архитектор:** Общая архитектура, ключевые решения, координация.
- **PM:** Управление задачами, сроками, коммуникация с заказчиком.
- **Embedded-инженер:** Работа с CAN-шиной, прошивки MCU, шлюз на Jetson (взаимодействие с бортом).
- **5 Backend-разработчиков (Go):** Разработка сервисов бэкенда. Я был одним из них, пришел Junior, к концу проекта - уверенный Mid.
- **DevOps:** Инфраструктура (Docker, k8s?), CI/CD, мониторинг (Prometheus/Grafana), управление MinIO/Kafka.

# 🏗️ АРХИТЕКТУРА КАМАЗ OTA СИСТЕМЫ

## 📋 **ОБЩАЯ СТРУКТУРА ПО СЛОЯМ**

### **External Users**

- **Операторы КАМАЗ** - основные пользователи системы
- **Системные API** - внешние интеграции (ERP, мониторинг автопарка)
- **Инженеры** - техническая поддержка и разработка

### **UI Layer**

- **Web UI Dashboard** - панель управления для операторов

### **API Layer**

- **API Gateway** - аутентификация, авторизация, rate limiting, load balancing
- **Redis (Sessions)** - JWT токены, пользовательские сессии

### **Business Logic** (3 кластера)

- **Job Management Cluster** 🔥
- **Device Communication Cluster** 🔥
- **Analytics & Testing Cluster**

### **Message Bus**

- **Apache Kafka** - event-driven архитектура между всеми сервисами

### **Data Layer**

- **Shared Data Infrastructure** - PostgreSQL, MinIO, мониторинг, etc.

### **External Systems**

- **MQTT Broker** - связь с устройствами
- **КАМАЗ Самосвалы** - 500+ устройств в карьерах
- **КАМАЗ ERP** - корпоративные системы

---

## 🔥 **JOB MANAGEMENT CLUSTER** (мой основной фокус)

### **Job Service**

- **Задача:** Центральная бизнес-логика управления заданиями на обновление
- **Функции:** CRUD API, валидация правил, управление жизненным циклом заданий
- **Технологии:** Go, gRPC, PostgreSQL (pgx)
- **Паттерны:** Outbox pattern, Optimistic locking
- **Мой вклад:** 🔥 Основной разработчик (Junior → Mid)

### **Scheduler**

- **Задача:** Планирование выполнения заданий по времени
- **Функции:** Cron expressions, timezone handling, массовые обновления
- **Технологии:** Go, Redis для кеша расписаний
- **Мой вклад:** Интеграция с Kafka, timezone handling

### **PostgreSQL Primary**

- **Задача:** Основное хранилище данных кластера
- **Данные:** jobs, devices, firmwares, event_outbox
- **Особенности:** ACID транзакции, миграции golang-migrate

### **Outbox Worker**

- **Задача:** Гарантированная публикация событий в Kafka
- **Функции:** Чтение outbox таблицы, публикация, retry логика
- **Паттерн:** Transactional outbox для exactly-once семантики

### **Redis (Schedule Cache)**

- **Задача:** Кеширование расписаний и временных меток
- **Функции:** Быстрый доступ к cron data, TTL управление

---

## 🔥 **DEVICE COMMUNICATION CLUSTER** (мой основной фокус)

### **Dispatcher**

- **Задача:** Отправка команд обновления устройствам
- **Функции:** Генерация команд, MQTT publishing, retry логика
- **Технологии:** Go, MQTT client, MinIO integration
- **Особенности:** Circuit breaker, presigned URLs

### **Event Processor**

- **Задача:** Высокопроизводительная обработка статусов от устройств
- **Функции:** Kafka consumer, парсинг сообщений, обновление статусов
- **Технологии:** Go, Kafka (sarama), Redis для идемпотентности
- **Производительность:** 1000+ сообщений/минуту, 50 партиций
- **Мой вклад:** 🔥 Основной разработчик (Этап 2)

### **MinIO Storage**

- **Задача:** S3-совместимое хранилище прошивок
- **Функции:** Бинарники firmware, presigned URLs, версионирование
- **Интеграция:** С Dispatcher для генерации ссылок скачивания

### **Redis (Message Cache)**

- **Задача:** Кеширование для идемпотентности обработки
- **Функции:** Дедубликация сообщений, временное хранение статусов

### **Kafka Consumer Groups**

- **command-dispatcher-group:** 5 инстансов для команд
- **status-processor-group:** 10 инстансов для статусов (высокая нагрузка)
- **telemetry-collector-group:** 3 инстанса для телеметрии

---

## 📊 **ANALYTICS & TESTING CLUSTER** (Этап 3)

### **Reporting Service**

- **Задача:** Аналитика и генерация отчетов
- **Функции:** ETL процессы, PDF/CSV generation, REST API
- **Технологии:** Go, PostgreSQL read replica, библиотеки для PDF
- **Мой вклад:** Схема таблиц, ETL, экспорт отчетов

### **Device Simulator**

- **Задача:** Эмуляция устройств для тестирования
- **Функции:** Virtual device behavior, нагрузочное тестирование, edge cases
- **Возможности:** До 1000 виртуальных устройств на инстанс
- **Мой вклад:** Разработка симулятора (Этап 3)

### **PostgreSQL Read Replica**

- **Задача:** Отдельная база для аналитических запросов
- **Функции:** Тяжелые аналитические запросы без влияния на основную систему
- **Репликация:** Асинхронная от primary

### **ETL Process**

- **Задача:** Обработка потока событий для аналитики
- **Функции:** Агрегация данных, материализованные представления
- **Технологии:** Go goroutines, worker pools

### **Redis (Report Cache)**

- **Задача:** Кеширование тяжелых аналитических запросов
- **Функции:** Cache invalidation, TTL для отчетов

---

## 🔄 **APACHE KAFKA MESSAGE BUS**

### **Топики:**

- **job.events** - 3 партиции, 7 дней (created, updated, completed)
- **device.commands** - 20 партиций, 24 часа (команды устройствам)
- **device.status** - 50 партиций, 72 часа (статусы от устройств)
- **device.telemetry** - 30 партиций, 24 часа (мониторинг устройств)

### **Dead Letter Queues:**

- **job.dlq** - проблемные события заданий
- **device.status.dlq** - некорректные статусы от устройств

### **Партиционирование:**

- **По device_id** для ordered processing статусов
- **По event_type** для параллельной обработки событий

---

## 🗄️ **SHARED DATA INFRASTRUCTURE**

### **Мониторинг:**

- **Prometheus** - сбор метрик, alerting rules
- **Grafana** - дашборды, визуализация

### **Логирование:**

- **ELK Stack** - централизованные логи всех сервисов

### **Безопасность:**

- **Vault** - управление секретами и сертификатами

### **Service Discovery:**

- **Consul** - обнаружение сервисов, health checks

### **Load Balancing:**

- **HAProxy** - распределение нагрузки

### **Backup & Archive:**

- **Система резервного копирования** - disaster recovery

---

## 🌐 **ВНЕШНИЕ СИСТЕМЫ**

### **MQTT Broker**

- **Задача:** Связь с устройствами в карьерах
- **Протокол:** MQTT v3.1.1, QoS=1
- **Топики:** commands/{device_id}/ota, status/{device_id}/update

### **КАМАЗ Самосвалы**

- **Количество:** 500+ устройств
- **Платформы:** STM32, ESP32, Jetson модули
- **Условия:** Нестабильная связь, удаленные карьеры

### **КАМАЗ ERP**

- **Задача:** Интеграция с корпоративными системами
- **API:** Fleet management, maintenance scheduling
- **Протокол:** REST API, OAuth2 аутентификация

---

## 🎯 **МОЯ РОЛЬ И ПРОГРЕССИЯ**

### **Этап 1 (1.5 года): Job Management - Junior**

- Job Service CRUD API, базовая интеграция с PostgreSQL

### **Этап 2 (1 год): Device Communication - Junior+/Mid-**

- Event Processor, Kafka integration, производительность

### **Этап 3 (1 год): Analytics & Testing - Mid**

- Reporting Service, Device Simulator, самостоятельные задачи

### **Технологический стек:**

- **Backend:** Go, gRPC, REST API
- **Базы данных:** PostgreSQL, Redis
- **Очереди:** Apache Kafka (sarama)
- **Хранилище:** MinIO (S3-compatible)
- **Мониторинг:** Prometheus, Grafana
- **Инфраструктура:** Docker, Kubernetes

# 🔄 WORKFLOW ОБНОВЛЕНИЯ ПРОШИВКИ

## **1. Создание задания**

```
Оператор (UI) → API Gateway → Job Service
```

- **Запрос:** `POST /jobs` (device_id, firmware_version, schedule_time)
- **Job Service:** Валидация → сохранение в PostgreSQL + event_outbox
- **Outbox Worker:** Читает outbox → публикует `job.created` в Kafka

## **2. Планирование**

```
Kafka (job.events) → Scheduler
```

- **Scheduler:** Получает `job.created` → добавляет в расписание
- **По времени:** Scheduler публикует `job.ready_to_execute` в Kafka

## **3. Отправка команды**

```
Kafka (ready_to_execute) → Dispatcher → MQTT → Устройство
```

- **Dispatcher:** Генерирует команду + presigned URL из MinIO
- **MQTT:** Отправка на `commands/{device_id}/ota`
- **Устройство:** Получает команду, начинает обновление

## **4. Выполнение обновления**

```
Устройство → MQTT → Kafka → Event Processor → Job Service
```

- **Устройство шлет статусы:** `downloading` → `flashing` → `done`/`error`
- **Event Processor:** Обрабатывает статусы → обновляет задание через gRPC
- **Job Service:** Финальный статус в PostgreSQL

## **5. Завершение**

```
Job Service → Kafka (job.completed) → Reporting Service
```

- **Reporting:** Получает событие завершения → обновляет аналитику
- **Оператор:** Видит результат в UI через API Gateway

---

**Время выполнения:** 15-60 минут в зависимости от размера прошивки и качества связи.

**Точки отказа:** MQTT связь, качество сети в карьере, совместимость прошивки.

# 🔥 СЛОЖНЫЕ ТЕХНИЧЕСКИЕ ЗАДАЧИ

## **1. Outbox Pattern - детально**

### **Проблема (Dual Write Problem):**

Когда создается задание, нужно сделать две операции:

1. Сохранить задание в PostgreSQL
2. Отправить событие `job.created` в Kafka

**Что может пойти не так:**

- Сохранили в базу → упали до отправки в Kafka = событие потеряно
- Отправили в Kafka → упали до записи в базу = другие сервисы получили событие о несуществующем задании
- Частичные сбои: база доступна, Kafka недоступна

### **Классические плохие решения:**

- **Сначала Kafka, потом база** - риск phantom events
- **Сначала база, потом Kafka** - риск lost events
- **Distributed transactions (2PC)** - сложно, медленно, Kafka не поддерживает

### **Outbox Pattern - правильное решение:**

**Как работает:**

1. **В одной транзакции:** сохраняем задание в таблицу `jobs` + событие в таблицу `event_outbox`
2. **Отдельный фоновый процесс (Outbox Worker):** читает `event_outbox`, отправляет в Kafka, помечает как "отправлено"
3. **Периодическая очистка:** удаляем старые отправленные события

**Ключевые принципы:**

- **Atomicity через единую транзакцию** - либо все сохранилось, либо ничего
- **At-least-once delivery** - событие точно дойдет до Kafka (может дойти дважды)
- **Idempotency** - получатели должны уметь обрабатывать дубликаты
- **Retry logic** - если Kafka недоступна, повторяем попытки

**Что получили:**

- **Гарантированная доставка** - события не теряются никогда
- **Eventual consistency** - событие дойдет, но с небольшой задержкой
- **Resilience** - система работает даже когда Kafka недоступна

### **Мои конкретные решения:**

- **Exponential backoff** для retry - 1s, 2s, 4s, 8s, потом в DLQ
- **Dead Letter Queue** для событий которые не удается отправить долго
- **Batch processing** - отправляем по 10 событий за раз для производительности
- **Monitoring** - метрики: `outbox_pending_total`, `outbox_errors_total`

---

## **2. Optimistic Locking** (кратко)

**Проблема:** Concurrent updates статуса задания от разных устройств
**Решение:** Добавил поле `version`, обновление только если версия совпадает
**Результат:** Race conditions исчезли, корректные финальные статусы

## **3. Deadlock в Event Processor** (кратко)

**Проблема:** Захватывал mutex перед медленными gRPC вызовами, блокировал всю обработку
**Решение:** Вынес обновление метрик за пределы критической секции
**Результат:** Throughput вырос в 5 раз, latency снизился

## **4. Kafka Partitioning Strategy** (кратко)

**Проблема:** Hot partitions, неравномерная нагрузка на consumer'ы
**Решение:** Партиционирование по `hash(device_id)`, 50 партиций для device.status
**Результат:** Равномерное распределение, linear scaling до 50 consumer'ов

## **5. Idempotency в Event Processor** (кратко)

**Проблема:** Дублированные сообщения от устройств при network glitches
**Решение:** Unique constraint на `(job_id, message_id, status)` + Redis кеш
**Результат:** Exactly-once processing семантика без потери производительности

## **6. Circuit Breaker для gRPC** (кратко)

**Проблема:** Когда Job Service недоступен, Event Processor роняет все сообщения
**Решение:** Circuit breaker pattern - при ошибках переходим в fallback режим
**Результат:** Graceful degradation, сообщения накапливаются в очереди для retry

## **7. Consumer Group Rebalancing** (кратко)

**Проблема:** При deploy constant rebalancing убивал processing
**Решение:** Увеличил session timeout, graceful shutdown, sticky partition assignment
**Результат:** Stable consumer group, minimal downtime при deploy

---

## **🎯 Главное на собесе:**

**Не код, а понимание:**

- **Зачем** нужен outbox pattern
- **Какие проблемы** он решает
- **Как работает** на высоком уровне
- **Какие trade-offs** - задержка vs consistency

**Outbox - это про:**

- Понимание distributed systems проблем
- Знание паттернов для их решения
- Опыт работы с event-driven архитектурой
- Production-ready мышление

Именно это ценят на собесах, а не знание конкретного SQL синтаксиса!

# 💥 ФАКАПЫ И ИХ РЕШЕНИЯ

## **1. Deadlock в Event Processor - детально**

### **Что произошло:**

Обычный вечер, Event Processor обрабатывал статусы от устройств. Внезапно **throughput упал до нуля**. Сервис живой, логи пишутся, но **сообщения не обрабатываются**. Consumer lag растет: 100... 500... 1000 сообщений.

### **Первые действия (паника):**

- Смотрю дашборд Grafana - **все метрики в красной зоне**
- Логи показывают: сообщения читаются из Kafka, но не обрабатываются
- Перезапускаю сервис - помогает на 10 минут, потом повторяется
- **Алерт в Slack**, тимлид подключается к разбору

### **Поиск причины:**

**Симптомы:**

- Processing rate = 0 сообщений/сек (было 50+)
- CPU usage = 100% (обычно 20-30%)
- Memory растет медленно (не memory leak)
- **Горутины висят** - их количество растет

**Debugging:**

- **pprof профиль** показал: 50+ горутин ждут один mutex
- Одна горутина держит mutex и **ждет network call**
- Это gRPC вызов в Job Service для обновления статуса

### **Root Cause Analysis:**

**Плохой код:**

```go
func (p *EventProcessor) processMessage(msg) {
    p.metricsMutex.Lock()  // ❌ ОШИБКА: lock ДО network call

    // Медленный network call - 100-500ms
    err := p.jobServiceClient.UpdateStatus(...)

    // Обновление метрик
    p.updateMetrics(...)

    p.metricsMutex.Unlock()
}
```

**Что происходило:**

1. Горутина A захватывает mutex
2. Идет делать gRPC call (медленно: 200ms)
3. Горутины B, C, D... встают в очередь на mutex
4. При высокой нагрузке все 50 горутин блокируются
5. **Полный deadlock системы**

### **Немедленное решение (hotfix):**

```go
func (p *EventProcessor) processMessage(msg) {
    // Сначала network call БЕЗ lock
    err := p.jobServiceClient.UpdateStatus(...)

    // ПОТОМ быстрое обновление метрик
    p.metricsMutex.Lock()
    p.updateMetrics(...)  // 1-2ms максимум
    p.metricsMutex.Unlock()
}
```

### **Долгосрочные улучшения:**

1. **Убрал mutex для метрик** - Prometheus client thread-safe для основных операций
2. **Atomic counters** для простых метрик
3. **Connection pooling** для gRPC - переиспользование соединений
4. **Circuit breaker** - если Job Service медленный, не блокируем обработку
5. **Async metrics** - отправляем метрики в отдельном канале

### **Мониторинг для предотвращения:**

- **Goroutine count** - алерт если >100
- **Mutex contention** - pprof в production
- **Processing latency P99** - алерт если >1 секунда
- **gRPC call duration** - алерт если >500ms

### **Что узнал:**

- **Никогда не держи lock во время network calls**
- **Critical sections должны быть максимально короткими**
- **Production profiling обязателен** для Go сервисов
- **Load testing** не покрыл этот сценарий - нужны **stress tests**

### **Результат:**

- **Throughput:** 0 → 200+ сообщений/сек
- **Latency P99:** ∞ → 50ms
- **CPU usage:** 100% → 15-25%
- **Система стабильна** уже полгода после фикса

---

## **2. Memory leak в production** (кратко)

**Что:** Event Processor жрал память, через 2 часа OOM kill
**Причина:** Map для idempotency рос бесконечно + goroutine leaks
**Решение:** LRU cache + proper goroutine cleanup + memory monitoring
**Урок:** Always profile memory usage в production

## **3. Connection pool exhaustion** (кратко)

**Что:** Job Service отдавал 500-ки "too many connections"  
**Причина:** Не закрывали transactions в error cases
**Решение:** defer tx.Rollback() + query timeouts + pool tuning
**Урок:** Proper resource cleanup критичен в Go

## **4. Kafka consumer group rebalancing storm** (кратко)

**Что:** После deploy constant rebalancing, processing stopped
**Причина:** Aggressive rolling deployment + короткие timeouts
**Решение:** Увеличили session timeout + graceful shutdown + blue-green deploy
**Урок:** Kafka требует аккуратного деплоя и настройки таймаутов

## **5. PostgreSQL query timeout cascade** (кратко)

**Что:** Один медленный запрос заблокировал всю базу
**Причина:** Long-running analytical query без LIMIT
**Решение:** Query timeouts + read replica для analytics + connection limits
**Урок:** Изолировать OLTP и OLAP нагрузку

## **6. Redis connection leak** (кратко)

**Что:** Event Processor переставал работать с кешем
**Причина:** Не закрывали Redis connections в error paths
**Решение:** Connection pooling + proper error handling
**Урок:** All external resources need proper lifecycle management

## **7. MQTT broker overload** (кратко)

**Что:** Устройства не получали команды обновления
**Причина:** Отправляли команды всем 500 устройствам одновременно
**Решение:** Rate limiting + batch sending + queue management
**Урок:** External systems have limits, respect them

---

## **🎯 Главное про факапы на собесе:**

**Показывает:**

- **Production experience** - работал с реальными проблемами
- **Troubleshooting skills** - умею искать причины
- **System thinking** - понимаю как компоненты взаимодействуют
- **Learning from mistakes** - анализирую и улучшаю

**Не стесняйся рассказывать факапы!**

- Все делают ошибки в production
- Важно как ты их исправляешь и что выносишь
- Это показывает **реальный опыт**, а не теорию

**Структура рассказа:**

1. **Что случилось** (симптомы)
2. **Как искал причину** (debugging process)
3. **Root cause** (техническая причина)
4. **Как исправил** (immediate + long-term)
5. **Что узнал** (lessons learned)
