package dnacenter

import (
	"context"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Config is the configuration structure used to instantiate a
// new Cisco DNA Center client.
type Config struct {
	BaseURL   string
	Username  string
	Password  string
	Debug     string
	SSLVerify string
}

// NewClient returns a new Cisco DNA Center client.
func (c *Config) NewClient() (*dnacentersdkgo.Client, error) {
	client, err := dnacentersdkgo.NewClientWithOptions(c.BaseURL,
		c.Username, c.Password,
		c.Debug, c.SSLVerify,
	)
	if err != nil {
		return client, err
	}
	client.RestyClient().SetLogger(createLogger())
	return client, err
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	config := Config{
		BaseURL:   d.Get("base_url").(string),
		Username:  d.Get("username").(string),
		Password:  d.Get("password").(string),
		Debug:     d.Get("debug").(string),
		SSLVerify: d.Get("ssl_verify").(string),
	}

	client, err := config.NewClient()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Cisco DNA Center client",
			Detail:   err.Error(),
		})
		return nil, diags
	}
	return client, diags
}
