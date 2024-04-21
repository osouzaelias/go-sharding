## Sobre projeto

Este projeto tem como objetivo demonstrar a técnica de sharding de dados. Para isso, disponibiliza uma API simples, 
desenvolvida em Go, que conta com três rotas principais.

---

<details>
 <summary> <code>POST</code> <code><b>/customer</b></code><code>(adiciona novo cliente)</code></summary>

##### Responses

> | http code | content-type       | response                                |
> |-----------|--------------------|-----------------------------------------|
> | `201`     | `application/json` | `{ "id": "000", "name": "xxxx" }`       |
> | `400`     | `application/json` | `{ "errors": ["operation error..."] }}` |

##### Example cURL

> ```shell
>  curl --location 'http://localhost:8080/customer' \
>       --header 'Content-Type: application/json' \
>       --data '{ "id": "81136522000117", "name": "Company Go Inc." }'
> ```

</details>

<details>
 <summary><code>GET</code> <code><b>/customer/{id}</b></code><code>(consulta o cliente por id)</code></summary>

##### Parameters

> | name |  type      | data type | description                                            |
> |------|------------|-----------|--------------------------------------------------------|
> | `id` |  required  | uint64    | Identificador único do cliente (ex.: CNPJ, CPF e etc.) |

##### Responses

> | http code | content-type       | response                                |
> |-----------|--------------------|-----------------------------------------|
> | `200`     | `application/json` | `{ "id": "000", "name": "xxxx" }`       |
> | `404`     | `application/json` | `{ "errors": ["customer not found"] }}` |
> | `400`     | `application/json` | `{ "errors": ["operation error..."] }}` |

##### Example cURL

> ```shell
> curl --location 'http://localhost:8080/customer/81136522000117'
> ```

</details>

<details>
 <summary><code>GET</code><code><b>/health</b></code><code>(verifica a saúde da aplicação)</code></summary>

##### Responses

> | http code | content-type       | response                  |
> |-----------|--------------------|---------------------------|
> | `200`     | `application/json` | `{ "Status": "Healthy" }` |

##### Example cURL

> ```shell
> curl --location 'http://localhost:8080/health'
> ```

</details>

---

Ao inserir um novo registro, a aplicação utiliza o algoritmo `Rendezvous Hashing` para determinar em qual nó 
(tabela DynamoDB) o registro será armazenado. O mesmo algoritmo é aplicado para localizar o nó onde o registro foi 
salvo quando uma consulta é realizada usando o ID do cliente.

## Pré requisitos

* Instale e configure o [AWS CLI](https://docs.aws.amazon.com/pt_br/cli/latest/userguide/getting-started-install.html)
para interagir com o DynamoDB via linha de comando. 
* Instale o [NoSQL Workbench para DynamoDB](https://docs.aws.amazon.com/pt_br/amazondynamodb/latest/developerguide/workbench.html), 
para operar o DynamoDB locamente na sua máquina.
* Instale o [Go](https://go.dev/dl/) com versão igual ou superior a 1.22
* Tenha disponível um editor de código. Algumas opções incluem [Neovim](https://neovim.io/), [Zed](https://zed.dev/) 
ou [VS Code](https://code.visualstudio.com/), No desenvolvimento deste projeto, utilizei o [GoLand](https://www.jetbrains.com/pt-br/go/promo/). 

## Como executar

**#1**: Clone o repositório no seu computador e acesse o diretório raiz da aplicação.

---

**#2**: Execute o script `create_tables.sh` localizado na pasta `scripts` para criar as tabelas DynamoDB localmente:

```shell (ocidogo)
bash scripts/create_tables.sh
```

---

**#3**: Configure as variáveis de ambiente necessárias para a execução da aplicação:

```shell
export PORT=8080
export REGION=us-west-2 
export TABLES=node1,node2,node3
```

---

**#4**: Inicie a aplicação utilizando o comando abaixo. Você verá a mensagem indicando que o servidor está ativo:

```shell
go run cmd/main.go
```

> [!NOTE]
> Se for a primeira vez que executa o projeto, instale as dependências com:
> ```shell
> go mod tidy
> ```

---

**#5**: Com a aplicação em funcionamento, abra uma nova aba do terminal e execute o script data_load.sh para enviar 
1.000 requisições com dados aleatórios:

```shell
bash scripts/data_load.sh
```

---

**#6**: Para verificar a distribuição dos dados entre as tabelas, execute o script `data_count.sh`:

```shell
bash scripts/data_count.sh
```

Os dados também podem ser visualizados usando a ferramenta `NoSQL Workbench`.

---

## Referências

- Como estruturar um projeto em Go: [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- Para escolher o algoritmo para distribuição de dados: [Consistent Hashing: Algorithmic Tradeoffs](https://dgryski.medium.com/consistent-hashing-algorithmic-tradeoffs-ef6b8e2fcae8)
- Caso real de implementação: [Client lib for redis](https://github.com/redis/go-redis/blob/21bd40a47e56e61c0598ea1bdf8e02e67d1aa651/ring.go#L28) 
- Para entender sobre [Rendezvous hashing](https://en.wikipedia.org/wiki/Rendezvous_hashing)
- Implementações do algoritmo de hashing: [xxHash non-cryptographic hash algorithm](https://xxhash.com/)
- Entender sobre geração de números aleatórios: [Xorshift random number generators,](https://pt.wikipedia.org/wiki/Xorshift)
