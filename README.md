# moscow-taxi-parking

API server that gets data from data.gov.ru and takes HTTP requests.
As a storage using Redis.

### ENV
````

````

### HTTP methods
````
GET /api/v1/parking/id/{ID}
Responses parking info
````
````
GET /api/v1/parking/global-id/{ID}
Responses parking info
````
````
POST /api/v1/parking/mode

Params required
page=       page
per_page    result per page

Body required
{"mode":""}

Responses parking info
````

### How to start
Run by
`docker-compose up -d`
after containers up by
`go run ./main.go`
