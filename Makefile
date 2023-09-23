PORT ?= 3000

## test: run tests
.PHONY: test
test:
	go test -v

## testfuzz: run tests with fuzz (30 seconds)
.PHONY: testfuzz
testfuzz:
	go test -v -race -fuzz=Fuzz -fuzztime 30s .

## bench: run benchmark tests
.PHONY: bench
bench:
	go test -v -run none -bench Reverse -benchtime 3s

## doc: run godoc server at 3000 unless PORT env-var is set
.PHONY: doc
doc:
	@echo "open http://localhost:$(PORT)/pkg/github.com/mkdemir/vigilsys/"
	godoc -http=:$(PORT)

.PHONY: help
all: help
help: Makefile
	@echo
	@echo "  make <command>"
	@echo
	@echo "  commands:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
