package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTagMemberCount() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTagMemberCountRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"member_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"member_association_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"level": &schema.Schema{
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

func dataSourceTagMemberCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	tagID := d.Get("id").(string)
	tagMemberCountQueryParams := dnac.GetTagMemberCountQueryParams{}

	if v, ok := d.GetOk("member_type"); ok {
		tagMemberCountQueryParams.MemberType = v.(string)
	}
	if v, ok := d.GetOk("member_association_type"); ok {
		tagMemberCountQueryParams.MemberAssociationType = v.(string)
	}
	if v, ok := d.GetOk("level"); ok {
		tagMemberCountQueryParams.Level = v.(string)
	}

	// Prepare Request
	response, _, err := client.Tag.GetTagMemberCount(tagID, &tagMemberCountQueryParams)
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
