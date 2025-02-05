package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceMapsImportPerform() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Sites.

- For a previously initatied import, approves the import to be performed, accepting that data loss may occur.  A Map
import will fully replace existing Maps data for the site(s) defined in the archive. The Map Archive Import Status API
/maps/import/${contextUuid}/status should always be checked to validate the pre-import validation output prior to
performing the import.
`,

		CreateContext: resourceMapsImportPerformCreate,
		ReadContext:   resourceMapsImportPerformRead,
		DeleteContext: resourceMapsImportPerformDelete,
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
							Description: `importContextUuid path parameter. The unique import context UUID given by a previous call of Start Import API
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

func resourceMapsImportPerformCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vImportContextUUID := resourceItem["import_context_uuid"]

	vvImportContextUUID := vImportContextUUID.(string)

	// has_unknown_response: None

	response1, restyResp1, err := client.Sites.ImportMapArchivePerformImport(vvImportContextUUID)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing ImportMapArchivePerformImport", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if err := d.Set("item", response1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ImportMapArchivePerformImport response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

	//Analizar verificacion.

}
func resourceMapsImportPerformRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceMapsImportPerformDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}
