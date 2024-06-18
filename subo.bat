REM Download the wkhtmltopdf binary
curl -L -o wkhtmltopdf.tar.xz https://github.com/wkhtmltopdf/wkhtmltopdf/releases/download/0.12.4/wkhtmltox-0.12.4_linux-generic-amd64.tar.xz

REM Extract the wkhtmltopdf binary
tar -xf wkhtmltopdf.tar.xz --strip-components 2 wkhtmltox/bin/wkhtmltopdf

git add .
git commit -m "Incluindo variavel ambiente"
git push
set GOOS=linux
set GOARCH=arm64
set CGO_ENABLED=0
go build -o bootstrap main.go
del main.zip
tar.exe -a -cf main.zip bootstrap wkhtmltopdf
