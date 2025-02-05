package dnacenter

import (
	"context"
	"errors"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceImagesDistributionServerSettings() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Software Image Management (SWIM).

- Add remote server for distributing software images. Upto two such distribution servers are supported.
`,

		CreateContext: resourceImagesDistributionServerSettingsCreate,
		ReadContext:   resourceImagesDistributionServerSettingsRead,
		UpdateContext: resourceImagesDistributionServerSettingsUpdate,
		DeleteContext: resourceImagesDistributionServerSettingsDelete,
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
							Description: `Unique identifier for the server
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"port_number": &schema.Schema{
							Description: `Port number
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"root_location": &schema.Schema{
							Description: `Server root location
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"server_address": &schema.Schema{
							Description: `FQDN or IP address of the server
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"username": &schema.Schema{
							Description: `Server username
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

						"password": &schema.Schema{
							Description: `Server password
`,
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"port_number": &schema.Schema{
							Description: `Port number
`,
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
						},
						"root_location": &schema.Schema{
							Description: `Server root location
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"server_address": &schema.Schema{
							Description: `FQDN or IP address of the server
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"username": &schema.Schema{
							Description: `Server username
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

func resourceImagesDistributionServerSettingsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestImagesDistributionServerSettingsAddImageDistributionServer(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	vServer := resourceItem["server_address"]
	vvServer := interfaceToString(vServer)
	item2, err := searchSoftwareImageManagementSwimRetrieveImageDistributionServers(m, vvServer)
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["server_address"] = vvServer
		d.SetId(joinResourceID(resourceMap))
		return resourceImagesDistributionServerSettingsRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.SoftwareImageManagementSwim.AddImageDistributionServer(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddImageDistributionServer", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddImageDistributionServer", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing AddImageDistributionServer", err))
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
				"Failure when executing AddImageDistributionServer", err1))
			return diags
		}
	}
	item3, err := searchSoftwareImageManagementSwimRetrieveImageDistributionServers(m, vvServer)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AddImageDistributionServer", err,
			"Failure at AddImageDistributionServer, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["server_address"] = vvServer
	d.SetId(joinResourceID(resourceMap))
	return resourceImagesDistributionServerSettingsRead(ctx, d, m)
}

func resourceImagesDistributionServerSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvServer := resourceMap["server_address"]
	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveImageDistributionServers")

		response1, err := searchSoftwareImageManagementSwimRetrieveImageDistributionServers(m, vvServer)

		if err != nil || response1 == nil {

			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items := []dnacentersdkgo.ResponseSoftwareImageManagementSwimRetrieveImageDistributionServersResponse{
			*response1,
		}

		// Review flatten function used
		vItem1 := flattenSoftwareImageManagementSwimRetrieveImageDistributionServersItems(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveImageDistributionServers search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceImagesDistributionServerSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceImagesDistributionServerSettingsRead(ctx, d, m)
}

func resourceImagesDistributionServerSettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := errors.New("Delete not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing ImagesDistributionServerSettings", err, "Delete method is not supported",
		"Failure at ImagesDistributionServerSettingsDelete, unexpected response", ""))
	return diags
}
func expandRequestImagesDistributionServerSettingsAddImageDistributionServer(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSoftwareImageManagementSwimAddImageDistributionServer {
	request := dnacentersdkgo.RequestSoftwareImageManagementSwimAddImageDistributionServer{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".server_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".server_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".server_address")))) {
		request.ServerAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port_number")))) {
		request.PortNumber = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".root_location")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".root_location")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".root_location")))) {
		request.RootLocation = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchSoftwareImageManagementSwimRetrieveImageDistributionServers(m interface{}, vServerAddress string) (*dnacentersdkgo.ResponseSoftwareImageManagementSwimRetrieveImageDistributionServersResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseSoftwareImageManagementSwimRetrieveImageDistributionServersResponse
	var ite *dnacentersdkgo.ResponseSoftwareImageManagementSwimRetrieveImageDistributionServers
	ite, _, err = client.SoftwareImageManagementSwim.RetrieveImageDistributionServers()
	if err != nil || ite == nil {
		return foundItem, err

	}
	items := ite
	itemsCopy := *items.Response
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if item.ServerAddress == vServerAddress {
			var getItem *dnacentersdkgo.ResponseSoftwareImageManagementSwimRetrieveImageDistributionServersResponse
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
