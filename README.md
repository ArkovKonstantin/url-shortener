### Запуск приложения
```
$ git clone https://github.com/ArkovKonstantin/chat-server.git
$ cd chat-server
$ make build
$ docker-compose up
```
Приложение будет доступно по адресу http://localhost:9000

API методы
---
#### Создание укороченной ссылки
_Request_
```
curl --request POST \
  --url http://localhost:9000/shorten \
  --header 'content-type: application/json' \
  --data '{
	"url": "https://mail.ru"
    }'
```
_Responses:_
* `HTTP/1.1 201 Created`
* `HTTP/1.1 400 Bad Request`
  * invalid URI for request
* `HTTP/1.1 500 Internal Server Error`

#### Создание укороченной ссылки c кастомным именем
_Request_
```
curl -s -D - --request POST \
  --url http://localhost:9000/create \
  --header 'content-type: application/json' \
  --data '{
	"url": "https://mail.ru",
	"name": "custom_url3"
    }'
```
_Responses:_

* `HTTP/1.1 201 Created`
* `HTTP/1.1 400 Bad Request`
  * url with this name already exists
  * name cannot be empty
  * invalid URI for request
* `HTTP/1.1 500 Internal Server Error`
#### Переход по кастомной ссылке
_Request_
```
curl  -s -D - --request GET \
  --url http://localhost:9000/custom_url \
  --header 'host: www.example.com'
```
_Responses:_

* `HTTP/1.1 301 Moved Permanently`
* `HTTP/1.1 404 Not Found`
  * url does not exists
* `HTTP/1.1 500 Internal Server Error`