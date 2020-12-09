package dnacenter

import (
	"context"
	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkCredentialSiteAssignment() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceNetworkCredentialSiteAssignmentCreate,
		ReadContext:   resourceNetworkCredentialSiteAssignmentRead,
		UpdateContext: resourceNetworkCredentialSiteAssignmentUpdate,
		DeleteContext: resourceNetworkCredentialSiteAssignmentDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"site_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cli": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comments": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"credential_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable_password": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"password": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"username": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"http_read": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"comments": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"credential_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"password": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"secure": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"username": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"http_write": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"comments": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"credential_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"password": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"secure": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"username": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"snmp_v2_read": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"comments": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"credential_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"read_community": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"snmp_v2_write": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comments": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"credential_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"write_community": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"snmp_v3": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"auth_password": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"auth_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"comments": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"credential_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"privacy_password": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"privacy_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_mode": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"username": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

//NetworkCredentialSiteAssignmentParams defines the properties of the request and Terraform
type NetworkCredentialSiteAssignmentParams struct {
	Cli         []interface{}
	HTTPRead    []interface{}
	HTTPWrite   []interface{}
	SNMPV2Read  []interface{}
	SNMPV2Write []interface{}
	SNMPV3      []interface{}
}

func constructNetworkCredentialSiteAssignment(params *NetworkCredentialSiteAssignmentParams) *dnac.AssignCredentialToSiteRequest {
	if params != nil {
		result := dnac.AssignCredentialToSiteRequest{}
		if params.Cli != nil && len(params.Cli) > 0 {
			itemCli := params.Cli[0].(map[string]interface{})
			if v, ok := itemCli["id"]; ok && v != nil {
				result.CliID = v.(string)
			}
		}
		if params.HTTPRead != nil && len(params.HTTPRead) > 0 {
			itemHTTPRead := params.HTTPRead[0].(map[string]interface{})
			if v, ok := itemHTTPRead["id"]; ok && v != nil {
				result.HTTPRead = v.(string)
			}
		}
		if params.HTTPWrite != nil && len(params.HTTPWrite) > 0 {
			itemHTTPWrite := params.HTTPWrite[0].(map[string]interface{})
			if v, ok := itemHTTPWrite["id"]; ok && v != nil {
				result.HTTPWrite = v.(string)
			}
		}
		if params.SNMPV2Read != nil && len(params.SNMPV2Read) > 0 {
			itemSNMPV2Read := params.SNMPV2Read[0].(map[string]interface{})
			if v, ok := itemSNMPV2Read["id"]; ok && v != nil {
				result.SNMPV2ReadID = v.(string)
			}
		}
		if params.SNMPV2Write != nil && len(params.SNMPV2Write) > 0 {
			itemSNMPV2Write := params.SNMPV2Write[0].(map[string]interface{})
			if v, ok := itemSNMPV2Write["id"]; ok && v != nil {
				result.SNMPV2WriteID = v.(string)
			}
		}
		if params.SNMPV3 != nil && len(params.SNMPV3) > 0 {
			itemSNMPV3 := params.SNMPV3[0].(map[string]interface{})
			if v, ok := itemSNMPV3["id"]; ok && v != nil {
				result.SNMPV3ID = v.(string)
			}
		}
		return &result
	}
	return nil
}

func resourceNetworkCredentialSiteAssignmentCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)
	var diags diag.Diagnostics

	siteID := d.Get("site_id").(string)

	requestData := NetworkCredentialSiteAssignmentParams{}
	if v, ok := d.GetOk("cli"); ok && v != nil {
		requestData.Cli = v.([]interface{})
	}
	if v, ok := d.GetOk("http_read"); ok && v != nil {
		requestData.HTTPRead = v.([]interface{})
	}
	if v, ok := d.GetOk("http_write"); ok && v != nil {
		requestData.HTTPWrite = v.([]interface{})
	}
	if v, ok := d.GetOk("snmp_v2_read"); ok && v != nil {
		requestData.SNMPV2Read = v.([]interface{})
	}
	if v, ok := d.GetOk("snmp_v2_write"); ok && v != nil {
		requestData.SNMPV2Write = v.([]interface{})
	}
	if v, ok := d.GetOk("snmp_v3"); ok && v != nil {
		requestData.SNMPV3 = v.([]interface{})
	}

	request := constructNetworkCredentialSiteAssignment(&requestData)
	createResponse, _, err := client.NetworkSettings.AssignCredentialToSite(siteID, request)
	if err != nil {
		return diag.FromErr(err)
	}
	log.Printf("createResponse %+v", createResponse)

	d.SetId(siteID)

	resourceNetworkCredentialSiteAssignmentRead(ctx, d, m)
	return diags
}

