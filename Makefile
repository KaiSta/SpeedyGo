# Build tooling for SpeedyGo.
#
#   make build     build the CLI for the host platform (./speedygo)
#   make release   cross-compile release binaries into dist/
#   make clean     remove build output

BINARY   := speedygo
MAIN     := ./traceReplay
DIST     := dist

# GOOS/GOARCH pairs built by "make release".
PLATFORMS ?= linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64 windows/arm64

GOFLAGS ?=
LDFLAGS ?= -s -w

EXE := $(if $(filter windows,$(GOOS)),.exe,)

.PHONY: all build release clean

all: build

build:
	go build $(GOFLAGS) -trimpath -ldflags '$(LDFLAGS)' -o $(BINARY)$(EXE) $(MAIN)

release:
	@mkdir -p $(DIST)
	@set -e; for p in $(PLATFORMS); do \
		os=$${p%/*}; arch=$${p#*/}; \
		out="$(DIST)/$(BINARY)-$$os-$$arch"; \
		if [ "$$os" = "windows" ]; then out="$$out.exe"; fi; \
		echo "building $$out"; \
		GOOS=$$os GOARCH=$$arch CGO_ENABLED=0 \
			go build $(GOFLAGS) -trimpath -ldflags '$(LDFLAGS)' -o "$$out" $(MAIN); \
	done

clean:
	rm -rf $(BINARY) $(BINARY).exe $(DIST)
