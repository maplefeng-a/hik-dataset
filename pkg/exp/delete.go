package exp

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Delete(name string) io.ReadCloser {
	url := "https://cn.bing.com/search?" + name
	var ret io.ReadCloser
	if resp, err := http.Get(url); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		ret = resp.Body
	}
	return ret
}
