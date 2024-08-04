# TokenWebLoginGoVue

🌍 *[Português](README.md) ∙ [English](README_en.md)*

TokenWebLoginGoVue é um projeto que demonstra um sistema de autenticação completo para aplicações web, utilizando Go para o backend e Vue.js para a interface do usuário. O projeto implementa autenticação baseada em tokens, usando JWT (JSON Web Tokens) e PASETO (Platform-Agnostic Security Tokens), permitindo login seguro e gerenciamento de sessões de usuário.

## Funcionalidades

- **Autenticação Segura**: Autenticação baseada em tokens usando JWT e PASETO.
- **Renovação de Tokens**: Tokens são renovados automaticamente para garantir sessões seguras e contínuas.
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
├── api/                     # Código do servidor Go
│   ├── main.go              # Entrada principal do aplicativo Go
│   ├── controller/          # Controladores da API
│   ├── middleware/          # Middleware de autenticação
│   ├── data/                # Repositório de dados
│   └── docs/                # Documentação Swagger gerenciada com swaggo
│
└── ui/                      # Código da interface Vue.js
    ├── src/
    │   ├── components/      # Componentes Vue
    │   ├── router/          # Configuração de rotas
    │   ├── views/           # Páginas Vue
    │   ├── App.vue          # Componente raiz
    │   ├── main.js          # Arquivo de entrada Vue.js
    |   └── config.js        # Configurações de acesso à api e outras constantes do experimento
    └── public/              # Arquivos estáticos
```

## Como Usar

1. **Login**: Navegue para `/login`, insira o ID do usuário e a senha para fazer login.
2. **Página Protegida**: Acesse `/protected` após login bem-sucedido para visualizar a página protegida.
3. **Renovação Automática**: Os tokens são renovados automaticamente enquanto o usuário estiver ativo.
4. **Logout por Inatividade**: O usuário será desconectado após 5 minutos de inatividade.

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou pull requests para melhorias e correções.

## Licença

Este projeto está licenciado sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## Contato

Para dúvidas ou sugestões, entre em contato com [seu-email@example.com](mailto:jcf_ssp@hotmail.com).
