package dnacenter

import (
	"context"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v5/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceConfigurationTemplateVersion() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create and read operations on Configuration Templates.

- API to version the current contents of the template.
`,

		CreateContext: resourceConfigurationTemplateVersionCreate,
		ReadContext:   resourceConfigurationTemplateVersionRead,
		UpdateContext: resourceConfigurationTemplateVersionUpdate,
		DeleteContext: resourceConfigurationTemplateVersionDelete,
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

						"composite": &schema.Schema{
							Description: `Is it composite template
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `Name of template
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"project_id": &schema.Schema{
							Description: `UUID of project
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"project_name": &schema.Schema{
							Description: `Name of project
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"template_id": &schema.Schema{
							Description: `UUID of template
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"versions_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"author": &schema.Schema{
										Description: `Author of version template
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": &schema.Schema{
										Description: `Description of template
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": &schema.Schema{
										Description: `UUID of template
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"version": &schema.Schema{
										Description: `Current version of template
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"version_comment": &schema.Schema{
										Description: `Version comment
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"version_time": &schema.Schema{
										Description: `Template version time
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
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

						"comments": &schema.Schema{
							Description: `Template version comments
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"template_id": &schema.Schema{
							Description: `UUID of template
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

func resourceConfigurationTemplateVersionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestConfigurationTemplateVersionVersionTemplate(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vTemplateID := resourceItem["template_id"]
	vvTemplateID := interfaceToString(vTemplateID)
	// if vvTemplateID != "" {
	// 	getResponse1, _, err := client.ConfigurationTemplates.GetsAllTheVersionsOfAGivenTemplate(vvTemplateID)
	// 	if err == nil && getResponse1 != nil {
	// 		resourceMap := make(map[string]string)
	// 		resourceMap["template_id"] = vvTemplateID
	// 		d.SetId(joinResourceID(resourceMap))
	// 		return resourceConfigurationTemplateVersionRead(ctx, d, m)
	// 	}
	// }
	resp1, restyResp1, err := client.ConfigurationTemplates.VersionTemplate(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing VersionTemplate", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing VersionTemplate", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["template_id"] = vvTemplateID
	d.SetId(joinResourceID(resourceMap))
	return resourceConfigurationTemplateVersionRead(ctx, d, m)
}

func resourceConfigurationTemplateVersionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vTemplateID := resourceMap["template_id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetsAllTheVersionsOfAGivenTemplate")
		vvTemplateID := vTemplateID

		item1, _, err := client.ConfigurationTemplates.GetsAllTheVersionsOfAGivenTemplate(vvTemplateID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		vItem1 := flattenConfigurationTemplatesGetsAllTheVersionsOfAGivenTemplateItems(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetsAllTheVersionsOfAGivenTemplate search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceConfigurationTemplateVersionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceConfigurationTemplateVersionRead(ctx, d, m)
}

func resourceConfigurationTemplateVersionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete ConfigurationTemplateVersion on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestConfigurationTemplateVersionVersionTemplate(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesVersionTemplate {
	request := dnacentersdkgo.RequestConfigurationTemplatesVersionTemplate{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".comments")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".comments")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".comments")))) {
		request.Comments = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_id")))) {
		request.TemplateID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
