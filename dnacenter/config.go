package dnacenter

import (
	"context"

	dnac "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	baseURL := d.Get("base_url").(string)
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	debug := d.Get("debug").(string)
	sslVerify := d.Get("ssl_verify").(string)

	var diags diag.Diagnostics

	if (username != "") && (password != "") && (baseURL != "") {
		c, err := dnac.NewClientWithOptions(baseURL, username, password, debug, sslVerify)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create DNA Center client",
				Detail:   "Unable to authorize user for DNA Center client",
			})
			return nil, diags
		}

		return c, diags
	}

	c, err := dnac.NewClient()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create DNA Center client",
			Detail:   "Unable to auth user for DNA Center client. Required values are baseURL, username, password",
		})
		return nil, diags
	}

	return c, diags
}
