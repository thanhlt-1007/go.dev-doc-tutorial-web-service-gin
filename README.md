# go.dev-doc-tutorial-web-service-gin

- Tutorial: Developing a RESTful API with Go and Gin

- Reference: https://go.dev/doc/tutorial/web-service-gin

## gvm

```sh
gvm install go1.23.5
gvm use go1.23.5
```

## go mod download

```sh
cd web-service-gin
go mod download
```

# go run

```sh
cd web-service-gin
go run .
```

## cURL

- GET /albums

```sh
curl http://localhost:8080/albums
```

- GET /albums/:id

```sh
curl http://localhost:8080/albums/1
```

- POST /albums

```sh
curl --location 'http://localhost:8080/albums' \
--header 'Content-Type: application/json' \
--data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```