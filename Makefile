.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

.PHONY: gaps
gaps:
	@GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap ./gaps/tree/.
	@zip ./builds/gaps/tree.zip -r bootstrap
	@rm bootstrap