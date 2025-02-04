package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceConfigFiles() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Configuration Archive.

- Retrieves the list of network device configuration file details, sorted by createdTime in descending order. Use
/intent/api/v1/networkDeviceConfigFiles/{id}/downloadMasked to download masked configurations, or
/intent/api/v1/networkDeviceConfigFiles/{id}/downloadUnmasked for unmasked configurations.
`,

		ReadContext: dataSourceNetworkDeviceConfigFilesRead,
		Schema: map[string]*schema.Schema{
			"file_type": &schema.Schema{
				Description: `fileType query parameter. Type of device configuration file.Available values : 'RUNNINGCONFIG', 'STARTUPCONFIG', 'VLAN'
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id query parameter. Unique identifier (UUID) of the configuration file.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to be retrieved defaults to 500 if not specified, with a maximum allowed limit of 500.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId query parameter. Unique identifier (UUID) of the network devices. The number of networkDeviceId(s) must not exceed 5.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page; the first record is numbered 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"created_by": &schema.Schema{
							Description: `The entity responsible for creating the configuration changes.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"created_time": &schema.Schema{
							Description: `The UNIX epoch timestamp in milliseconds marking when the resource was created.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"file_type": &schema.Schema{
							Description: `Type of configuration file. Config File Type can be 'RUNNINGCONFIG' or 'STARTUPCONFIG' or 'VLAN'.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `Unique identifier (UUID) of the configuration file.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"network_device_id": &schema.Schema{
							Description: `Unique identifier (UUID) of the network devices.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"version_id": &schema.Schema{
							Description: `The version unique identifier triggered after any config change.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkDeviceConfigFilesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID, okID := d.GetOk("id")
	vNetworkDeviceID, okNetworkDeviceID := d.GetOk("network_device_id")
	vFileType, okFileType := d.GetOk("file_type")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkDeviceConfigurationFileDetails")
		queryParams1 := dnacentersdkgo.GetNetworkDeviceConfigurationFileDetailsQueryParams{}

		if okID {
			queryParams1.ID = vID.(string)
		}
		if okNetworkDeviceID {
			queryParams1.NetworkDeviceID = vNetworkDeviceID.(string)
		}
		if okFileType {
			queryParams1.FileType = vFileType.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.ConfigurationArchive.GetNetworkDeviceConfigurationFileDetails(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetNetworkDeviceConfigurationFileDetails", err,
				"Failure at GetNetworkDeviceConfigurationFileDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenConfigurationArchiveGetNetworkDeviceConfigurationFileDetailsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkDeviceConfigurationFileDetails response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenConfigurationArchiveGetNetworkDeviceConfigurationFileDetailsItems(items *[]dnacentersdkgo.ResponseConfigurationArchiveGetNetworkDeviceConfigurationFileDetailsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["network_device_id"] = item.NetworkDeviceID
		respItem["version_id"] = item.VersionID
		respItem["file_type"] = item.FileType
		respItem["created_by"] = item.CreatedBy
		respItem["created_time"] = item.CreatedTime
		respItems = append(respItems, respItem)
	}
	return respItems
}
