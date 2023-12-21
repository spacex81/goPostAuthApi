docker-compose down; env GOOS=linux GOARCH=amd64 go build -o main main.go; docker-compose up --build -d;
