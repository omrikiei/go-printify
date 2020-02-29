package go_printify

import (
	"fmt"
	"net/http"
)

const (
	blueprintsPath                = "catalog/blueprints.json"
	blueprintPath                 = "catalog/blueprints/%d.json"
	blueprintProvidersPath        = "catalog/blueprints/%d/print_providers.json"
	BlueprintProviderVariantsPath = "catalog/blueprints/%d/print_providers/%d/variants.json"
	BluePrintProviderShippingPath = "catalog/blueprints/%d/print_providers/%d/shipping.json"
	PrintProvidersPath            = "catalog/print_providers.json"
	PrintProviderPath             = "catalog/print_providers/%d.json"
)

type Blueprint struct {
	Id     int      `json:"id"`
	Title  string   `json:"title"`
	Brand  string   `json:"brand"`
	Model  string   `json:"model"`
	Images []string `json:"images"`
}

type PrintProvider struct {
	Id       int               `json:"id"`
	Title    string            `json:"title"`
	Location map[string]string `json:"location"`
	Variants []CatalogVariant  `json:"variants"`
}

type CatalogVariant struct {
	Id           int                    `json:"id"`
	Title        string                 `json:"title"`
	Options      []CatalogVariantOption `json:"options"`
	Placeholders []CatalogPlaceholder   `json:"placeholders"`
}

type CatalogVariantOption struct {
	Color string `json:"color"`
	Size  string `json:"size"`
}

type CatalogPlaceholder struct {
	Position string `json:"position"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}

type ShippingProperties struct {
	HandlingTime struct {
		Value int    `json:"value"`
		Unit  string `json:"unit"`
	} `json:"handling_time"`
	Profiles struct {
		VariantIds      []int    `json:"variant_ids"`
		FirstItem       priceTag `json:"first_item"`
		AdditionalItems priceTag `json:"additional_items"`
		Countries       []string `json:"countries"`
	} `json:"profiles"`
}

type priceTag struct {
	Currency string  `json:"currency"`
	Cost     float32 `json:"cost"`
}

/*
Retrieve a list of available blueprints
*/
func (c *Client) ListBluePrints() ([]*Blueprint, error) {
	req, err := c.newRequest(http.MethodGet, blueprintsPath, nil)
	if err != nil {
		return nil, err
	}
	blueprintList := make([]*Blueprint, 0)
	_, err = c.do(req, blueprintList)
	return blueprintList, err
}

/*
Retrieve a specific blueprint
*/
func (c *Client) GetBlueprint(Id int) (*Blueprint, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf(blueprintPath, Id), nil)
	if err != nil {
		return nil, err
	}
	blueprint := &Blueprint{}
	_, err = c.do(req, blueprint)
	return blueprint, err
}

/*
Retrieve a list of all print providers that fulfill orders for a specific blueprint
*/
func (c *Client) GetPrintProviders(b *Blueprint) ([]*PrintProvider, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf(blueprintProvidersPath, b.Id), nil)
	if err != nil {
		return nil, err
	}
	providers := make([]*PrintProvider, 0)
	_, err = c.do(req, providers)
	return providers, err
}

/*
Retrieve a list of variants of a blueprint from a specific print provider
*/
func (c *Client) GetVariants(b *Blueprint, p *PrintProvider) ([]*CatalogVariant, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf(BlueprintProviderVariantsPath, b.Id, p.Id), nil)
	if err != nil {
		return nil, err
	}
	variants := make([]*CatalogVariant, 0)
	_, err = c.do(req, variants)
	return variants, err
}

/*
Retrieve shipping information
*/
func (c *Client) GetShippingInformation(b *Blueprint, p *PrintProvider) (*ShippingProperties, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf(BluePrintProviderShippingPath, b.Id, p.Id), nil)
	if err != nil {
		return nil, err
	}
	shippingInfo := &ShippingProperties{}
	_, err = c.do(req, shippingInfo)
	return shippingInfo, err
}

/*
Retrieve a list of available print providers
*/
func (c *Client) GetAvailablePrintProviders() ([]*PrintProvider, error) {
	req, err := c.newRequest(http.MethodGet, PrintProvidersPath, nil)
	if err != nil {
		return nil, err
	}
	providers := make([]*PrintProvider, 0)
	_, err = c.do(req, providers)
	return providers, err
}

/*
Retrieve a specific print provider and a list of associated blueprint offerings
*/
func (c *Client) GetPrintProvider(Id int) (*PrintProvider, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf(PrintProviderPath, Id), nil)
	if err != nil {
		return nil, err
	}
	provider := &PrintProvider{}
	_, err = c.do(req, provider)
	return provider, err
}
