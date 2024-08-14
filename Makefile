build:
	go build main.go controllers.go helpers.go middleware.go
	cd ./client && npm install
	cd ./client && npm run build