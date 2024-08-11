build:
	go build main.go controllers.go
	cd ./client && npm install
	cd ./client && npm run build