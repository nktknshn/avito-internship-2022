# Avito Internship Task 1

Реализация тестового задания на позицию стажера-бэкендера В авито

https://github.com/avito-tech/internship_backend_2022

- реализованы все требования
- слоистая архитектура
- покрытие тестами 70%
- доменная логика реализована в технике DDD
- GRPC сервер
- cli инструмент
- метрики для prometheus
- трассировка opentrace
- структурное логирование 
- сборщик логов grafana/loki
- интеграционные тесты через dockertest
- утилита для бенчмарка

Что лучше было бы сделать иначе:
- openapi спецификация генерируется при помощи swaggo из аннотаций. Лучше было бы генерировать сервер и клиент из спецификации, а не наоборот.

## Архитектура

Упрощенная диаграмма архитектуры:

![](doc/simple.png)

Более детальная диаграмма архитектуры: [doc/diagram.png](doc/diagram.png)

Использована слоистая гексагональная архитектура из четырех слоев:

- доменная логика
- приложение
- адаптеры
- инфраструктура

### Доменная логика

* [/internal/balance/domain](/internal/balance/domain)

Ядром является доменная логика, выполненная в технике DDD. Для надежности все доменные типы создаются при помощи конструкторов, которые проверяют соотвествие значений бизнес-правилам. Например: баланс пользователя не может быть отрицательным, и для него введен отдельный тип `Amount`, значение которого является приватным и доступно для инициации лишь в рамках доменной логики, а его изменение только через методы, проверяющие валидность диапазона значений [0..math.MaxInt64]. 

```go
// Неотрициательное кол-во копеек
type Amount struct {
	amount int64
}

func (a Amount) Add(b AmountPositive) (Amount, error) {
    // ...
}

func (a Amount) Sub(b AmountPositive) (Amount, error) {
    // ...
}
```

Баланс пользователя представлен value-object типом `AccountBalance`, для которого также реализован ряд методов, реализующих перемещение средств между доступным и резервированным балансами, и возвращающих доменную ошибку при попытке некорректной операции.

```go
type AccountBalance struct {
    available amount.Amount
	reserved  amount.Amount
}

func (ac AccountBalance) Deposit(a amount.AmountPositive) (AccountBalance, error) {
    // ...
}
func (ac AccountBalance) Reserve(a amount.AmountPositive) (AccountBalance, error) {
    // ...
}
func (ac AccountBalance) Withdraw(a amount.AmountPositive) (AccountBalance, error) {
    // ...
}
// ...
```

Баланс пользователя является частью сущности Account, для которой описаны методы, изменяющие её состояние через pointer receiver, и соответствующие требуемым юзкейсам: Deposit, Reserve, ReserveCancel, ReserveConfirm, Transfer.

Транзакции представлены тремя сущностями: TransactionDeposit (зачисление на счет), TransactionSpend (оплата продукта) и TransactionTransfer (перевод между пользователями). Статус операции оплаты продукта определяется наличием транзакции Spend со статусом, отличным от Canceled: Reserved, если деньги зарезервированы; Confirmed, если продукт оплачен. 

В доменной зоне также описаны интерфейсы репозиториев для каждой из сущностей: Account и Transactions.

### Приложение

* [/internal/balance/app](/internal/balance/app)
* [/internal/balance/app_impl](/internal/balance/app_impl)

В этом слое реализованы сценарии использования (use-cases). Сценарий должен реализовывать один из двух интерфейсов:

```go
// сценарий не возвращает результат, например Deposit
type UseCase0[In any] interface {
	Handle(ctx context.Context, in In) error
    // имя сценария используется для логирования, метрик и для разграничения прав 
    // доступа пользователей микросервиса
	GetName() string
}

// сценарий возвращает результат, например GetBalance
type UseCase1[In any, Out any] interface {
	Handle(ctx context.Context, in In) (Out, error)
	GetName() string
}
```

Сценарий характеризуется набором зависимостей, которые передаются конструктору `New`, типом запроса `In` и опциональным типом результата `Out`. 

Так как сценарию может потребоваться изменение как баланса пользователя, так и списка его транзакций, необходимо предусмотреть атомарность изменения в двух репозиториях. Для этого применятся библиотека [go-transaction-manager](github.com/avito-tech/go-transaction-manager), которая позволяет сделать это при помощи высокоуровненего стабильного API, сохраняющего слой приложения от низкоуровневых деталей реализации репозиториев. 

Для этого сценарию, использующему репозиторий передается в качестве зависимостей объект типа `trm.Manager`

```go
// Manager manages a transaction from Begin to Commit or Rollback.
type Manager interface {
	// Do processes a transaction inside a closure.
	Do(context.Context, func(ctx context.Context) error) error
	// DoWithSettings processes a transaction inside a closure with custom trm.Settings.
	DoWithSettings(context.Context, Settings, func(ctx context.Context) error) error
}
```

Это также позволит использвать данный сценарий в других сценариях в рамках одной транзакции ([Статья на хабре](https://habr.com/ru/companies/avito/articles/727168/)), если это понадобится.

#### Декораторы use-case

То, что все сценарии использования реализуют единый интрефейс, позволяет добавить в архитектуру слоя приложения такую функцию, как декораторы, которые по функциональности представляют собой аналог middleware в других областях.

Например, декоратор логирования:

```go
type Decorator0Logging[T any] struct {
	base   UseCase0Handler[T]
	logger logging.Logger
}

func (d *Decorator0Logging[T]) Handle(ctx context.Context, in T) (err error) {
	defer func() {
		if err != nil {
			d.logger.Error(d.base.GetName(), "use_case", d.base.GetName(), "error", err)
		}
	}()
	d.logger.Info(d.base.GetName(), "use_case", d.base.GetName())
	return d.base.Handle(ctx, in)
}
```

В данном микросервисе реализованы декораторы: логирование, метрики, panic-recover.

## Запуск

### Тесты

```bash
make cover

make cover-html
```

## Использованные инструменты

- [go-transaction-manager](github.com/avito-tech/go-transaction-manager) для кросс-репозиторных транзакций
- [chi]()
- [sqlx](https://github.com/jmoiron/sqlx) обертка для упрощения работы с базой данных
- [pgx](https://github.com/jackc/pgx) для работы с базой данных
- [cobra](https://github.com/spf13/cobra) для cli инструмента
- [swaggo](https://github.com/swaggo/swag) генерация openapi спецификации из аннотаций
- [openapitools/openapi-generator](https://github.com/openapitools/openapi-generator) генерация http-клиента для утилиты бенчмарка
- [testify](https://github.com/stretchr/testify) для тестирования
- [dockertest](https://github.com/ory/dockertest) интеграционные тесты с базой данных
- [reflex](https://github.com/cespare/reflex) для перекомпиляции приложения при изменении кода
- [go-ergo-handler](https://github.com/nktknshn/go-ergo-handler) для создания http-хэндлеров
- [docker](https://www.docker.com/) Docker
- [prometheus](https://prometheus.io/) для сбора метрик
<!-- - [opentracing](https://opentracing.io/) для трассировки -->
- [grafana](https://grafana.com/) для визуализации метрик и логов
- [loki](https://grafana.com/oss/loki/) для сбора логов
