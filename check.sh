#~/bin/sh
echo '[FirmaChain-Test start]'
go test ./...
echo '[FirmaChain-Test finish]'
echo '[golangci-lint start]'
golangci-lint run
echo '[golangci-lint finish]'