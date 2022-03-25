package dnacenter

import (
	"context"
	"errors"
	"io"
	"os"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAuthenticationImportCertificateP12() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Authentication Management.
		- This method is used to upload a PKCS#12 file.
		Upload the file to the **p12FileUpload** form data field
	
`,

		CreateContext: resourceAuthenticationImportCertificateP12Create,
		ReadContext:   resourceAuthenticationImportCertificateP12Read,
		DeleteContext: resourceAuthenticationImportCertificateP12Delete,

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
						"list_of_users": &schema.Schema{
							Description: `listOfUsers query parameter.`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"p12_file_path": &schema.Schema{
							Description: `P12 file absolute path.`,
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"p12_password": &schema.Schema{
							Description: `p12Password query parameter. P12 Passsword
			`,
							Type:      schema.TypeString,
							Optional:  true,
							ForceNew:  true,
							Sensitive: true,
						},
						"pk_password": &schema.Schema{
							Description: `pkPassword query parameter. Private Key Passsword
			`,
							Type:      schema.TypeString,
							Optional:  true,
							ForceNew:  true,
							Sensitive: true,
						},
					},
				},
			},
		},
	}
}

func resourceAuthenticationImportCertificateP12Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	resourceItem := *getResourceItem(d.Get("parameters"))
	vP12FilePath := resourceItem["p12_file_path"]
	vFileName := resourceItem["file_name"]
	vP12Password, okP12Password := d.GetOk("parameters.0.p12_password")
	vPkPassword, okPkPassword := d.GetOk("parameters.0.pk_password")
	vListOfUsers, okListOfUsers := d.GetOk("parameters.0.list_of_users")
	var diags diag.Diagnostics

	log.Printf("[DEBUG] Selected method 1: ImportCertificateP12")

	queryParams1 := dnacentersdkgo.ImportCertificateP12QueryParams{}

	if okP12Password {
		queryParams1.P12Password = vP12Password.(string)
	}
	if okPkPassword {
		queryParams1.PkPassword = vPkPassword.(string)
	}
	if okListOfUsers {
		queryParams1.ListOfUsers = interfaceToSliceString(vListOfUsers)
	}

	isDir, err := IsDirectory(vP12FilePath.(string))
	if err != nil || isDir {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing ImportCertificateP12", err,
			"Failure at ImportCertificateP12, unexpected response", ""))
		return diags
	}

	f, err := os.Open(vP12FilePath.(string))
	if err != nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing ImportCertificateP12", err,
			"Failure at ImportCertificateP12, unexpected response", ""))
		return diags
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Printf("File close error %s", err.Error())
		}
	}()

	var r io.Reader
	r = f

	response1, restyResp1, err := client.AuthenticationManagement.ImportCertificateP12(
		&queryParams1,
		&dnacentersdkgo.ImportCertificateP12MultipartFields{
			P12FileUpload:     r,
			P12FileUploadName: vFileName.(string),
		},
	)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing ImportCertificateP12", err,
			"Failure at ImportCertificateP12, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing ImportCertificateP12", err))
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
			errorMsg := response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing ImportCertificateP12", err1))
			return diags
		}
	}
	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	vItem1 := flattenAuthenticationManagementImportCertificateP12Item(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ImportCertificateP12 response",
			err))
		return diags
	}
	log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
	d.SetId(getUnixTimeString())
	return resourceAuthenticationImportCertificateP12Read(ctx, d, m)
}

func resourceAuthenticationImportCertificateP12Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	return diags
}

func resourceAuthenticationImportCertificateP12Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func flattenAuthenticationManagementImportCertificateP12Item(item *dnacentersdkgo.ResponseAuthenticationManagementImportCertificateP12Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["task_id"] = item.TaskID
	respItem["url"] = item.URL
	return []map[string]interface{}{
		respItem,
	}
}
