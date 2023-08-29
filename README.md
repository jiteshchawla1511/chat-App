# chat-App
Chat application using GO Lang, React and Redis 

## Prerequisites

- [Go](https://golang.org/dl/) installed on your system.
- [Free Redis Cloud Account](https://redis.com/try-free/)
- Node Js

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/jiteshchawla1511/chat-App.git
   ```

2. Navigate to the project directory:
   
   ```bash
   cd chat-App
   ```

3. Install the necessary dependencies:
   
   ```bash
   go mod download
   ```
   ```bash
   cd client
   npm install
   ```

5. Run the main file

  1) Terminal 1
     ```bash
     go run main.go --server=http
     ```

  2) Terminal 2
     ```bash
     go run main.go --server=websocket
     ```

  3) Terminal 3
     ```bash
     cd client
     ```
     ```bash
     npm start
     ```
     
     
##  Explanation

1. User: RedisJSON, SET, SortedSET, Key-Value pair (features: register, login, verify-contact, contact list)
2. Chat: RedisJSON, Redisearch (features: new chat, chat-history)


## ENV FILE 

REDIS_CONNECTION_STRING = from redis cloud

REDIS_PASSWORD = from redis cloud


Thankyou 😄
