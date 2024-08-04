# TokenWebLoginGoVue

🌍 *[Português](README.md) ∙ [English](README_en.md)*

TokenWebLoginGoVue is a project that demonstrates a complete authentication system for web applications, using Go for the backend and Vue.js for the user interface. The project implements token-based authentication using JWT (JSON Web Tokens) and PASETO (Platform-Agnostic Security Tokens), allowing secure login and user session management.

## Features

- **Secure Authentication**: Token-based authentication using JWT and PASETO.
- **Token Renewal**: Tokens are automatically renewed to ensure secure and continuous sessions.
- **Idle Logout**: Users are automatically logged out after a period of inactivity.
- **User-Friendly Interface**: User interface built with Vue.js for testing login flows.
- **API Documentation**: Clear and accessible documentation using Swagger to explore the authentication API.

## Technologies Used

- **Backend**: Go
- **Frontend**: Vue.js
- **Libraries**:
  - `gorilla/mux` for routing in Go
  - `swaggo/http-swagger` for API documentation
  - `vue-router` for navigation in Vue.js
  - `crypto-js` for password hashing on the frontend

## Requirements

- **Node.js** (version 18 or higher)
- **Go** (version 1.21 or higher)
- **Vue CLI**

## Setup and Execution

### Backend (Go)

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/TokenWebLoginGoVue.git
   cd TokenWebLoginGoVue
   cd api
   ```

2. Install the dependencies:

   ```bash
   go mod tidy
   ```

3. Run the Go server:

   ```bash
   go run main.go
   ```

   The server will be available at `http://localhost:8082`.

### Frontend (Vue.js)

1. Navigate to the frontend directory:

   ```bash
   cd ui
   ```

2. Install Vue.js dependencies:

   ```bash
   npm install
   ```

3. Run the Vue.js application:

   ```bash
   npm run serve
   ```

   The application will be available at `http://localhost:8080`.

## Project Structure

```plaintext
TokenWebLoginGoVue/
│
├── api/                     # Go server code
│   ├── main.go              # Main entry point of the Go application
│   ├── controller/          # API controllers
│   ├── middleware/          # Authentication middleware
│   ├── data/                # Data repository
│   └── docs/                # Swagger documentation managed with swaggo
│
└── ui/                      # Vue.js interface code
    ├── src/
    │   ├── components/      # Vue components
    │   ├── router/          # Router configuration
    │   ├── views/           # Vue pages
    │   ├── App.vue          # Root component
    │   ├── main.js          # Vue.js entry file
    │   └── config.js        # API access configurations and other experiment constants
    └── public/              # Static files
```

## Usage

1. **Login**: Navigate to `/login`, enter the user ID and password to log in.
2. **Protected Page**: Access `/protected` after a successful login to view the protected page.
3. **Automatic Renewal**: Tokens are automatically renewed while the user is active.
4. **Idle Logout**: The user will be logged out after 5 minutes of inactivity.

## Contribution

Contributions are welcome! Feel free to open issues or pull requests for improvements and fixes.

## License

This project is licensed under the MIT license. See the [LICENSE](LICENSE) file for more details.

## Contact

For questions or suggestions, contact [your-email@example.com](mailto:jcf_ssp@hotmail.com).
