# moscow-taxi-parking

API server that gets data from data.gov.ru and takes HTTP requests.
As a storage using Redis.

### ENV
````
BIND_PORT=  HTTP server's listen port (Default is :8080)
DB_CONN=    Required. Redis DB connection string like "localhost:6379"
DB_PWD=     Redis DB password (Default is "")
FILE_NAME=  In case if data source is unavailable put json file to ./local/ and set file name
(Default is "data-20190906T0100.json")
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
page=       page (default is 1)
per_page    result per page (default is 5)

Body required
{"mode":""}

Responses parking info
````

### How to start
Run by
`docker-compose up -d`
after containers up by
`go run ./main.go`

### How to test
Run by `make test`
