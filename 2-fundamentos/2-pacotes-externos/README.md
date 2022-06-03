```go
// gerar modulo
go mod init modulo 

// criar arquivo binário com o conteúdo projeto
go build

// para rodar/ler o arquivo compilado
./modulo

// para rodar o executavel
go run main.go

// instalando dependencia externa
// go get <url>
go get github.com/badoux/checkmail

// remove todas dependencias que não estão sendo utilizadas
go mod tidy
```