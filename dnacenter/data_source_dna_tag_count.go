package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTagCount() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTagCountRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name_space": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"attribute_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"size": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"system_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"response": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceTagCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	tagCountQueryParams := dnac.GetTagCountQueryParams{}
	if v, ok := d.GetOk("name"); ok {
		tagCountQueryParams.Name = v.(string)
	}
	if v, ok := d.GetOk("name_space"); ok {
		tagCountQueryParams.NameSpace = v.(string)
	}
	if v, ok := d.GetOk("attribute_name"); ok {
		tagCountQueryParams.AttributeName = v.(string)
	}
	if v, ok := d.GetOk("level"); ok {
		tagCountQueryParams.Level = v.(string)
	}
	if v, ok := d.GetOk("size"); ok {
		tagCountQueryParams.Size = v.(string)
	}
	if v, ok := d.GetOk("system_tag"); ok {
		tagCountQueryParams.SystemTag = v.(string)
	}

	// Prepare Request
	response, _, err := client.Tag.GetTagCount(&tagCountQueryParams)
	if err != nil {
		return diag.FromErr(err)
	}

	// set response to Terraform data source
	if err := d.Set("response", response.Response); err != nil {
		return diag.FromErr(err)
	}

	// always run, Set resource id
	// Unix time  forces this resource to refresh during every Terraform apply
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
