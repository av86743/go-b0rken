package	inner

import(
	"log"

	// compiles
// _ldb	"github.com/syndtr/goleveldb/leveldb"
// _opt	"github.com/syndtr/goleveldb/leveldb/opt"
	// compilation fails, go1.12.5
 _ldb	"github.com/av86743/go-leveldb/leveldb"
 _opt	"github.com/av86743/go-leveldb/leveldb/opt"
)

func Foo() {
	dsn := "/tmp/no-no-no"
	opt := &_opt.Options{}
	db, eopen := _ldb.OpenFile(dsn, opt)
	if eopen != nil {
		log.Printf("dsn %#v\n", dsn)
		log.Fatalf("leveldb.OpenFile: %s", eopen.Error())
	}
	defer db.Close()
}
