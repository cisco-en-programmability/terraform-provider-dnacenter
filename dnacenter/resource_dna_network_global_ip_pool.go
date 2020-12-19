package dnacenter

import (
	"context"
	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkGlobalIPPool() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceNetworkGlobalIPPoolCreate,
		ReadContext:   resourceNetworkGlobalIPPoolRead,
		UpdateContext: resourceNetworkGlobalIPPoolUpdate,
		DeleteContext: resourceNetworkGlobalIPPoolDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{

			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateStringHasValueFunc([]string{"Generic"}), //REVIEW: .
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ip_address_space": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateStringHasValueFunc([]string{"IPv4", "IPv6"}),
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_options": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"configure_external_dhcp": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"context": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"context_key": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"context_value": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"owner": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"create_time": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"dhcp_server_ips": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"dns_server_ips": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"gateways": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"ip_pool_cidr": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"ip_pool_name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"ipv6": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"last_update_time": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"overlapping": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"owner": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"parent_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"shared": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"total_ip_address_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"used_ip_address_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"used_percentage": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

type globalPoolCompare struct {
	ID            string
	IPPoolName    string
	IPPoolCidr    string
	Gateways      []string
	DhcpServerIPs *[]string
	DNSServerIPs  *[]string
}

func hasGlobalPool(response *dnac.GetGlobalPoolResponse, expected *globalPoolCompare) (bool, bool, *dnac.GetGlobalPoolResponseResponse) {
	if response != nil {
		var foundValue *dnac.GetGlobalPoolResponseResponse
		for _, item := range response.Response {
			if expected.ID != "" {
				if expected.ID == item.ID {
					foundValue = &item
					break
				}
			} else {
				if expected.IPPoolName == item.IPPoolName {
					foundValue = &item
					break
				}
			}
		}
		if foundValue != nil {
			cmpResult := (expected.IPPoolName != foundValue.IPPoolName) ||
				(expected.IPPoolCidr != foundValue.IPPoolCidr) ||
				!hasSameSliceString(expected.Gateways, foundValue.Gateways) ||
				(expected.DhcpServerIPs != nil && !hasSameSliceString(*expected.DhcpServerIPs, foundValue.DhcpServerIPs)) ||
				(expected.DNSServerIPs != nil && !hasSameSliceString(*expected.DNSServerIPs, foundValue.DNSServerIPs))
			return false, cmpResult, foundValue
		}
		return true, false, nil
	}
	return true, false, nil
}

func resourceNetworkGlobalIPPoolCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)
	var diags diag.Diagnostics

	items := d.Get("item").([]interface{})[0]
	item := items.(map[string]interface{})

	var id string
	if v, ok := item["id"]; ok {
		id = v.(string)
	}
	ipPoolName := item["ip_pool_name"].(string)
	ipPoolCIDR := item["ip_pool_cidr"].(string)
	gateways := d.Get("gateway").(string)
	typeV := d.Get("type").(string)
	ipAddressSpace := d.Get("ip_address_space").(string)

	var dhcpServerIPs *[]string
	if v, ok := item["dhcp_server_ips"]; ok {
		dhcpServerIPtmp := convertSliceInterfaceToSliceString(v.([]interface{}))
		dhcpServerIPs = &dhcpServerIPtmp
	}
	var dnsServerIPs *[]string
	if v, ok := item["dns_server_ips"]; ok {
		dnsServerIPstmp := convertSliceInterfaceToSliceString(v.([]interface{}))
		dnsServerIPs = &dnsServerIPstmp
	}

	userRequest := globalPoolCompare{
		ID:            id,
		IPPoolName:    ipPoolName,
		IPPoolCidr:    ipPoolCIDR,
		Gateways:      strings.Split(gateways, ","),
		DhcpServerIPs: dhcpServerIPs,
		DNSServerIPs:  dnsServerIPs,
	}

	searchResponse, _, err := client.NetworkSettings.GetGlobalPool(&dnac.GetGlobalPoolQueryParams{})
	if err == nil && searchResponse != nil {
		// Check if element already exists
		_, performUpdate, foundValue := hasGlobalPool(searchResponse, &userRequest)
		log.Printf("performUpdate %v, foundValue %v", performUpdate, foundValue)
		if performUpdate {
			id = foundValue.ID
			updateRequest := dnac.UpdateGlobalPoolRequest{
				Settings: dnac.UpdateGlobalPoolRequestSettings{
					IPpool: []dnac.UpdateGlobalPoolRequestSettingsIPpool{
						{
							DhcpServerIPs: dhcpServerIPs,
							DNSServerIPs:  dnsServerIPs,
							Gateway:       gateways,
							ID:            id,
							IPPoolName:    ipPoolName,
						},
					},
				},
			}

			log.Printf("updateRequest %+v", updateRequest)
			updateResponse, _, err := client.NetworkSettings.UpdateGlobalPool(&updateRequest)
			if err != nil {
				return diag.FromErr(err)
			}
			log.Printf("updateResponse %+v", updateResponse)

			// Wait for execution status to complete
			time.Sleep(5 * time.Second)

			// Update resource id

			d.SetId(foundValue.ID)
			resourceNetworkGlobalIPPoolRead(ctx, d, m)
			return diags
		}
	}

	var dhcpServerIPsCreate []string
	var dnsServerIPsCreate []string
	if dhcpServerIPs != nil {
		dhcpServerIPsCreate = *dhcpServerIPs
	}
	if dnsServerIPs != nil {
		dnsServerIPsCreate = *dnsServerIPs
	}
	// Construct payload from resource schema (item)
	createRequest := dnac.CreateGlobalPoolRequest{
		Settings: dnac.CreateGlobalPoolRequestSettings{
			IPpool: []dnac.CreateGlobalPoolRequestSettingsIPpool{
				{
					IPAddressSpace: ipAddressSpace,
					DhcpServerIPs:  dhcpServerIPsCreate,
					DNSServerIPs:   dnsServerIPsCreate,
					Gateway:        gateways,
					IPPoolCidr:     ipPoolCIDR,
					IPPoolName:     ipPoolName,
					Type:           typeV,
				},
			},
		},
	}
	_, _, err = client.NetworkSettings.CreateGlobalPool(&createRequest)
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for execution status to complete
	time.Sleep(5 * time.Second)

	searchResponse2, _, err := client.NetworkSettings.GetGlobalPool(&dnac.GetGlobalPoolQueryParams{})
	if err != nil {
		return diag.FromErr(err)
	}

	if err == nil && searchResponse2 != nil {
		_, _, foundValue2 := hasGlobalPool(searchResponse2, &userRequest)

		// Update resource id
		d.SetId(foundValue2.ID)
		resourceNetworkGlobalIPPoolRead(ctx, d, m)
		return diags
	}

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Unable to get created global pool",
	})
	return diags
}

func resourceNetworkGlobalIPPoolRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	poolID := d.Id()
	searchResponse, _, err := client.NetworkSettings.GetGlobalPool(&dnac.GetGlobalPoolQueryParams{})
	if err != nil {
		d.SetId("")
		return diags
	}
	userRequest := globalPoolCompare{ID: poolID}
	_, _, foundValue := hasGlobalPool(searchResponse, &userRequest)

	if foundValue == nil {
		d.SetId("")
		return diags
	}

	ipPool := flattenNetworkGlobalIPPoolReadItem(foundValue)
	if err := d.Set("item", ipPool); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceNetworkGlobalIPPoolUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	id := d.Id()

	poolID := id
	searchResponse, _, err := client.NetworkSettings.GetGlobalPool(&dnac.GetGlobalPoolQueryParams{})
	if err != nil {
		d.SetId("")
		return diags
	}
	userRequest := globalPoolCompare{ID: poolID}
	_, _, foundValue := hasGlobalPool(searchResponse, &userRequest)
	if foundValue == nil {
		d.SetId("")
		return diags
	}

	// Check if properties inside resource has changes
	if d.HasChange("item") {

		items := d.Get("item").([]interface{})[0]
		item := items.(map[string]interface{})

		ipPoolName := item["ip_pool_name"].(string)
		gateways := d.Get("gateway").(string)

		var dhcpServerIPs *[]string
		if v, ok := item["dhcp_server_ips"]; ok {
			dhcpServerIPtmp := convertSliceInterfaceToSliceString(v.([]interface{}))
			dhcpServerIPs = &dhcpServerIPtmp
		}
		var dnsServerIPs *[]string
		if v, ok := item["dns_server_ips"]; ok {
			dnsServerIPstmp := convertSliceInterfaceToSliceString(v.([]interface{}))
			dnsServerIPs = &dnsServerIPstmp
		}
		updateRequest := dnac.UpdateGlobalPoolRequest{
			Settings: dnac.UpdateGlobalPoolRequestSettings{
				IPpool: []dnac.UpdateGlobalPoolRequestSettingsIPpool{
					{
						DhcpServerIPs: dhcpServerIPs,
						DNSServerIPs:  dnsServerIPs,
						Gateway:       gateways,
						ID:            id,
						IPPoolName:    ipPoolName,
					},
				},
			},
		}

		log.Printf("updateRequest %+v", updateRequest)
		updateResponse, _, err := client.NetworkSettings.UpdateGlobalPool(&updateRequest)
		if err != nil {
			return diag.FromErr(err)
		}
		log.Printf("updateResponse %+v", updateResponse)
		// Wait for execution status to complete
		time.Sleep(5 * time.Second)

		// Update resource last_updated
		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	// Update resource on Terraform by using read function
	return resourceNetworkGlobalIPPoolRead(ctx, d, m)
}

func resourceNetworkGlobalIPPoolDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	poolID := d.Id()
	searchResponse, _, err := client.NetworkSettings.GetGlobalPool(&dnac.GetGlobalPoolQueryParams{})
	if err == nil && searchResponse != nil {
		// Check if element already exists
		userRequest := globalPoolCompare{ID: poolID}
		_, _, foundValue := hasGlobalPool(searchResponse, &userRequest)
		if foundValue == nil {
			return diags
		}
	}

	// Call function to delete resource
	_, _, err = client.NetworkSettings.DeleteGlobalIPPool(poolID)
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for execution status to complete
	time.Sleep(5 * time.Second)

	searchResponse, _, err = client.NetworkSettings.GetGlobalPool(&dnac.GetGlobalPoolQueryParams{})
	if err == nil && searchResponse != nil {
		// Check if element already exists
		userRequest := globalPoolCompare{ID: poolID}
		_, _, foundValue := hasGlobalPool(searchResponse, &userRequest)
		if foundValue != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to delete Global IP Pool",
			})
			return diags
		}
	}

	return diags
}
