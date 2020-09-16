### Запуск приложения
```
$ git clone https://github.com/ArkovKonstantin/chat-server.git
$ cd chat-server
$ make build
$ docker-compose up
```
Приложение будет доступно по адресу http://localhost:9000

### API методы

#### Создание укороченной ссылки
```
curl --request POST \
  --url http://localhost:9000/shorten \
  --header 'content-type: application/json' \
  --data '{
	"url": "https://mail.ru"
}'
```
#### Создание укороченной ссылки c кастомным именем
```
curl --request POST \
  --url http://localhost:9000/create \
  --header 'content-type: application/json' \
  --data '{
	"url": "https://mail.ru",
	"name": "custom_url"
}'
```
#### Переход по кастомной ссылке
```
curl --request GET \
  --url http://localhost:9000/custom_url \
  --header 'host: www.example.com'
```