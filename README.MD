## Тестовое задание Dvigus

Rate Limmiter && Clickhouse

### Технологии:
- Go
- PostgreSQL
- Docker

### Библиотеки:
- gin-gonic/gin 1.8.2
- swaggo/swag v1.3.3
- sirupsen/logrus 1.9.0
- ilyakaznacheev/cleanEnv 1.4.2
- github.com/patrickmn/go-cache v2.1.0
- github.com/stretchr/testify v1.8.1

### Диаграмма Базы Данных:
![Untitled](https://user-images.githubusercontent.com/65400970/227274734-d2c11376-83d2-4c56-b09b-f12ac8388bde.png)


### Запуск приложения локально (выполнять по порядку!):
- git clone  
- создать свою .env переменную или использовать мою.
- make docker
- make migrateup
- Часть 2 находится в одноименной папке.

### Тестирование:
- make test