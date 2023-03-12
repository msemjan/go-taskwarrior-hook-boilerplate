all: build

build:
	go build .

install: build
	mv on-modify.log-changes ~/.task/hooks/

uninstall:
	rm ~/.task/hooks/on-modify.log-changes 

clean:
	$(RM) on-modify.log-changes
