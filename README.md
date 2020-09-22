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

### Нагрузочное тестирование
Нагрузочное тестирование было проведено при помощи фреймворка [Locust](https://locust.io/). <br>
Для проведения тестов, необходимо описать сценарий в файле locustfile.py и запустить модуль командой `locust -f locustfile.py`. Затем результаты тестирования можно наблюдать на графиках в веб-интерфейсе или таблице результатов в консоли. <br>
Ниже приведен листинг файла `locustfile.py` <br>
Сценарий проведения теста: <br>
* Методом `shorten` было сгенерировано 5000 записей в базе данных.
* Далее был запущен метод `redirect`, с нагрузкой 1000 одновременных чтений

```python3
import random
from locust import HttpUser, task, between, TaskSet


def convert_base(num, to_base=10, from_base=10):
    if isinstance(num, str):
        n = int(num, from_base)
    else:
        n = int(num)
    alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-._~"
    if n < to_base:
        return alphabet[n]
    else:
        return convert_base(n // to_base, to_base) + alphabet[n % to_base]


class ShortenerTaskSet(TaskSet):

    # Генерация укороченных ссылок для адреса https://ya.ru
    def shorten(self):
        for _ in range(25):
            self.client.post("/shorten", json={"url": "https://ya.ru"})

    # Запрос по укороченной ссылке
    def redirect(self):
        name = convert_base(random.choice(range(1, 5001)), 66, 10)
        self.client.get(f"/{name}")

    @task
    def workflow(self):
        # self.shorten()
        self.redirect()


class WebsiteUser(HttpUser):
    tasks = [ShortenerTaskSet]
    wait_time = between(1, 2)

```

*RPS*
Среднее число обрабатываемых запросов в секунду составило 150 при нагрузке 1000 одновременных чтений. <br>

![Image](https://github.com/ArkovKonstantin/url-shortener/raw/master/assets/total_requests_per_second.png) <br>

*Response time* <br>
* Желтый график - это 95% percentil
* Зеленый график - это среднее время ответа, кторое составило 4,5 - 4,7 сек.

![Image](https://github.com/ArkovKonstantin/url-shortener/raw/master/assets/response_times_(ms).png) <br>

*Number of users* <br>

![Image](https://github.com/ArkovKonstantin/url-shortener/raw/master/assets/number_of_users.png) <br>
