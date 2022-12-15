# Environment variables
TINY_GO_PACKAGE = https://github.com/tinygo-org/tinygo/releases/download/v0.26.0/tinygo_0.26.0_amd64.deb

user = Ignema
repo = goseek


# Declaring phony targets
.PHONY: init env wasm wasi dev publish clean


# Recipe definitions
all: init clean wasm env dev

init:
	npm ci
	wget -O ./tinygo.deb ${TINY_GO_PACKAGE}
	sudo dpkg -i tinygo.deb
	rm *.deb

env: env_user env_repo env_token

env_user:
	sed -i "s/\(GITHUB_USER = \"\)..*/\1$(user)\"/g" wrangler.toml

env_repo:
	sed -i "s/\(GITHUB_REPO = \"\)..*/\1$(repo)\"/g" wrangler.toml

env_token:
	sed -i "s/\(GITHUB_TOKEN = \"\)..*/\1$(token)\"/g" wrangler.toml

wasm:
	mkdir -p bin
	cd src/function && go get -d . && tinygo build -o ../../bin/worker.wasm -target=wasm -gc=leaking -no-debug -opt=2 ./worker.go

wasi:
	mkdir -p bin
	cd src/function && go get -d . && tinygo build -o ../../bin/wasi.wasm -wasm-abi=generic -target=wasi ./worker.go

dev:
	npx wrangler dev src/util/wrapper.mjs

publish:
	npx wrangler publish

clean:
	rm -f bin/*