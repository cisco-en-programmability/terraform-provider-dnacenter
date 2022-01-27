package dnacenter

import (
	"context"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSdaFabric() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on SDA.

- Delete SDA Fabric

- Add SDA Fabric
`,

		CreateContext: resourceSdaFabricCreate,
		ReadContext:   resourceSdaFabricRead,
		UpdateContext: resourceSdaFabricUpdate,
		DeleteContext: resourceSdaFabricDelete,
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

						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"execution_id": &schema.Schema{
							Description: `Execution Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"fabric_domain_type": &schema.Schema{
							Description: `Fabric Domain type
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"fabric_name": &schema.Schema{
							Description: `Fabric name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"fabric_type": &schema.Schema{
							Description: `Fabric type
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"status": &schema.Schema{
							Description: `Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"fabric_name": &schema.Schema{
							Description: `Fabric Name (from DNAC2.2.3 onwards following default fabric name  Default_LAN_Fabric)
`,
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceSdaFabricCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSdaFabricAddFabric(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vFabricName := resourceItem["fabric_name"]
	vvFabricName := interfaceToString(vFabricName)

	queryParams1 := dnacentersdkgo.GetSdaFabricInfoQueryParams{}

	queryParams1.FabricName = vvFabricName

	getResponse2, _, err := client.Sda.GetSdaFabricInfo(&queryParams1)
	if getResponse2.Status != "failed" {
		resourceMap := make(map[string]string)
		resourceMap["fabric_name"] = vvFabricName
		d.SetId(joinResourceID(resourceMap))
		return resourceSdaFabricRead(ctx, d, m)
	}

	response1, restyResp1, err := client.Sda.AddFabric(request1)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddFabric", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddFabric", err))
		return diags
	}
	executionId := response1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionId)
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp1, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetExecutionByID", err,
				"Failure at GetExecutionByID, unexpected response", ""))
			return diags
		}
		for response2.Status == "IN_PROGRESS" {
			time.Sleep(10 * time.Second)
			response2, restyResp1, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp1 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}
		if response2.Status == "FAILURE" {
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing AddFabric", err))
			return diags
		}
	}
	resourceMap := make(map[string]string)
	resourceMap["fabric_name"] = vvFabricName
	d.SetId(joinResourceID(resourceMap))
	return resourceSdaFabricRead(ctx, d, m)
}

func resourceSdaFabricRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vFabricName := resourceMap["fabric_name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetSdaFabricInfo")
		queryParams1 := dnacentersdkgo.GetSdaFabricInfoQueryParams{}

		queryParams1.FabricName = vFabricName

		response1, restyResp1, err := client.Sda.GetSdaFabricInfo(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			// diags = append(diags, diagErrorWithAlt(
			// 	"Failure when executing GetSdaFabricInfo", err,
			// 	"Failure at GetSdaFabricInfo, unexpected response", ""))
			// return diags
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetSdaFabricInfoItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSdaFabricInfo response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceSdaFabricUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceSdaFabricRead(ctx, d, m)
}

func resourceSdaFabricDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vFabricName := resourceMap["fabric_name"]

	queryParams1 := dnacentersdkgo.GetSdaFabricInfoQueryParams{}
	queryParams1.FabricName = vFabricName
	item, restyResp1, err := client.Sda.GetSdaFabricInfo(&queryParams1)
	if err != nil || item == nil || item.Status == "failed" {
		/*diags = append(diags, diagErrorWithAlt(
		"Failure when executing GetSDAFabricInfo", err,
		"Failure at GetSDAFabricInfo, unexpected response", ""))*/
		d.SetId("")
		return diags
	}

	queryParams2 := dnacentersdkgo.DeleteSdaFabricQueryParams{}
	queryParams2.FabricName = vFabricName
	response1, restyResp1, err := client.Sda.DeleteSdaFabric(&queryParams2)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteSdaFabric", err, restyResp1.String(),
				"Failure at DeleteSdaFabric, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteSdaFabric", err,
			"Failure at DeleteSdaFabric, unexpected response", ""))
		return diags
	}
	executionId := response1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionId)
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp1, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetExecutionByID", err,
				"Failure at GetExecutionByID, unexpected response", ""))
			return diags
		}
		for response2.Status == "IN_PROGRESS" {
			time.Sleep(10 * time.Second)
			response2, restyResp1, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp1 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}
		if response2.Status == "FAILURE" {
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing DeleteSdaFabric", err))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSdaFabricAddFabric(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddFabric {
	request := dnacentersdkgo.RequestSdaAddFabric{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fabric_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fabric_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fabric_name")))) {
		request.FabricName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
