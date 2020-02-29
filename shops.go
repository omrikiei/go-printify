package go_printify

import (
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"net/http"
)

const (
	shopsPath          = "shops.json"
	disconnectShopPath = "shops/%d/connection.json"
)

type Shop struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	SalesChannel string `json:"sales_channel"`
}

/*
Retrieve list of shops in a Printify account
*/
func (c *Client) ListShops() ([]*Shop, error) {
	req, err := c.newRequest(http.MethodGet, shopsPath, nil)
	if err != nil {
		return nil, err
	}
	shopList := make([]*Shop, 0)
	_, err = c.do(req, shopList)
	return shopList, err
}

/*
Disconnect a shop
*/
func (c *Client) DeleteShop(Id int) error {
	req, err := c.newRequest(http.MethodDelete, fmt.Sprint(disconnectShopPath, Id), nil)
	if err != nil {
		return err
	}
	_, err = c.do(req, struct{}{})
	return err
}
