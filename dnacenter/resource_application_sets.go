package dnacenter

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceApplicationSets() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on Application Policy.

- Delete existing application-set by it's id

- Create new custom application-set/s
`,

		CreateContext: resourceApplicationSetsCreate,
		ReadContext:   resourceApplicationSetsRead,
		UpdateContext: resourceApplicationSetsUpdate,
		DeleteContext: resourceApplicationSetsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"application_set": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id_ref": &schema.Schema{
										Description: `Id Ref`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"indicative_network_identity": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"display_name": &schema.Schema{
										Description: `displayName`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"id": &schema.Schema{
										Description: `id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"lower_port": &schema.Schema{
										Description: `lowerPort`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"ports": &schema.Schema{
										Description: `ports`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"protocol": &schema.Schema{
										Description: `protocol`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"upper_port": &schema.Schema{
										Description: `upperPort`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"network_applications": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"app_protocol": &schema.Schema{
										Description: `App Protocol`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"application_sub_type": &schema.Schema{
										Description: `Application Sub Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"application_type": &schema.Schema{
										Description: `Application Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"category_id": &schema.Schema{
										Description: `Category Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"display_name": &schema.Schema{
										Description: `Display Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"dscp": &schema.Schema{
										Description: `Dscp`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"engine_id": &schema.Schema{
										Description: `Engine Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"help_string": &schema.Schema{
										Description: `Help String`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ignore_conflict": &schema.Schema{
										Description: `Ignore Conflict`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"long_description": &schema.Schema{
										Description: `Long Description`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"popularity": &schema.Schema{
										Description: `Popularity`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"rank": &schema.Schema{
										Description: `Rank`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"server_name": &schema.Schema{
										Description: `Server Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"traffic_class": &schema.Schema{
										Description: `Traffic Class`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"url": &schema.Schema{
										Description: `Url`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"network_identity": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"display_name": &schema.Schema{
										Description: `Display Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"lower_port": &schema.Schema{
										Description: `Lower Port`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"ports": &schema.Schema{
										Description: `Ports`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"protocol": &schema.Schema{
										Description: `Protocol`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"upper_port": &schema.Schema{
										Description: `Upper Port`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestApplicationPolicyCreateApplicationSet`,
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"id": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func resourceApplicationSetsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning ApplicationSets Create")
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vName := resourceItem["name"]

	vvName := interfaceToString(vName)

	queryParams1 := dnacentersdkgo.GetApplicationSetsQueryParams{}

	queryParams1.Name = vvName

	item, err := searchApplicationPolicyGetApplicationSets(m, queryParams1)
	if err != nil || item != nil {
		resourceMap := make(map[string]string)
		resourceMap["name"] = vvName
		d.SetId(joinResourceID(resourceMap))
		return resourceApplicationSetsRead(ctx, d, m)
	}
	request1 := expandRequestApplicationSetsCreateApplicationSet(ctx, "parameters", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	resp1, restyResp1, err := client.ApplicationPolicy.CreateApplicationSet(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateApplicationSet", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateApplicationSet", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing CreateApplicationSet", err))
		return diags
	}
	taskId := resp1.Response.TaskID
	log.Printf("[DEBUG] TASKID => %s", taskId)
	if taskId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetTaskByID(taskId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTaskByID", err,
				"Failure at GetTaskByID, unexpected response", ""))
			return diags
		}
		if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
			log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
			diags = append(diags, diagError(
				"Failure when executing CreateApplicationSet", err))
			return diags
		}
	}

	resourceMap := make(map[string]string)
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return resourceApplicationSetsRead(ctx, d, m)
}

func resourceApplicationSetsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName := resourceMap["name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetApplicationSets")
		queryParams1 := dnacentersdkgo.GetApplicationSetsQueryParams{}

		queryParams1.Name = vName

		response1, err := searchApplicationPolicyGetApplicationSets(m, queryParams1)

		if err != nil || response1 == nil {
			// diags = append(diags, diagErrorWithAlt(
			// 	"Failure when executing GetApplicationSets", err,
			// 	"Failure at GetApplicationSets, unexpected response", ""))
			// return diags
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenApplicationPolicyGetApplicationSetsItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetApplicationSets search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceApplicationSetsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceApplicationSetsRead(ctx, d, m)
}

func resourceApplicationSetsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName := resourceMap["name"]
	var vID string
	selectedMethod := 1
	queryParams1 := dnacentersdkgo.GetApplicationSetsQueryParams{}

	queryParams1.Name = vName

	queryParams2 := dnacentersdkgo.DeleteApplicationSetQueryParams{}

	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 1 {

		item1, err := searchApplicationPolicyGetApplicationSets(m, queryParams1)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		vID = item1.ID
	}

	queryParams2.ID = vID
	response1, restyResp1, err := client.ApplicationPolicy.DeleteApplicationSet(&queryParams2)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteApplicationSet", err, restyResp1.String(),
				"Failure at DeleteApplicationSet, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteApplicationSet", err,
			"Failure at DeleteApplicationSet, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestApplicationSetsCreateApplicationSet(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationSet {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationSet{}
	if v := expandRequestApplicationSetsCreateApplicationSetItemArray(ctx, key, d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestApplicationSetsCreateApplicationSetItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationSet {
	request := []dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationSet{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestApplicationSetsCreateApplicationSetItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestApplicationSetsCreateApplicationSetItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationSet {
	request := dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationSet{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func searchApplicationPolicyGetApplicationSets(m interface{}, queryParams dnacentersdkgo.GetApplicationSetsQueryParams) (*dnacentersdkgo.ResponseApplicationPolicyGetApplicationSetsResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseApplicationPolicyGetApplicationSetsResponse
	var ite *dnacentersdkgo.ResponseApplicationPolicyGetApplicationSets
	ite, _, err = client.ApplicationPolicy.GetApplicationSets(&queryParams)

	if ite == nil {
		return foundItem, err
	}

	if ite.Response == nil {
		return foundItem, err
	}

	items := ite.Response
	itemsCopy := *items
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if item.Name == queryParams.Name {
			var getItem *dnacentersdkgo.ResponseApplicationPolicyGetApplicationSetsResponse
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
