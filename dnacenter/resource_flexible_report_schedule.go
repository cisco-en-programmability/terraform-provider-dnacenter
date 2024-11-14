package dnacenter

import (
	"context"
	"errors"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFlexibleReportSchedule() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and update operations on Reports.

- Update schedule of flexible report
`,

		CreateContext: resourceFlexibleReportScheduleCreate,
		ReadContext:   resourceFlexibleReportScheduleRead,
		UpdateContext: resourceFlexibleReportScheduleUpdate,
		DeleteContext: resourceFlexibleReportScheduleDelete,
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
				Elem:     schema.TypeString,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"report_id": &schema.Schema{
							Description: `reportId path parameter. Id of the report
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"schedule": &schema.Schema{
							Description: `Schedule information
`,
							Type:     schema.TypeString, //TEST,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceFlexibleReportScheduleCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	// TODO: Add the path params to `item` schema
	//       & return it individually
	resourceMap["id"] = interfaceToString(resourceItem["report_id"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceFlexibleReportScheduleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vReportID := resourceMap["report_id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetFlexibleReportScheduleByReportID")
		vvReportID := vReportID

		response1, restyResp1, err := client.Reports.GetFlexibleReportScheduleByReportID(vvReportID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := restyResp1.String()
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetFlexibleReportScheduleByReportID response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceFlexibleReportScheduleUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvReportID := resourceMap["report_id"]
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvReportID)
		request1 := expandRequestFlexibleReportScheduleUpdateScheduleOfFlexibleReport(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Reports.UpdateScheduleOfFlexibleReport(vvReportID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateScheduleOfFlexibleReport", err, restyResp1.String(),
					"Failure at UpdateScheduleOfFlexibleReport, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateScheduleOfFlexibleReport", err,
				"Failure at UpdateScheduleOfFlexibleReport, unexpected response", ""))
			return diags
		}

		//TODO REVIEW

	}

	return resourceFlexibleReportScheduleRead(ctx, d, m)
}

func resourceFlexibleReportScheduleDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := errors.New("Delete not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing FlexibleReportSchedule", err, "Delete method is not supported",
		"Failure at FlexibleReportScheduleDelete, unexpected response", ""))
	return diags
}
func expandRequestFlexibleReportScheduleUpdateScheduleOfFlexibleReport(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestReportsUpdateScheduleOfFlexibleReport {
	request := dnacentersdkgo.RequestReportsUpdateScheduleOfFlexibleReport{}
	request.Schedule = expandRequestFlexibleReportScheduleUpdateScheduleOfFlexibleReportSchedule(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestFlexibleReportScheduleUpdateScheduleOfFlexibleReportSchedule(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestReportsUpdateScheduleOfFlexibleReportSchedule {
	var request dnacentersdkgo.RequestReportsUpdateScheduleOfFlexibleReportSchedule
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
