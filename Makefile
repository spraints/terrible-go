PKGS=enum
GO=$(CURDIR)/bin/run go

.PHONY: test
test:
	@set -e; \
	for d in $(PKGS); do \
		echo + test $$d; \
		(cd $$d && $(GO) test .); \
	done

.PHONY: fmt
fmt:
	@echo + gofmt ./...
	@$(GO) fmt ./...

.PHONY: tidy
tidy:
	@set -e; \
	for d in $(PKGS); do \
		echo + tidy $$d; \
		(cd $$d && $(GO) mod tidy); \
	done
