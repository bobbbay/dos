# Check if GoLang is installed
command -v go >/dev/null 2>&1 || { echo >&2 "I require GoLang but it's not installed.  Aborting."; exit 1; }
# If it is, build the script
go build .
# Be nice
echo -e "\e[32m\e[1m!\e[0m\e[32m Installation Complete!\e[0m"
# Lastly, tell the user to add this to their PATH
echo -e "\e[33m\e[1m! Please add $PWD to your \$PATH. If you don't know how to do this, it is recommended against using this tool.\e[0m"