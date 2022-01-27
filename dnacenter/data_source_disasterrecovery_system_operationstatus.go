package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDisasterrecoverySystemOperationstatus() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Disaster Recovery.

- Returns the status of Disaster Recovery operation performed on the system.
`,

		ReadContext: dataSourceDisasterrecoverySystemOperationstatusRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"end_timestamp": &schema.Schema{
							Description: `End timestamp of the DR event
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"initiated_by": &schema.Schema{
							Description: `Who initiated this event. Is it a system triggered one or user triggered one.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ipconfig": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"interface": &schema.Schema{
										Description: `Enterprise or Management interface
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ip": &schema.Schema{
										Description: `This floating virtual IP address automatically moves to and operates on the site that is currently acting as your network's active site. If the interface is enterprise, then this address manages traffic between your disaster recovery system and your Enterprise network. If the interface is management, then this address manages traffic between your disaster recovery system and your Management network.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"vip": &schema.Schema{
										Description: `Is this interface a Virtual IP address or not. This is true for Site VIP
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"message": &schema.Schema{
							Description: `Detailed Description about the DR event
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"severity": &schema.Schema{
							Description: `Severity of the DR Event.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"site": &schema.Schema{
							Description: `Site of the DR in which this event occurred.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"start_timestamp": &schema.Schema{
							Description: `Starting timestamp of the DR event
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"status": &schema.Schema{
							Description: `Status of the DR Event.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"tasks": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"end_timestamp": &schema.Schema{
										Description: `End timestamp of the DR event
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ipconfig": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"interface": &schema.Schema{
													Description: `Enterprise or Management interface
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"ip": &schema.Schema{
													Description: `This floating virtual IP address automatically moves to and operates on the site that is currently acting as your network's active site. If the interface is enterprise, then this address manages traffic between your disaster recovery system and your Enterprise network. If the interface is management, then this address manages traffic between your disaster recovery system and your Management network.
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"vip": &schema.Schema{
													Description: `Is this interface a Virtual IP address or not. This is true for Site VIP
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"message": &schema.Schema{
										Description: `Detailed description about the DR event
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"site": &schema.Schema{
										Description: `Site of the DR in which this event occured
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"start_timestamp": &schema.Schema{
										Description: `Starting timestamp of the DR event
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"status": &schema.Schema{
										Description: `Status of the DR event. 
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"title": &schema.Schema{
										Description: `DR Event Summary
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"title": &schema.Schema{
							Description: `DR Event Summary
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

func dataSourceDisasterrecoverySystemOperationstatusRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: DisasterRecoveryOperationalStatus")

		response1, restyResp1, err := client.DisasterRecovery.DisasterRecoveryOperationalStatus()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing DisasterRecoveryOperationalStatus", err,
				"Failure at DisasterRecoveryOperationalStatus, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDisasterRecoveryDisasterRecoveryOperationalStatusItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting DisasterRecoveryOperationalStatus response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDisasterRecoveryDisasterRecoveryOperationalStatusItem(item *dnacentersdkgo.ResponseDisasterRecoveryDisasterRecoveryOperationalStatus) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["severity"] = item.Severity
	respItem["status"] = item.Status
	respItem["initiated_by"] = item.InitiatedBy
	respItem["ipconfig"] = flattenDisasterRecoveryDisasterRecoveryOperationalStatusItemIPconfig(item.IPconfig)
	respItem["tasks"] = flattenDisasterRecoveryDisasterRecoveryOperationalStatusItemTasks(item.Tasks)
	respItem["title"] = item.Title
	respItem["site"] = item.Site
	respItem["start_timestamp"] = item.StartTimestamp
	respItem["message"] = item.Message
	respItem["end_timestamp"] = item.EndTimestamp
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDisasterRecoveryDisasterRecoveryOperationalStatusItemIPconfig(items *[]dnacentersdkgo.ResponseDisasterRecoveryDisasterRecoveryOperationalStatusIPconfig) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["interface"] = item.Interface
		respItem["vip"] = item.Vip
		respItem["ip"] = item.IP
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDisasterRecoveryDisasterRecoveryOperationalStatusItemTasks(items *[]dnacentersdkgo.ResponseDisasterRecoveryDisasterRecoveryOperationalStatusTasks) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["status"] = item.Status
		respItem["ipconfig"] = flattenDisasterRecoveryDisasterRecoveryOperationalStatusItemTasksIPconfig(item.IPconfig)
		respItem["title"] = item.Title
		respItem["site"] = item.Site
		respItem["start_timestamp"] = item.StartTimestamp
		respItem["message"] = item.Message
		respItem["end_timestamp"] = item.EndTimestamp
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDisasterRecoveryDisasterRecoveryOperationalStatusItemTasksIPconfig(items *[]dnacentersdkgo.ResponseDisasterRecoveryDisasterRecoveryOperationalStatusTasksIPconfig) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["interface"] = item.Interface
		respItem["vip"] = item.Vip
		respItem["ip"] = item.IP
		respItems = append(respItems, respItem)
	}
	return respItems
}
