CC=go

make:
	$(CC) install

compile:
	$(CC) build

test:
	$(CC) test ./...
