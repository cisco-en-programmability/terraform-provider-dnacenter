package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFileNamespaceFiles() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on File.

- Returns list of files under a specific namespace
`,

		ReadContext: dataSourceFileNamespaceFilesRead,
		Schema: map[string]*schema.Schema{
			"name_space": &schema.Schema{
				Description: `nameSpace path parameter. A listing of fileId's
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"attribute_info": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"download_path": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"encrypted": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"file_format": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"file_size": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"md5_checksum": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"name_space": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"sftp_server_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"sha1_checksum": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"task_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceFileNamespaceFilesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vNameSpace := d.Get("name_space")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetListOfFiles")
		vvNameSpace := vNameSpace.(string)

		response1, restyResp1, err := client.File.GetListOfFiles(vvNameSpace)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetListOfFiles", err,
				"Failure at GetListOfFiles, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenFileGetListOfFilesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetListOfFiles response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenFileGetListOfFilesItems(items *[]dnacentersdkgo.ResponseFileGetListOfFilesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["attribute_info"] = flattenFileGetListOfFilesItemsAttributeInfo(item.AttributeInfo)
		respItem["download_path"] = item.DownloadPath
		respItem["encrypted"] = boolPtrToString(item.Encrypted)
		respItem["file_format"] = item.FileFormat
		respItem["file_size"] = item.FileSize
		respItem["id"] = item.ID
		respItem["md5_checksum"] = item.Md5Checksum
		respItem["name"] = item.Name
		respItem["name_space"] = item.NameSpace
		respItem["sftp_server_list"] = flattenFileGetListOfFilesItemsSftpServerList(item.SftpServerList)
		respItem["sha1_checksum"] = item.Sha1Checksum
		respItem["task_id"] = item.TaskID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenFileGetListOfFilesItemsAttributeInfo(item *dnacentersdkgo.ResponseFileGetListOfFilesResponseAttributeInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenFileGetListOfFilesItemsSftpServerList(items *[]dnacentersdkgo.ResponseFileGetListOfFilesResponseSftpServerList) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}
