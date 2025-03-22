# Pentakill Backend 🎮

Este é um backend em Go para armazenar e gerenciar jogos, usando MySQL e Docker.

## 🔧 Tecnologias

- Golang
- MySQL
- Docker
- Gorm (ORM para Go)

## 🚀 Como rodar o projeto

1. Clone o repositório:

   ```sh
   git clone https://github.com/JoaumVictor/pentakill-backend.git

   ```

2. Rode os serviços com Docker:

   ```sh
     docker-compose up -d
   ```

3. Acesse o backend:

   ```sh
     http://localhost:8080
   ```

📌 Rotas disponíveis

    GET /games → Lista todos os jogos

    POST /games → Adiciona um jogo

    PUT /games/:id → Edita um jogo

    DELETE /games/:id → Remove um jogo

By JoaumVictor
