package indexeddb_test

import (
	"github.com/gowebapi/webapi"
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/indexeddb"
)

func Example_Open() {
	request := webapi.GetWindow().IndexedDB().Open("test", nil)
	onsucess := func(event *domcore.Event) interface{} {
		return nil
	}
	request.SetOnsuccess(domcore.EventHandlerToJS(onsucess))
	request.SetOnupgradeneeded(domcore.EventHandlerToJS(func(event *domcore.Event) interface{} {
		db := indexeddb.IDBDatabaseFromJS(request.Result())
		db.CreateObjectStore("names", &indexeddb.IDBObjectStoreParameters{
			AutoIncrement: true,
		})
		/* javascript version is:
		   var db = event.target.result;
		   var objStore = db.createObjectStore("names", { autoIncrement : true });
		*/
		return nil
	}))
}
