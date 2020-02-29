package go_printify

import (
	"fmt"
	"net/http"
	"time"
)

const (
	productsPath       = "shops/%d/products.json"
	productPath        = "shops/%d/products/%d.json"
	publishProductPath = "shops/%d/products/%d/publish.json"
	publishSuccessPath = "shops/%d/products/%d/publishing_succeeded.json"
	publishFailedPath  = "shops/%d/products/%d/publishing_failed.json"
	unpublishPath = "shops/%d/products/%d/unpublish.json"
)

type Product struct {
	Id                     int                      `json:"id"`
	Title                  string                   `json:"title"`
	Description            string                   `json:"description"`
	Tags                   []string                 `json:"tags"`
	Options                []map[string]interface{} `json:"options"`
	Variants               []ProductVariant         `json:"variants"`
	Images                 []ProductMockUpImage     `json:"images"`
	CreatedAt              time.Time                `json:"created_at"`
	UpdatedAt              time.Time                `json:"updated_at"`
	Visible                bool                     `json:"visible"`
	BlueprintId            int                      `json:"blueprint_id"`
	PrintProviderId        int                      `json:"print_provider_id"`
	UserId                 int                      `json:"user_id"`
	ShopId                 int                      `json:"shop_id"`
	PrintAreas             []PrintArea              `json:"print_areas"`
	PrintDetails           PrintDetails             `json:"print_details"`
	External               []External               `json:"external"`
	IsLocked               bool                     `json:"is_locked"`
	SalesChannelProperties []string                 `json:"sales_channel_properties"`
}

type ProductVariant struct {
	Id          int     `json:"id"`
	Sku         string  `json:"sku"`
	Price       float32 `json:"price"`
	Cost        float32 `json:"cost"`
	Title       string  `json:"title"`
	Grams       int     `json:"grams"`
	IsEnabled   bool    `json:"is_enabled"`
	InStock     bool    `json:"in_stock"` // Deprecated
	IsDefault   bool    `json:"is_default"`
	IsAvailable bool    `json:"is_available"`
	Options     []int   `json:"options"`
}

type ProductMockUpImage struct {
	Src        string `json:"src"`
	VariantIds int    `json:"variant_ids"`
	Position   string `json:"position"`
	IsDefault  bool   `json:"is_default"`
}

type ProducePlaceholder struct {
	Position string         `json:"position"`
	Images   []ProductImage `json:"images"`
}

type ProductImage struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Scale  int    `json:"scale"`
	Angle  int    `json:"angle"`
}

type PrintArea struct {
}

type PrintDetails struct {
	PrintOnSide string `json:"print_on_side"`
}

type PublishingProperties struct {
	Images      bool `json:"images"`
	Variants    bool `json:"variants"`
	Title       bool `json:"title"`
	Description bool `json:"description"`
	Tags        bool `json:"tags"`
}

type External struct {
	Id     int    `json:"id"`
	Handle string `json:"handle"`
}

/*
Retrieve a list of products
*/
func (c *Client) GetProducts(shopId int, page int) ([]*Product, error) {
	path := fmt.Sprintf(productsPath, shopId)
	if page != 1 {
		path = fmt.Sprintf("%s?page=%d", path, page)
	}
	req, err := c.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	products := make([]*Product, 0)
	_, err = c.do(req, products)
	return products, err
}

/*
Retrieve a product
*/
func (c *Client) GetProduct(shopId, productId int) (*Product, error) {
	path := fmt.Sprintf(productPath, shopId, productId)
	req, err := c.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	product := &Product{}
	_, err = c.do(req, product)
	return product, err
}

/*
Create a new product
*/
func (c *Client) CreateProduct(product Product) error {
	req, err := c.newRequest(http.MethodPost, productsPath, product)
	if err != nil {
		return err
	}
	_, err = c.do(req, product)
	return err
}

/*
Update a product
*/
func (c *Client) UpdateProduct(shopId int, product Product) (*Product, error) {
	path := fmt.Sprintf(productPath, shopId, product.Id)
	req, err := c.newRequest(http.MethodPut, path, product)
	if err != nil {
		return nil, err
	}
	updatedProduct := &Product{}
	_, err = c.do(req, updatedProduct)
	return updatedProduct, err
}

/*
Delete a product
*/
func (c *Client) DeleteProduct(shopId int, productId int) error {
	path := fmt.Sprintf(productPath, shopId, productId)
	req, err := c.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}
	_, err = c.do(req, nil)
	return err
}

/*
Publish a product
*/
func (c *Client) PublishProduct(shopId, productId int, publishProperties PublishingProperties) error {
	path := fmt.Sprintf(publishProductPath, shopId, productId)
	req, err := c.newRequest(http.MethodPost, path, publishProperties)
	if err != nil {
		return err
	}
	_, err = c.do(req, nil)
	return err
}

/*
Set product publish status to succeeded
*/
func (c *Client) SetProductPublishSuccess(shopId, productId int, external External) error {
	path := fmt.Sprintf(publishSuccessPath, shopId, productId)
	req, err := c.newRequest(http.MethodPost, path, external)
	if err != nil {
		return err
	}
	_, err = c.do(req, nil)
	return err
}

/*
Set product publish status to failed
 */
func (c *Client) SetProductPublishFailre(shopId, productId int, reason string) error {
	path := fmt.Sprintf(publishFailedPath, shopId, productId)
	req, err := c.newRequest(http.MethodPost, path, map[string]string{"reason": reason})
	if err != nil {
		return err
	}
	_, err = c.do(req, nil)
	return err
}

/*
Notify that a product has been unpublished
 */
func (c *Client) Unpublish(shopId, productId int) error {
	path := fmt.Sprintf(unpublishPath, shopId, productId)
	req, err := c.newRequest(http.MethodPost, path, nil)
	if err != nil {
		return err
	}
	_, err = c.do(req, nil)
	return err
}

