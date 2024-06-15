git add .
git commit -m "teste 3"
git push
set GOOS=linux
set GOARCH=arm64
set CGO_ENABLED=0
go build -o bootstrap main.go
del main.zip
tar.exe -a -cf main.zip bootstrap