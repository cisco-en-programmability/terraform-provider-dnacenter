package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceApplicationVisibilityNetworkDevices() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Application Policy.

- This data source retrieves the list of network devices with their application visibility status. The list can be
filtered using the query parameters. Multiple filters can be applied.
`,

		ReadContext: dataSourceApplicationVisibilityNetworkDevicesRead,
		Schema: map[string]*schema.Schema{
			"app_telemetry_deployment_status": &schema.Schema{
				Description: `appTelemetryDeploymentStatus query parameter. Status of the application telemetry deployment on the network device. Available values: SCHEDULED, IN_PROGRESS, COMPLETED, FAILED, NOT_DEPLOYED.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"app_telemetry_readiness_status": &schema.Schema{
				Description: `appTelemetryReadinessStatus query parameter. Indicates whether the network device is ready for application telemetry enablement or not. Available values: ENABLED, READY, NOT_READY, NOT_SUPPORTED
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"application_registry_sync_status": &schema.Schema{
				Description: `applicationRegistrySyncStatus query parameter. Indicates whether the latest definitions from application registry have been synchronized with the network device or not. Available values: SYNCING, IN_SYNC, OUT_OF_SYNC, NOT_APPLICABLE
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"cbar_deployment_status": &schema.Schema{
				Description: `cbarDeploymentStatus query parameter. Status of the CBAR deployment on the network device. Available values: SCHEDULED, IN_PROGRESS, COMPLETED, FAILED, NOT_DEPLOYED
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"cbar_readiness_status": &schema.Schema{
				Description: `cbarReadinessStatus query parameter. Indicates whether the network device is ready for CBAR enablement or not. Available values: ENABLED, READY, NOT_READY, NOT_SUPPORTED
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"hostname": &schema.Schema{
				Description: `hostname query parameter. The host name of the network device.
Partial search is supported. For example, searching for *switch* will include *edge-switch1.domain.com*, *switch25*, etc.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"ids": &schema.Schema{
				Description: `ids query parameter. List of network devices ids. If this parameter is not provided, all network devices will be included in the response. Multiple network device IDs can be provided.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page. Minimum value is: 1, Maximum value is: 500
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"management_address": &schema.Schema{
				Description: `managementAddress query parameter. The management address for the network device. This is normally IP address of the device. But it could be hostname in some cases like Meraki devices.
Partial search is supported. For example, searching for *25.* would include *10.25.1.1*, *25.5.10.1*, *225.225.1.0*, *10.10.10.125*, etc.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page; the first record is numbered 1. Default value is: 1.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"order": &schema.Schema{
				Description: `order query parameter. Whether ascending or descending order should be used to sort the response. Available values are: asc, desc. Default value is: asc
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol_pack_status": &schema.Schema{
				Description: `protocolPackStatus query parameter. Indicates whether the NBAR protocol pack is up-to-date or not on the network device. Available values: LATEST, OUTDATED, UNKNOWN
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol_pack_update_status": &schema.Schema{
				Description: `protocolPackUpdateStatus query parameter. Status of the NBAR protocol pack update on the network device. Available values: SCHEDULED, IN_PROGRESS, SUCCESS, FAILED, NONE
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId query parameter. The site ID where the network device is assigned.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. A property within the response to sort by.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"app_telemetry_deployment_status": &schema.Schema{
							Description: `Status of the application telemetry deployment on the network device.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"app_telemetry_readiness_status": &schema.Schema{
							Description: `Indicates whether the network device is ready for application telemetry enablement or not.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"application_registry_sync_status": &schema.Schema{
							Description: `Indicates whether the latest definitions from application registry have been synchronized with the network device or not.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"cbar_deployment_status": &schema.Schema{
							Description: `Status of the CBAR deployment on the network device.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"cbar_readiness_status": &schema.Schema{
							Description: `Indicates whether the network device is ready for CBAR enablement or not.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"hostname": &schema.Schema{
							Description: `The host name of the network device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `The network device id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"management_address": &schema.Schema{
							Description: `The management address for the network device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"protocol_pack_status": &schema.Schema{
							Description: `Indicates whether the NBAR protocol pack is up-to-date or not on the network device.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"protocol_pack_update_status": &schema.Schema{
							Description: `Status of the NBAR protocol pack update on the network device.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"site_id": &schema.Schema{
							Description: `The site ID where the network device is assigned.
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

func dataSourceApplicationVisibilityNetworkDevicesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vIDs, okIDs := d.GetOk("ids")
	vManagementAddress, okManagementAddress := d.GetOk("management_address")
	vHostname, okHostname := d.GetOk("hostname")
	vSiteID, okSiteID := d.GetOk("site_id")
	vAppTelemetryDeploymentStatus, okAppTelemetryDeploymentStatus := d.GetOk("app_telemetry_deployment_status")
	vAppTelemetryReadinessStatus, okAppTelemetryReadinessStatus := d.GetOk("app_telemetry_readiness_status")
	vCbarDeploymentStatus, okCbarDeploymentStatus := d.GetOk("cbar_deployment_status")
	vCbarReadinessStatus, okCbarReadinessStatus := d.GetOk("cbar_readiness_status")
	vProtocolPackStatus, okProtocolPackStatus := d.GetOk("protocol_pack_status")
	vProtocolPackUpdateStatus, okProtocolPackUpdateStatus := d.GetOk("protocol_pack_update_status")
	vApplicationRegistrySyncStatus, okApplicationRegistrySyncStatus := d.GetOk("application_registry_sync_status")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveTheListOfNetworkDevicesWithTheirApplicationVisibilityStatus")
		queryParams1 := dnacentersdkgo.RetrieveTheListOfNetworkDevicesWithTheirApplicationVisibilityStatusQueryParams{}

		if okIDs {
			queryParams1.IDs = vIDs.(string)
		}
		if okManagementAddress {
			queryParams1.ManagementAddress = vManagementAddress.(string)
		}
		if okHostname {
			queryParams1.Hostname = vHostname.(string)
		}
		if okSiteID {
			queryParams1.SiteID = vSiteID.(string)
		}
		if okAppTelemetryDeploymentStatus {
			queryParams1.AppTelemetryDeploymentStatus = vAppTelemetryDeploymentStatus.(string)
		}
		if okAppTelemetryReadinessStatus {
			queryParams1.AppTelemetryReadinessStatus = vAppTelemetryReadinessStatus.(string)
		}
		if okCbarDeploymentStatus {
			queryParams1.CbarDeploymentStatus = vCbarDeploymentStatus.(string)
		}
		if okCbarReadinessStatus {
			queryParams1.CbarReadinessStatus = vCbarReadinessStatus.(string)
		}
		if okProtocolPackStatus {
			queryParams1.ProtocolPackStatus = vProtocolPackStatus.(string)
		}
		if okProtocolPackUpdateStatus {
			queryParams1.ProtocolPackUpdateStatus = vProtocolPackUpdateStatus.(string)
		}
		if okApplicationRegistrySyncStatus {
			queryParams1.ApplicationRegistrySyncStatus = vApplicationRegistrySyncStatus.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(string)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(string)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}

		response1, restyResp1, err := client.ApplicationPolicy.RetrieveTheListOfNetworkDevicesWithTheirApplicationVisibilityStatus(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveTheListOfNetworkDevicesWithTheirApplicationVisibilityStatus", err,
				"Failure at RetrieveTheListOfNetworkDevicesWithTheirApplicationVisibilityStatus, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenApplicationPolicyRetrieveTheListOfNetworkDevicesWithTheirApplicationVisibilityStatusItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveTheListOfNetworkDevicesWithTheirApplicationVisibilityStatus response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenApplicationPolicyRetrieveTheListOfNetworkDevicesWithTheirApplicationVisibilityStatusItems(items *[]dnacentersdkgo.ResponseApplicationPolicyRetrieveTheListOfNetworkDevicesWithTheirApplicationVisibilityStatusResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["management_address"] = item.ManagementAddress
		respItem["hostname"] = item.Hostname
		respItem["site_id"] = item.SiteID
		respItem["app_telemetry_deployment_status"] = item.AppTelemetryDeploymentStatus
		respItem["app_telemetry_readiness_status"] = item.AppTelemetryReadinessStatus
		respItem["cbar_deployment_status"] = item.CbarDeploymentStatus
		respItem["cbar_readiness_status"] = item.CbarReadinessStatus
		respItem["protocol_pack_status"] = item.ProtocolPackStatus
		respItem["protocol_pack_update_status"] = item.ProtocolPackUpdateStatus
		respItem["application_registry_sync_status"] = item.ApplicationRegistrySyncStatus
		respItems = append(respItems, respItem)
	}
	return respItems
}
