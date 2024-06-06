### Getting Started
1. create a `.env` file and attach your OpenAI API key
2. make sure you have docker setup and run `docker-compose up`
3. call `/recommendations`

example: 
```
curl http://localhost:8080/recommendations?prompt="i want mexican food"
```