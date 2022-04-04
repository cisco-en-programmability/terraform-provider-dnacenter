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

func resourceAuthenticationImportCertificate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Authentication Management.
		- This method is used to upload a certificate.
		Upload the files to the **certFileUpload** and **pkFileUpload** form data fields
	
`,

		CreateContext: resourceAuthenticationImportCertificateCreate,
		ReadContext:   resourceAuthenticationImportCertificateRead,
		DeleteContext: resourceAuthenticationImportCertificateDelete,

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
						"cert_file_path": &schema.Schema{
							Description: `Cert file absolute path.`,
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"pk_file_name": &schema.Schema{
							Description: `File name.`,
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"cert_file_name": &schema.Schema{
							Description: `File name.`,
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
						},
						"list_of_users": &schema.Schema{
							Description: `listOfUsers query parameter.`,
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"pk_file_path": &schema.Schema{
							Description: `Pk file absolute path.`,
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
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

func resourceAuthenticationImportCertificateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	resourceItem := *getResourceItem(d.Get("parameters"))
	vCertFilePath := resourceItem["cert_file_path"]
	vPkFileName := resourceItem["pk_file_name"]
	vPkFilePath := resourceItem["pk_file_path"]
	vCertFileName := resourceItem["pk_file_path"]
	vPkPassword, okPkPassword := d.GetOk("parameters.0.pk_password")
	vListOfUsers, okListOfUsers := d.GetOk("parameters.0.list_of_users")
	var diags diag.Diagnostics

	log.Printf("[DEBUG] Selected method 1: ImportCertificate")
	queryParams1 := dnacentersdkgo.ImportCertificateQueryParams{}

	if okPkPassword {
		queryParams1.PkPassword = vPkPassword.(string)
	}
	if okListOfUsers {
		queryParams1.ListOfUsers = interfaceToSliceString(vListOfUsers)
	}

	isDir, err := IsDirectory(vCertFilePath.(string))
	if err != nil || isDir {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CertFile", err,
			"Failure at CertFile, Path is a directory", ""))
		return diags
	}

	isDir2, err := IsDirectory(vPkFilePath.(string))
	if err != nil || isDir2 {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing PkFile", err,
			"Failure at PkFile, Path is a directory", ""))
		return diags
	}

	first_file, err := os.Open(vPkFilePath.(string))
	if err != nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing PkFile", err,
			"Failure at PkFile, unexpected response", ""))
		return diags
	}
	second_file, err := os.Open(vCertFilePath.(string))
	if err != nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CertFile", err,
			"Failure at CertFile, unexpected response", ""))
		return diags
	}
	defer func() {
		if err = first_file.Close(); err != nil {
			log.Printf("File close error %s", err.Error())
		}
	}()

	defer func() {
		if err = second_file.Close(); err != nil {
			log.Printf("File close error %s", err.Error())
		}
	}()

	var r io.Reader
	var r2 io.Reader
	r = first_file
	r2 = second_file
	response1, restyResp1, err := client.AuthenticationManagement.ImportCertificate(
		&queryParams1,
		&dnacentersdkgo.ImportCertificateMultipartFields{
			PkFileUploadName:   vPkFileName.(string),
			PkFileUpload:       r,
			CertFileUploadName: vCertFileName.(string),
			CertFileUpload:     r2,
		},
	)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing ImportCertificate", err,
			"Failure at ImportCertificate, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing ImportCertificate", err))
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
				"Failure when executing ImportCertificate", err1))
			return diags
		}
	}
	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	vItem1 := flattenAuthenticationManagementImportCertificateItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ImportCertificate response",
			err))
		return diags
	}
	log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
	d.SetId(getUnixTimeString())
	return resourceAuthenticationImportCertificateRead(ctx, d, m)
}

func resourceAuthenticationImportCertificateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	return diags
}

func resourceAuthenticationImportCertificateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func flattenAuthenticationManagementImportCertificateItem(item *dnacentersdkgo.ResponseAuthenticationManagementImportCertificateResponse) []map[string]interface{} {
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
