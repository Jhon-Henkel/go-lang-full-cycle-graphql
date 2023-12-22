# Go-lang Full Cycle Graph-Ql
Repositório para armazenar os exemplos de graph-ql do curso de Go da Full Cycle

## Rodando o projeto
Para rodar o projeto é necessário dar start no server e criar as tabelas no banco de dados
- Rodando o server:
  ```sh
  go run server.go
  ```
- Criar as tabelas no banco de dados:
  ```sh
  sqlite3 database.db;
  CREATE TABLE categories (id string, name string, description string);
  CREATE TABLE courses (id string, name string, description string, category_id string);
  ```

## Documentação API 
- https://github.com/Jhon-Henkel/go-lang-full-cycle-graphql/blob/main/.docs/graphQlDocs.md

## Problemas comuns
- Erro ao instalar o pacote do sqlite3
    ```sh
    sudo apt install sqlite
    ```
- Erro ao rodar o projeto pois requer o CGO_ENABLED
    ```sh
    export CGO_ENABLED=1
    ```
- Erro com GCC
    ```sh
    sudo apt install gcc
    ```