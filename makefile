.PHONY: init
init:
	wget https://github.com/tinygo-org/tinygo/releases/download/v0.26.0/tinygo_0.26.0_amd64.deb
	sudo dpkg -i tinygo_0.26.0_amd64.deb
	rm *.deb

.PHONY: env
env:
	sed -i "s/\(GITHUB_TOKEN = \"\)..*/\1$(token)\"/g" wrangler.toml

.PHONY: wasm
wasm:
	mkdir -p bin
	cd src/function && go get -d . && tinygo build -o ../../bin/worker.wasm -target=wasm -gc=leaking -no-debug -opt=2 ./worker.go

.PHONY: wasi
wasi:
	mkdir -p bin
	cd src/function && go get -d . && tinygo build -o ../../bin/wasi.wasm -wasm-abi=generic -target=wasi ./worker.go

.PHONY: dev
dev:
	npx wrangler dev src/util/wrapper.mjs

.PHONY: publish
publish:
	npx wrangler publish

.PHONY: clean
clean:
	rm -f bin/*