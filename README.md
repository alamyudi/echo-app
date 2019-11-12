Echo-app is an example of golang rest api using echo framework, included :

- [x] Rest API 
    - [x] using echo framework to managing content and product
    - [x] Splitting between model business and application presentation
    - [x] Configurable the app to running in certain environment
    - [x] request and response serializer and validation
    - [ ] ORM
    - [ ] Unit test
- [x] Vendoring a golang dependencies using glide as a package management
- [x] Hot reloading using realize
- [x] Dockerizing the app and mysql
- [ ] CI/CD
- [ ] Generate postman collection

### Running the app using docker

```
docker-compose up --build
```

access it using curl or postman in `http://localhost:8081/v1` as base url. for example :

```
curl -X GET \
  http://localhost:8081/v1/product \
  -H 'Accept: */*' \
  -H 'Accept-Encoding: gzip, deflate' \
  -H 'Authorization: Basic ZWNobzplY2hvZWg='
```

### Building the executable

```
env GOOS=linux GOARCH=amd64 go build -v .
```