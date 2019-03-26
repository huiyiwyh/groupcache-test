package groupcachetest

import (
	"fmt"

	"github.com/golang/groupcache"
)

func getDataFromDatabase(ctx groupcache.Context, key string, dest groupcache.Sink) error {
	fmt.Println("get data from database")
	dest.SetString("lalalalala")
	return nil
}

// client means just gets data from databse and other peers, but not provide.
func client() {

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

	macbookPro := groupcache.NewGroup("MacbookPro", 1<<20, groupcache.GetterFunc(getDataFromDatabase))

	var data []byte
	macbookPro.Get(nil, "ytwer", groupcache.AllocatingByteSliceSink(&data))
	fmt.Println("the value is ", string(data))
}
