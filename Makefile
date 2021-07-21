APP=aak

serve:
	go run ./cmd/$(APP)/...

dev:
	reflex -s -r '\.go$$' make serve
