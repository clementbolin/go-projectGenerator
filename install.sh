if [[ $EUID -ne 0 ]]; then
    echo "You must run this with superuser priviliges.  Try \"sudo ./install\"" 2>&1
    exit 1
else
    echo "Installing go-ProjectGenerator..."
fi
go build -o go-ProjectGenerator
sudo cp ./go-projectGenerator /usr/local/bin
rm go-ProjectGenerator
go clean --modcache
echo "go-ProjectGenerator install"
