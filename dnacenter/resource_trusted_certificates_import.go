package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceTrustedCertificatesImport() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Cisco Trusted Certificates.

- Imports trusted certificate into a truststore. Accepts .pem or .der file as input.
`,

		CreateContext: resourceTrustedCertificatesImportCreate,
		ReadContext:   resourceTrustedCertificatesImportRead,
		DeleteContext: resourceTrustedCertificatesImportDelete,
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
					Schema: map[string]*schema.Schema{},
				},
			},
		},
	}
}

func resourceTrustedCertificatesImportCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	// has_unknown_response: True

	response1, err := client.CiscoTrustedCertificates.ImportTrustedCertificate()

	if err != nil || response1 == nil {
		d.SetId("")
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %s", response1.String())

	//Analizar verificacion.

	if err := d.Set("item", response1.String()); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ImportTrustedCertificate response",
			err))
		return diags
	}
	d.SetId(getUnixTimeString())
	return diags

}
func resourceTrustedCertificatesImportRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceTrustedCertificatesImportDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}
