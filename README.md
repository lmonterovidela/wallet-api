# wallet-api

wallet-api allows the debit and credit from accounts

Run api:
- Download the code
- Run in a terminal: 'docker-compose up' (that up a mysqldatabase and redis)
- Run in a terminal: go build cmd/web/*.go

Run linter:
- golangci-lint run ./...

Run test:
- docker-compose -f docker-compose.test.yml up (up database to test repository)
- go test ./...
- docker-compose -f docker-compose.test.yml down (down database to test repository)