func networkCredentialSiteAssignmentSimplify(response *dnac.GetDeviceCredentialDetailsResponse) *NetworkCredentialSiteAssignmentParams {
	result := NetworkCredentialSiteAssignmentParams{}
	if response != nil {
		result.Cli = flattenNetworkDeviceCredentialReadItemsCli(&response.Cli)
		result.HTTPRead = flattenNetworkDeviceCredentialReadItemsHTTPRead(&response.HTTPRead)
		result.HTTPWrite = flattenNetworkDeviceCredentialReadItemsHTTPWrite(&response.HTTPWrite)
		result.SNMPV2Read = flattenNetworkDeviceCredentialReadItemsSNMPV2Read(&response.SNMPV2Read)
		result.SNMPV2Write = flattenNetworkDeviceCredentialReadItemsSNMPV2Write(&response.SNMPV2Write)
		result.SNMPV3 = flattenNetworkDeviceCredentialReadItemsSNMPV3(&response.SNMPV3)
	}
	return &result
}

func resourceNetworkCredentialSiteAssignmentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	siteID := d.Id()
	queryParams := dnac.GetDeviceCredentialDetailsQueryParams{SiteID: siteID}
	searchResponse, _, err := client.NetworkSettings.GetDeviceCredentialDetails(&queryParams)
	if err != nil {
		return diag.FromErr(err)
	}

	networkCredentialSiteAssignmentSimplified := networkCredentialSiteAssignmentSimplify(searchResponse)

	if err := d.Set("cli", networkCredentialSiteAssignmentSimplified.Cli); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("http_read", networkCredentialSiteAssignmentSimplified.HTTPRead); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("http_write", networkCredentialSiteAssignmentSimplified.HTTPWrite); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("snmp_v2_read", networkCredentialSiteAssignmentSimplified.SNMPV2Read); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("snmp_v2_write", networkCredentialSiteAssignmentSimplified.SNMPV2Write); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("snmp_v3", networkCredentialSiteAssignmentSimplified.SNMPV3); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceNetworkCredentialSiteAssignmentUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	siteID := d.Id()
	requestData := NetworkCredentialSiteAssignmentParams{}
	if d.HasChanges("cli", "http_read", "http_write", "snmp_v2_read", "snmp_v2_write", "snmp_v3") {
		if v, ok := d.GetOk("cli"); ok && v != nil {
			requestData.Cli = v.([]interface{})
		}
		if v, ok := d.GetOk("http_read"); ok && v != nil {
			requestData.HTTPRead = v.([]interface{})
		}
		if v, ok := d.GetOk("http_write"); ok && v != nil {
			requestData.HTTPWrite = v.([]interface{})
		}
		if v, ok := d.GetOk("snmp_v2_read"); ok && v != nil {
			requestData.SNMPV2Read = v.([]interface{})
		}
		if v, ok := d.GetOk("snmp_v2_write"); ok && v != nil {
			requestData.SNMPV2Write = v.([]interface{})
		}
		if v, ok := d.GetOk("snmp_v3"); ok && v != nil {
			requestData.SNMPV3 = v.([]interface{})
		}
	}

	request := constructNetworkCredentialSiteAssignment(&requestData)
	updateResponse, _, err := client.NetworkSettings.AssignCredentialToSite(siteID, request)
	if err != nil {
		return diag.FromErr(err)
	}
	log.Printf("updateResponse %+v", updateResponse)

	// Wait for execution status to complete
	time.Sleep(5 * time.Second)

	// Update resource last_updated
	d.Set("last_updated", time.Now().Format(time.RFC850))

	return resourceNetworkCredentialSiteAssignmentRead(ctx, d, m)
}

func resourceNetworkCredentialSiteAssignmentDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}
