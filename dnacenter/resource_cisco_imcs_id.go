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

func resourceCiscoImcsID() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read, update and delete operations on Cisco IMC.

- This resource removes a specific Cisco Integrated Management Controller (IMC) configuration from a Catalyst Center
node using the provided identifier.
The Cisco IMC configuration APIs enable the management of connections between Cisco IMC and Cisco Catalyst Center. By
providing the Cisco IMC IP address and credentials to Catalyst Center, Catalyst Center can access and report the health
status of hardware components within the Cisco appliance.
More data about the Cisco IMC can be retrieved using the APIs exposed directly by Cisco IMC. Details are available in
the Cisco IMC documentation https://www.cisco.com/c/en/us/support/servers-unified-computing/ucs-c-series-integrated-
management-controller/series.html#~tab-documents
The Cisco IMC configuration is relevant only for Catalyst Center deployments based on UCS appliances. In cases where
Cisco IMC configuration is not supported by the deployment, these APIs will respond with a *404 Not Found* status code.

- This resource updates the Cisco Integrated Management Controller (IMC) configuration for a Catalyst Center node,
identified by the specified ID.
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

		CreateContext: resourceCiscoImcsIDCreate,
		ReadContext:   resourceCiscoImcsIDRead,
		UpdateContext: resourceCiscoImcsIDUpdate,
		DeleteContext: resourceCiscoImcsIDDelete,
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

						"id": &schema.Schema{
							Description: `id path parameter. The unique identifier for this Cisco IMC configuration
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"ip_address": &schema.Schema{
							Description: `IP address of the Cisco IMC
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

func resourceCiscoImcsIDCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	// TODO: Add the path params to `item` schema
	//       & return it individually
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	resourceMap["name"] = interfaceToString(resourceItem["name"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceCiscoImcsIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesTheCiscoIMCConfigurationForACatalystCenterNode")
		vvID := vID

		response1, restyResp1, err := client.CiscoIMC.RetrievesTheCiscoIMCConfigurationForACatalystCenterNode(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenCiscoIMCRetrievesTheCiscoIMCConfigurationForACatalystCenterNodeItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesTheCiscoIMCConfigurationForACatalystCenterNode response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceCiscoImcsIDUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestCiscoImcsIDUpdatesTheCiscoIMCConfigurationForACatalystCenterNode(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.CiscoIMC.UpdatesTheCiscoIMCConfigurationForACatalystCenterNode(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdatesTheCiscoIMCConfigurationForACatalystCenterNode", err, restyResp1.String(),
					"Failure at UpdatesTheCiscoIMCConfigurationForACatalystCenterNode, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdatesTheCiscoIMCConfigurationForACatalystCenterNode", err,
				"Failure at UpdatesTheCiscoIMCConfigurationForACatalystCenterNode, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdatesTheCiscoIMCConfigurationForACatalystCenterNode", err))
			return diags
		}
		taskId := response1.Response.TaskID
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
					"Failure when executing UpdatesTheCiscoIMCConfigurationForACatalystCenterNode", err1))
				return diags
			}
		}

	}

	return resourceCiscoImcsIDRead(ctx, d, m)
}

func resourceCiscoImcsIDDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["id"]
	restyResp1, err := client.CiscoIMC.DeletesTheCiscoIMCConfigurationForACatalystCenterNode(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeletesTheCiscoIMCConfigurationForACatalystCenterNode", err, restyResp1.String(),
				"Failure at DeletesTheCiscoIMCConfigurationForACatalystCenterNode, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeletesTheCiscoIMCConfigurationForACatalystCenterNode", err,
			"Failure at DeletesTheCiscoIMCConfigurationForACatalystCenterNode, unexpected response", ""))
		return diags
	}

	//TODO REVIEW

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestCiscoImcsIDUpdatesTheCiscoIMCConfigurationForACatalystCenterNode(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestCiscoIMCUpdatesTheCiscoIMCConfigurationForACatalystCenterNode {
	request := dnacentersdkgo.RequestCiscoIMCUpdatesTheCiscoIMCConfigurationForACatalystCenterNode{}
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
