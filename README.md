# 🚀 Golang MVC Boilerplate

A **professional Golang MVC boilerplate** built using **Gin**, designed to help developers quickly start scalable backend APIs with a clean and maintainable project structure.

This boilerplate includes **Swagger API documentation**, environment configuration, routing structure, and a modular architecture suitable for real-world backend development.

---

## ✨ Features

- ⚡ MVC Architecture (Controllers, Models, Routes)
- 🌐 Gin Web Framework
- 📄 Swagger API Documentation
- 🔐 Environment Configuration Support
- 🧩 Modular & Scalable Folder Structure
- 🛠️ Makefile Commands
- 🚀 Production-ready project setup

---

## 📁 Project Structure
```
Golang-MVC-Boilerplate/
│
├── cmd/server/ # Application entry point (main.go)
├── config/ # Environment & app configuration
├── controllers/ # API controllers
├── database/ # Database connection setup
├── routes/ # Route definitions
├── docs/ # Swagger generated files
├── .env.example # Example environment variables
├── Makefile
└── README.md
```

---

## ⚙️ Installation & Setup

### 1️⃣ Clone the Repository
```
git clone https://github.com/hrishabpachange/golang-mvc-boilerplate.git

cd golang-mvc-boilerplate
```
---

### 2️⃣ Install Dependencies
```
go mod tidy
```


---

### 3️⃣ Setup Environment Variables

Create a `.env` file using `.env.example`:

---

### 4️⃣ Run the Server

Using Makefile:
`make run`


OR manually:
`go run cmd/server/main.go`


Server will start on:
http://localhost:8080

---

## 📘 Swagger API Documentation

Generate Swagger docs:
`make swagger`


Open Swagger UI:
http://localhost:8080/swagger/index.html

## 🛠️ Makefile Commands
```
make run       # Start the server
make swagger   # Generate Swagger documentation
```

## 📌 Technologies Used
- Go (Golang)
- Gin Web Framework
- Swaggo (Swagger)
- godotenv

## 🤝 Contributing
Contributions, suggestions, and improvements are welcome.
Feel free to fork the repository and submit a pull request.

## 📄 License

This project is open-source and available under the MIT License.

## 👨‍💻 Author

## 👨‍💻 Author

**Hrishab Pachange**  
AI/ML & Backend Developer | Plainsurf Solutions Pvt. Ltd.

- 💻 Backend Development (Golang, REST APIs, MVC)
- 🤖 AI/ML Engineering
- 🧩 Scalable system & API design
- 📄 Open-source & backend tooling enthusiast

GitHub: https://github.com/hrishabpachange