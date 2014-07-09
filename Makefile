alldirs=$(shell find . \( -path ./Godeps -o -path ./.git \) -prune -o -type d -print)
GODIRS=$(foreach dir, $(alldirs), $(if $(wildcard $(dir)/*.go),$(dir)))
GOFILES=$(foreach dir, $(alldirs), $(wildcard $(dir)/*.go))

ifeq ("$(WERCKER)", "true")
BUILD_TAGS="wercker"	
endif

all: build

lint:
	golint $(GOFILES)
	go vet ./...

test:
	godep go test -tags $(BUILD_TAGS) -v ./...

coverage:
	@echo "mode: set" > acc.out
	@for dir in $(GODIRS); do \
		cmd="godep go test -tags $(BUILD_TAGS) -v -coverprofile=profile.out $$dir"; \
		eval $$cmd; \
		if test $$? -ne 0; then \
			exit 1; \
		fi; \
		if test -f profile.out; then \
			cat profile.out | grep -v "mode: set" >> acc.out; \
		fi; \
	done
	@rm -f ./profile.out

build:
	godep go build -tags $(BUILD_TAGS) -v ./...

clean:
	rm -f ./goscribe ./gin-bin ./acc.out

save:
	godep save ./...
