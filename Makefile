# GO minimum version check.
GO_MIN_VERSION := 11600 # 1.16
GO_VERSION_CHECK := \
  $(shell expr \
    $(shell go version | \
      awk '{print $$3}' | \
      cut -do -f2 | \
      sed -e 's/\.\([0-9][0-9]\)/\1/g' -e 's/\.\([0-9]\)/0\1/g' -e 's/^[0-9]\{3,4\}$$/&00/' \
    ) \>= $(GO_MIN_VERSION) \
  )


test: check-go
	@go clean -testcache
	go test -race -timeout=30s $(TESTARGS) ./...
.PHONY: test

cover: check-go
	@$(MAKE) test TESTARGS="-tags test -coverprofile=coverage.out"
	@go tool cover -html=coverage.out
	@rm -f coverage.out
.PHONY: cover

gen: check-go
	go generate ./...
.PHONY: gen

clean:
	@ls build | grep -v .gitkeep | xargs -I {} rm build/{}
.PHONY: clean

check-go:
ifeq ($(GO_VERSION_CHECK),0)
$(error go1.16 or higher is required)
endif
.PHONY: check-go