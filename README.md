# Banner project
### Запуск тестов
```bash
docker-compose -f docker-compose.yml -f docker-compose.test.yml up
Или
docker-compose -f docker-compose.yml -f docker-compose.test.yml run --rm rotation_banner make test
```
### Запуск проекта
```bash
docker-compose up
```
### Доступ к API(для проверки)

- Rest api [localhost:8080](http://localhost:8080/api)
- Grpc api [localhost:8000](http://localhost:8000)
- Prometheus [localhost:9090](http://localhost:9090/graph)
- RabbitMQ [localhost:15672](http://localhost:15672)(login: guest, pass: guest)

---
## Ручное тестирование
### REST API
#### Добавление баннера в ротацию
```bash
curl -X POST \
  http://localhost:8080/api/add_banner/ \
  -H 'Content-Type: application/javascript' \
  -d '{
	"id_banner": 1,
	"id_slot": 1
}'
```
#### Удаление баннера из ротации
```bash
curl -X DELETE \
  http://localhost:8080/api/del_banner/1/ \
  -H 'cache-control: no-cache'
```
#### Засчитать преход(клик) по баннеру
```bash
curl -X POST \
  http://localhost:8080/api/count_transition/ \
  -H 'Content-Type: application/json' \
  -d '{"id_banner": 1,"id_slot": 1,"id_soc_dem": 1}'
```
#### Получение баннера для показа
```bash
curl -X GET http://localhost:8080/api/get_banner/1/1/
```