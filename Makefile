build: compile
	wasm-tools component embed ./wit poll.module.wasm --output poll.embed.wasm
	wasm-tools component new poll.embed.wasm -o poll.wasm --adapt wit/adapters/tier2/wasi_snapshot_preview1.wasm

bindings:
	wit-bindgen tiny-go --out-dir poll_out ./wit

compile: bindings
	tinygo build -target=wasi -o poll.module.wasm main.go

clean:
	rm -rf poll_out
	rm *.wasm

delete-worker:
	golem-cli worker delete -p golem-poll -t poll --worker-name poll-tester

upload:
	golem-cli template add -p golem-poll -t golempoll poll.wasm
	golem-cli worker add -p golem-poll -t golempoll --worker-name poll-tester