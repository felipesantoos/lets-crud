# Lets CRUD
Um CRUD feito em Go.

# Primeiros passos
- Crie uma nova pasta para o projeto com o comando `mkdir lets-crud`.
- Entre na nova pasta criada com o comando `cd lets-crud`.
- Execute o comando `go mod init letscrud` pra definir o nome do módulo do projeto.
- Instale o echo com o comando `go get github.com/labstack/echo/v4`.
- Crie o arquivo main.go.
- Execute o servidor com o comando `go run main.go`.
- Teste as rotas usando a URL `http://localhost:3000`.

# Rotas
## Criar um novo cliente
> **POST** http://localhost:3000/customer
```json
{
	"cpf": "759.547.300-40",
	"name": "Felipe da Silva Santos",
	"birthDate": "2001-11-18"
}

```
## Listar todos os clientes
> **GET** http://localhost:3000/customers
## Lista um cliente específico
> **GET** http://localhost:3000/customer/:id
## Atualizar um cliente específico
> **PUT** http://localhost:3000/customer/:id
```json
{
	"cpf": "995.694.610-98",
	"name": "João Pedro da Silva Santos",
	"birthDate": "2020-06-11"
}
```
## Deletar um cliente específico
> **DELETE** http://localhost:3000/customer/:id

# Objetivos
- [ ] Criar esquema do banco de dados.
- [ ] Criar endpoints.
	- [x] Criar rotas.
	- [ ] Criar handlers.
- [ ] Criar service.
	- [ ] Criar interfaces.
	- [ ] Implementar métodos.
- [ ] Criar repository.
	- [ ] Criar interface.
	- [ ] Implementar métodos.

# Configuração do banco de dados
```bash
go get -u github.com/go-sql-driver/mysql
sudo mysql
```
```sql
DROP DATABASE IF EXISTS letscrud;
CREATE DATABASE letscrud;
USE letscrud;

CREATE TABLE customer (
    id SERIAL PRIMARY KEY,
    cpf VARCHAR(11) UNIQUE NOT NULL,
    name VARCHAR(200) NOT NULL,
    birthDate DATE
);
```
