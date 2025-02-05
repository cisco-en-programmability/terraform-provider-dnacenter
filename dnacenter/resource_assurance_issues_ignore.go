package dnacenter

import (
	"context"

	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceAssuranceIssuesIgnore() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Issues.

- Ignores the given list of issues. The response contains the list of issues which were successfully ignored as well as
the issues which are failed to ignore. For detailed information about the usage of the API, please refer to the Open API
specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-IssuesLifecycle-1.0.0-resolved.yaml
`,

		CreateContext: resourceAssuranceIssuesIgnoreCreate,
		ReadContext:   resourceAssuranceIssuesIgnoreRead,
		DeleteContext: resourceAssuranceIssuesIgnoreDelete,
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

						"failure_issue_ids": &schema.Schema{
							Description: `Failure Issue Ids`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"successful_issue_ids": &schema.Schema{
							Description: `Successful Issue Ids`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
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
						"xca_lle_rid": &schema.Schema{
							Description: `X-CALLER-ID header parameter. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"issue_ids": &schema.Schema{
							Description: `Issue Ids`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func resourceAssuranceIssuesIgnoreCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vXCaLLERID := resourceItem["xca_lle_rid"]

	request1 := expandRequestAssuranceIssuesIgnoreIgnoreTheGivenListOfIssues(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	headerParams1 := dnacentersdkgo.IgnoreTheGivenListOfIssuesHeaderParams{}

	headerParams1.XCaLLERID = vXCaLLERID.(string)

	// has_unknown_response: None

	response1, restyResp1, err := client.Issues.IgnoreTheGivenListOfIssues(request1, &headerParams1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing IgnoreTheGivenListOfIssues", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenIssuesIgnoreTheGivenListOfIssuesItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting IgnoreTheGivenListOfIssues response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

	//Analizar verificacion.

}
func resourceAssuranceIssuesIgnoreRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceAssuranceIssuesIgnoreDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestAssuranceIssuesIgnoreIgnoreTheGivenListOfIssues(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestIssuesIgnoreTheGivenListOfIssues {
	request := dnacentersdkgo.RequestIssuesIgnoreTheGivenListOfIssues{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".issue_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".issue_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".issue_ids")))) {
		request.IssueIDs = interfaceToSliceString(v)
	}
	return &request
}

func flattenIssuesIgnoreTheGivenListOfIssuesItem(item *dnacentersdkgo.ResponseIssuesIgnoreTheGivenListOfIssuesResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["successful_issue_ids"] = item.SuccessfulIssueIDs
	respItem["failure_issue_ids"] = item.FailureIssueIDs
	return []map[string]interface{}{
		respItem,
	}
}
