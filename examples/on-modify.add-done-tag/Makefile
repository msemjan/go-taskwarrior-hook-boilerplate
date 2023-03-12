BIN_NAME=on-modify.add-done-tag

all: build

build:
	go build .

install: build
	mv "${BIN_NAME}" ~/.task/hooks/

uninstall:
	rm "~/.task/hooks/${BIN_NAME}"

clean:
	$(RM) "${BIN_NAME}"
