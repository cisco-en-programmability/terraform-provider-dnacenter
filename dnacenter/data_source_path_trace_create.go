package dnacenter

import (
	"context"

	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourcePathTraceCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Path Trace.

- Initiates a new flow analysis with periodic refresh and stat collection options. Returns a request id and a task id to
get results and follow progress.
`,

		ReadContext: dataSourcePathTraceCreateRead,
		Schema: map[string]*schema.Schema{
			"control_path": &schema.Schema{
				// Type:     schema.TypeBool,
				Type:         schema.TypeString,
				ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
				Optional:     true,
			},
			"dest_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dest_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"inclusions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"flow_analysis_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
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
			"periodic_refresh": &schema.Schema{
				// Type:     schema.TypeBool,
				Type:         schema.TypeString,
				ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
				Optional:     true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourcePathTraceCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: InitiateANewPathtrace")
		request1 := expandRequestPathTraceCreateInitiateANewPathtrace(ctx, "", d)

		response1, restyResp1, err := client.PathTrace.InitiateANewPathtrace(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing InitiateANewPathtrace", err,
				"Failure at InitiateANewPathtrace, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenPathTraceInitiateANewPathtraceItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting InitiateANewPathtrace response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestPathTraceCreateInitiateANewPathtrace(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestPathTraceInitiateANewPathtrace {
	request := dnacentersdkgo.RequestPathTraceInitiateANewPathtrace{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".control_path")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".control_path")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".control_path")))) {
		request.ControlPath = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dest_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dest_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dest_ip")))) {
		request.DestIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dest_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dest_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dest_port")))) {
		request.DestPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".inclusions")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".inclusions")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".inclusions")))) {
		request.Inclusions = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".periodic_refresh")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".periodic_refresh")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".periodic_refresh")))) {
		request.PeriodicRefresh = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".source_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".source_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".source_ip")))) {
		request.SourceIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".source_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".source_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".source_port")))) {
		request.SourcePort = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenPathTraceInitiateANewPathtraceItem(item *dnacentersdkgo.ResponsePathTraceInitiateANewPathtraceResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["flow_analysis_id"] = item.FlowAnalysisID
	respItem["task_id"] = item.TaskID
	respItem["url"] = item.URL
	return []map[string]interface{}{
		respItem,
	}
}
