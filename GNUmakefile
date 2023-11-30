default: testacc

# Run acceptance tests
.PHONY: testacc, cover
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

cover:
	go test -v -coverprofile cover.out ./...
	go tool cover -html cover.out -o cover.html
