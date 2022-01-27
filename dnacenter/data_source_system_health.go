package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSystemHealth() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Health and Performance.

- This data source retrieves the latest system events
`,

		ReadContext: dataSourceSystemHealthRead,
		Schema: map[string]*schema.Schema{
			"domain": &schema.Schema{
				Description: `domain query parameter. Fetch system events with this domain. Possible values of domain are listed here : /dna/platform/app/consumer-portal/developer-toolkit/events
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter.`,
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter.`,
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"subdomain": &schema.Schema{
				Description: `subdomain query parameter. Fetch system events with this subdomain. Possible values of subdomain are listed here : /dna/platform/app/consumer-portal/developer-toolkit/events
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"summary": &schema.Schema{
				Description: `summary query parameter. Fetch the latest high severity event
`,
				Type:     schema.TypeBool,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"health_events": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"description": &schema.Schema{
										Description: `Details of the event
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"domain": &schema.Schema{
										Description: `Domain of the event
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"hostname": &schema.Schema{
										Description: `Hostname of the event
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance": &schema.Schema{
										Description: `Instance of the event
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"severity": &schema.Schema{
										Description: `Severity of the event
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"state": &schema.Schema{
										Description: `State of the event
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"status": &schema.Schema{
										Description: `Event status
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"sub_domain": &schema.Schema{
										Description: `Sub domain of the event
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"timestamp": &schema.Schema{
										Description: `Time of the event occurance
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"host_name": &schema.Schema{
							Description: `Cluster name 
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"version": &schema.Schema{
							Description: `API version
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

func dataSourceSystemHealthRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSummary, okSummary := d.GetOk("summary")
	vDomain, okDomain := d.GetOk("domain")
	vSubdomain, okSubdomain := d.GetOk("subdomain")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: SystemHealthApI")
		queryParams1 := dnacentersdkgo.SystemHealthApIQueryParams{}

		if okSummary {
			queryParams1.Summary = vSummary.(bool)
		}
		if okDomain {
			queryParams1.Domain = vDomain.(string)
		}
		if okSubdomain {
			queryParams1.Subdomain = vSubdomain.(string)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}

		response1, restyResp1, err := client.HealthAndPerformance.SystemHealthApI(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing SystemHealthApI", err,
				"Failure at SystemHealthApI, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenHealthAndPerformanceSystemHealthApIItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting SystemHealthApI response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenHealthAndPerformanceSystemHealthApIItem(item *dnacentersdkgo.ResponseHealthAndPerformanceSystemHealthApI) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["health_events"] = flattenHealthAndPerformanceSystemHealthApIItemHealthEvents(item.HealthEvents)
	respItem["version"] = item.Version
	respItem["host_name"] = item.HostName
	return []map[string]interface{}{
		respItem,
	}
}

func flattenHealthAndPerformanceSystemHealthApIItemHealthEvents(items *[]dnacentersdkgo.ResponseHealthAndPerformanceSystemHealthAPIHealthEvents) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["severity"] = item.Severity
		respItem["hostname"] = item.Hostname
		respItem["instance"] = item.Instance
		respItem["sub_domain"] = item.SubDomain
		respItem["domain"] = item.Domain
		respItem["description"] = item.Description
		respItem["state"] = item.State
		respItem["timestamp"] = item.Timestamp
		respItem["status"] = item.Status
		respItems = append(respItems, respItem)
	}
	return respItems
}
