#setup go

echo "export GOROOT=/usr/local/go" >> ~/.bash_profile
echo "export GOPATH=\$HOME/go" >> ~/.bash_profile
echo "export PATH=\$PATH:\$GOROOT/bin:\$GOPATH/bin" >> ~/.bash_profile
. ~/.bash_profile

go get github.com/revel/revel
go get github.com/revel/cmd/revel
revel help
