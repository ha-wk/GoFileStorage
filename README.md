# Wobot File Storage API

 How anyone can run the Project?????

1. Clone the repo:
git clone https://github.com/ha-wk/WobotAIAssignment
and then do cd wobot-file-storage

2. Initialize the Go module:
go mod init wobot-file-storage
go mod tidy

3. Run the server:
go run cmd/main.go

4. Use Postman or any API tool to test the following endpoints:

- `POST /register` – to create a user  
  (Send JSON: `{ "username": "yourname", "password": "pass" }`)

- `POST /login` – to get the JWT token

- Use the JWT token in the `Authorization` header for:
  - `POST /upload` – upload a file (as `multipart/form-data`)
  - `GET /storage/remaining` – check your storage usage
  - `GET /files` – list your uploaded files

We can discuss if you find any issues in running above project and make request.
