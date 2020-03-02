package go_printify

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func TestClient_ListShops(t *testing.T) {
	shops := []*Shop{
		{
			5432,
			"My new store",
			"My Sales Channel",
		},
		{
			9876,
			"My other new store",
			"disconnected",
		},
	}

	s := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(200)
		_, _ = rw.Write([]byte(`[
			{
			  "id": 5432,
			  "title": "My new store",
			  "sales_channel": "My Sales Channel"
			},
			{
			  "id": 9876,
			  "title": "My other new store",
			  "sales_channel": "disconnected"
			}
		]
		`))
	}))
	serverUrl, _ := url.Parse(s.URL)
	client := NewClient("bla")
	client.BaseURL = serverUrl
	defer s.Close()
	shopsRes, err := client.ListShops()
	if err != nil {
		fmt.Println(err)
	}
	if !reflect.DeepEqual(shopsRes, shops) {
		fmt.Println(shopsRes, shops)
		t.Fail()
	}

}

func TestClient_DeleteShop(t *testing.T) {
	// nothing to test here really
}
