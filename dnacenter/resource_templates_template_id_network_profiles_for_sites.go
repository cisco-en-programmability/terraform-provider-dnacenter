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

func resourceTemplatesTemplateIDNetworkProfilesForSites() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create and read operations on Configuration Templates.

- Attaches a network profile to a Day-N CLI template by passing the profile ID and template ID.
`,

		CreateContext: resourceTemplatesTemplateIDNetworkProfilesForSitesCreate,
		ReadContext:   resourceTemplatesTemplateIDNetworkProfilesForSitesRead,
		UpdateContext: resourceTemplatesTemplateIDNetworkProfilesForSitesUpdate,
		DeleteContext: resourceTemplatesTemplateIDNetworkProfilesForSitesDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"profile_id": &schema.Schema{
							Description: `The id of the network profile, retrievable from /intent/api/v1/networkProfilesForSites
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"template_id": &schema.Schema{
							Description: `templateId path parameter. The id of the template, retrievable from GET /intent/api/v1/templates
`,
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceTemplatesTemplateIDNetworkProfilesForSitesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestTemplatesTemplateIDNetworkProfilesForSitesAttachNetworkProfileToADayNCliTemplate(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vTemplateID := resourceItem["template_id"]
	vvTemplateID := interfaceToString(vTemplateID)
	vProfileID := resourceItem["profile_id"]
	vvProfileID := interfaceToString(vProfileID)
	item2, err := client.ConfigurationTemplates.RetrieveTheNetworkProfilesAttachedToACLITemplate(vvTemplateID)
	if err != nil || item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["template_id"] = vvTemplateID
		resourceMap["profile_id"] = vvProfileID
		d.SetId(joinResourceID(resourceMap))
		return resourceTemplatesTemplateIDNetworkProfilesForSitesRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.ConfigurationTemplates.AttachNetworkProfileToADayNCliTemplate(vvTemplateID, request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AttachNetworkProfileToADayNCliTemplate", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AttachNetworkProfileToADayNCliTemplate", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing AttachNetworkProfileToADayNCliTemplate", err))
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
				"Failure when executing AttachNetworkProfileToADayNCliTemplate", err1))
			return diags
		}
	}
	item3, err := client.ConfigurationTemplates.RetrieveTheNetworkProfilesAttachedToACLITemplate(vvTemplateID)
	if err != nil || item3 != nil {
		resourceMap := make(map[string]string)
		resourceMap["template_id"] = vvTemplateID
		resourceMap["profile_id"] = vvProfileID
		d.SetId(joinResourceID(resourceMap))
		return resourceTemplatesTemplateIDNetworkProfilesForSitesRead(ctx, d, m)
	}
	item4, err := client.ConfigurationTemplates.RetrieveTheNetworkProfilesAttachedToACLITemplate(vvTemplateID)
	if err != nil || item4 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AttachNetworkProfileToADayNCliTemplate", err,
			"Failure at AttachNetworkProfileToADayNCliTemplate, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["template_id"] = vvTemplateID
	resourceMap["profile_id"] = vvProfileID
	d.SetId(joinResourceID(resourceMap))
	return resourceTemplatesTemplateIDNetworkProfilesForSitesRead(ctx, d, m)
}

func resourceTemplatesTemplateIDNetworkProfilesForSitesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vTemplateID := resourceMap["template_id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveTheNetworkProfilesAttachedToACLITemplate")
		vvTemplateID := vTemplateID

		restyResp1, err := client.ConfigurationTemplates.RetrieveTheNetworkProfilesAttachedToACLITemplate(vvTemplateID)

		if err != nil || restyResp1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*restyResp1))

		if err := d.Set("item", restyResp1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveTheNetworkProfilesAttachedToACLITemplate response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceTemplatesTemplateIDNetworkProfilesForSitesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceTemplatesTemplateIDNetworkProfilesForSitesRead(ctx, d, m)
}

func resourceTemplatesTemplateIDNetworkProfilesForSitesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete TemplatesTemplateIDNetworkProfilesForSites on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestTemplatesTemplateIDNetworkProfilesForSitesAttachNetworkProfileToADayNCliTemplate(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesAttachNetworkProfileToADayNCliTemplate {
	request := dnacentersdkgo.RequestConfigurationTemplatesAttachNetworkProfileToADayNCliTemplate{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_id")))) {
		request.ProfileID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
