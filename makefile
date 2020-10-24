build:
	command -v go >/dev/null 2>&1 || { echo >&2 "I require GoLang but it's not installed.  Aborting."; exit 1; }
	go build .
	@echo -e "\e[32m\e[1m!\e[0m\e[32m Installation Complete! Run \`make install` to complete the installation.\e[0m"
install:
	cp ./dos /usr/bin/dos
update:
	@git pull origin master && make && make install
