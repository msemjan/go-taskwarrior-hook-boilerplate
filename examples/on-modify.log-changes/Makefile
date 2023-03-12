all: build

build:
	go build .

install: build
	mv on-modify.log-changes ~/.task/hooks/

clean:
	$(RM) on-modify.log-changes
