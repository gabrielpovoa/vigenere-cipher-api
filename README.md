# 🔐 API Vigenère – Criptografia e Descriptografia com Go

Esta é uma API simples escrita em Go que permite **codificar e decodificar palavras** usando o **código de Vigenère**. Também salva os resultados em um arquivo `.txt`.

---

## 🚀 Como rodar o projeto

### Pré-requisitos

- [Go instalado](https://golang.org/doc/install)
- Git instalado (opcional, se for clonar o projeto)

### Passos

```bash
# Clone o repositório (caso não tenha feito isso ainda)
git clone https://github.com/gabrielpovoa/vigenere-cipher-api

# Acesse a pasta do projeto
cd vigenere-api

# Inicialize o módulo (caso ainda não tenha)
go mod tidy

# Rode a API
go run main.go

A API estará disponível em: http://localhost:8080
```
## 📡 Rotas disponíveis
🔐 POST /vigenere
<br>Codifica uma palavra usando o código de Vigenère.<br>

- Body (JSON):
    ```
    {
    "text": "HELLO",
    "key": "KEY"
    }
    ```

- Resposta

    ```
    {
    "encrypted": "RIJVS"
    }
    ```

____
## 🔓 POST /vigenere/decode
Decodifica uma palavra cifrada usando a mesma chave.

- BODY (JSON)

    ```
    {
    "text": "RIJVS",
    "key": "KEY"
    }
    ```

- Resposta

    ```
    {
      "encrypted": "HELLO"
    }
    ```
___

## 📝 Resultado em arquivo .txt
Após cada requisição, a API grava os resultados no arquivo result.txt com o seguinte formato:

- 📦 Dependências
    - A API utiliza apenas pacotes padrão da linguagem Go:

    1. net/http
    2. encoding/json
    3. strings
    4. fmt
    5. os

👨‍💻 Author
João Gabriel Póvoa

GitHub: @gabrielpovoa