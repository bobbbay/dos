build:
	command -v go >/dev/null 2>&1 || { echo >&2 "I require GoLang but it's not installed.  Aborting."; exit 1; }
	go build .
	@echo "Run make install to complete the installation."
install:
	cp ./dos /usr/bin/dos
update:
	@git pull origin master && make && make install
