echo "changing directory to Tasks"
cd $GOPATH/src/github.com/pbillerot/graduel/resources
echo "creating table"
cat schema.sql | sqlite3 /mnt/nas/data/graduel/picsou.sqlite
echo "building the go binary"
go build -o ../bin/graduel

echo "starting the binary"
./bin/graduel
