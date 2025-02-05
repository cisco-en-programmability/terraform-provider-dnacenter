package dnacenter

import (
	"context"
	"io"
	"os"

	"errors"

	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceSwimImageFile() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Software Image Management (SWIM).

- Fetches a software image from local file system and uploads to DNA Center. Supported software image files extensions
are bin, img, tar, smu, pie, aes, iso, ova, tar_gz and qcow2
`,

		CreateContext: resourceSwimImageFileCreate,
		ReadContext:   resourceSwimImageFileRead,
		DeleteContext: resourceSwimImageFileDelete,
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

						"task_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"file_name": &schema.Schema{
							Description: `File name.`,
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"file_path": &schema.Schema{
							Description: `File absolute path.`,
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"is_third_party": &schema.Schema{
							Description: `isThirdParty query parameter. Third party Image check
`,
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
						},
						"third_party_application_type": &schema.Schema{
							Description: `thirdPartyApplicationType query parameter. Third Party Application Type
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"third_party_image_family": &schema.Schema{
							Description: `thirdPartyImageFamily query parameter. Third Party image family
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"third_party_vendor": &schema.Schema{
							Description: `thirdPartyVendor query parameter. Third Party Vendor
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
		},
	}
}

func resourceSwimImageFileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	vIsThirdParty, okIsThirdParty := resourceItem["is_third_party"]
	vThirdPartyVendor, okThirdPartyVendor := resourceItem["third_party_vendor"]
	vThirdPartyImageFamily, okThirdPartyImageFamily := resourceItem["third_party_image_family"]
	vThirdPartyApplicationType, okThirdPartyApplicationType := resourceItem["third_party_application_type"]
	vFileName := resourceItem["file_name"]
	vFilePath := resourceItem["file_path"]

	if vFileName.(string) != "" {
		query := dnacentersdkgo.GetSoftwareImageDetailsQueryParams{
			Name: vFileName.(string),
		}
		item, err := searchSoftwareImageManagementSwimGetSoftwareImageDetailsFile(m, query)

		if item != nil && err == nil {
			resourceMap := make(map[string]string)
			resourceMap["file_name"] = vFileName.(string)
			resourceMap["file_path"] = vFilePath.(string)
			d.SetId(joinResourceID(resourceMap))
			return resourceSwimImageFileRead(ctx, d, m)
		}
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ImportLocalSoftwareImage")
		queryParams1 := dnacentersdkgo.ImportLocalSoftwareImageQueryParams{}

		if okIsThirdParty {
			queryParams1.IsThirdParty = vIsThirdParty.(bool)
		}
		if okThirdPartyVendor {
			queryParams1.ThirdPartyVendor = vThirdPartyVendor.(string)
		}
		if okThirdPartyImageFamily {
			queryParams1.ThirdPartyImageFamily = vThirdPartyImageFamily.(string)
		}
		if okThirdPartyApplicationType {
			queryParams1.ThirdPartyApplicationType = vThirdPartyApplicationType.(string)
		}

		isDir, err := IsDirectory(vFilePath.(string))
		if err != nil || isDir {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing File", err,
				"Failure at File, Path is a directory", ""))
			return diags
		}
		f, err := os.Open(vFilePath.(string))
		if err != nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ImportLocalSoftwareImage", err,
				"Failure at ImportLocalSoftwareImage, unexpected response", ""))
			return diags
		}
		defer func() {
			if err = f.Close(); err != nil {
				log.Printf("File close error %s", err.Error())
			}
		}()

		var r io.Reader
		r = f

		response1, restyResp1, err := client.SoftwareImageManagementSwim.ImportLocalSoftwareImage(
			&queryParams1,
			&dnacentersdkgo.ImportLocalSoftwareImageMultipartFields{
				File:     r,
				FileName: vFileName.(string),
			},
		)
		log.Printf("[DEBUG] File name => %s", vFileName.(string))
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			if err != nil {
				log.Printf("[DEBUG] Error response => %s", err.Error())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ImportLocalSoftwareImage", err,
				"Failure at ImportLocalSoftwareImage, unexpected response", ""))
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
				restyResp3, err := client.CustomCall.GetCustomCall(response2.Response.AdditionalStatusURL, nil)
				if err != nil {
					diags = append(diags, diagErrorWithAlt(
						"Failure when executing GetCustomCall", err,
						"Failure at GetCustomCall, unexpected response", ""))
					return diags
				}
				var errorMsg string
				if restyResp3 == nil {
					errorMsg = response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
				} else {
					errorMsg = restyResp3.String()
				}
				err1 := errors.New(errorMsg)
				diags = append(diags, diagError(
					"Failure when executing ImportLocalSoftwareImage", err1))
				return diags
			}
		}
	}
	resourceMap := make(map[string]string)
	resourceMap["file_name"] = vFileName.(string)
	resourceMap["file_path"] = vFilePath.(string)
	d.SetId(joinResourceID(resourceMap))
	return resourceSwimImageFileRead(ctx, d, m)
}
func resourceSwimImageFileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName := resourceMap["file_name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetSoftwareImageDetails")
		queryParams1 := dnacentersdkgo.GetSoftwareImageDetailsQueryParams{
			Name: vName,
		}

		response1, err := searchSoftwareImageManagementSwimGetSoftwareImageDetailsFile(m, queryParams1)

		if err != nil {
			diags = append(diags, diagError(
				"Failure when setting searchSoftwareImageManagementSwimGetSoftwareImageDetailsFile search response",
				err))
			return diags
		}
		if response1 == nil {
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
		items := []dnacentersdkgo.ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponse{
			*response1,
		}
		vItem1 := flattenSoftwareImageManagementSwimGetSoftwareImageDetailsItems(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ImportLocalSoftwareImage response",
				err))
			return diags
		}

	}
	return diags
}

func resourceSwimImageFileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	err := errors.New("Delete not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing SwimImageFileDelete", err, "Delete method is not supported",
		"Failure at SwimImageFileDelete, unexpected response", ""))

	return diags
}

func searchSoftwareImageManagementSwimGetSoftwareImageDetailsFile(m interface{}, queryParams dnacentersdkgo.GetSoftwareImageDetailsQueryParams) (*dnacentersdkgo.ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponse
	var ite *dnacentersdkgo.ResponseSoftwareImageManagementSwimGetSoftwareImageDetails
	ite, _, err = client.SoftwareImageManagementSwim.GetSoftwareImageDetails(&queryParams)
	if err != nil {
		return foundItem, err
	}
	items := ite.Response
	if items == nil {
		return foundItem, err
	}
	itemsCopy := *items
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if item.Name == queryParams.Name {
			var getItem *dnacentersdkgo.ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponse
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
