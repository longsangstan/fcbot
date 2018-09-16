rm main
rm main.zip
GOOS=linux GOARCH=amd64 go build -o main *.go
zip main.zip main