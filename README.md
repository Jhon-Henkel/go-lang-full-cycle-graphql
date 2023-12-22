# go-lang-full-cycle-graphql
Repositório para armazenar os exemplos de graph-ql do curso de Go da Full Cycle

# Rodando o projeto
```sh
go run server.go
```
- Criar as tabelas no banco de dados
```sh
sqlite3 database.db;
create table categories (id string, name string, description string);
```

# Problemas comuns
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