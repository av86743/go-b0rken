module github.com/av86743/go-b0rken

go 1.12

replace (
	_ => ./_
	//	github.com/av86743/go-b0rken => ../go-b0rken
	github.com/av86743/go-leveldb => github.com/syndtr/goleveldb v1.0.1-0.20190318030020-c3a204f8e965
)

require (
	github.com/av86743/go-leveldb v0.0.0-00010101000000-000000000000
	github.com/syndtr/goleveldb v1.0.0 // indirect
)
