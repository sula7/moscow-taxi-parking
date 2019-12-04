# moscow-taxi-parking

API server that gets data from data.gov.ru and takes HTTP requests.
As a storage using Redis.

### ENV
If some env is changed, need to edit docker-compose.test.yml
````
BIND_PORT=  HTTP server's listen port (Default is :8080)
DB_CONN=    Required. Redis DB connection string like "localhost:6379"
DB_PWD=     Redis DB password (Default is "")
SRC_URL=    Required. Data source URL. Example https://data.gov.ru/opendata/7704786030-taxiparking/data-20190906T0100.json?encoding=UTF-8
FILE_NAME=  In case if data source is unavailable put json file to ./local/ and set file name
(Default is "data-20190906T0100.json")
````

### Run
Run by
`docker-compose up -d`
after containers up by
`go run ./main.go`

### Test
Run by `make test`


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
