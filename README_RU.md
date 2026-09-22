# Vetmessager

Backend для [Vetmessager](https://github.com/vetalana28/svetmessager).

## Установка

### Клонирование репозитория

```bash
git clone https://github.com/maximyunak/vetmessager.git
cd vetmessager
```

### Запуск PostgreSQL

```bash
make env-up
```

### Запуск миграций базы данных

```bash
make migrate-up
```

### Запуск приложения в режиме разработки

```bash
make app-run
```

Приложение будет доступно по адресу:

```text
http://localhost:5050
```

### Сборка и запуск с помощью Docker

```bash
make app-deploy
```

## API

### Базовый URL

```text
http://localhost:5050/api/{version}/{route}
```

Например:

```text
http://localhost:5050/api/v1/users
```

### Swagger

Документация API Swagger доступна по адресу:

```text
http://localhost:5050/swagger
```

## Структура проекта

```text
.
├── cmd/                         # Точки входа приложения
│   └── vetmessager/
│       ├── main.go              # Точка входа приложения
│       └── Dockerfile           # Конфигурация Docker-образа
│
├── docs/                        # Документация API Swagger
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
│
├── internal/                    # Внутренний код приложения
│   ├── core/                    # Общие компоненты приложения
│   │   ├── auth/                # JWT-аутентификация и управление токенами
│   │   ├── domain/              # Доменные модели
│   │   ├── errors/              # Общие ошибки приложения
│   │   ├── logger/              # Настройка и реализация логирования
│   │   ├── password/            # Хеширование и проверка паролей
│   │   ├── repository/          # Общая инфраструктура репозиториев
│   │   │   └── postgres/
│   │   │       └── pull/        # Подключение и конфигурация PostgreSQL
│   │   └── transport/           # Общая HTTP-инфраструктура
│   │       └── http/
│   │           ├── middleware/  # HTTP middleware
│   │           ├── request/     # Декодирование HTTP-запросов
│   │           ├── response/    # Обработка HTTP-ответов
│   │           ├── server/      # HTTP-сервер и маршрутизация
│   │           └── utils/       # HTTP-вспомогательные функции
│   │
│   └── features/                # Функциональные модули приложения
│       └── users/               # Функциональность пользователей
│           ├── repository/      # Работа с данными пользователей
│           │   └── postgres/    # Реализация для PostgreSQL
│           ├── service/         # Бизнес-логика пользователей
│           └── transport/       # HTTP API пользователей
│               └── http/
│
├── migrations/                  # Миграции базы данных
│   ├── 000001_init.up.sql
│   └── 000001_init.down.sql
│
├── Makefile                     # Команды для разработки и деплоя
├── docker-compose.yml           # Конфигурация Docker-инфраструктуры
├── go.mod                       # Модуль Go и зависимости
└── go.sum                       # Контрольные суммы зависимостей
```
