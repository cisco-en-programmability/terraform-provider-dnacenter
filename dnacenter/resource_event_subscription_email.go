package dnacenter

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEventSubscriptionEmail() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and update operations on Event Management.

- Create Email Subscription Endpoint for list of registered events.

- Update Email Subscription Endpoint for list of registered events
`,

		CreateContext: resourceEventSubscriptionEmailCreate,
		ReadContext:   resourceEventSubscriptionEmailRead,
		UpdateContext: resourceEventSubscriptionEmailUpdate,
		DeleteContext: resourceEventSubscriptionEmailDelete,
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

						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"filter": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"categories": &schema.Schema{
										Description: `Categories`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"domains_subdomains": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"domain": &schema.Schema{
													Description: `Domain`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"sub_domains": &schema.Schema{
													Description: `Sub Domains`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"event_ids": &schema.Schema{
										Description: `Event Ids`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"others": &schema.Schema{
										Description: `Others`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"severities": &schema.Schema{
										Description: `Severities`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"site_ids": &schema.Schema{
										Description: `Site Ids`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"sources": &schema.Schema{
										Description: `Sources`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"types": &schema.Schema{
										Description: `Types`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"is_private": &schema.Schema{
							Description: `Is Private`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"subscription_endpoints": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"connector_type": &schema.Schema{
										Description: `Connector Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"instance_id": &schema.Schema{
										Description: `Instance Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"subscription_details": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"connector_type": &schema.Schema{
													Description: `Connector Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"description": &schema.Schema{
													Description: `Description`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"from_email_address": &schema.Schema{
													Description: `From Email Address`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"instance_id": &schema.Schema{
													Description: `Instance Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"subject": &schema.Schema{
													Description: `Subject`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"to_email_addresses": &schema.Schema{
													Description: `To Email Addresses`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
						"subscription_id": &schema.Schema{
							Description: `Subscription Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"tenant_id": &schema.Schema{
							Description: `Tenant Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"version": &schema.Schema{
							Description: `Version`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestEventManagementCreateEmailEventSubscription`,
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"payload": &schema.Schema{
							Description: `Array of RequestEventManagementCreateEmailEventSubscription`,
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"description": &schema.Schema{
										Description: `Description
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"filter": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"categories": &schema.Schema{
													Description: `Categories`,
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"domains_subdomains": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"domain": &schema.Schema{
																Description: `Domain`,
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"sub_domains": &schema.Schema{
																Description: `Sub Domains`,
																Type:        schema.TypeList,
																Optional:    true,
																Computed:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},
												"event_ids": &schema.Schema{
													Description: `Event Ids`,
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"severities": &schema.Schema{
													Description: `Severities`,
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeInt,
													},
												},
												"site_ids": &schema.Schema{
													Description: `Site Ids`,
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"sources": &schema.Schema{
													Description: `Sources`,
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"types": &schema.Schema{
													Description: `Types`,
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"name": &schema.Schema{
										Description: `Name
`,
										Type:     schema.TypeString,
										Optional: true,
										Default:  "",
									},
									"subscription_endpoints": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"instance_id": &schema.Schema{
													Description: `(From Get Email Subscription Details --> pick InstanceId)
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"subscription_details": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"connector_type": &schema.Schema{
																Description: `Connector Type (Must be EMAIL)
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"description": &schema.Schema{
																Description: `Description`,
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"from_email_address": &schema.Schema{
																Description: `Senders Email Address
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"name": &schema.Schema{
																Description: `Name`,
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"subject": &schema.Schema{
																Description: `Email Subject
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"to_email_addresses": &schema.Schema{
																Description: `Recipient's Email Addresses (Comma separated)
`,
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},
											},
										},
									},
									"subscription_id": &schema.Schema{
										Description: `Subscription Id (Unique UUID)
`,
										Type:     schema.TypeString,
										Optional: true,
										Default:  "",
									},
									"version": &schema.Schema{
										Description: `Version
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								}}}},
				},
			},
		},
	}
}

func resourceEventSubscriptionEmailCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters.0.payload"))
	request1 := expandRequestEventSubscriptionEmailCreateEmailEventSubscription(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vName := resourceItem["name"]
	vvName := interfaceToString(vName)
	vSubscriptionID := resourceItem["subscription_id"]
	vvSubscriptionID := interfaceToString(vSubscriptionID)
	queryParamImport := dnacentersdkgo.GetEmailEventSubscriptionsQueryParams{}
	queryParamImport.Name = vvName
	item2, err := searchEventManagementGetEmailEventSubscriptions(m, queryParamImport, vvSubscriptionID)
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["name"] = item2.Name
		resourceMap["subscription_id"] = item2.SubscriptionID
		d.SetId(joinResourceID(resourceMap))
		return resourceEventSubscriptionEmailRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.EventManagement.CreateEmailEventSubscription(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateEmailEventSubscription", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateEmailEventSubscription", err))
		return diags
	}
	// TODO REVIEW
	queryParamValidate := dnacentersdkgo.GetEmailEventSubscriptionsQueryParams{}
	queryParamValidate.Name = vvName
	item3, err := searchEventManagementGetEmailEventSubscriptions(m, queryParamValidate, vvSubscriptionID)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CreateEmailEventSubscription", err,
			"Failure at CreateEmailEventSubscription, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["name"] = vvName
	resourceMap["subscription_id"] = item3.SubscriptionID
	d.SetId(joinResourceID(resourceMap))
	return resourceEventSubscriptionEmailRead(ctx, d, m)
}

func resourceEventSubscriptionEmailRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetEmailEventSubscriptions")
		queryParams1 := dnacentersdkgo.GetEmailEventSubscriptionsQueryParams{}

		item1, err := searchEventManagementGetEmailEventSubscriptions(m, queryParams1, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		items := []dnacentersdkgo.ResponseItemEventManagementGetEmailEventSubscriptions{
			*item1,
		}
		// Review flatten function used
		vItem1 := flattenEventManagementGetEmailEventSubscriptionsItems2(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetEmailEventSubscriptions search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceEventSubscriptionEmailUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]
	if d.HasChange("parameters") {
		request1 := expandRequestEventSubscriptionEmailUpdateEmailEventSubscription(ctx, "parameters", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		// Add SubscriptionID to update
		if request1 != nil && len(*request1) > 0 {
			req := *request1
			req[0].SubscriptionID = vvID
			request1 = &req
		}

		response1, restyResp1, err := client.EventManagement.UpdateEmailEventSubscription(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateEmailEventSubscription", err, restyResp1.String(),
					"Failure at UpdateEmailEventSubscription, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateEmailEventSubscription", err,
				"Failure at UpdateEmailEventSubscription, unexpected response", ""))
			return diags
		}

		//TODO REVIEW

	}

	return resourceEventSubscriptionEmailRead(ctx, d, m)
}

func resourceEventSubscriptionEmailDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete EventSubscriptionEmail on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestEventSubscriptionEmailCreateEmailEventSubscription(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateEmailEventSubscription {
	request := dnacentersdkgo.RequestEventManagementCreateEmailEventSubscription{}
	if v := expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemEventManagementCreateEmailEventSubscription {
	request := []dnacentersdkgo.RequestItemEventManagementCreateEmailEventSubscription{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementCreateEmailEventSubscription {
	request := dnacentersdkgo.RequestItemEventManagementCreateEmailEventSubscription{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_id")))) {
		request.SubscriptionID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_endpoints")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_endpoints")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_endpoints")))) {
		request.SubscriptionEndpoints = expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemEventManagementCreateEmailEventSubscriptionSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestItemEventManagementCreateEmailEventSubscriptionSubscriptionEndpoints{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementCreateEmailEventSubscriptionSubscriptionEndpoints {
	request := dnacentersdkgo.RequestItemEventManagementCreateEmailEventSubscriptionSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementCreateEmailEventSubscriptionSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestItemEventManagementCreateEmailEventSubscriptionSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".from_email_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".from_email_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".from_email_address")))) {
		request.FromEmailAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".to_email_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".to_email_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".to_email_addresses")))) {
		request.ToEmailAddresses = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subject")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subject")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subject")))) {
		request.Subject = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementCreateEmailEventSubscriptionFilter {
	request := dnacentersdkgo.RequestItemEventManagementCreateEmailEventSubscriptionFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".domains_subdomains")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".domains_subdomains")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".domains_subdomains")))) {
		request.DomainsSubdomains = expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemFilterDomainsSubdomainsArray(ctx, key+".domains_subdomains", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".types")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".types")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".types")))) {
		request.Types = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".categories")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".categories")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".categories")))) {
		request.Categories = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".severities")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".severities")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".severities")))) {
		request.Severities = interfaceToSliceInt(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sources")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sources")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sources")))) {
		request.Sources = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_ids")))) {
		request.SiteIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemFilterDomainsSubdomainsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemEventManagementCreateEmailEventSubscriptionFilterDomainsSubdomains {
	request := []dnacentersdkgo.RequestItemEventManagementCreateEmailEventSubscriptionFilterDomainsSubdomains{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemFilterDomainsSubdomains(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionEmailCreateEmailEventSubscriptionItemFilterDomainsSubdomains(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementCreateEmailEventSubscriptionFilterDomainsSubdomains {
	request := dnacentersdkgo.RequestItemEventManagementCreateEmailEventSubscriptionFilterDomainsSubdomains{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".domain")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".domain")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".domain")))) {
		request.Domain = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sub_domains")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sub_domains")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sub_domains")))) {
		request.SubDomains = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscription(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscription {
	request := dnacentersdkgo.RequestEventManagementUpdateEmailEventSubscription{}
	if v := expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemEventManagementUpdateEmailEventSubscription {
	request := []dnacentersdkgo.RequestItemEventManagementUpdateEmailEventSubscription{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementUpdateEmailEventSubscription {
	request := dnacentersdkgo.RequestItemEventManagementUpdateEmailEventSubscription{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_id")))) {
		request.SubscriptionID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_endpoints")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_endpoints")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_endpoints")))) {
		request.SubscriptionEndpoints = expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemEventManagementUpdateEmailEventSubscriptionSubscriptionEndpoints {
	request := []dnacentersdkgo.RequestItemEventManagementUpdateEmailEventSubscriptionSubscriptionEndpoints{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementUpdateEmailEventSubscriptionSubscriptionEndpoints {
	request := dnacentersdkgo.RequestItemEventManagementUpdateEmailEventSubscriptionSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementUpdateEmailEventSubscriptionSubscriptionEndpointsSubscriptionDetails {
	request := dnacentersdkgo.RequestItemEventManagementUpdateEmailEventSubscriptionSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".from_email_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".from_email_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".from_email_address")))) {
		request.FromEmailAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".to_email_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".to_email_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".to_email_addresses")))) {
		request.ToEmailAddresses = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subject")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subject")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subject")))) {
		request.Subject = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemFilter(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementUpdateEmailEventSubscriptionFilter {
	request := dnacentersdkgo.RequestItemEventManagementUpdateEmailEventSubscriptionFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".domains_subdomains")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".domains_subdomains")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".domains_subdomains")))) {
		request.DomainsSubdomains = expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemFilterDomainsSubdomainsArray(ctx, key+".domains_subdomains", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".types")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".types")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".types")))) {
		request.Types = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".categories")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".categories")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".categories")))) {
		request.Categories = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".severities")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".severities")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".severities")))) {
		request.Severities = interfaceToSliceInt(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sources")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sources")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sources")))) {
		request.Sources = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_ids")))) {
		request.SiteIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemFilterDomainsSubdomainsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemEventManagementUpdateEmailEventSubscriptionFilterDomainsSubdomains {
	request := []dnacentersdkgo.RequestItemEventManagementUpdateEmailEventSubscriptionFilterDomainsSubdomains{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemFilterDomainsSubdomains(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionEmailUpdateEmailEventSubscriptionItemFilterDomainsSubdomains(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemEventManagementUpdateEmailEventSubscriptionFilterDomainsSubdomains {
	request := dnacentersdkgo.RequestItemEventManagementUpdateEmailEventSubscriptionFilterDomainsSubdomains{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".domain")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".domain")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".domain")))) {
		request.Domain = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sub_domains")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sub_domains")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sub_domains")))) {
		request.SubDomains = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchEventManagementGetEmailEventSubscriptions(m interface{}, queryParams dnacentersdkgo.GetEmailEventSubscriptionsQueryParams, vID string) (*dnacentersdkgo.ResponseItemEventManagementGetEmailEventSubscriptions, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseItemEventManagementGetEmailEventSubscriptions
	var ite *dnacentersdkgo.ResponseEventManagementGetEmailEventSubscriptions
	if vID != "" {
		queryParams.Offset = 1
		nResponse, _, err := client.EventManagement.GetEmailEventSubscriptions(nil)
		maxPageSize := len(*nResponse)
		for len(*nResponse) > 0 {
			time.Sleep(15 * time.Second)
			for _, item := range *nResponse {
				if vID == item.SubscriptionID {
					foundItem = &item
					return foundItem, err
				}
			}
			queryParams.Limit = float64(maxPageSize)
			queryParams.Offset += float64(maxPageSize)
			nResponse, _, err = client.EventManagement.GetEmailEventSubscriptions(&queryParams)
		}
		return nil, err
	} else if queryParams.Name != "" {
		ite, _, err = client.EventManagement.GetEmailEventSubscriptions(&queryParams)
		if err != nil || ite == nil {
			return foundItem, err
		}
		itemsCopy := *ite
		if itemsCopy == nil {
			return foundItem, err
		}
		for _, item := range itemsCopy {
			if item.Name == queryParams.Name {
				foundItem = &item
				return foundItem, err
			}
		}
		return foundItem, err
	}
	return foundItem, err
}
