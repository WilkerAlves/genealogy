# Genealogy

- Instalação
    * Git
    * Go
    * Docker
    * docker-compose

## Utilização:

```bash
    git clone https://github.com/pedrocmart/crud-go

    cd genealogy
```
- Renomear o arquivo `.env.exemplo` para `.env` e preencher as chaves com as seguintes informações:
  - DB_HOST=localhost
  - DB_PORT=3307
  - DB_NAME=MainUser
  - DB_USER=MainPassword
  - DB_PASS=genealogy
  - SERVER_PORT=8080
```bash
    docker-compose up
```

## Rotas disponiveis:

- http://localhost:8080/persons/{personID}
  - Esse endpoint busca uma pessoa usando como parametro o id da mesma
  - GET
- http://localhost:8080/persons/{personID}
  - Esse endpoint atualiza o nome de uma pessoa usando como parametro o id da mesma
  - PUT
  - Necessario envio um payload json ``{ "name": "Bruce" }``
- http://localhost:8080/persons/
  - Esse endpoint cria uma pessoa
  - POST
  - Necessario envio um payload json ``{ "name": "Bruce" }``
- http://localhost:8080/persons/{personID}
  - Esse endpoint busca uma pessoa usando como parametro o id da mesma
  - DELETE
- http://localhost:8080/relationship/
  - Esse endpoint cria a relação de uma pessoa com a outra
  - POST
  - Necessario envio um payload json ``{ "parent": personID, "children": personID }``
  - personID deve ser um valor inteiro maior que zero
- http://localhost:8080/relationship/find?id=1&findrelationship=9
  - Esse endpoint busca o relacionamento entre duas pessoas usando como parametro os ids das mesmas
  - GET
- http://localhost:8080/relationship/4
  - Esse endpoint busca a arvore genealogia de uma pessoa usando como parametro os id da mesma
  - GET