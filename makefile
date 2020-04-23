install:
	chmod u+x && ./install.sh
update:
	git pull origin master && chmod u+x && ./install.sh