# 🔐 Go Auth API (JWT)

API de autenticação desenvolvida em Go utilizando JWT (JSON Web Token), com foco em boas práticas de backend, segurança e organização em camadas.

---

## 🚀 Tecnologias

* Go (Golang)
* Gin (framework HTTP)
* GORM (ORM)
* SQLite (banco de dados)
* JWT (autenticação)
* bcrypt (hash de senha)

---

## 🧠 Arquitetura

O projeto segue uma estrutura modular baseada em separação de responsabilidades:

```
cmd/
internal/
  controller/
  service/
  repository/
  model/
  dto/
  middleware/
pkg/
  database/
configs/
```

---

## 🔐 Funcionalidades

* ✅ Cadastro de usuário
* ✅ Login com validação de credenciais
* ✅ Geração de token JWT
* ✅ Hash seguro de senha com bcrypt
* ✅ Estrutura preparada para rotas protegidas

---

## 📡 Endpoints

### 📌 Register

```http
POST /register
```

**Body:**

```json
{
  "email": "user@email.com",
  "password": "123456"
}
```

---

### 📌 Login

```http
POST /login
```

**Response:**

```json
{
  "token": "jwt_token_aqui"
}
```

---

## ▶️ Como executar

```bash
# clonar o repositório
git clone https://github.com/seu-usuario/go-auth-jwt-api

# entrar na pasta
cd go-auth-jwt-api

# instalar dependências
go mod tidy

# rodar o projeto
go run cmd/main.go
```

A API estará disponível em:

```
http://localhost:8080
```

---

## 🧪 Testando a API

Você pode utilizar:

* Thunder Client (VS Code)
* Postman
* Insomnia

---

## 🔒 Segurança

* Senhas armazenadas com hash (bcrypt)
* Tokens JWT com expiração
* Estrutura pronta para middleware de autenticação

---

## 💡 Objetivo do projeto

Este projeto foi desenvolvido com foco em prática de:

* Construção de APIs REST em Go
* Autenticação com JWT
* Organização de código em camadas
* Boas práticas de segurança

---

## 👨‍💻 Autor

Desenvolvido por João Victor 🚀
