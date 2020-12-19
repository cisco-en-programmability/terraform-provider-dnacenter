package dnacenter

import (
	"context"
	"fmt"
	"strconv"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTagMemberQuery() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTagMemberQueryRead,
		Schema: map[string]*schema.Schema{
			"tag_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"member_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"member_association_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"limit": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceTagMemberQueryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	tagID := d.Get("tag_id").(string)

	queryParams := dnac.GetTagMembersByIDQueryParams{}
	if v, ok := d.GetOk("member_type"); ok {
		queryParams.MemberType = v.(string)
	}
	if v, ok := d.GetOk("member_association_type"); ok {
		queryParams.MemberAssociationType = v.(string)
	}
	if v, ok := d.GetOk("level"); ok {
		queryParams.Level = v.(string)
	}
	if v, ok := d.GetOk("offset"); ok {
		queryParams.Offset = fmt.Sprintf("%.0f", v.(float64))
	}
	if v, ok := d.GetOk("limit"); ok {
		queryParams.Limit = fmt.Sprintf("%.0f", v.(float64))
	}

	_, restyResponse, err := client.Tag.GetTagMembersByID(tagID, &queryParams)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("items", restyResponse.String()); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
