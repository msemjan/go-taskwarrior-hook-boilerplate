all: build

build:
	go build .

install: build
	mv on-modify.add-done-tag ~/.task/hooks/

uninstall:
	rm ~/.task/hooks/on-modify.add-done-tag

clean:
	$(RM) on-modify.add-done-tag
