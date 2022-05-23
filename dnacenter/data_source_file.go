package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFile() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on File.

- Downloads a file specified by fileId
`,

		ReadContext: dataSourceFileRead,
		Schema: map[string]*schema.Schema{
			"dirpath": &schema.Schema{
				Description: `Directory absolute path in which to save the file.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"file_id": &schema.Schema{
				Description: `fileId path parameter. File Identification number
`,
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceFileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vFileID := d.Get("file_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: DownloadAFileByFileID")
		vvFileID := vFileID.(string)

		response1, _, err := client.File.DownloadAFileByFileID(vvFileID)

		if err != nil {
			diags = append(diags, diagError(
				"Failure when executing DownloadAFileByFileID", err))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response")

		vvDirpath := d.Get("dirpath").(string)
		err = response1.SaveDownload(vvDirpath)
		if err != nil {
			diags = append(diags, diagError(
				"Failure when downloading file", err))
			return diags
		}
		log.Printf("[DEBUG] Downloaded file %s", vvDirpath)

	}
	return diags
}
