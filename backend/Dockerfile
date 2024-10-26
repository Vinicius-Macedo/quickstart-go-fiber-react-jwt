# Use a imagem base do Go 1.23
FROM golang:1.23 AS build

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie o arquivo go.mod e go.sum para o diretório de trabalho
COPY go.mod go.sum ./
RUN go mod download

# Copie o restante do código-fonte para o diretório de trabalho
COPY . .

# Compile o código
RUN go build -o server ./cmd/api


# Exponha a porta que o aplicativo irá usar
EXPOSE 3000

# Comando para executar o aplicativo
CMD ["./server"]