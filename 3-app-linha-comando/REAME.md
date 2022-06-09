## Instalações
```bash
go mod init linha-de-comando 
go get github.com/urfave/cli
go run main.go 
```


## Como rodar o App
```bash
# rodar app em desenvolvimento
go run main.go ip 
go run main.go ip --host google.com.br  
go run main.go servidores --host google.com.br  

# gerar build do app
go build

# exibir IP / server
./linha-de-comando ip --host google.com.br
./linha-de-comando servidores --host google.com.br 

# exibir IP padrão
./linha-de-comando ip
```