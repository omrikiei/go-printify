package go_printify

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func TestClient_ListBluePrints(t *testing.T) {
	blueprints := []*Blueprint{
		{
			Id:          3,
			Title:       "Kids Regular Fit Tee",
			Description: "Description goes here",
			Brand:       "Delta",
			Model:       "11736",
			Images: []string{
				"https://images.printify.com/5853fe7dce46f30f8327f5cd",
				"https://images.printify.com/5c487ee2a342bc9b8b2fc4d2",
			},
		},
		{
			Id:          5,
			Title:       "Men's Cotton Crew Tee",
			Description: "Description goes here",
			Brand:       "Next Level",
			Model:       "3600",
			Images: []string{
				"https://images.printify.com/5a2ffc81b8e7e3656268fb44",
				"https://images.printify.com/5cdc0126b97b6a00091b58f7",
			},
		},
		{
			Id:          6,
			Title:       "Unisex Heavy Cotton Tee",
			Description: "Description goes here",
			Brand:       "Gildan",
			Model:       "5000",
			Images: []string{
				"https://images.printify.com/5a2fd7d9b8e7e36658795dc0",
				"https://images.printify.com/5c595436a342bc1670049902",
				"https://images.printify.com/5c595427a342bc166b6d3002",
				"https://images.printify.com/5a2fd022b8e7e3666c70623a",
			},
		},
	}
	s := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(200)
		_, _ = rw.Write([]byte(`
[
    {
        "id": 3,
        "title": "Kids Regular Fit Tee",
        "description": "Description goes here",
        "brand": "Delta",
        "model": "11736",
        "images": [
            "https://images.printify.com/5853fe7dce46f30f8327f5cd",
            "https://images.printify.com/5c487ee2a342bc9b8b2fc4d2"
        ]
    },
    {
        "id": 5,
        "title": "Men's Cotton Crew Tee",
        "description": "Description goes here",
        "brand": "Next Level",
        "model": "3600",
        "images": [
            "https://images.printify.com/5a2ffc81b8e7e3656268fb44",
            "https://images.printify.com/5cdc0126b97b6a00091b58f7"
        ]
    },
    {
        "id": 6,
        "title": "Unisex Heavy Cotton Tee",
        "description": "Description goes here",
        "brand": "Gildan",
        "model": "5000",
        "images": [
            "https://images.printify.com/5a2fd7d9b8e7e36658795dc0",
            "https://images.printify.com/5c595436a342bc1670049902",
            "https://images.printify.com/5c595427a342bc166b6d3002",
            "https://images.printify.com/5a2fd022b8e7e3666c70623a"
        ]
    }]
		`))
	}))
	serverUrl, _ := url.Parse(s.URL)
	client := NewClient("bla")
	client.BaseURL = serverUrl
	defer s.Close()
	resBlueprints, err := client.ListBluePrints()
	if err != nil {
		fmt.Println(err)
	}
	if !reflect.DeepEqual(resBlueprints, blueprints) {
		fmt.Println(resBlueprints, blueprints)
		t.Fail()
	}
}

func TestClient_GetBlueprint(t *testing.T) {
	blueprint := &Blueprint{
		3,
		"Kids Regular Fit Tee",
		"Description goes here",
		"Delta",
		"11736",
		[]string{
			"https://images.printify.com/5853fe7dce46f30f8327f5cd",
			"https://images.printify.com/5c487ee2a342bc9b8b2fc4d2",
		},
	}
	s := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(200)
		_, _ = rw.Write([]byte(`{
     "id": 3,
     "title": "Kids Regular Fit Tee",
     "description": "Description goes here",
     "brand": "Delta",
     "model": "11736",
     "images": [
        "https://images.printify.com/5853fe7dce46f30f8327f5cd",
        "https://images.printify.com/5c487ee2a342bc9b8b2fc4d2"
     ]
}`))
	}))
	client := NewClient("bla")
	serverUrl, _ := url.Parse(s.URL)
	client.BaseURL = serverUrl
	defer s.Close()
	resBlueprint, err := client.GetBlueprint(blueprint.Id)
	if err != nil {
		fmt.Println(err)
	}
	if !reflect.DeepEqual(resBlueprint, blueprint) {
		fmt.Println(resBlueprint, blueprint)
		t.Fail()
	}
}

func TestClient_GetPrintProviders(t *testing.T) {
	printProviders := []*PrintProvider{
		{
			Id:    3,
			Title: "DJ",
		},
		{
			Id:    8,
			Title: "Fifth Sun",
		},
		{
			Id:    16,
			Title: "MyLocker",
		},
		{
			Id:    24,
			Title: "Inklocker",
		},
	}
	s := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(200)
		_, _ = rw.Write([]byte(`[
    {
        "id": 3,
        "title": "DJ"
    },
    {
        "id": 8,
        "title": "Fifth Sun"
    },
    {
        "id": 16,
        "title": "MyLocker"
    },
    {
        "id": 24,
        "title": "Inklocker"
    }
]`))
	}))
	client := NewClient("bla")
	serverUrl, _ := url.Parse(s.URL)
	client.BaseURL = serverUrl
	defer s.Close()
	resPrintProviders, err := client.GetAvailablePrintProviders()
	if err != nil {
		fmt.Println(err)
	}
	if !reflect.DeepEqual(resPrintProviders, printProviders) {
		fmt.Println(resPrintProviders, printProviders)
		t.Fail()
	}
}
