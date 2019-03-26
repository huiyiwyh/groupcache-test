package groupcachetest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang/groupcache"
)

func serverGetData(ctx groupcache.Context, key string, dest groupcache.Sink) error {
	fmt.Println("get data from database")
	dest.SetString("yuyuyuyu")
	return nil
}

func server() {

	local_addr := "http://191.167.1.37:8080"

	peers_addr := []string{
		"http://191.167.1.37:8080",
		"http://191.167.1.111:8080",
		"http://191.167.1.38:8080",
		"http://191.167.1.39:8080",
		"http://191.167.1.112:8080",
	}

	peers := groupcache.NewHTTPPool(local_addr)
	peers.Set(peers_addr...)

	macbookPro := groupcache.NewGroup("MacbookPro", 1<<20, groupcache.GetterFunc(serverGetData))

	var data []byte
	macbookPro.Get(nil, "ytwer", groupcache.AllocatingByteSliceSink(&data))
	fmt.Println("the value is ", string(data))

	time.Sleep(time.Second * 2)

	var dataNext []byte
	macbookPro.Get(nil, "ytwer", groupcache.AllocatingByteSliceSink(&dataNext))
	fmt.Println("the value is ", string(dataNext))

	go http.ListenAndServe(":8080", nil)
	select {}
}
