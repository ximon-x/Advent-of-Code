goony-run:
	@echo "Executing Go Solution... \n"
	cd goony && go run cmd/main.go

rusty-run:
	@echo "Executing Rust Solution... \n"
	cd rusty && cargo run src/main.rs

ziggy-run:
	@echo "Executing Zig Solution... \n"
	cd ziggy && zig run src/main.zig
