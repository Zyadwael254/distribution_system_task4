# distribution_system_task4
distribution_system_task4
# Simple Chatroom using Go RPC
https://drive.google.com/file/d/1zsEDXJ3JPEt2PgdAGLDz-b1bcrHhSHZV/view?usp=sharing

## Description
This project is a simple chatroom built with **Go** using **RPC (Remote Procedure Call)**.  
The server receives messages from clients and displays them on its terminal.  
Clients can send messages but do not see the chat history.

---

## How It Works
- **Server:** Runs on port `1234`, receives messages, and prints them.
- **Client:** Connects to the server and sends messages until typing `exit`.

---

## Run Instructions

### 1️⃣ Start the server
```bash
go run server.go
```
### 1️⃣ Start the client on diff terminals
```bash
go run client.go
```
## video
https://drive.google.com/file/d/1zsEDXJ3JPEt2PgdAGLDz-b1bcrHhSHZV/view?usp=sharing
