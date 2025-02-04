package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceApplicationVisibilityNetworkDevicesCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Application Policy.

- This data source retrieves the count of network devices for the given application visibility status filters.
`,

		ReadContext: dataSourceApplicationVisibilityNetworkDevicesCountRead,
		Schema: map[string]*schema.Schema{
			"app_telemetry_deployment_status": &schema.Schema{
				Description: `appTelemetryDeploymentStatus query parameter. Status of the application telemetry deployment on the network device. Available values: SCHEDULED, IN_PROGRESS, COMPLETED, FAILED, NOT_DEPLOYED
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
			"management_address": &schema.Schema{
				Description: `managementAddress query parameter. The management address for the network device. This is normally IP address of the device. But it could be hostname in some cases like Meraki devices.
Partial search is supported. For example, searching for *25.* would include *10.25.1.1*, *25.5.10.1*, *225.225.1.0*, *10.10.10.125*, etc.
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

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Description: `Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceApplicationVisibilityNetworkDevicesCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveTheCountOfNetworkDevicesForTheGivenApplicationVisibilityStatusFilters")
		queryParams1 := dnacentersdkgo.RetrieveTheCountOfNetworkDevicesForTheGivenApplicationVisibilityStatusFiltersQueryParams{}

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

		response1, restyResp1, err := client.ApplicationPolicy.RetrieveTheCountOfNetworkDevicesForTheGivenApplicationVisibilityStatusFilters(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveTheCountOfNetworkDevicesForTheGivenApplicationVisibilityStatusFilters", err,
				"Failure at RetrieveTheCountOfNetworkDevicesForTheGivenApplicationVisibilityStatusFilters, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenApplicationPolicyRetrieveTheCountOfNetworkDevicesForTheGivenApplicationVisibilityStatusFiltersItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveTheCountOfNetworkDevicesForTheGivenApplicationVisibilityStatusFilters response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenApplicationPolicyRetrieveTheCountOfNetworkDevicesForTheGivenApplicationVisibilityStatusFiltersItem(item *dnacentersdkgo.ResponseApplicationPolicyRetrieveTheCountOfNetworkDevicesForTheGivenApplicationVisibilityStatusFiltersResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
