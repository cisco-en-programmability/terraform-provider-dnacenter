package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceMapsImport() *schema.Resource {
	return &schema.Resource{
		Description: `It performs delete operation on Sites.

- Cancels a previously initatied import, allowing the system to cleanup cached resources about that import data, and
ensures the import cannot accidentally be performed / approved at a later time.
`,

		CreateContext: resourceMapsImportCreate,
		ReadContext:   resourceMapsImportRead,
		DeleteContext: resourceMapsImportDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"import_context_uuid": &schema.Schema{
							Description: `importContextUuid path parameter. The unique import context UUID given by a previous call to Start Import API
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

func resourceMapsImportCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vImportContextUUID := resourceItem["import_context_uuid"]

	vvImportContextUUID := vImportContextUUID.(string)

	// has_unknown_response: None

	response1, restyResp1, err := client.Sites.ImportMapArchiveCancelAnImport(vvImportContextUUID)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing ImportMapArchiveCancelAnImport", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if err := d.Set("item", response1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ImportMapArchiveCancelAnImport response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

	//Analizar verificacion.

}
func resourceMapsImportRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceMapsImportDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}
