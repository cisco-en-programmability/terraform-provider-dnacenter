package dnacenter

import (
	"context"
	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSDAFabricPortAssignmentForUserDevice() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceSDAFabricPortAssignmentForUserDeviceCreate,
		ReadContext:   resourceSDAFabricPortAssignmentForUserDeviceRead,
		UpdateContext: resourceSDAFabricPortAssignmentForUserDeviceUpdate,
		DeleteContext: resourceSDAFabricPortAssignmentForUserDeviceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{

			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"site_name_hierarchy": &schema.Schema{
				Type:     schema.TypeString,
				Required: true, //REVIEW: It may be only Optional & Computed
			},
			"device_management_ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"interface_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"data_ip_address_pool_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true, //REVIEW: It may be only Optional & Computed
			},
			"voice_ip_address_pool_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true, //REVIEW: It may be only Optional & Computed
			},
			"authenticate_template_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true, //REVIEW: It may be only Optional & Computed
			},
			"scalable_group_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceSDAFabricPortAssignmentForUserDeviceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)
	var diags diag.Diagnostics

	deviceManagementIPAddress := d.Get("device_management_ip_address").(string)
	interfaceName := d.Get("interface_name").(string)

	searchResponse, _, err := client.SDA.GetPortAssignmentForUserDeviceInSDAFabric(&dnac.GetPortAssignmentForUserDeviceInSDAFabricQueryParams{
		Devicp:        deviceManagementIPAddress,
		InterfaceName: interfaceName,
	})
	if err == nil && searchResponse != nil {
		if deviceManagementIPAddress == searchResponse.DeviceManagementIPAddress {
			// Update resource id
			d.SetId(strings.Join([]string{interfaceName, deviceManagementIPAddress}, "_/_"))
			resourceSDAFabricPortAssignmentForUserDeviceRead(ctx, d, m)
			return diags
		}
	}

	var requests []dnac.AddPortAssignmentForUserDeviceInSDAFabricRequest
	var request dnac.AddPortAssignmentForUserDeviceInSDAFabricRequest
	request.DeviceManagementIPAddress = deviceManagementIPAddress
	request.InterfaceName = interfaceName
	if v, ok := d.GetOk("authenticate_template_name"); ok {
		request.AuthenticateTemplateName = v.(string)
	}
	if v, ok := d.GetOk("data_ip_address_pool_name"); ok {
		request.DataIPAddressPoolName = v.(string)
	}
	if v, ok := d.GetOk("site_name_hierarchy"); ok {
		request.SiteNameHierarchy = v.(string)
	}
	if v, ok := d.GetOk("voice_ip_address_pool_name"); ok {
		request.VoiceIPAddressPoolName = v.(string)
	}
	requests = append(requests, request)
	_, _, err = client.SDA.AddPortAssignmentForUserDeviceInSDAFabric(&requests)
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for execution status to complete
	time.Sleep(5 * time.Second)

	// Update resource id
	d.SetId(strings.Join([]string{interfaceName, deviceManagementIPAddress}, "_/_"))
	resourceSDAFabricPortAssignmentForUserDeviceRead(ctx, d, m)
	return diags
}

func resourceSDAFabricPortAssignmentForUserDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	resourceIDs := strings.Split(d.Id(), "_/_")
	interfaceName, deviceManagementIPAddress := resourceIDs[0], resourceIDs[1]

	searchResponse, _, err := client.SDA.GetPortAssignmentForUserDeviceInSDAFabric(&dnac.GetPortAssignmentForUserDeviceInSDAFabricQueryParams{
		Devicp:        deviceManagementIPAddress,
		InterfaceName: interfaceName,
	})
	if err != nil || searchResponse == nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "May have been unable to get SDA fabric port assignment for access point",
		})
		// REVIEW:.
		return diags
	}
	if deviceManagementIPAddress == searchResponse.DeviceManagementIPAddress {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "May have been unable to get SDA fabric port assignment for access point",
		})
		// REVIEW:.
		return diags
	}

	if err := d.Set("authenticate_template_name", searchResponse.AuthenticateTemplateName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("data_ip_address_pool_name", searchResponse.DataIPAddressPoolName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("device_management_ip_address", searchResponse.DeviceManagementIPAddress); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("interface_name", searchResponse.InterfaceName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("scalable_group_name", searchResponse.ScalableGroupName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("site_name_hierarchy", searchResponse.SiteNameHierarchy); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("voice_ip_address_pool_name", searchResponse.VoiceIPAddressPoolName); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceSDAFabricPortAssignmentForUserDeviceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceSDAFabricPortAssignmentForUserDeviceRead(ctx, d, m)
}

func resourceSDAFabricPortAssignmentForUserDeviceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	resourceIDs := strings.Split(d.Id(), "_/_")
	interfaceName, deviceManagementIPAddress := resourceIDs[0], resourceIDs[1]

	searchResponse, _, err := client.SDA.GetPortAssignmentForUserDeviceInSDAFabric(&dnac.GetPortAssignmentForUserDeviceInSDAFabricQueryParams{
		Devicp:        deviceManagementIPAddress,
		InterfaceName: interfaceName,
	})
	if err != nil || searchResponse == nil {
		return diags
	}
	if deviceManagementIPAddress == searchResponse.DeviceManagementIPAddress {
		return diags
	}

	// Call function to delete resource
	deleteRequest := []dnac.DeletePortAssignmentForUserDeviceInSDAFabricRequest{}
	_, _, err = client.SDA.DeletePortAssignmentForUserDeviceInSDAFabric(&dnac.DeletePortAssignmentForUserDeviceInSDAFabricQueryParams{
		Devicp:        deviceManagementIPAddress,
		InterfaceName: interfaceName,
	}, &deleteRequest)
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for execution status to complete
	time.Sleep(5 * time.Second)

	searchResponse, _, err = client.SDA.GetPortAssignmentForUserDeviceInSDAFabric(&dnac.GetPortAssignmentForUserDeviceInSDAFabricQueryParams{
		Devicp:        deviceManagementIPAddress,
		InterfaceName: interfaceName,
	})
	if err == nil && searchResponse != nil {
		// Check if element already exists
		if deviceManagementIPAddress == searchResponse.DeviceManagementIPAddress {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to delete SDA fabric port assignment for access point",
			})
		}
		return diags
	}

	return diags
}
