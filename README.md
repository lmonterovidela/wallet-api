# wallet-api

wallet-api allows the debit and credit from accounts

Run api:
- Download the code
- Run in a terminal: 'docker-compose -f docker-compose.db.yml up' (that up a mysqldatabase)
- Run make up or go build cmd/web/*.go


Run whole api on docker:
- Download the code
- Run in a terminal: docker-compose up

Run linter:
- golangci-lint run ./...

Run test:
- docker-compose -f docker-compose.test.yml up
- @go test ./...
- docker-compose -f docker-compose.test.yml down
