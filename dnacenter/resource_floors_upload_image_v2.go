package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceFloorsUploadImageV2() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Site Design.

- Uploads floor image.
`,

		CreateContext: resourceFloorsUploadImageV2Create,
		ReadContext:   resourceFloorsUploadImageV2Read,
		DeleteContext: resourceFloorsUploadImageV2Delete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Description: `id path parameter. Floor Id
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
		},
	}
}

func resourceFloorsUploadImageV2Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vID := resourceItem["id"]

	vvID := vID.(string)

	// has_unknown_response: True

	response1, err := client.SiteDesign.UploadsFloorImageV2(vvID)

	if err != nil || response1 == nil {
		d.SetId("")
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %s", response1.String())

	//Analizar verificacion.

	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting UploadsFloorImageV2 response",
			err))
		return diags
	}
	d.SetId(getUnixTimeString())
	return diags

}
func resourceFloorsUploadImageV2Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceFloorsUploadImageV2Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}
