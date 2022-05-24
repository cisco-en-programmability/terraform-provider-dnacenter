package dnacenter

import (
	"context"


	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceFileImport() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on File.

- Uploads a new file within a specific nameSpace
`,

		ReadContext:   dataSourceFileImportRead,
		Schema: map[string]*schema.Schema{
      "file_name": &schema.Schema{
        Description: `File name.`,
        Type:        schema.TypeString,
        Required:    true,
      },
      "file_path": &schema.Schema{
        Description: `File absolute path.`,
        Type:        schema.TypeString,
        Required:    true,
      },
      "name_space": &schema.Schema{
        Description: `nameSpace path parameter.`,
        Type:        schema.TypeString,
        Required:    true,
      },
      "item": &schema.Schema{
        Type:     schema.TypeList,
        Computed: true,
      },
		},
	}
}

func dataSourceFileImportRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vNameSpace := d.Get("name_space")


	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: UploadFile")
		vvNameSpace := vNameSpace.(string)

		response1, restyResp1, err := client.File.UploadFile(vvNameSpace)

	
	
		if err != nil || response1 == nil {
		  if restyResp1 != nil {
		    log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		  }
		  diags = append(diags, diagErrorWithAlt(
		    "Failure when executing UploadFile", err,
		    "Failure at UploadFile, unexpected response", ""))
		  return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenFileUploadFileItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting UploadFile response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
  return diags
}



func flattenFileUploadFileItem(item *dnacentersdkgo.) []interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	return []map[string]interface{}{
		respItem,
	}
}
