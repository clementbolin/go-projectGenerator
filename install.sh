if [[ $EUID -ne 0 ]]; then
    echo "You must run this with superuser priviliges.  Try \"sudo ./install\"" 2>&1
    exit 1
else
    echo "Installing go-ProjectGenerator..."
fi
make build
sudo cp ./bin/go-projectGenerator /usr/local/bin
make clean
make clean-cache-mod
echo "go-ProjectGenerator install"
