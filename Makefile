GC=go

make:
	$(GC) install

compile:
	$(GC) build

test:
	$(GC) test ./...

.PHONY: clean
clean:
	-rm feel.exe
