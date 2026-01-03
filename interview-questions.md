# Как я изучал Go

- Официальный [a tour of go](https://go.dev/tour/welcome/1) 
(плохие примеры примерно с трети, идиотские запутанные математические примеры, но советую прорваться через него в любом случае)

- Далее с чатом гпт читаю [go roadmap](https://roadmap.sh/golang)

- Смотрю базу красавчика:
    - Базовые вопросы
    - Интересные техсобесы, выписываю вопросы с них

## Уточнения к базовым вопросам красавчика

- What are `/internal`, `/pkg`, `cmd` ? Which one is setup, wiring ? Which one is main business logic: handlers, services, interfaces, adapters, repo ?  
- What are the import rules for stuff inside `/internal` ?
    - Can/should you import `/internal/user/service` from `/internal/user/repo` ?
    - Can/should you import `/internal/one` from `/internal/two` ?
    - Can/should you import `/internal/one` from `/pkg` ?
    - Can/should you import `/internal/one` from `/cmd` ?
- What exactly is `init()` where does it run and how many times ?
- Are Both `sync.Map` and `context` type safe ? 
- What is `checksum` ?
- Basic auth methods and ways to handle security problems they produce.
    - CSRF for cookies 
    - XSS for JWT tokens
- `Distributed transaction` (protect correctness or distributed actions and rollback correctly in case of a fail)
- `distributed lock` (who can act now)
- What is the difference between the two above ?  
Easy real-world example: lets say 10 users come and want to book 1 last room.  
How does the system behave in this case ?  
What part of this behavior is distributed transaction and which one is distributed lock?

## Вопросы с собесов красавчика 

Questions from interviews *(may contain duplicates with basic questions)*

- var wg sync.WaitGroup - что такое wait group ? Зачем нужны wg.Add(), wg.Done(), wg.Wait()
- Clean architecture / Hexagonal architecture. Are they the same ?
    - Handler - ?
    - Services - ?
    - Interface (Port) -> 
    - Adapter - ?
    - Repository - ?
- Layer of architecture - logical grouping with dependency direction
- Kafka: Topics, Partitions, Consumers, Consumer groups. Where is logic handled ? Consumer side ? Producer side ? Kafka itself ? What is idempotent producer ?
- Transactional outbox / Transactional inbox
- Clouds / DevOPS terms, other DBs and Queues on a high level: EventBridge, SQS, NATS, DynamoDB, BigQuery, Memcached, Cassandra, Terraform, Lambda and Edge-functions
- Memcached vs Redis. Which commands have you used apart from SET / GET / DEL ?
- gRPC, gRPC streams.
- Does gRPC care about field names?
    - Can you change field names or types ?
    - Will it break gRPC compatibility ?
    - Why if you have a gRPC gateway for using gRPC via REST, will it break if you change field names ?
- Redis cluster, sentinel - what are those ?
- Types of testing: Unit, Integration, E2E, load, smoke
- Goroutines:
    - GMP Model
    Scheduler
    stealing.
    States of a goroutine: executing, blocked, waiting in a queue.
    Why does go need GMP to be efficient ?
    Are goroutines executed on P or M ? What is P ?
    What is the difference between concurrency and parallelism ?
    What is GOMAXPROCS. What is the default value of it ?
    Can we set it to 1 on a 8 core machine ?
    runtime.NumCPU() - what is it (self descriptive)
    Deadlock, Lovelock
    Why are goroutines cheap to create, switch ?
    (Fact) goroutines start with 2kb, grow 250mb(32bit), 1GB(64bit)
- Mutex vs Atomic. Which one is better ? Can we compare them ?
- Why does error interface in go often contains a pointer to an underlying error value ? Why not just return the struct instead of pointer to it ?
- What is IPC ? Can threads and processes share memory and access each other’s variables ? Where do threads run ?Can I open file in one thread and pass the handler to another ? Does it apply to system threads or go routines ?
- Synchronisation objects in go. Channels, Mutex, WaitGroup, Atomic. How would you download files in parallel and count how many finished, what wi.l be used for it - is Mutex needed in this case ?
- (Under the hood question) What is the difference between cooperative and preemptive scheduling ? How does modern go operate ?
- Error handling in Go.
    - What is go philosophy about errors ?
    - Do we get a stack trace by default with panic ?
    - Do we have it when normal error happens ?
- What is return error “up the stack” ?
    - What is panic and where is it recovered ?
    - Wrap error, sentinel error, typed error - ? 
- Smart lock: many goroutines reading, but one writing. sync.Mutex vs sync.RWMutex
- 3 pillars of observability
    * Basic setup: OTEL, OTEL Collector
    * Metrics: Prometheus, Grafana UI
    * Traces: Jaeger
    * Logs: Loki
- Do you need separate libs to handle errors in Go ? (No)
- What do you see if you print a nil pointer ? <nil>
- Are nil pointers of different values the same ? (No)
- Pointer size ? (Basically a “Machine word” 4bytes in 32, 8 bytes in 64)
- Why are optional fields often represented with pointers ? (In order to distinguish 0, nil)
- (Geek stuff) size of struct with one float32.
- Context in go. What is it used for ? Context hierarchy. context.WithValue, context.WithTimeout
- Is it fine to inject dependencies via context.WithValue ? Do you have to type-assert when getting values from context.WithValue ?
- Go Garbage Collector (GC)
    - Does it stop-the-world ? Is this a problem ?
    - Is GC concurrent ?
    - Mark-and-sweep
- encoding/json, reflect, easyjson.
- Generics vs reflect. Before people used interface{} + reflect, now generics exist. Generics - compile time, reflect - runtime (slower)
- Interfaces vs generic. (Interfaces - accept any type with certain behaviour, generic - same algorithm for many concrete types)
- Fucking SOLID
- Circular imports in go
    1. High level code depends on lower level code, never vice-versa. Keep domain with types only
    2. Types are low-level, behaviour is high level
    3. Interfaces are on the consumer side. A service is the consumer of a port/interface, and adapters are plugged in to satisfy that port.
    4. If packages need each other - extract to a shared one
    5. God package everyone imports - bad
- OOP in go.
    - Does go have a special constructor ?
    - Does go have private/public ?
    - How do you define getters for private fields ?
    - Suppose you have a structure with members[]string inside. How do you return members immutably to the outside world ?
- Memory leaks in go (something you don’t need but keep running or keep a reference of)
    - goroutines running web socket
    - handlers
    retry loops
- Let’s say we have tickets and segments. We have multiple tickets for each segment.  
What would be better choice for shard key: ticket_id or segment_id ? Why ?
- We have two services: A, B. A has a SLA (<5 sec). A->B and needs to obtain data like (does user exist ?). If service B starts to take longer time, it affects A. Tight coupling, how to handle that ?
