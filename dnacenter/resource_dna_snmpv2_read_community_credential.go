package dnacenter

import (
	"context"
	"fmt"
	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSNMPReadCommunityCredential() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceSNMPReadCommunityCredentialCreate,
		ReadContext:   resourceSNMPReadCommunityCredentialRead,
		UpdateContext: resourceSNMPReadCommunityCredentialUpdate,
		DeleteContext: resourceSNMPReadCommunityCredentialDelete,
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
				MaxItems: 1,
				Required: true,

				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comments": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"credential_type": &schema.Schema{
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validateStringHasValueFunc([]string{"GLOBAL", "APP"}),
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"instance_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"read_community": &schema.Schema{
							Type:      schema.TypeString,
							Required:  true,
							Sensitive: true,
						},
					},
				},
			},
		},
	}
}

func constructUpdateSNMPReadCommunityCredentialsRequest(prevID string, credential map[string]interface{}) *dnac.UpdateSNMPReadCommunityRequest {
	credentialRequest := dnac.UpdateSNMPReadCommunityRequest{}
	if v, ok := credential["comments"]; ok {
		credentialRequest.Comments = v.(string)
	}
	if v, ok := credential["credential_type"]; ok {
		credentialRequest.CredentialType = v.(string)
	}
	credentialRequest.ID = prevID
	if v, ok := credential["instance_tenant_id"]; ok {
		credentialRequest.InstanceTenantID = v.(string)
	}
	if v, ok := credential["instance_uuid"]; ok {
		credentialRequest.InstanceUUID = v.(string)
	}
	credentialRequest.ReadCommunity = credential["read_community"].(string)
	credentialRequest.Description = credential["description"].(string)
	return &credentialRequest
}

func resourceSNMPReadCommunityCredentialCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)
	var diags diag.Diagnostics

	credentialSubTypeCompare := "SNMPv2ReadCommunity"

	item := d.Get("item").([]interface{})[0]
	credential := item.(map[string]interface{})

	// Check if element already exists
	v1, ok1 := credential["id"]
	v2, ok2 := credential["instance_uuid"]
	var userGaveID bool
	var prevID string
	if !userGaveID && ok1 {
		prevID = v1.(string)
		userGaveID = true
	}
	if !userGaveID && ok2 {
		prevID = v2.(string)
		userGaveID = true
	}

	if userGaveID {
		searchResponse, _, err := client.Discovery.GetCredentialSubTypeByCredentialID(prevID)
		if err == nil && searchResponse != nil && strings.HasPrefix(searchResponse.Response, credentialSubTypeCompare) {
			// It already exists on DNAC. Update resource id
			d.SetId(prevID)

			// Construct payload from resource schema (item)
			credentialRequest := constructUpdateSNMPReadCommunityCredentialsRequest(prevID, credential)
			// Call function to update tag resource
			response, _, err := client.Discovery.UpdateSNMPReadCommunity(credentialRequest)
			if err != nil {
				return diag.FromErr(err)
			}

			// Call function to check task
			taskID := response.Response.TaskID
			taskResponse, _, err := client.Task.GetTaskByID(taskID)
			if err != nil {
				return diag.FromErr(err)
			}

			// Check if task was completed successfully
			if taskResponse.Response.IsError {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Unable to create updated SNMPReadCommunity credential",
					Detail:   taskResponse.Response.FailureReason,
				})
				return diags
			}

			// Update resource data (on Terraform and DNAC), other data is unknown
			resourceSNMPReadCommunityCredentialRead(ctx, d, m)
			return diags
		}
	}

	// Construct payload from resource schema (item)
	credentialRequest := dnac.CreateSNMPReadCommunityRequest{}
	if v, ok := credential["comments"]; ok {
		credentialRequest.Comments = v.(string)
	}
	if v, ok := credential["credential_type"]; ok {
		credentialRequest.CredentialType = v.(string)
	}
	if v, ok := credential["id"]; ok {
		credentialRequest.ID = v.(string)
	}
	if v, ok := credential["instance_tenant_id"]; ok {
		credentialRequest.InstanceTenantID = v.(string)
	}
	if v, ok := credential["instance_uuid"]; ok {
		credentialRequest.InstanceUUID = v.(string)
	}
	credentialRequest.ReadCommunity = credential["read_community"].(string)
	credentialRequest.Description = credential["description"].(string)

	credentialRequests := []dnac.CreateSNMPReadCommunityRequest{credentialRequest}
	// Call function to create tag resource
	response, _, err := client.Discovery.CreateSNMPReadCommunity(&credentialRequests)
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for execution status to complete
	time.Sleep(10 * time.Second)

	// Call function to check task
	taskID := response.Response.TaskID
	taskResponse, _, err := client.Task.GetTaskByID(taskID)
	if err != nil {
		return diag.FromErr(err)
	}

	// Check if task was completed successfully
	if taskResponse.Response.IsError {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create SNMPReadCommunity credential",
			Detail:   taskResponse.Response.FailureReason,
		})
		return diags
	}

	// Update resource id
	idFound := false

	if !idFound && taskResponse.Response.Data != "" {
		d.SetId(taskResponse.Response.Data)
		idFound = true
	}
	if !idFound && taskResponse.Response.Progress != "" {
		d.SetId(taskResponse.Response.Progress)
		idFound = true
	}
	if idFound {
		resourceSNMPReadCommunityCredentialRead(ctx, d, m)
		return diags
	}
	return diag.FromErr(fmt.Errorf("Unable to retrieve id of object created"))
}

func resourceSNMPReadCommunityCredentialRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	credentialSubTypeCompare := "SNMPv2ReadCommunity"
	credentialSubTypeSearch := "SNMPV2_READ_COMMUNITY"

	// Get resource id (that's also the value of tag.id)
	credentialID := d.Id()

	// Call function to read tag.id
	searchResponse, _, err := client.Discovery.GetCredentialSubTypeByCredentialID(credentialID)
	if err != nil || searchResponse == nil {
		// Resource does not exist
		d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
		return diags
	}
	if !strings.HasPrefix(searchResponse.Response, credentialSubTypeCompare) {
		// it does not have the same credentialSubType
		d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
		return diags
	}

	response, _, err := client.Discovery.GetGlobalCredentials(&dnac.GetGlobalCredentialsQueryParams{CredentialSubType: credentialSubTypeSearch})
	if err != nil {
		return diag.FromErr(err)
	}

	if response != nil {
		var foundCredential *dnac.GetGlobalCredentialsResponseResponse
		for _, fCredential := range response.Response {
			if fCredential.ID == credentialID {
				foundCredential = &fCredential
				break
			}
		}

		credential := flattenCredentialReadItem(foundCredential)
		if err := d.Set("item", credential); err != nil {
			return diag.FromErr(err)
		}

	}

	return diags
}

func resourceSNMPReadCommunityCredentialUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	credentialID := d.Id()

	credentialSubTypeCompare := "SNMPv2ReadCommunity"
	// Call function to read tag.id
	searchResponse, _, err := client.Discovery.GetCredentialSubTypeByCredentialID(credentialID)
	if err != nil || searchResponse == nil {
		// Resource does not exist
		d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
		return diags
	}
	if !strings.HasPrefix(searchResponse.Response, credentialSubTypeCompare) {
		// it does not have the same credentialSubType
		d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
		return diags
	}

	// Check if properties inside resource has changes
	if d.HasChange("item") {
		item := d.Get("item").([]interface{})[0]
		credential := item.(map[string]interface{})

		// Construct payload from resource schema (item)
		credentialRequest := constructUpdateSNMPReadCommunityCredentialsRequest(credentialID, credential)
		// Call function to update tag resource
		response, _, err := client.Discovery.UpdateSNMPReadCommunity(credentialRequest)
		if err != nil {
			return diag.FromErr(err)
		}

		// Call function to check task
		taskID := response.Response.TaskID
		taskResponse, _, err := client.Task.GetTaskByID(taskID)
		if err != nil {
			return diag.FromErr(err)
		}

		// Check if task was completed successfully
		if taskResponse.Response.IsError {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to update SNMPReadCommunity credential",
				Detail:   taskResponse.Response.FailureReason,
			})
			return diags
		}

		// Update resource last_updated
		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	// Update resource on Terraform by using read function
	return resourceSNMPReadCommunityCredentialRead(ctx, d, m)
}

func resourceSNMPReadCommunityCredentialDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	credentialID := d.Id()
	searchResponse, _, err := client.Discovery.GetCredentialSubTypeByCredentialID(credentialID)
	if err != nil || searchResponse == nil {
		return diags
	}

	// Call function to delete resource
	response, _, err := client.Discovery.DeleteGlobalCredentialsByID(credentialID)
	if err != nil {
		return diag.FromErr(err)
	}

	// Call function to check task
	taskID := response.Response.TaskID
	taskResponse, _, err := client.Task.GetTaskByID(taskID)
	if err != nil {
		return diag.FromErr(err)
	}

	if taskResponse != nil && taskResponse.Response.IsError {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to delete SNMPReadCommunity credential",
			Detail:   taskResponse.Response.FailureReason,
		})
		return diags
	}

	return diags
}
