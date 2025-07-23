# Desafio Clean Architecture - Order Management

Este repositório contém a implementação de um sistema de gerenciamento de pedidos utilizando Clean Architecture. O sistema inclui uma API REST, um serviço GRPC e uma interface GraphQL para listar pedidos.

## Requisitos

- [x] Implementar usecase de listagem de orders
- [x] Endpoint REST (GET /order)
- [x] Service ListOrders com GRPC
- [x] Query ListOrders GraphQL
- [x] Migrations do banco de dados
- [x] Arquivo api.http com requests
- [x] Docker/Docker Compose configurado
- [x] README.md com instruções

## Como executar

1. Clone o repositório
2. Execute o comando:
   ```bash
   docker compose up
   ```

## Portas dos Serviços

- REST API: porta 8080
- GRPC Service: porta 50051
- GraphQL: porta 8080/graphql

## Docker Compose

### Run

```bash
docker compose up --build
```

> **Atenção**:
>
> Aguarde a inicialização dos serviços, incluindo o banco de dados e RabbitMQ.
>
> A aplicação será iniciada automaticamente em cerca de 15/30 segundos.
>
> **Logs**:
> ordersystem | O sistema está iniciado, aguardando alguns segundos...
> ordersystem | Starting web server on port :8000  
> ordersystem | Starting gRPC server on port 50051  
> ordersystem | Starting GraphQL server on port 8080
>
> --

## Teste

### API REST

para teste via rest, use o arquivo `api/` para fazer as requisições.

### GRPC

use ferramentas `evans` para realizar GRPC

```bash
evans --proto internal/infra/grpc/protofiles/order.proto repl

call ListOrders
```

### GraphQL

Acesse a URL `http://localhost:8080/` e utilize a query:

```graphql
query {
  ListOrders {
    id
    Price
    Tax
    FinalPrice
  }
}
```
