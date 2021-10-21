# GO GraphQL


## [GraphQL](https://graphql.org/learn/)

  Uma linguagem de consulta para sua API
  GraphQL é uma linguagem de consulta para APIs e um tempo de execução 
  para atender a essas consultas com seus dados existentes. 
  GraphQL fornece uma descrição completa e compreensível dos dados em sua API, 
  dá aos clientes o poder de pedir exatamente o que precisam e nada mais, 
  torna mais fácil desenvolver APIs ao longo do tempo e habilita ferramentas poderosas de desenvolvedor.
  
  ![graphql](https://upload.wikimedia.org/wikipedia/commons/1/17/GraphQL_Logo.svg)

## Instalações

### Criar módulo go do projeto:
````bash
go mod init github.com/flaviossantana/go-graphql 
````

### gqlgen
gqlgen é uma biblioteca Go para construir servidores GraphQL sem qualquer problema.

#### Instalar gqlgen
```bash
go get github.com/99designs/gqlgen
```

### Criar os arquivos do projeto com base no schema.graphqls
```bash
go run github.com/99designs/gqlgen init
```
