# gRPC GO

## Atualizar Stubs gRPC

Sempre que o arquivo `.proto` for atualizado, é necessário regenerar os stubs para Go.  
Execute o comando abaixo:

protoc --go_out=. --go-grpc_out=. proto/course_category.proto

## Criar tabelas sqlite3

sqlite3 db.sqlite
create table categories (id string, name string, description string);

## Instalar dependências

go mod tidy

## Rodar o servidor gRPC

go run cmd/grpcServer/main.go

## Rodar o client gRPC com Evans

Para interagir com o servidor gRPC usando o Evans, execute o comando abaixo. Certifique-se de que o servidor esteja em execução antes de iniciar o client.

evans -r repl

No prompt do Evans, você pode selecionar o pacote e visualizar os serviços disponíveis:

- package pb
- show service
- service CategoryService
- call CreateCategory
- call ListCategories
- call GetCategory
- call CreateCategoryStream
    - ctrl + d para parar a stream

Consulte a documentação do Evans: https://github.com/ktr0731/evans para mais comandos e opções.

## Ambiente de Desenvolvimento

Este projeto inclui um arquivo Dockerfile e suporte a Dev Containers do VS Code.  
Você pode abrir o projeto diretamente em um dev container para ter todas as dependências pré-instaladas e um ambiente consistente.

## Requisitos

- Go instalado e configurado no PATH (caso não use o dev container)
- Evans instalado (go install github.com/ktr0731/evans@latest)
- Dependências do projeto instaladas (go mod tidy)

## Referências

- Documentação do gRPC em Go: https://grpc.io/docs/languages/go/
- Evans CLI: https://github.com/ktr0731/evans
