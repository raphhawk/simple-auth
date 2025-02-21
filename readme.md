# Usage of the api

## Run api
```bash 
    # Start Application
    docker-compose up -d

    # Stop Application
    docker-compose down
```
## Test api
```bash
    # Sign Up
    curl -X POST http://localhost:8080/signup -d '{"email":"user@example.com", "password":"securepass"}' -H "Content-Type: application/json"
    
    # Sign In
    curl -X POST http://localhost:8080/signin -d '{"email":"user@example.com", "password":"securepass"}' -H "Content-Type: application/json"
    
    # Refresh Token
    curl -X POST http://localhost:8080/refresh -H "Authorization: <your_jwt_token>"

    # Revoke Token
    curl -X POST http://localhost:8080/revoke -H "Authorization: <your_jwt_token>"
```
