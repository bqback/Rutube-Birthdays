# Установка

Для пользования прокси надо сгенерировать .env-файл

## Скрипт

`sudo ./install.sh postgres postgres postgres BirthdayStorage birthday-db`

Докер-контейнер с БД и беком также собирается и поднимается скриптом

## Вручную

1. Создать файл `.env` в папке config

`touch config/.env`

2. Добавить туда следующий текст

```
MIGRATOR_PASSWORD="postgres"
POSTGRES_USER="postgres"
POSTGRES_PASSWORD="postgres"
POSTGRES_DB="BirthdayStorage"
POSTGRES_HOST="birthday-db"
```

3. Поднять докер

`docker-compose up --build --detach`
