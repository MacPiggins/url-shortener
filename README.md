# url-shortener


<!-- ToC start -->
# Содержание
1. [Описание задачи](#Описание-задачи)
1. [Реализация](#Реализация)
1. [API](#API)
1. [Сборка и запуск](#Сборка-и-запуск)
1. [Тестирование](#Тестирование)
<!-- ToC end -->

# Описание задачи
Реализовать сервис, предоставляющий API по созданию сокращённых ссылок.

Ссылка должна быть:
- Уникальной; на один оригинальный URL должна ссылаться только одна сокращенная ссылка;
- Длиной 10 символов;
- Из символов латинского алфавита в нижнем и верхнем регистре, цифр и символа _ (подчеркивание).

Сервис должен быть написан на Go и принимать следующие запросы по http:
1. Метод Post, который будет сохранять оригинальный URL в базе и возвращать сокращённый.
2. Метод Get, который будет принимать сокращённый URL и возвращать оригинальный.



Решение должно соответствовать условиям:
- Сервис распространён в виде Docker-образа; 
- В качестве хранилища ожидаем in-memory решение и PostgreSQL. Какое хранилище использовать, указывается параметром при запуске сервиса; 
- Реализованный функционал покрыт Unit-тестами.

# Реализация
### структура проекта
```
.
├── cmd                 	
│   ├── url-shortener       // точка входа в приложение
├── internal
│   ├── app           	    // http сервер
│   ├── database         	// 
│       ├── noDB            // интегрированное хранилище
│       ├── postgres        // взаимодействие с постгрес
│       └── redis           // взаимодействие с редис
│   ├── transport
│       └── rest            // типы и функции для rest архитектуры
├── pkg
│   ├── Base63    	        // кодировка Base62 + "_"
│   └── token_generator     // генератор уникальных токенов
```

# API

### POST /api/create
тело запроса - json формата:
{
    "link": "somelink"
}

пример: 
``` curl -X POST -H "Content-Type: application/json" -d '{"link": "somelink"}' https://address.ru/api/create ```

тело ответа - строка:
"unique token"



### GET /api/retrieve
запрос c query параметром token:

/api/retrieve?token=sometoken

пример: ```curl GET https://address.ru/api/retrive?token=sometoken```

# Сборка и запуск
### Из репозитория
docker build `<name>`

docker run `<name>`

c postgres: docker run `<name>` postgres `<connlink>`

c redis: docker run `<name>` redis `<connlink>`

c integrated storage: docker run
### из docker hub
docker run macpiggins/url-shortener

c postgres: docker run macpiggins/url-shortener postgres `<connlink>`

c redis: docker run macpiggins/url-shortener redis `<connlink>`

c integrated storage: docker run macpiggins/url-shortener

# Тестирование
go test ./...

