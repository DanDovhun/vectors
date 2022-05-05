install:
	go get fyne.io/fyne/v2

init:
	go mod init Vectors

build:
	go mod tidy
	go build

buildClean:
	go mod init Vectors
	go mod tidy
	go build

run:
	go run main.go

buildAndRun:
	go mod tidy && go build
	./Vectors

cleanBuildAndRun:
	go mod init Vectors && go mod tidy && go build
	./Vectors