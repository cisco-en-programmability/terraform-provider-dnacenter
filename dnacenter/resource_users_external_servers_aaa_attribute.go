package dnacenter

import (
	"context"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUsersExternalServersAAAAttribute() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on User and Roles.

- Add or update the custom AAA attribute for external authentication. Note that if you decide not to set the custom AAA
attribute, a default AAA attribute will be used for authentication based on the protocol supported by your server. For
TACACS servers it will be "cisco-av-pair" and for RADIUS servers it will be "Cisco-AVPair".

- Delete the custom AAA attribute that was added. Note that by deleting the AAA attribute, a default AAA attribute will
be used for authentication based on the protocol supported by your server. For TACACS servers it will be "cisco-av-pair"
and for RADIUS servers it will be "Cisco-AVPair".
`,

		CreateContext: resourceUsersExternalServersAAAAttributeCreate,
		ReadContext:   resourceUsersExternalServersAAAAttributeRead,
		UpdateContext: resourceUsersExternalServersAAAAttributeUpdate,
		DeleteContext: resourceUsersExternalServersAAAAttributeDelete,
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

						"aaa_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"attribute_name": &schema.Schema{
										Description: `Value of the custom AAA attribute name
`,
										Type:     schema.TypeString,
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

						"attribute_name": &schema.Schema{
							Description: `name of the custom AAA attribute.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceUsersExternalServersAAAAttributeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	request1 := expandRequestUsersExternalServersAAAAttributeAddAndUpdateAAAAttributeAPI(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	item2, _, err := client.UserandRoles.GetAAAAttributeAPI()
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		d.SetId(joinResourceID(resourceMap))
		return resourceUsersExternalServersAAAAttributeRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.UserandRoles.AddAndUpdateAAAAttributeAPI(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddAndUpdateAAAAttributeAPI", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddAndUpdateAAAAttributeAPI", err))
		return diags
	}
	// TODO REVIEW
	item3, _, err := client.UserandRoles.GetAAAAttributeAPI()
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AddAndUpdateAAAAttributeAPI", err,
			"Failure at AddAndUpdateAAAAttributeAPI, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)

	d.SetId(joinResourceID(resourceMap))
	return resourceUsersExternalServersAAAAttributeRead(ctx, d, m)
}

func resourceUsersExternalServersAAAAttributeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAAAAttributeAPI")

		response1, restyResp1, err := client.UserandRoles.GetAAAAttributeAPI()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenUserandRolesGetAAAAttributeAPIItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAAAAttributeAPI response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceUsersExternalServersAAAAttributeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceUsersExternalServersAAAAttributeRead(ctx, d, m)
}

func resourceUsersExternalServersAAAAttributeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	response1, restyResp1, err := client.UserandRoles.DeleteAAAAttributeAPI()
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteAAAAttributeAPI", err, restyResp1.String(),
				"Failure at DeleteAAAAttributeAPI, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteAAAAttributeAPI", err,
			"Failure at DeleteAAAAttributeAPI, unexpected response", ""))
		return diags
	}

	//TODO REVIEW

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestUsersExternalServersAAAAttributeAddAndUpdateAAAAttributeAPI(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestUserandRolesAddAndUpdateAAAAttributeAPI {
	request := dnacentersdkgo.RequestUserandRolesAddAndUpdateAAAAttributeAPI{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attribute_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attribute_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attribute_name")))) {
		request.AttributeName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
