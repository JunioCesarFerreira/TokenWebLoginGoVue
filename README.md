# TokenWebLoginGoVue

🌍 *[**Português**](README.md) ∙ [English](README_en.md)*

TokenWebLoginGoVue é um projeto que demonstra um sistema de autenticação completo para aplicações web, utilizando Go para o backend e Vue.js para a interface do usuário. O projeto implementa autenticação baseada em tokens, usando JWT (JSON Web Tokens) e PASETO (Platform-Agnostic Security Tokens), permitindo login seguro e gerenciamento de sessões de usuário.

## Funcionalidades

- **Autenticação Segura**: Autenticação baseada em tokens usando JWT e PASETO.
- **Renovação de Tokens**: Tokens são renovados automaticamente para garantir sessões seguras e contínuas.
- **Inclusão de dados adicionais nos tokens**: Incluímos um ID de instância nos exemplos, o que pode ser útil para a gestão de sessões e para cenários de multi-tenancy.
- **Logout por Inatividade**: Usuários são automaticamente desconectados após um período de inatividade.
- **Interface Amigável**: Interface de usuário construída com Vue.js para testes de fluxo de login.
- **Documentação da API**: Documentação clara e acessível usando Swagger para explorar a API de autenticação.

## Tecnologias Utilizadas

- **Backend**: Go
- **Frontend**: Vue.js
- **Bibliotecas**:
  - `gorilla/mux` para roteamento no Go
  - `swaggo/http-swagger` para documentação da API
  - `vue-router` para navegação no Vue.js
  - `crypto-js` para hashing de senhas no frontend

## Requisitos

- **Node.js** (versão 18 ou superior)
- **Go** (versão 1.21 ou superior)
- **Vue CLI**

## Configuração e Execução

### Backend (Go)

1. Clone o repositório:

   ```bash
   git clone https://github.com/seu-usuario/TokenWebLoginGoVue.git
   cd TokenWebLoginGoVue
   cd api
   ```

2. Instale as dependências:

   ```bash
   go mod tidy
   ```

3. Execute o servidor Go:

   ```bash
   go run main.go
   ```

   O servidor estará disponível em `http://localhost:8082`.

### Frontend (Vue.js)

1. Navegue para o diretório do frontend:

   ```bash
   cd ui
   ```

2. Instale as dependências do Vue.js:

   ```bash
   npm install
   ```

3. Execute a aplicação Vue.js:

   ```bash
   npm run serve
   ```

   A aplicação estará disponível em `http://localhost:8080`.

## Estrutura do Projeto

```plaintext
TokenWebLoginGoVue/
│
├── api/                     # Go server code
│   ├── main.go              # Main entry point of the Go application
│   ├── controller/          # API controllers
│   ├── middleware/          # Authentication middleware
│   ├── data/                # Data repository
│   ├── docs/                # Swagger documentation managed with swaggo
│   └── Dockerfile           # Arquivo docker para gerar imagem da api
│
└── ui/                      # Vue.js interface code
    ├── src/
    │   ├── components/      # Vue components
    │   ├── router/          # Router configuration
    │   ├── views/           # Vue pages
    │   ├── App.vue          # Root component
    │   ├── main.js          # Vue.js entry file
    │   └── config.js        # API access configurations and other experiment constants
    ├── public/              # Static files
    └── Dockerfile           # Arquivo docker para gerar imagem da ui
```

## Como Usar

1. **Login**: Navegue para `/login`, insira o ID do usuário e a senha para fazer login.
2. **Página Protegida**: Acesse `/protected` após login bem-sucedido para visualizar a página protegida.
3. **Renovação Automática**: Os tokens são renovados automaticamente enquanto o usuário estiver ativo.
4. **Logout por Inatividade**: O usuário será desconectado após 5 minutos de inatividade.

### Utilizando Docker

No diretório principal execute:

```bash
docker compose build
docker compose up -d
```

## Notas

- [Características de Segurança](./notes/feature_pt.md)

- [Breve Comparação JWT vs PASETO](./notes/token_compare_pt.md)


## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou pull requests para melhorias e correções.

## Licença

Este projeto está licenciado sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
