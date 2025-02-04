package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAssuranceTasks() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Task.

- returns all existing tasks in a paginated list
default sorting of list is *startTime*, *asc*
valid field to sort by are [*startTime*,*endTime*,*updateTime*,*status*] For detailed information about the usage of the
API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-
api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceTasks-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceAssuranceTasksRead,
		Schema: map[string]*schema.Schema{
			"limit": &schema.Schema{
				Description: `limit query parameter. Maximum number of records to return
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"order": &schema.Schema{
				Description: `order query parameter. The sort order of the field ascending or descending.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. A field within the response to sort by.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Description: `status query parameter. used to get a subset of tasks by their status
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"xca_lle_rid": &schema.Schema{
				Description: `X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"data": &schema.Schema{
							Description: `Data`,
							Type:        schema.TypeString, //TEST,
							Computed:    true,
						},

						"end_time": &schema.Schema{
							Description: `End Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"error_code": &schema.Schema{
							Description: `Error Code`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"failure_reason": &schema.Schema{
							Description: `Failure Reason`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"progress": &schema.Schema{
							Description: `Progress`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"request_type": &schema.Schema{
							Description: `Request Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"result_url": &schema.Schema{
							Description: `Result Url`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"start_time": &schema.Schema{
							Description: `Start Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"status": &schema.Schema{
							Description: `Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"update_time": &schema.Schema{
							Description: `Update Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceAssuranceTasksRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")
	vStatus, okStatus := d.GetOk("status")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveAListOfAssuranceTasks")

		headerParams1 := dnacentersdkgo.RetrieveAListOfAssuranceTasksHeaderParams{}
		queryParams1 := dnacentersdkgo.RetrieveAListOfAssuranceTasksQueryParams{}

		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}
		if okStatus {
			queryParams1.Status = vStatus.(string)
		}
		headerParams1.XCaLLERID = vXCaLLERID.(string)

		response1, restyResp1, err := client.Task.RetrieveAListOfAssuranceTasks(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveAListOfAssuranceTasks", err,
				"Failure at RetrieveAListOfAssuranceTasks, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenTaskRetrieveAListOfAssuranceTasksItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveAListOfAssuranceTasks response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenTaskRetrieveAListOfAssuranceTasksItems(items *[]dnacentersdkgo.ResponseTaskRetrieveAListOfAssuranceTasksResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["status"] = item.Status
		respItem["start_time"] = item.StartTime
		respItem["end_time"] = item.EndTime
		respItem["update_time"] = item.UpdateTime
		respItem["progress"] = item.Progress
		respItem["failure_reason"] = item.FailureReason
		respItem["error_code"] = item.ErrorCode
		respItem["request_type"] = item.RequestType
		respItem["data"] = flattenTaskRetrieveAListOfAssuranceTasksItemsData(item.Data)
		respItem["result_url"] = item.ResultURL
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenTaskRetrieveAListOfAssuranceTasksItemsData(item *dnacentersdkgo.ResponseTaskRetrieveAListOfAssuranceTasksResponseData) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
