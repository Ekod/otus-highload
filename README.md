# highload-otus-server

**Запуск сервиса**
1. Убедиться, что установлен Docker-Compose
2. Из корневой директории проекта запустить 
   ```shell 
   make up
   ```
3. Подождать секунд 5 пока докер всё подтянет и запустит
4. Запустить миграции
   ```shell 
   make migrate
   ```
5. Запустить сервис
   ```shell 
   make run
   ```

**Миграции**

```shell
# Установка библиотеки миграции
brew install golang-migrate
# Создание миграции
migrate create -ext sql -dir zarf/migrations <migration_name_change_me>
# Применение миграции (на локальную БД)
make migrate
```

**Остановка сервиса**
1. Из корневой директории проекта запустить
 ```shell 
   make down
 ```