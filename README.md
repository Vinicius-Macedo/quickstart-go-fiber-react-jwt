# Projeto Quickstart com Autenticação JWT

Este é um projeto quickstart que implementa um sistema de autenticação básico utilizando JWT. O frontend é desenvolvido em React e o backend em Go (Golang). O projeto é containerizado usando Docker Compose.

## Pré-requisitos

- Docker
- Docker Compose

## Instalação

1. Clone o repositório:
    ```sh
    git clone https://github.com/Vinicius-Macedo/quickstart-go-fiber-react-jwt.git
    cd quickstart-go-fiber-react-jwt/
    ```

2. Inicie os containers:
    ```sh
    docker-compose up -d
    ```

## Uso

- O frontend estará disponível em: `http://localhost:5173`
- O backend estará disponível em: `http://localhost:3000`
- O banco de dados PostgreSQL estará disponível em: `localhost:5432`

## Estrutura do Projeto

- `frontend/`: Contém o código do frontend desenvolvido em React.
- `backend/`: Contém o código do backend desenvolvido em Go.
- `database/`: Contém os dados persistidos do PostgreSQL.
- `sql/`: Contém os scripts SQL para inicialização do banco de dados.

## Tecnologias Utilizadas

- React
- Go (Golang)
- PostgreSQL
- Docker
- Docker Compose

## Licença

Este projeto está licenciado sob a Licença MIT.
