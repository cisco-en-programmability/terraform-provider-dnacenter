package dnacenter

import (
	"context"
	"errors"
	"reflect"
	"time"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCiscoImcs() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create and read operations on Cisco IMC.

- This resource adds a Cisco Integrated Management Controller (IMC) configuration to a Cisco Catalyst Center node,
identified by its *nodeId*. Obtain the *nodeId* from the *id* attribute in the response of the
*/dna/intent/api/v1/nodes-config* API.
The Cisco IMC configuration APIs enable the management of connections between Cisco IMC and Cisco Catalyst Center. By
providing the Cisco IMC IP address and credentials to Catalyst Center, Catalyst Center can access and report the health
status of hardware components within the Cisco appliance.
More data about the Cisco IMC can be retrieved using the APIs exposed directly by Cisco IMC. Details are available in
the Cisco IMC documentation https://www.cisco.com/c/en/us/support/servers-unified-computing/ucs-c-series-integrated-
management-controller/series.html#~tab-documents
The Cisco IMC configuration is relevant only for Catalyst Center deployments based on UCS appliances. In cases where
Cisco IMC configuration is not supported by the deployment, these APIs will respond with a *404 Not Found* status code.
When Cisco IMC configuration is supported, this API responds with the URL of a diagnostic task.
`,

		CreateContext: resourceCiscoImcsCreate,
		ReadContext:   resourceCiscoImcsRead,
		UpdateContext: resourceCiscoImcsUpdate,
		DeleteContext: resourceCiscoImcsDelete,
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

						"id": &schema.Schema{
							Description: `The unique identifier for this Cisco IMC configuration
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_address": &schema.Schema{
							Description: `IP address of the Cisco IMC
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"node_id": &schema.Schema{
							Description: `The UUID that represents the Catalyst Center node. Its value can be obtained from the *id* attribute of the response of the */dna/intent/api/v1/nodes-config* API.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"username": &schema.Schema{
							Description: `Username of the Cisco IMC
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ip_address": &schema.Schema{
							Description: `IP address of the Cisco IMC
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"node_id": &schema.Schema{
							Description: `The UUID that represents the Catalyst Center node. Its value can be obtained from the *id* attribute of the response of the */dna/intent/api/v1/nodes-config* API.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"password": &schema.Schema{
							Description: `Password of the Cisco IMC
`,
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"username": &schema.Schema{
							Description: `Username of the Cisco IMC
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceCiscoImcsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestCiscoImcsAddsCiscoIMCConfigurationToACatalystCenterNode(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	vNodeId := resourceItem["node_id"]
	vvNodeId := interfaceToString(vNodeId)

	item2, err := searchCiscoIMCRetrievesCiscoIMCConfigurationsForCatalystCenterNodes(m, vvNodeId)
	if err != nil || item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["node_id"] = item2.NodeID
		d.SetId(joinResourceID(resourceMap))
		return resourceCiscoImcsRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.CiscoIMC.AddsCiscoIMCConfigurationToACatalystCenterNode(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddsCiscoIMCConfigurationToACatalystCenterNode", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddsCiscoIMCConfigurationToACatalystCenterNode", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing AddsCiscoIMCConfigurationToACatalystCenterNode", err))
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
			errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing AddsCiscoIMCConfigurationToACatalystCenterNode", err1))
			return diags
		}
	}
	item3, err := searchCiscoIMCRetrievesCiscoIMCConfigurationsForCatalystCenterNodes(m, vvNodeId)
	if err != nil || item3 != nil {
		resourceMap := make(map[string]string)
		resourceMap["node_id"] = item3.NodeID
		d.SetId(joinResourceID(resourceMap))
		return resourceCiscoImcsRead(ctx, d, m)
	}
	item4, err := searchCiscoIMCRetrievesCiscoIMCConfigurationsForCatalystCenterNodes(m, vvNodeId)
	if err != nil || item4 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AddsCiscoIMCConfigurationToACatalystCenterNode", err,
			"Failure at AddsCiscoIMCConfigurationToACatalystCenterNode, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["node_id"] = vvNodeId
	d.SetId(joinResourceID(resourceMap))
	return resourceCiscoImcsRead(ctx, d, m)
}

func resourceCiscoImcsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvNodeId := resourceMap["node_id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesCiscoIMCConfigurationsForCatalystCenterNodes")
		response1, restyResp1, err := client.CiscoIMC.RetrievesCiscoIMCConfigurationsForCatalystCenterNodes()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		item1, err := searchCiscoIMCRetrievesCiscoIMCConfigurationsForCatalystCenterNodes(m, vvNodeId)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		vItem1 := flattenCiscoIMCRetrievesCiscoIMCConfigurationsForCatalystCenterNodesByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesCiscoIMCConfigurationsForCatalystCenterNodes search response",
				err))
			return diags
		}

	}
	return diags
}

func flattenCiscoIMCRetrievesCiscoIMCConfigurationsForCatalystCenterNodesByIDItem(item *dnacentersdkgo.ResponseCiscoIMCRetrievesCiscoIMCConfigurationsForCatalystCenterNodesResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["nodeId"] = item.NodeID
	respItem["ipAddress"] = item.IPAddress
	respItem["username"] = item.Username
	return []map[string]interface{}{
		respItem,
	}
}

func resourceCiscoImcsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCiscoImcsRead(ctx, d, m)
}

func resourceCiscoImcsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete CiscoImcs on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestCiscoImcsAddsCiscoIMCConfigurationToACatalystCenterNode(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestCiscoIMCaddsCiscoIMCConfigurationToACatalystCenterNode {
	request := dnacentersdkgo.RequestCiscoIMCaddsCiscoIMCConfigurationToACatalystCenterNode{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".node_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".node_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".node_id")))) {
		request.NodeID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_address")))) {
		request.IPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchCiscoIMCRetrievesCiscoIMCConfigurationsForCatalystCenterNodes(m interface{}, vID string) (*dnacentersdkgo.ResponseCiscoIMCRetrievesCiscoIMCConfigurationsForCatalystCenterNodesResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseCiscoIMCRetrievesCiscoIMCConfigurationsForCatalystCenterNodesResponse
	var ite *dnacentersdkgo.ResponseCiscoIMCRetrievesCiscoIMCConfigurationsForCatalystCenterNodes
	ite, _, err = client.CiscoIMC.RetrievesCiscoIMCConfigurationsForCatalystCenterNodes()
	if err != nil || ite == nil {
		return foundItem, err

	}
	items := ite
	itemsCopy := *items
	for _, item := range *itemsCopy.Response {
		// Call get by _ method and set value to foundItem and return
		if item.NodeID == vID {
			var getItem *dnacentersdkgo.ResponseCiscoIMCRetrievesCiscoIMCConfigurationsForCatalystCenterNodesResponse
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
