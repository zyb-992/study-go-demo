package client

import (
	"io"
	"net/http"
)

type SelfClient struct {
}

func UseClient() {
	defaultClient := http.DefaultClient
	defaultClient.Post()
	var b io.Reader
	r, err := http.NewRequest(http.MethodPost, "/hello", b)

}
