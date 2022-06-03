package dnacenter

import (
	"context"
	"io"
	"os"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileImport() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on File.
- Uploads a new file within a specific nameSpace
`,

		CreateContext: resourceFileImportCreate,
		ReadContext:   resourceFileImportRead,
		DeleteContext: resourceFileImportDelete,
		Schema: map[string]*schema.Schema{
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
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
					},
				},
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceFileImportRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceFileImportCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	vFileName := resourceItem["file_name"]
	vFilePath := resourceItem["file_path"]
	vNameSpace := resourceItem["name_space"]
	vvNameSpace := vNameSpace.(string)

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: UploadFile")

		isDir, err := IsDirectory(vFilePath.(string))
		if err != nil || isDir {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing File", err,
				"Failure at File, Path is a directory", ""))
			return diags
		}
		f, err := os.Open(vFilePath.(string))
		if err != nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UploadFile", err,
				"Failure at UploadFile, unexpected response", ""))
			return diags
		}
		defer func() {
			if err = f.Close(); err != nil {
				log.Printf("File close error %s", err.Error())
			}
		}()

		var r io.Reader
		r = f

		response1, restyResp1, err := client.File.UploadFile(
			vvNameSpace,
			&dnacentersdkgo.UploadFileMultipartFields{
				File:     r,
				FileName: vFileName.(string),
			},
		)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UploadFile", err,
				"Failure at UploadFile, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %s", restyResp1.String())

		if err := d.Set("item", restyResp1.String()); err != nil {
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

func resourceFileImportDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}
