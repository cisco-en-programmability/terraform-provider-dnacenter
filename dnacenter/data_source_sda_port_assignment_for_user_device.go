package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaPortAssignmentForUserDevice() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Get Port assignment for user device in SDA Fabric.
`,

		ReadContext: dataSourceSdaPortAssignmentForUserDeviceRead,
		Schema: map[string]*schema.Schema{
			"device_management_ip_address": &schema.Schema{
				Description: `deviceManagementIpAddress query parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"interface_name": &schema.Schema{
				Description: `interfaceName query parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"authenticate_template_name": &schema.Schema{
							Description: `Authenticate Template Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"data_ip_address_pool_name": &schema.Schema{
							Description: `Data Ip Address Pool Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_management_ip_address": &schema.Schema{
							Description: `Device Management Ip Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"interface_name": &schema.Schema{
							Description: `Interface Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"scalable_group_name": &schema.Schema{
							Description: `Scalable Group Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"site_name_hierarchy": &schema.Schema{
							Description: `Site Name Hierarchy`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"status": &schema.Schema{
							Description: `Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"voice_ip_address_pool_name": &schema.Schema{
							Description: `Voice Ip Address Pool Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSdaPortAssignmentForUserDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceManagementIPAddress := d.Get("device_management_ip_address")
	vInterfaceName := d.Get("interface_name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetPortAssignmentForUserDeviceInSdaFabric")
		queryParams1 := dnacentersdkgo.GetPortAssignmentForUserDeviceInSdaFabricQueryParams{}

		queryParams1.DeviceManagementIPAddress = vDeviceManagementIPAddress.(string)

		queryParams1.InterfaceName = vInterfaceName.(string)

		response1, restyResp1, err := client.Sda.GetPortAssignmentForUserDeviceInSdaFabric(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetPortAssignmentForUserDeviceInSdaFabric", err,
				"Failure at GetPortAssignmentForUserDeviceInSdaFabric, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetPortAssignmentForUserDeviceInSdaFabricItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPortAssignmentForUserDeviceInSdaFabric response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetPortAssignmentForUserDeviceInSdaFabricItem(item *dnacentersdkgo.ResponseSdaGetPortAssignmentForUserDeviceInSdaFabric) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["status"] = item.Status
	respItem["description"] = item.Description
	respItem["site_name_hierarchy"] = item.SiteNameHierarchy
	respItem["device_management_ip_address"] = item.DeviceManagementIPAddress
	respItem["interface_name"] = item.InterfaceName
	respItem["data_ip_address_pool_name"] = item.DataIPAddressPoolName
	respItem["voice_ip_address_pool_name"] = item.VoiceIPAddressPoolName
	respItem["scalable_group_name"] = item.ScalableGroupName
	respItem["authenticate_template_name"] = item.AuthenticateTemplateName
	return []map[string]interface{}{
		respItem,
	}
}
