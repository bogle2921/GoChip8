all:
	go build -o ./bin/main.exe ./src/main.go

clean:
	del .\bin\*.exe