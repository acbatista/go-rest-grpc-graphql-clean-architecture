# Serviço de Pedidos com Clean Architecture

Este projeto demonstra uma implementação de Clean Architecture em Go, oferecendo um serviço de gerenciamento de pedidos através de múltiplas interfaces:

- API REST
- Serviço gRPC
- API GraphQL

## Portas dos Serviços

- API REST: http://localhost:8080
- gRPC: localhost:50051
- GraphQL: http://localhost:8080/graphql

## Pré-requisitos

- Docker e Docker Compose
- Go 1.21 ou superior

## Como Iniciar

1. Clone o repositório:

```bash
git clone https://github.com/acbatista/go-rest-grpc-graphql-clean-architecture
cd go-rest-grpc-graphql-clean-architecture
```

2. Inicie a aplicação e o banco de dados usando Docker Compose:

```bash
docker compose up
```

3. A aplicação irá automaticamente:
   - Configurar o banco de dados MySQL
   - Executar as migrações necessárias
   - Iniciar todos os serviços (REST, gRPC e GraphQL)

## Documentação da API

### Endpoints REST

- GET /order - Lista todos os pedidos
- POST /order - Cria um novo pedido

Exemplo de requisição para criar um pedido:

```json
POST http://localhost:8080/order
Content-Type: application/json

{
    "customer_name": "João Silva",
    "total": 99.99,
    "status": "pendente"
}
```

### GraphQL

Acesse o playground GraphQL em http://localhost:8080/graphql

Exemplo de consulta:

```graphql
query {
  listOrders {
    id
    customerName
    total
    status
    createdAt
  }
}
```

### gRPC

O serviço gRPC está disponível na porta 50051 com os seguintes métodos:

- ListOrders

Para testar usando grpcurl:

```bash
grpcurl -plaintext localhost:50051 order.OrderService/ListOrders
```

## Estrutura do Projeto

.
├── cmd/ 

## Testes

Para executar os testes:

```bash
go test ./...
```

## Exemplos de Uso

Você pode encontrar exemplos de requisições no arquivo `api.http`, que pode ser executado usando clientes REST como a extensão REST Client do VS Code ou Postman.

## Banco de Dados

O projeto utiliza MySQL como banco de dados. A estrutura do banco é criada automaticamente através das migrações quando a aplicação é iniciada.

### Configuração do Banco de Dados

- Host: localhost
- Porta: 3306
- Usuário: root
- Senha: root
- Banco de dados: orders_db

## Contribuindo

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/nova-feature`)
3. Faça commit das suas alterações (`git commit -am 'Adiciona nova feature'`)
4. Faça push para a branch (`git push origin feature/nova-feature`)
5. Crie um novo Pull Request 