package dnacenter

import (
	dnac "dnacenter-go-sdk/sdk"
	"strconv"
	"strings"
)

///// start application

func flattenApplicationsReadItems(appItems *[]dnac.GetApplicationsResponseResponse) []interface{} {
	if appItems != nil {
		ais := make([]interface{}, len(*appItems), len(*appItems))

		for i, appItem := range *appItems {
			ai := make(map[string]interface{})

			ai["application_set_id"] = appItem.ApplicationSet.IDRef
			ai["application_id"] = appItem.ID
			nais := make([]interface{}, len(appItem.NetworkApplications), len(appItem.NetworkApplications))
			niis := make([]interface{}, len(appItem.NetworkIDentity), len(appItem.NetworkIDentity))

			for j, networkAppItem := range appItem.NetworkApplications {
				nai := make(map[string]interface{})
				nai["app_protocol"] = networkAppItem.AppProtocol
				nai["application_subtype"] = networkAppItem.ApplicationSubType
				nai["application_type"] = networkAppItem.ApplicationType
				nai["category_id"] = networkAppItem.CategoryID
				nai["display_name"] = networkAppItem.DisplayName
				nai["dscp"] = networkAppItem.Dscp
				nai["engine_id"] = networkAppItem.EngineID
				nai["help_string"] = networkAppItem.HelpString
				nai["id"] = networkAppItem.ID
				nai["ignore_conflict"] = networkAppItem.IgnoreConflict
				nai["long_description"] = networkAppItem.LongDescription
				nai["name"] = networkAppItem.Name
				nai["popularity"] = networkAppItem.Popularity
				nai["rank"] = networkAppItem.Rank
				nai["server_name"] = networkAppItem.ServerName
				nai["traffic_class"] = networkAppItem.TrafficClass
				nai["url"] = networkAppItem.URL

				nais[j] = nai
			}

			for k, networkIdentityItem := range appItem.NetworkIDentity {
				nii := make(map[string]interface{})

				nii["display_name"] = networkIdentityItem.DisplayName
				nii["id"] = networkIdentityItem.ID
				nii["lower_port"] = networkIdentityItem.LowerPort
				nii["ports"] = networkIdentityItem.Ports
				nii["protocol"] = networkIdentityItem.Protocol
				nii["upper_port"] = networkIdentityItem.UpperPort

				niis[k] = nii
			}

			ai["application_network_applications"] = nais
			ai["application_network_identity"] = niis

			ais[i] = ai
		}
		return ais
	}

	return make([]interface{}, 0)
}

///// end application
///// start tag

func flattenTagQueryReadItems(tags *dnac.GetTagResponse) []interface{} {
	if tags != nil {
		ois := make([]interface{}, len(tags.Response), len(tags.Response))

		for i, tagItem := range tags.Response {
			oi := make(map[string]interface{})
			oi["system_tag"] = tagItem.SystemTag
			oi["description"] = tagItem.Description
			oi["name"] = tagItem.Name
			oi["instance_tenant_id"] = tagItem.InstanceTenantID

			dynamicRulesLen := len(tagItem.DynamicRules)
			if dynamicRulesLen > 0 {
				dis := make([]interface{}, dynamicRulesLen, dynamicRulesLen)

				for i, dynamicRule := range tagItem.DynamicRules {
					di := make(map[string]interface{})
					di["member_type"] = dynamicRule.MemberType

					// Check if empty
					if dynamicRule.Rules.Name != "" {
						dR := make(map[string]interface{})
						if len(dynamicRule.Rules.Items) > 0 {
							dR["items"] = dynamicRule.Rules.Items
						}
						if len(dynamicRule.Rules.Values) > 0 {
							dR["values"] = dynamicRule.Rules.Values
						}
						dR["name"] = dynamicRule.Rules.Name
						dR["operation"] = dynamicRule.Rules.Operation
						dR["value"] = dynamicRule.Rules.Value

						di["rules"] = []interface{}{dR}
					}
					dis[i] = di
				}
				oi["dynamic_rules"] = dis
			}
			ois[i] = oi
		}

		return ois
	}

	return make([]interface{}, 0)
}

func flattenTagReadItem(tag *dnac.GetTagByIDResponse) []interface{} {
	if tag != nil {
		ois := make([]interface{}, 1, 1)

		oi := make(map[string]interface{})
		oi["system_tag"] = tag.Response.SystemTag
		oi["description"] = tag.Response.Description
		oi["name"] = tag.Response.Name
		oi["instance_tenant_id"] = tag.Response.InstanceTenantID

		dynamicRulesLen := len(tag.Response.DynamicRules)
		if dynamicRulesLen > 0 {
			dis := make([]interface{}, dynamicRulesLen, dynamicRulesLen)

			for i, dynamicRule := range tag.Response.DynamicRules {
				di := make(map[string]interface{})
				di["member_type"] = dynamicRule.MemberType

				// Check if empty
				if dynamicRule.Rules.Name != "" {
					dR := make(map[string]interface{})
					if len(dynamicRule.Rules.Items) > 0 {
						dR["items"] = dynamicRule.Rules.Items
					}
					if len(dynamicRule.Rules.Values) > 0 {
						dR["values"] = dynamicRule.Rules.Values
					}
					dR["name"] = dynamicRule.Rules.Name
					dR["operation"] = dynamicRule.Rules.Operation
					dR["value"] = dynamicRule.Rules.Value

					di["rules"] = []interface{}{dR}
				}
				dis[i] = di
			}
			oi["dynamic_rules"] = dis
		}
		ois[0] = oi

		return ois
	}

	return make([]interface{}, 0)
}

///// end tag
///// start site

func flattenSiteReadItem(site *dnac.GetSiteResponse) []interface{} {
	if site != nil {
		ois := make([]interface{}, len(site.Response), len(site.Response))

		if len(site.Response) > 0 {
			for i, siteItem := range site.Response {
				oi := make(map[string]interface{})
				oi["id"] = siteItem.ID
				oi["name"] = siteItem.Name
				siteHierarchy := siteItem.SiteNameHierarchy
				if oi["name"] == siteHierarchy {
					oi["parent_name"] = oi["name"] // Possibly Global
				} else {
					oi["parent_name"] = strings.TrimSuffix(siteHierarchy, "/"+siteItem.Name)
				}

				if len(siteItem.AdditionalInfo) > 0 {
					for _, info := range siteItem.AdditionalInfo {
						if info.Namespace == "Location" {
							if typeS := info.Attributes.Type; typeS != "" {
								if typeS == "area" {
									oi["type"] = typeS
								} else if typeS == "building" {
									oi["type"] = typeS
									if v := info.Attributes.Address; v != "" {
										oi["address"] = v
									}
									if v := info.Attributes.Latitude; v != "" {
										if latitude, err := strconv.ParseFloat(v, 64); err == nil {
											oi["latitude"] = latitude
										}
									}
									if v := info.Attributes.Longitude; v != "" {
										if longitude, err := strconv.ParseFloat(v, 64); err == nil {
											oi["longitude"] = longitude
										}
									}
								} else if typeS == "floor" {
									oi["type"] = typeS
								}
							}
						}
						if info.Namespace == "mapsSummary" {
							if v := info.Attributes.RfModel; v != "" {
								oi["rf_model"] = v
							}
						}
						if info.Namespace == "mapGeometry" {
							if v := info.Attributes.Width; v != "" {
								if width, err := strconv.ParseFloat(v, 64); err == nil {
									oi["width"] = width
								}
							}
							if v := info.Attributes.Length; v != "" {
								if length, err := strconv.ParseFloat(v, 64); err == nil {
									oi["length"] = length
								}
							}
							if v := info.Attributes.Height; v != "" {
								if height, err := strconv.ParseFloat(v, 64); err == nil {
									oi["height"] = height
								}
							}
						}
					}
				}

				ois[i] = oi
			}
			return ois
		}
	}

	return make([]interface{}, 0)
}

func flattenSiteHealthReadItem(siteHealth *dnac.GetSiteHealthResponse) []interface{} {
	if siteHealth != nil {
		ois := make([]interface{}, len(siteHealth.Response), len(siteHealth.Response))
		for i, siteResponse := range siteHealth.Response {
			oi := make(map[string]interface{})
			oi["access_good_count"] = siteResponse.AccessGoodCount
			oi["access_total_count"] = siteResponse.AccessTotalCount
			oi["application_bytes_total_count"] = siteResponse.ApplicationBytesTotalCount
			oi["application_good_count"] = siteResponse.ApplicationGoodCount
			oi["application_health"] = siteResponse.ApplicationHealth
			oi["application_total_count"] = siteResponse.ApplicationTotalCount
			oi["client_health_wired"] = siteResponse.ClientHealthWired
			oi["client_health_wireless"] = siteResponse.ClientHealthWireless
			oi["core_good_count"] = siteResponse.CoreGoodCount
			oi["core_total_count"] = siteResponse.CoreTotalCount
			oi["distribution_good_count"] = siteResponse.DistributionGoodCount
			oi["distribution_total_count"] = siteResponse.DistributionTotalCount
			oi["dnac_info"] = siteResponse.DnacInfo
			oi["healthy_clients_percentage"] = siteResponse.HealthyClientsPercentage
			oi["healthy_network_device_percentage"] = siteResponse.HealthyNetworkDevicePercentage
			oi["latitude"] = siteResponse.Latitude
			oi["longitude"] = siteResponse.Longitude
			oi["network_health_access"] = siteResponse.NetworkHealthAccess
			oi["network_health_average"] = siteResponse.NetworkHealthAverage
			oi["network_health_core"] = siteResponse.NetworkHealthCore
			oi["network_health_distribution"] = siteResponse.NetworkHealthDistribution
			oi["network_health_others"] = siteResponse.NetworkHealthOthers
			oi["network_health_router"] = siteResponse.NetworkHealthRouter
			oi["network_health_wireless"] = siteResponse.NetworkHealthWireless
			oi["number_of_clients"] = siteResponse.NumberOfClients
			oi["number_of_network_device"] = siteResponse.NumberOfNetworkDevice
			oi["number_of_wired_clients"] = siteResponse.NumberOfWiredClients
			oi["number_of_wireless_clients"] = siteResponse.NumberOfWirelessClients
			oi["overall_good_devices"] = siteResponse.OverallGoodDevices
			oi["parent_site_id"] = siteResponse.ParentSiteID
			oi["parent_site_name"] = siteResponse.ParentSiteName
			oi["router_good_count"] = siteResponse.RouterGoodCount
			oi["router_total_count"] = siteResponse.RouterTotalCount
			oi["site_id"] = siteResponse.SiteID
			oi["site_name"] = siteResponse.SiteName
			oi["site_type"] = siteResponse.SiteType
			oi["total_number_of_active_wireless_clients"] = siteResponse.TotalNumberOfActiveWirelessClients
			oi["total_number_of_connected_wired_clients"] = siteResponse.TotalNumberOfConnectedWiredClients
			oi["wired_good_clients"] = siteResponse.WiredGoodClients
			oi["wireless_device_good_count"] = siteResponse.WirelessDeviceGoodCount
			oi["wireless_device_total_count"] = siteResponse.WirelessDeviceTotalCount
			oi["wireless_good_clients"] = siteResponse.WirelessGoodClients

			stats := make([]interface{}, 1, 1)
			stat := make(map[string]interface{})
			stat["app_total_count"] = siteResponse.ApplicationHealthStats.AppTotalCount
			stat["business_irrelevant_app_fair"] = siteResponse.ApplicationHealthStats.BusinessIrrelevantAppCount.Fair
			stat["business_irrelevant_app_good"] = siteResponse.ApplicationHealthStats.BusinessIrrelevantAppCount.Good
			stat["business_irrelevant_app_poor"] = siteResponse.ApplicationHealthStats.BusinessIrrelevantAppCount.Poor
			stat["business_relevant_app_fair"] = siteResponse.ApplicationHealthStats.BusinessRelevantAppCount.Fair
			stat["business_relevant_app_good"] = siteResponse.ApplicationHealthStats.BusinessRelevantAppCount.Good
			stat["business_relevant_app_poor"] = siteResponse.ApplicationHealthStats.BusinessRelevantAppCount.Poor
			stat["default_health_app_fair"] = siteResponse.ApplicationHealthStats.DefaultHealthAppCount.Fair
			stat["default_health_app_good"] = siteResponse.ApplicationHealthStats.DefaultHealthAppCount.Good
			stat["default_health_app_poor"] = siteResponse.ApplicationHealthStats.DefaultHealthAppCount.Poor
			stats[0] = stat

			oi["application_health_stats"] = stats
			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

///// end site
///// start template project

func flattenTemplateProjectsReadItems(templateProjects *[]dnac.GetProjectsResponse) []interface{} {
	if templateProjects != nil {
		projects := *templateProjects
		ois := make([]interface{}, len(projects), len(projects))
		for i, project := range projects {
			ois[i] = flattenTemplateProjectReadItem(&project)[0]
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenTemplateProjectReadItem(templateProject *dnac.GetProjectsResponse) []interface{} {
	if templateProject != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})
		oi["id"] = templateProject.ID
		oi["name"] = templateProject.Name
		oi["is_deletable"] = templateProject.IsDeletable

		if len(templateProject.Templates) > 0 {
			tis := make([]interface{}, len(templateProject.Templates), len(templateProject.Templates))
			for j, template := range templateProject.Templates {
				ti := make(map[string]interface{})

				ti["composite"] = template.Composite
				ti["id"] = template.ID
				ti["name"] = template.Name
				ti["language"] = template.Language
				ti["custom_params_order"] = template.CustomParamsOrder
				ti["last_update_time"] = template.LastUpdateTime
				ti["latest_version_time"] = template.LatestVersionTime
				ti["project_associated"] = template.ProjectAssociated
				ti["document_database"] = template.DocumentDatabase

				tis[j] = ti
			}
			oi["templates"] = tis
		}
		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenTemplateReadItem(template *dnac.GetTemplateDetailsResponse) []interface{} {
	if template != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["author"] = template.Author
		oi["composite"] = template.Composite

		if len(template.ContainingTemplates) > 0 {
			ctis := make([]interface{}, len(template.ContainingTemplates), len(template.ContainingTemplates))
			for j, containingTemplate := range template.ContainingTemplates {
				cti := make(map[string]interface{})
				cti["composite"] = containingTemplate.Composite
				cti["id"] = containingTemplate.ID
				cti["name"] = containingTemplate.Name
				cti["version"] = containingTemplate.Version
				ctis[j] = cti
			}
			oi["containing_templates"] = ctis
		}

		oi["create_time"] = template.CreateTime
		oi["description"] = template.Description

		if len(template.DeviceTypes) > 0 {
			ctis := make([]interface{}, len(template.DeviceTypes), len(template.DeviceTypes))
			for j, deviceType := range template.DeviceTypes {
				cti := make(map[string]interface{})
				cti["product_family"] = deviceType.ProductFamily
				cti["product_series"] = deviceType.ProductSeries
				cti["product_type"] = deviceType.ProductType
				ctis[j] = cti
			}
			oi["device_types"] = ctis
		}

		oi["failure_policy"] = template.FailurePolicy
		oi["id"] = template.ID
		oi["last_update_time"] = template.LastUpdateTime
		oi["name"] = template.Name
		oi["parent_template_id"] = template.ParentTemplateID
		oi["project_id"] = template.ProjectID
		oi["project_name"] = template.ProjectName
		oi["rollback_template_content"] = template.RollbackTemplateContent

		if len(template.RollbackTemplateParams) > 0 {
			ctis := make([]interface{}, len(template.RollbackTemplateParams), len(template.RollbackTemplateParams))
			for j, templateParam := range template.RollbackTemplateParams {
				cti := make(map[string]interface{})
				cti["binding"] = templateParam.Binding
				cti["data_type"] = templateParam.DataType
				cti["default_value"] = templateParam.DefaultValue
				cti["description"] = templateParam.Description
				cti["display_name"] = templateParam.DisplayName
				cti["group"] = templateParam.Group
				cti["id"] = templateParam.ID
				cti["instruction_text"] = templateParam.InstructionText
				cti["key"] = templateParam.Key
				cti["not_param"] = templateParam.NotParam
				cti["order"] = templateParam.Order
				cti["param_array"] = templateParam.ParamArray
				cti["parameter_name"] = templateParam.ParameterName
				cti["provider"] = templateParam.Provider

				if len(templateParam.Range) > 0 {
					ptis := make([]interface{}, len(templateParam.Range), len(templateParam.Range))
					for k, trange := range templateParam.Range {
						pti := make(map[string]interface{})
						pti["id"] = trange.ID
						pti["max_value"] = trange.MaxValue
						pti["min_value"] = trange.MinValue
						ptis[k] = pti
					}
					cti["range"] = ptis
				}

				cti["required"] = templateParam.Required
				ptis := make([]interface{}, 1, 1)
				tselection := templateParam.Selection
				pti := make(map[string]interface{})
				pti["id"] = tselection.ID
				pti["selection_type"] = tselection.SelectionType
				// REVIEW: SelectionValues type
				// pti["selection_values"] = tselection.SelectionValues
				ptis[0] = pti
				cti["selection"] = ptis

				ctis[j] = cti
			}
			oi["rollback_template_params"] = ctis
		}

		oi["software_type"] = template.SoftwareType
		oi["software_variant"] = template.SoftwareVariant
		oi["software_version"] = template.SoftwareVersion
		oi["template_content"] = template.TemplateContent

		if len(template.TemplateParams) > 0 {
			ctis := make([]interface{}, len(template.TemplateParams), len(template.TemplateParams))
			for j, templateParam := range template.TemplateParams {
				cti := make(map[string]interface{})
				cti["binding"] = templateParam.Binding
				cti["data_type"] = templateParam.DataType
				cti["default_value"] = templateParam.DefaultValue
				cti["description"] = templateParam.Description
				cti["display_name"] = templateParam.DisplayName
				cti["group"] = templateParam.Group
				cti["id"] = templateParam.ID
				cti["instruction_text"] = templateParam.InstructionText
				cti["key"] = templateParam.Key
				cti["not_param"] = templateParam.NotParam
				cti["order"] = templateParam.Order
				cti["param_array"] = templateParam.ParamArray
				cti["parameter_name"] = templateParam.ParameterName
				cti["provider"] = templateParam.Provider

				if len(templateParam.Range) > 0 {
					ptis := make([]interface{}, len(templateParam.Range), len(templateParam.Range))
					for k, trange := range templateParam.Range {
						pti := make(map[string]interface{})
						pti["id"] = trange.ID
						pti["max_value"] = trange.MaxValue
						pti["min_value"] = trange.MinValue
						ptis[k] = pti
					}
					cti["range"] = ptis
				}

				cti["required"] = templateParam.Required
				ptis := make([]interface{}, 1, 1)
				tselection := templateParam.Selection
				pti := make(map[string]interface{})
				pti["id"] = tselection.ID
				pti["selection_type"] = tselection.SelectionType
				// REVIEW: SelectionValues type
				// pti["selection_values"] = tselection.SelectionValues
				ptis[0] = pti
				cti["selection"] = ptis
				ctis[j] = cti
			}
			oi["rollback_template_params"] = ctis
		}

		oi["version"] = template.Version
		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenTemplatesAvailableReadItems(templatesAvailable *[]dnac.GetsTheTemplatesAvailableResponse) []interface{} {
	if templatesAvailable != nil {
		templates := *templatesAvailable
		ois := make([]interface{}, len(templates), len(templates))
		for i, template := range templates {
			// -
			oi := make(map[string]interface{})

			oi["composite"] = template.Composite
			oi["name"] = template.Name
			oi["project_id"] = template.ProjectID
			oi["project_name"] = template.ProjectName
			oi["template_id"] = template.TemplateID

			vis := make([]interface{}, len(template.VersionsInfo), len(template.VersionsInfo))
			for j, versionInfo := range template.VersionsInfo {
				vi := make(map[string]interface{})

				vi["description"] = versionInfo.Description
				vi["id"] = versionInfo.ID
				vi["version_time"] = versionInfo.VersionTime
				vi["author"] = versionInfo.Author
				vi["version"] = versionInfo.Version
				vi["version_comment"] = versionInfo.VersionComment

				vis[j] = vi
			}
			oi["versions_info"] = vis

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenTemplateVersionsReadItem(templateVersions *[]dnac.GetTemplateVersionsResponse) []interface{} {
	if templateVersions != nil {
		templates := *templateVersions
		ois := make([]interface{}, len(templates), len(templates))
		for i, template := range templates {
			// -
			oi := make(map[string]interface{})

			oi["composite"] = template.Composite
			oi["name"] = template.Name
			oi["project_id"] = template.ProjectID
			oi["project_name"] = template.ProjectName
			oi["template_id"] = template.TemplateID

			vis := make([]interface{}, len(template.VersionsInfo), len(template.VersionsInfo))
			for j, versionInfo := range template.VersionsInfo {
				vi := make(map[string]interface{})

				vi["description"] = versionInfo.Description
				vi["id"] = versionInfo.ID
				vi["version_time"] = versionInfo.VersionTime
				vi["author"] = versionInfo.Author
				vi["version"] = versionInfo.Version
				vi["version_comment"] = versionInfo.VersionComment

				vis[j] = vi
			}
			oi["versions_info"] = vis

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPreviewTemplateReadItem(templatePreview *dnac.PreviewTemplateResponse) []interface{} {
	if templatePreview != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["cli_preview"] = templatePreview.CliPreview
		oi["template_id"] = templatePreview.TemplateID
		vis := make([]interface{}, len(templatePreview.ValidationErrors), len(templatePreview.ValidationErrors))
		for j, validationError := range templatePreview.ValidationErrors {
			vi := make(map[string]interface{})
			vi["type"] = validationError.Type
			vi["message"] = validationError.Message
			vis[j] = vi
		}
		oi["validation_errors"] = vis
		ois[0] = oi

		return ois
	}
	return make([]interface{}, 0)
}

func flattenDeployTemplateReadItem(templateDeployment *dnac.DeployTemplateResponse) []interface{} {
	if templateDeployment != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["deployment_id"] = templateDeployment.DeploymentID
		oi["deployment_name"] = templateDeployment.DeploymentName
		oi["duration"] = templateDeployment.Duration
		oi["end_time"] = templateDeployment.EndTime
		oi["project_name"] = templateDeployment.ProjectName
		oi["start_time"] = templateDeployment.StartTime
		oi["status"] = templateDeployment.Status
		oi["template_name"] = templateDeployment.TemplateName
		oi["template_version"] = templateDeployment.TemplateVersion

		dis := make([]interface{}, len(templateDeployment.Devices), len(templateDeployment.Devices))
		for j, device := range templateDeployment.Devices {
			di := make(map[string]interface{})
			di["device_id"] = device.DeviceID
			di["duration"] = device.Duration
			di["end_time"] = device.EndTime
			di["ip_address"] = device.IPAddress
			di["name"] = device.Name
			di["start_time"] = device.StartTime
			di["status"] = device.Status

			dis[j] = di
		}
		oi["devices"] = dis

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenTemplateDeployStatusReadItem(deployStatus *dnac.GetTemplateDeploymentStatusResponse) []interface{} {
	if deployStatus != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["deployment_id"] = deployStatus.DeploymentID
		oi["deployment_name"] = deployStatus.DeploymentName
		oi["duration"] = deployStatus.Duration
		oi["end_time"] = deployStatus.EndTime
		oi["project_name"] = deployStatus.ProjectName
		oi["start_time"] = deployStatus.StartTime
		oi["status"] = deployStatus.Status
		oi["template_name"] = deployStatus.TemplateName
		oi["template_version"] = deployStatus.TemplateVersion

		dis := make([]interface{}, len(deployStatus.Devices), len(deployStatus.Devices))
		for j, device := range deployStatus.Devices {
			di := make(map[string]interface{})
			di["device_id"] = device.DeviceID
			di["duration"] = device.Duration
			di["end_time"] = device.EndTime
			di["ip_address"] = device.IPAddress
			di["name"] = device.Name
			di["start_time"] = device.StartTime
			di["status"] = device.Status
			dis[j] = di
		}
		oi["devices"] = dis

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

///// end template project
///// start credentials

func flattenCredentialReadItem(credential *dnac.GetGlobalCredentialsResponseResponse) []interface{} {
	if credential != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["comments"] = credential.Comments
		oi["credential_type"] = credential.CredentialType
		oi["description"] = credential.Description
		oi["id"] = credential.ID
		oi["instance_tenant_id"] = credential.InstanceTenantID
		oi["instance_uuid"] = credential.InstanceUUID

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

///// end credentials
///// start discovery

func flattenDiscoveryReadItem(discoveryResponse *dnac.GetDiscoveryByIDResponse) []interface{} {
	if discoveryResponse != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		// REVIEW: AttributeInfo type
		//oi["attribute_info"] = discoveryResponse.Response.AttributeInfo
		oi["cdp_level"] = discoveryResponse.Response.CdpLevel
		oi["device_ids"] = discoveryResponse.Response.DeviceIDs
		oi["discovery_condition"] = discoveryResponse.Response.DiscoveryCondition
		oi["discovery_status"] = discoveryResponse.Response.DiscoveryStatus
		oi["discovery_type"] = discoveryResponse.Response.DiscoveryType
		oi["enable_password_list"] = strings.Split(discoveryResponse.Response.EnablePasswordList, ",") // Tf expects TypeList, change to []string
		oi["global_credential_id_list"] = discoveryResponse.Response.GlobalCredentialIDList            // Already []string
		oi["id"] = discoveryResponse.Response.ID
		oi["ip_address_list"] = discoveryResponse.Response.IPAddressList
		oi["ip_filter_list"] = strings.Split(discoveryResponse.Response.IPFilterList, ",")
		oi["is_auto_cdp"] = discoveryResponse.Response.IsAutoCdp
		oi["lldp_level"] = discoveryResponse.Response.LldpLevel
		oi["name"] = discoveryResponse.Response.Name
		oi["netconf_port"] = discoveryResponse.Response.NetconfPort
		oi["num_devices"] = discoveryResponse.Response.NumDevices
		oi["parent_discovery_id"] = discoveryResponse.Response.ParentDiscoveryID
		oi["password_list"] = strings.Split(discoveryResponse.Response.PasswordList, ",")
		oi["preferred_mgmt_ip_method"] = discoveryResponse.Response.PreferredMgmtIPMethod
		oi["protocol_order"] = discoveryResponse.Response.ProtocolOrder
		oi["retry"] = discoveryResponse.Response.RetryCount
		oi["snmp_auth_passphrase"] = discoveryResponse.Response.SNMPAuthPassphrase
		oi["snmp_auth_protocol"] = discoveryResponse.Response.SNMPAuthProtocol
		oi["snmp_mode"] = discoveryResponse.Response.SNMPMode
		oi["snmp_priv_passphrase"] = discoveryResponse.Response.SNMPPrivPassphrase
		oi["snmp_priv_protocol"] = discoveryResponse.Response.SNMPPrivProtocol
		oi["snmp_ro_community"] = discoveryResponse.Response.SNMPRoCommunity
		oi["snmp_ro_community_desc"] = discoveryResponse.Response.SNMPRoCommunityDesc
		oi["snmp_rw_community"] = discoveryResponse.Response.SNMPRwCommunity
		oi["snmp_rw_community_desc"] = discoveryResponse.Response.SNMPRwCommunityDesc
		oi["snmp_user_name"] = discoveryResponse.Response.SNMPUserName
		oi["timeout"] = discoveryResponse.Response.TimeOut
		oi["update_mgmt_ip"] = discoveryResponse.Response.UpdateMgmtIP
		oi["user_name_list"] = strings.Split(discoveryResponse.Response.UserNameList, ",")

		httpReadCredentials := make([]interface{}, 1, 1)
		httpWriteCredentials := make([]interface{}, 1, 1)
		httpReadCredential := make(map[string]interface{})
		httpWriteCredential := make(map[string]interface{})

		httpReadCredential["comments"] = discoveryResponse.Response.HTTPReadCredential.Comments
		httpReadCredential["credential_type"] = discoveryResponse.Response.HTTPReadCredential.CredentialType
		httpReadCredential["description"] = discoveryResponse.Response.HTTPReadCredential.Description
		httpReadCredential["id"] = discoveryResponse.Response.HTTPReadCredential.ID
		httpReadCredential["instance_tenant_id"] = discoveryResponse.Response.HTTPReadCredential.InstanceTenantID
		httpReadCredential["instance_uuid"] = discoveryResponse.Response.HTTPReadCredential.InstanceUUID
		httpReadCredential["password"] = discoveryResponse.Response.HTTPReadCredential.Password
		httpReadCredential["port"] = discoveryResponse.Response.HTTPReadCredential.Port
		httpReadCredential["secure"] = discoveryResponse.Response.HTTPReadCredential.Secure
		httpReadCredential["username"] = discoveryResponse.Response.HTTPReadCredential.Username

		httpWriteCredential["comments"] = discoveryResponse.Response.HTTPWriteCredential.Comments
		httpWriteCredential["credential_type"] = discoveryResponse.Response.HTTPWriteCredential.CredentialType
		httpWriteCredential["description"] = discoveryResponse.Response.HTTPWriteCredential.Description
		httpWriteCredential["id"] = discoveryResponse.Response.HTTPWriteCredential.ID
		httpWriteCredential["instance_tenant_id"] = discoveryResponse.Response.HTTPWriteCredential.InstanceTenantID
		httpWriteCredential["instance_uuid"] = discoveryResponse.Response.HTTPWriteCredential.InstanceUUID
		httpWriteCredential["password"] = discoveryResponse.Response.HTTPWriteCredential.Password
		httpWriteCredential["port"] = discoveryResponse.Response.HTTPWriteCredential.Port
		httpWriteCredential["secure"] = discoveryResponse.Response.HTTPWriteCredential.Secure
		httpWriteCredential["username"] = discoveryResponse.Response.HTTPWriteCredential.Username

		httpReadCredentials[0] = httpReadCredential
		httpWriteCredentials[0] = httpWriteCredential
		oi["http_read_credential"] = httpReadCredentials
		oi["http_write_credential"] = httpWriteCredentials

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

///// end discovery
///// start pnp device

func flattenPnPDeviceReadItemDayZeroConfig(response *dnac.GetDeviceByIDResponseDayZeroConfig) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})
		oi["config"] = response.Config
		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceReadItemDeviceInfoAAACredentials(response *dnac.GetDeviceByIDResponseDeviceInfoAAACredentials) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["password"] = response.Password
		oi["username"] = response.Username

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceReadItemDeviceInfoFileSystemList(response *[]dnac.GetDeviceByIDResponseDeviceInfoFileSystemList) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["freespace"] = element.Freespace
			oi["name"] = element.Name
			oi["readable"] = element.Readable
			oi["size"] = element.Size
			oi["type"] = element.Type
			oi["writeable"] = element.Writeable

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceReadItemDeviceInfoHTTPHeaders(response *[]dnac.GetDeviceByIDResponseDeviceInfoHTTPHeaders) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})
			oi["key"] = element.Key
			oi["value"] = element.Value
			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceReadItemDeviceInfoIPInterfaces(response *[]dnac.GetDeviceByIDResponseDeviceInfoIPInterfaces) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["ipv4_address"] = element.IPv4Address
			oi["ipv6_address_list"] = element.IPv6AddressList
			oi["mac_address"] = element.MacAddress
			oi["name"] = element.Name
			oi["status"] = element.Status

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceReadItemDeviceInfoLocation(response *dnac.GetDeviceByIDResponseDeviceInfoLocation) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["address"] = response.Address
		oi["altitude"] = response.Altitude
		oi["latitude"] = response.Latitude
		oi["longitude"] = response.Longitude
		oi["site_id"] = response.SiteID

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceReadItemDeviceInfoNeighborLinks(response *[]dnac.GetDeviceByIDResponseDeviceInfoNeighborLinks) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})
			oi["local_interface_name"] = element.LocalInterfaceName
			oi["local_mac_address"] = element.LocalMacAddress
			oi["local_short_interface_name"] = element.LocalShortInterfaceName
			oi["remote_device_name"] = element.RemoteDeviceName
			oi["remote_interface_name"] = element.RemoteInterfaceName
			oi["remote_mac_address"] = element.RemoteMacAddress
			oi["remote_platform"] = element.RemotePlatform
			oi["remote_short_interface_name"] = element.RemoteShortInterfaceName
			oi["remote_version"] = element.RemoteVersion
			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceReadItemDeviceInfoPnpProfileListPrimaryEndpoint(response *dnac.GetDeviceByIDResponseDeviceInfoPnpProfileListPrimaryEndpoint) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["certificate"] = response.Certificate
		oi["fqdn"] = response.Fqdn
		oi["ipv4_address"] = response.IPv4Address
		oi["ipv6_address"] = response.IPv6Address
		oi["port"] = response.Port
		oi["protocol"] = response.Protocol

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceReadItemDeviceInfoPnpProfileListSecondaryEndpoint(response *dnac.GetDeviceByIDResponseDeviceInfoPnpProfileListSecondaryEndpoint) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["certificate"] = response.Certificate
		oi["fqdn"] = response.Fqdn
		oi["ipv4_address"] = response.IPv4Address
		oi["ipv6_address"] = response.IPv6Address
		oi["port"] = response.Port
		oi["protocol"] = response.Protocol

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceReadItemDeviceInfoPnpProfileList(response *[]dnac.GetDeviceByIDResponseDeviceInfoPnpProfileList) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["created_by"] = element.CreatedBy
			oi["discovery_created"] = element.DiscoveryCreated
			oi["primary_endpoint"] = flattenPnPDeviceReadItemDeviceInfoPnpProfileListPrimaryEndpoint(&element.PrimaryEndpoint)
			oi["profile_name"] = element.ProfileName
			oi["secondary_endpoint"] = flattenPnPDeviceReadItemDeviceInfoPnpProfileListSecondaryEndpoint(&element.SecondaryEndpoint)

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceReadItemDeviceInfoPreWorkflowCliOuputs(response *[]dnac.GetDeviceByIDResponseDeviceInfoPreWorkflowCliOuputs) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["cli"] = element.Cli
			oi["cli_output"] = element.CliOutput

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceReadItemDeviceInfoStackInfoStackMemberList(response *[]dnac.GetDeviceByIDResponseDeviceInfoStackInfoStackMemberList) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["hardware_version"] = element.HardwareVersion
			oi["license_level"] = element.LicenseLevel
			oi["license_type"] = element.LicenseType
			oi["mac_address"] = element.MacAddress
			oi["pid"] = element.Pid
			oi["priority"] = element.Priority
			oi["role"] = element.Role
			oi["serial_number"] = element.SerialNumber
			oi["software_version"] = element.SoftwareVersion
			oi["stack_number"] = element.StackNumber
			oi["state"] = element.State
			oi["sudi_serial_number"] = element.SudiSerialNumber

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceReadItemDeviceInfoStackInfo(response *dnac.GetDeviceByIDResponseDeviceInfoStackInfo) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["is_full_ring"] = response.IsFullRing
		oi["stack_member_list"] = flattenPnPDeviceReadItemDeviceInfoStackInfoStackMemberList(&response.StackMemberList)
		oi["stack_ring_protocol"] = response.StackRingProtocol
		oi["supports_stack_workflows"] = response.SupportsStackWorkflows
		oi["total_member_count"] = response.TotalMemberCount
		oi["valid_license_levels"] = response.ValidLicenseLevels

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceReadItemDeviceInfo(response *dnac.GetDeviceByIDResponseDeviceInfo) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["aaa_credentials"] = flattenPnPDeviceReadItemDeviceInfoAAACredentials(&response.AAACredentials)
		oi["added_on"] = response.AddedOn
		oi["addn_mac_addrs"] = response.AddnMacAddrs
		oi["agent_type"] = response.AgentType
		oi["auth_status"] = response.AuthStatus
		oi["authenticated_mic_number"] = response.AuthenticatedMicNumber
		oi["authenticated_sudi_serial_no"] = response.AuthenticatedSudiSerialNo
		oi["capabilities_supported"] = response.CapabilitiesSupported
		oi["cm_state"] = response.CmState
		oi["description"] = response.Description
		oi["device_sudi_serial_nos"] = response.DeviceSudiSerialNos
		oi["device_type"] = response.DeviceType
		oi["features_supported"] = response.FeaturesSupported
		oi["file_system_list"] = flattenPnPDeviceReadItemDeviceInfoFileSystemList(&response.FileSystemList)
		oi["first_contact"] = response.FirstContact
		oi["hostname"] = response.Hostname
		oi["http_headers"] = flattenPnPDeviceReadItemDeviceInfoHTTPHeaders(&response.HTTPHeaders)
		oi["image_file"] = response.ImageFile
		oi["image_version"] = response.ImageVersion
		oi["http_headers"] = flattenPnPDeviceReadItemDeviceInfoIPInterfaces(&response.IPInterfaces)
		oi["last_contact"] = response.LastContact
		oi["last_sync_time"] = response.LastSyncTime
		oi["last_update_on"] = response.LastUpdateOn
		oi["location"] = flattenPnPDeviceReadItemDeviceInfoLocation(&response.Location)
		oi["mac_address"] = response.MacAddress
		oi["mode"] = response.Mode
		oi["name"] = response.Name
		oi["neighbor_links"] = flattenPnPDeviceReadItemDeviceInfoNeighborLinks(&response.NeighborLinks)
		oi["onb_state"] = response.OnbState
		oi["pid"] = response.Pid
		oi["pnp_profile_list"] = flattenPnPDeviceReadItemDeviceInfoPnpProfileList(&response.PnpProfileList)
		oi["populate_inventory"] = response.PopulateInventory
		oi["pre_workflow_cli_ouputs"] = flattenPnPDeviceReadItemDeviceInfoPreWorkflowCliOuputs(&response.PreWorkflowCliOuputs)
		oi["project_id"] = response.ProjectID
		oi["project_name"] = response.ProjectName
		oi["reload_requested"] = response.ReloadRequested
		oi["serial_number"] = response.SerialNumber
		oi["site_id"] = response.SiteID
		oi["site_name"] = response.SiteName
		oi["smart_account_id"] = response.SmartAccountID
		oi["source"] = response.Source
		oi["stack"] = response.Stack
		oi["stack_info"] = flattenPnPDeviceReadItemDeviceInfoStackInfo(&response.StackInfo)
		oi["state"] = response.State
		oi["sudi_required"] = response.SudiRequired
		oi["tags"] = response.Tags
		oi["user_mic_numbers"] = response.UserMicNumbers
		oi["user_sudi_serial_nos"] = response.UserSudiSerialNos
		oi["virtual_account_id"] = response.VirtualAccountID
		oi["workflow_id"] = response.WorkflowID
		oi["workflow_name"] = response.WorkflowName

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceReadItemRunSummaryListHistoryTaskInfoAddnDetails(response *[]dnac.GetDeviceByIDResponseRunSummaryListHistoryTaskInfoAddnDetails) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})
			oi["key"] = element.Key
			oi["value"] = element.Value
			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceReadItemRunSummaryListHistoryTaskInfoWorkItemList(response *[]dnac.GetDeviceByIDResponseRunSummaryListHistoryTaskInfoWorkItemList) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["command"] = element.Command
			oi["end_time"] = element.EndTime
			oi["output_str"] = element.OutputStr
			oi["start_time"] = element.StartTime
			oi["state"] = element.State
			oi["time_taken"] = element.TimeTaken

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceReadItemRunSummaryListHistoryTaskInfo(response *dnac.GetDeviceByIDResponseRunSummaryListHistoryTaskInfo) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["addn_details"] = flattenPnPDeviceReadItemRunSummaryListHistoryTaskInfoAddnDetails(&response.AddnDetails)
		oi["name"] = response.Name
		oi["time_taken"] = response.TimeTaken
		oi["type"] = response.Type
		oi["work_item_list"] = flattenPnPDeviceReadItemRunSummaryListHistoryTaskInfoWorkItemList(&response.WorkItemList)

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceReadItemRunSummaryList(response *[]dnac.GetDeviceByIDResponseRunSummaryList) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, summary := range *response {
			oi := make(map[string]interface{})
			oi["details"] = summary.Details
			oi["error_flag"] = summary.ErrorFlag
			oi["history_task_info"] = flattenPnPDeviceReadItemRunSummaryListHistoryTaskInfo(&summary.HistoryTaskInfo)
			oi["timestamp"] = summary.Timestamp

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceReadItemSystemResetWorkflowTasksWorkItemList(response *[]dnac.GetDeviceByIDResponseSystemResetWorkflowTasksWorkItemList) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["command"] = element.Command
			oi["end_time"] = element.EndTime
			oi["output_str"] = element.OutputStr
			oi["start_time"] = element.StartTime
			oi["state"] = element.State
			oi["time_taken"] = element.TimeTaken

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceReadItemSystemResetWorkflowTasks(response *[]dnac.GetDeviceByIDResponseSystemResetWorkflowTasks) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["curr_work_item_idx"] = element.CurrWorkItemIDx
			oi["end_time"] = element.EndTime
			oi["name"] = element.Name
			oi["start_time"] = element.StartTime
			oi["state"] = element.State
			oi["task_seq_no"] = element.TaskSeqNo
			oi["time_taken"] = element.TimeTaken
			oi["type"] = element.Type
			oi["work_item_list"] = flattenPnPDeviceReadItemSystemResetWorkflowTasksWorkItemList(&element.WorkItemList)

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceReadItemSystemResetWorkflow(response *dnac.GetDeviceByIDResponseSystemResetWorkflow) []interface{} {
	ois := make([]interface{}, 1, 1)
	if response != nil {
		oi := make(map[string]interface{})

		oi["id"] = response.TypeID
		oi["add_to_inventory"] = response.AddToInventory
		oi["added_on"] = response.AddedOn
		oi["config_id"] = response.ConfigID
		oi["curr_task_idx"] = response.CurrTaskIDx
		oi["description"] = response.Description
		oi["end_time"] = response.EndTime
		oi["exec_time"] = response.ExecTime
		oi["image_id"] = response.ImageID
		oi["instance_type"] = response.InstanceType
		oi["lastupdate_on"] = response.LastupdateOn
		oi["name"] = response.Name
		oi["start_time"] = response.StartTime
		oi["state"] = response.State
		oi["tasks"] = flattenPnPDeviceReadItemSystemResetWorkflowTasks(&response.Tasks)
		oi["tenant_id"] = response.TenantID
		oi["type"] = response.Type
		oi["use_state"] = response.UseState
		oi["version"] = response.Version

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceReadItemSystemWorkflowTasksWorkItemList(response *[]dnac.GetDeviceByIDResponseSystemWorkflowTasksWorkItemList) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["command"] = element.Command
			oi["end_time"] = element.EndTime
			oi["output_str"] = element.OutputStr
			oi["start_time"] = element.StartTime
			oi["state"] = element.State
			oi["time_taken"] = element.TimeTaken

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceReadItemSystemWorkflowTasks(response *[]dnac.GetDeviceByIDResponseSystemWorkflowTasks) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["curr_work_item_idx"] = element.CurrWorkItemIDx
			oi["end_time"] = element.EndTime
			oi["name"] = element.Name
			oi["start_time"] = element.StartTime
			oi["state"] = element.State
			oi["task_seq_no"] = element.TaskSeqNo
			oi["time_taken"] = element.TimeTaken
			oi["type"] = element.Type
			oi["work_item_list"] = flattenPnPDeviceReadItemSystemWorkflowTasksWorkItemList(&element.WorkItemList)

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceReadItemSystemWorkflow(response *dnac.GetDeviceByIDResponseSystemWorkflow) []interface{} {
	ois := make([]interface{}, 1, 1)
	if response != nil {
		oi := make(map[string]interface{})

		oi["id"] = response.TypeID
		oi["add_to_inventory"] = response.AddToInventory
		oi["added_on"] = response.AddedOn
		oi["config_id"] = response.ConfigID
		oi["curr_task_idx"] = response.CurrTaskIDx
		oi["description"] = response.Description
		oi["end_time"] = response.EndTime
		oi["exec_time"] = response.ExecTime
		oi["image_id"] = response.ImageID
		oi["instance_type"] = response.InstanceType
		oi["lastupdate_on"] = response.LastupdateOn
		oi["name"] = response.Name
		oi["start_time"] = response.StartTime
		oi["state"] = response.State
		oi["tasks"] = flattenPnPDeviceReadItemSystemWorkflowTasks(&response.Tasks)
		oi["tenant_id"] = response.TenantID
		oi["type"] = response.Type
		oi["use_state"] = response.UseState
		oi["version"] = response.Version

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceReadItemWorkflowTasksWorkItemList(response *[]dnac.GetDeviceByIDResponseWorkflowTasksWorkItemList) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["command"] = element.Command
			oi["end_time"] = element.EndTime
			oi["output_str"] = element.OutputStr
			oi["start_time"] = element.StartTime
			oi["state"] = element.State
			oi["time_taken"] = element.TimeTaken

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceReadItemWorkflowTasks(response *[]dnac.GetDeviceByIDResponseWorkflowTasks) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["curr_work_item_idx"] = element.CurrWorkItemIDx
			oi["end_time"] = element.EndTime
			oi["name"] = element.Name
			oi["start_time"] = element.StartTime
			oi["state"] = element.State
			oi["task_seq_no"] = element.TaskSeqNo
			oi["time_taken"] = element.TimeTaken
			oi["type"] = element.Type
			oi["work_item_list"] = flattenPnPDeviceReadItemWorkflowTasksWorkItemList(&element.WorkItemList)

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceReadItemWorkflow(response *dnac.GetDeviceByIDResponseWorkflow) []interface{} {
	ois := make([]interface{}, 1, 1)
	if response != nil {
		oi := make(map[string]interface{})

		oi["id"] = response.TypeID
		oi["add_to_inventory"] = response.AddToInventory
		oi["added_on"] = response.AddedOn
		oi["config_id"] = response.ConfigID
		oi["curr_task_idx"] = response.CurrTaskIDx
		oi["description"] = response.Description
		oi["end_time"] = response.EndTime
		oi["exec_time"] = response.ExecTime
		oi["image_id"] = response.ImageID
		oi["instance_type"] = response.InstanceType
		oi["lastupdate_on"] = response.LastupdateOn
		oi["name"] = response.Name
		oi["start_time"] = response.StartTime
		oi["state"] = response.State
		oi["tasks"] = flattenPnPDeviceReadItemWorkflowTasks(&response.Tasks)
		oi["tenant_id"] = response.TenantID
		oi["type"] = response.Type
		oi["use_state"] = response.UseState
		oi["version"] = response.Version

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceReadItemWorkflowParametersConfigListConfigParameters(response *[]dnac.GetDeviceByIDResponseWorkflowParametersConfigListConfigParameters) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})
			oi["key"] = element.Key
			oi["value"] = element.Value
			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceReadItemWorkflowParametersConfigList(response *[]dnac.GetDeviceByIDResponseWorkflowParametersConfigList) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, summary := range *response {
			oi := make(map[string]interface{})

			oi["config_id"] = summary.ConfigID
			oi["config_parameters"] = flattenPnPDeviceReadItemWorkflowParametersConfigListConfigParameters(&summary.ConfigParameters)

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceReadItemWorkflowParameters(response *dnac.GetDeviceByIDResponseWorkflowParameters) []interface{} {
	ois := make([]interface{}, 1, 1)
	if response != nil {
		oi := make(map[string]interface{})

		oi["config_list"] = flattenPnPDeviceReadItemWorkflowParametersConfigList(&response.ConfigList)
		oi["license_level"] = response.LicenseLevel
		oi["license_type"] = response.LicenseType
		oi["top_of_stack_serial_number"] = response.TopOfStackSerialNumber

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceReadItem(deviceResponse *dnac.GetDeviceByIDResponse) []interface{} {
	if deviceResponse != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["id"] = deviceResponse.TypeID

		oi["day_zero_config"] = flattenPnPDeviceReadItemDayZeroConfig(&deviceResponse.DayZeroConfig)

		oi["day_zero_config_preview"] = deviceResponse.DayZeroConfigPreview

		oi["device_info"] = flattenPnPDeviceReadItemDeviceInfo(&deviceResponse.DeviceInfo)

		oi["run_summary_list"] = flattenPnPDeviceReadItemRunSummaryList(&deviceResponse.RunSummaryList)

		oi["system_reset_workflow"] = flattenPnPDeviceReadItemSystemResetWorkflow(&deviceResponse.SystemResetWorkflow)

		oi["system_workflow"] = flattenPnPDeviceReadItemSystemWorkflow(&deviceResponse.SystemWorkflow)

		oi["tenant_id"] = deviceResponse.TenantID
		oi["version"] = deviceResponse.Version

		oi["workflow"] = flattenPnPDeviceReadItemWorkflow(&deviceResponse.Workflow)

		oi["workflow_parameters"] = flattenPnPDeviceReadItemWorkflowParameters(&deviceResponse.WorkflowParameters)

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

///// start pnp workflow

func flattenPnPWorkflowReadItemTasksWorkItemList(response *[]dnac.GetWorkflowByIDResponseTasksWorkItemList) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["command"] = element.Command
			oi["end_time"] = element.EndTime
			oi["output_str"] = element.OutputStr
			oi["start_time"] = element.StartTime
			oi["state"] = element.State
			oi["time_taken"] = element.TimeTaken

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPWorkflowReadItemTasks(response *[]dnac.GetWorkflowByIDResponseTasks) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["curr_work_item_idx"] = element.CurrWorkItemIDx
			oi["end_time"] = element.EndTime
			oi["name"] = element.Name
			oi["start_time"] = element.StartTime
			oi["state"] = element.State
			oi["task_seq_no"] = element.TaskSeqNo
			oi["time_taken"] = element.TimeTaken
			oi["type"] = element.Type
			oi["work_item_list"] = flattenPnPWorkflowReadItemTasksWorkItemList(&element.WorkItemList)

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPWorkflowReadItem(response *dnac.GetWorkflowByIDResponse) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["id"] = response.TypeID
		oi["add_to_inventory"] = response.AddToInventory
		oi["added_on"] = response.AddedOn
		oi["config_id"] = response.ConfigID
		oi["curr_task_idx"] = response.CurrTaskIDx
		oi["description"] = response.Description
		oi["end_time"] = response.EndTime
		oi["exec_time"] = response.ExecTime
		oi["image_id"] = response.ImageID
		oi["instance_type"] = response.InstanceType
		oi["lastupdate_on"] = response.LastupdateOn
		oi["name"] = response.Name
		oi["start_time"] = response.StartTime
		oi["state"] = response.State
		oi["tasks"] = flattenPnPWorkflowReadItemTasks(&response.Tasks)
		oi["tenant_id"] = response.TenantID
		oi["type"] = response.Type
		oi["use_state"] = response.UseState
		oi["version"] = response.Version

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

///// end pnp workflow
///// start pnp global settings

func flattenPnPGlobalSettingsReadItemsAAACredentials(response *dnac.GetPnPGlobalSettingsResponseAAACredentials) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["password"] = response.Password
		oi["username"] = response.Username

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPGlobalSettingsReadItemsDefaultProfile(response *dnac.GetPnPGlobalSettingsResponseDefaultProfile) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["cert"] = response.Cert
		oi["fqdn_addresses"] = response.FqdnAddresses
		oi["ip_addresses"] = response.IPAddresses
		oi["port"] = response.Port
		oi["proxy"] = response.Proxy

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPGlobalSettingsReadItemsSavaMappingListProfile(response *dnac.GetPnPGlobalSettingsResponseSavaMappingListProfile) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["address_fqdn"] = response.AddressFqdn
		oi["address_ip_v4"] = response.AddressIPV4
		oi["cert"] = response.Cert
		oi["make_default"] = response.MakeDefault
		oi["name"] = response.Name
		oi["port"] = response.Port
		oi["profile_id"] = response.ProfileID
		oi["proxy"] = response.Proxy

		ois[0] = oi
	}
	return make([]interface{}, 0)
}

func flattenPnPGlobalSettingsReadItemsSavaMappingListSyncResultSyncList(response *[]dnac.GetPnPGlobalSettingsResponseSavaMappingListSyncResultSyncList) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["device_sn_list"] = element.DeviceSnList
			oi["sync_type"] = element.SyncType

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPGlobalSettingsReadItemsSavaMappingListSyncResult(response *dnac.GetPnPGlobalSettingsResponseSavaMappingListSyncResult) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["sync_list"] = flattenPnPGlobalSettingsReadItemsSavaMappingListSyncResultSyncList(&response.SyncList)
		oi["sync_msg"] = response.SyncMsg

		ois[0] = oi
	}
	return make([]interface{}, 0)
}

func flattenPnPGlobalSettingsReadItemsSavaMappingList(response *[]dnac.GetPnPGlobalSettingsResponseSavaMappingList) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, item := range *response {
			oi := make(map[string]interface{})

			oi["auto_sync_period"] = item.AutoSyncPeriod
			oi["cco_user"] = item.CcoUser
			oi["expiry"] = item.Expiry
			oi["last_sync"] = item.LastSync
			oi["profile"] = flattenPnPGlobalSettingsReadItemsSavaMappingListProfile(&item.Profile)
			oi["smart_account_id"] = item.SmartAccountID
			oi["sync_result"] = flattenPnPGlobalSettingsReadItemsSavaMappingListSyncResult(&item.SyncResult)
			oi["sync_result_str"] = item.SyncResultStr
			oi["sync_start_time"] = item.SyncStartTime
			oi["sync_status"] = item.SyncStatus
			oi["tenant_id"] = item.TenantID
			oi["token"] = item.Token
			oi["virtual_account_id"] = item.VirtualAccountID

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPGlobalSettingsReadItemsTaskTimeOuts(response *dnac.GetPnPGlobalSettingsResponseTaskTimeOuts) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["config_time_out"] = response.ConfigTimeOut
		oi["general_time_out"] = response.GeneralTimeOut
		oi["image_download_time_out"] = response.ImageDownloadTimeOut

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPGlobalSettingsReadItems(response *dnac.GetPnPGlobalSettingsResponse) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["id"] = response.TypeID
		oi["aaa_credentials"] = flattenPnPGlobalSettingsReadItemsAAACredentials(&response.AAACredentials)
		oi["accept_eula"] = response.AcceptEula
		oi["default_profile"] = flattenPnPGlobalSettingsReadItemsDefaultProfile(&response.DefaultProfile)
		oi["sava_mapping_list"] = flattenPnPGlobalSettingsReadItemsSavaMappingList(&response.SavaMappingList)
		oi["task_time_outs"] = flattenPnPGlobalSettingsReadItemsTaskTimeOuts(&response.TaskTimeOuts)
		oi["tenant_id"] = response.TenantID
		oi["version"] = response.Version

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

///// end pnp global settings
///// start other pnp elements

func flattenPnPDeviceSyncResultVacctItemProfile(response *dnac.GetSyncResultForVirtualAccountResponseProfile) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["address_fqdn"] = response.AddressFqdn
		oi["address_ip_v4"] = response.AddressIPV4
		oi["cert"] = response.Cert
		oi["make_default"] = response.MakeDefault
		oi["name"] = response.Name
		oi["port"] = response.Port
		oi["profile_id"] = response.ProfileID
		oi["proxy"] = response.Proxy

		ois[0] = oi
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceSyncResultVacctItemSyncResultSyncList(response *[]dnac.GetSyncResultForVirtualAccountResponseSyncResultSyncList) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["device_sn_list"] = element.DeviceSnList
			oi["sync_type"] = element.SyncType

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceSyncResultVacctItemSyncResult(response *dnac.GetSyncResultForVirtualAccountResponseSyncResult) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["sync_list"] = flattenPnPDeviceSyncResultVacctItemSyncResultSyncList(&response.SyncList)
		oi["sync_msg"] = response.SyncMsg

		ois[0] = oi
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceSyncResultVacctItem(response *dnac.GetSyncResultForVirtualAccountResponse) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["auto_sync_period"] = response.AutoSyncPeriod
		oi["cco_user"] = response.CcoUser
		oi["expiry"] = response.Expiry
		oi["last_sync"] = response.LastSync
		oi["profile"] = flattenPnPDeviceSyncResultVacctItemProfile(&response.Profile)
		oi["smart_account_id"] = response.SmartAccountID
		oi["sync_result"] = flattenPnPDeviceSyncResultVacctItemSyncResult(&response.SyncResult)
		oi["sync_result_str"] = response.SyncResultStr
		oi["sync_start_time"] = response.SyncStartTime
		oi["sync_status"] = response.SyncStatus
		oi["tenant_id"] = response.TenantID
		oi["token"] = response.Token
		oi["virtual_account_id"] = response.VirtualAccountID

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPWorkflowsReadItemsTasksWorkItemList(response *[]dnac.GetWorkflowsResponseTasksWorkItemList) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["command"] = element.Command
			oi["end_time"] = element.EndTime
			oi["output_str"] = element.OutputStr
			oi["start_time"] = element.StartTime
			oi["state"] = element.State
			oi["time_taken"] = element.TimeTaken

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPWorkflowsReadItemsTasks(response *[]dnac.GetWorkflowsResponseTasks) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["curr_work_item_idx"] = element.CurrWorkItemIDx
			oi["end_time"] = element.EndTime
			oi["name"] = element.Name
			oi["start_time"] = element.StartTime
			oi["state"] = element.State
			oi["task_seq_no"] = element.TaskSeqNo
			oi["time_taken"] = element.TimeTaken
			oi["type"] = element.Type
			oi["work_item_list"] = flattenPnPWorkflowsReadItemsTasksWorkItemList(&element.WorkItemList)

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPWorkflowsReadItems(response *[]dnac.GetWorkflowsResponse) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))

		for i, item := range *response {
			oi := make(map[string]interface{})

			oi["id"] = item.TypeID
			oi["add_to_inventory"] = item.AddToInventory
			oi["added_on"] = item.AddedOn
			oi["config_id"] = item.ConfigID
			oi["curr_task_idx"] = item.CurrTaskIDx
			oi["description"] = item.Description
			oi["end_time"] = item.EndTime
			oi["exec_time"] = item.ExecTime
			oi["image_id"] = item.ImageID
			oi["instance_type"] = item.InstanceType
			oi["lastupdate_on"] = item.LastupdateOn
			oi["name"] = item.Name
			oi["start_time"] = item.StartTime
			oi["state"] = item.State
			oi["tasks"] = flattenPnPWorkflowsReadItemsTasks(&item.Tasks)
			oi["tenant_id"] = item.TenantID
			oi["type"] = item.Type
			oi["use_state"] = item.UseState
			oi["version"] = item.Version

			ois[i] = oi
		}

		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItemsDayZeroConfig(response *dnac.GetPnpDeviceListResponseDayZeroConfig) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})
		oi["config"] = response.Config
		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItemsDeviceInfoAAACredentials(response *dnac.GetPnpDeviceListResponseDeviceInfoAAACredentials) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["password"] = response.Password
		oi["username"] = response.Username

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItemsDeviceInfoFileSystemList(response *[]dnac.GetPnpDeviceListResponseDeviceInfoFileSystemList) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["freespace"] = element.Freespace
			oi["name"] = element.Name
			oi["readable"] = element.Readable
			oi["size"] = element.Size
			oi["type"] = element.Type
			oi["writeable"] = element.Writeable

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItemsDeviceInfoHTTPHeaders(response *[]dnac.GetPnpDeviceListResponseDeviceInfoHTTPHeaders) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})
			oi["key"] = element.Key
			oi["value"] = element.Value
			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItemsDeviceInfoIPInterfaces(response *[]dnac.GetPnpDeviceListResponseDeviceInfoIPInterfaces) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["ipv4_address"] = element.IPv4Address
			oi["ipv6_address_list"] = element.IPv6AddressList
			oi["mac_address"] = element.MacAddress
			oi["name"] = element.Name
			oi["status"] = element.Status

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItemsDeviceInfoLocation(response *dnac.GetPnpDeviceListResponseDeviceInfoLocation) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["address"] = response.Address
		oi["altitude"] = response.Altitude
		oi["latitude"] = response.Latitude
		oi["longitude"] = response.Longitude
		oi["site_id"] = response.SiteID

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItemsDeviceInfoNeighborLinks(response *[]dnac.GetPnpDeviceListResponseDeviceInfoNeighborLinks) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})
			oi["local_interface_name"] = element.LocalInterfaceName
			oi["local_mac_address"] = element.LocalMacAddress
			oi["local_short_interface_name"] = element.LocalShortInterfaceName
			oi["remote_device_name"] = element.RemoteDeviceName
			oi["remote_interface_name"] = element.RemoteInterfaceName
			oi["remote_mac_address"] = element.RemoteMacAddress
			oi["remote_platform"] = element.RemotePlatform
			oi["remote_short_interface_name"] = element.RemoteShortInterfaceName
			oi["remote_version"] = element.RemoteVersion
			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItemsDeviceInfoPnpProfileListPrimaryEndpoint(response *dnac.GetPnpDeviceListResponseDeviceInfoPnpProfileListPrimaryEndpoint) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["certificate"] = response.Certificate
		oi["fqdn"] = response.Fqdn
		oi["ipv4_address"] = response.IPv4Address
		oi["ipv6_address"] = response.IPv6Address
		oi["port"] = response.Port
		oi["protocol"] = response.Protocol

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItemsDeviceInfoPnpProfileListSecondaryEndpoint(response *dnac.GetPnpDeviceListResponseDeviceInfoPnpProfileListSecondaryEndpoint) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["certificate"] = response.Certificate
		oi["fqdn"] = response.Fqdn
		oi["ipv4_address"] = response.IPv4Address
		oi["ipv6_address"] = response.IPv6Address
		oi["port"] = response.Port
		oi["protocol"] = response.Protocol

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItemsDeviceInfoPnpProfileList(response *[]dnac.GetPnpDeviceListResponseDeviceInfoPnpProfileList) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["created_by"] = element.CreatedBy
			oi["discovery_created"] = element.DiscoveryCreated
			oi["primary_endpoint"] = flattenPnPDevicesReadItemsDeviceInfoPnpProfileListPrimaryEndpoint(&element.PrimaryEndpoint)
			oi["profile_name"] = element.ProfileName
			oi["secondary_endpoint"] = flattenPnPDevicesReadItemsDeviceInfoPnpProfileListSecondaryEndpoint(&element.SecondaryEndpoint)

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItemsDeviceInfoPreWorkflowCliOuputs(response *[]dnac.GetPnpDeviceListResponseDeviceInfoPreWorkflowCliOuputs) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["cli"] = element.Cli
			oi["cli_output"] = element.CliOutput

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItemsDeviceInfoStackInfoStackMemberList(response *[]dnac.GetPnpDeviceListResponseDeviceInfoStackInfoStackMemberList) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["hardware_version"] = element.HardwareVersion
			oi["license_level"] = element.LicenseLevel
			oi["license_type"] = element.LicenseType
			oi["mac_address"] = element.MacAddress
			oi["pid"] = element.Pid
			oi["priority"] = element.Priority
			oi["role"] = element.Role
			oi["serial_number"] = element.SerialNumber
			oi["software_version"] = element.SoftwareVersion
			oi["stack_number"] = element.StackNumber
			oi["state"] = element.State
			oi["sudi_serial_number"] = element.SudiSerialNumber

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItemsDeviceInfoStackInfo(response *dnac.GetPnpDeviceListResponseDeviceInfoStackInfo) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["is_full_ring"] = response.IsFullRing
		oi["stack_member_list"] = flattenPnPDevicesReadItemsDeviceInfoStackInfoStackMemberList(&response.StackMemberList)
		oi["stack_ring_protocol"] = response.StackRingProtocol
		oi["supports_stack_workflows"] = response.SupportsStackWorkflows
		oi["total_member_count"] = response.TotalMemberCount
		oi["valid_license_levels"] = response.ValidLicenseLevels

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItemsDeviceInfo(response *dnac.GetPnpDeviceListResponseDeviceInfo) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["aaa_credentials"] = flattenPnPDevicesReadItemsDeviceInfoAAACredentials(&response.AAACredentials)
		oi["added_on"] = response.AddedOn
		oi["addn_mac_addrs"] = response.AddnMacAddrs
		oi["agent_type"] = response.AgentType
		oi["auth_status"] = response.AuthStatus
		oi["authenticated_mic_number"] = response.AuthenticatedMicNumber
		oi["authenticated_sudi_serial_no"] = response.AuthenticatedSudiSerialNo
		oi["capabilities_supported"] = response.CapabilitiesSupported
		oi["cm_state"] = response.CmState
		oi["description"] = response.Description
		oi["device_sudi_serial_nos"] = response.DeviceSudiSerialNos
		oi["device_type"] = response.DeviceType
		oi["features_supported"] = response.FeaturesSupported
		oi["file_system_list"] = flattenPnPDevicesReadItemsDeviceInfoFileSystemList(&response.FileSystemList)
		oi["first_contact"] = response.FirstContact
		oi["hostname"] = response.Hostname
		oi["http_headers"] = flattenPnPDevicesReadItemsDeviceInfoHTTPHeaders(&response.HTTPHeaders)
		oi["image_file"] = response.ImageFile
		oi["image_version"] = response.ImageVersion
		oi["http_headers"] = flattenPnPDevicesReadItemsDeviceInfoIPInterfaces(&response.IPInterfaces)
		oi["last_contact"] = response.LastContact
		oi["last_sync_time"] = response.LastSyncTime
		oi["last_update_on"] = response.LastUpdateOn
		oi["location"] = flattenPnPDevicesReadItemsDeviceInfoLocation(&response.Location)
		oi["mac_address"] = response.MacAddress
		oi["mode"] = response.Mode
		oi["name"] = response.Name
		oi["neighbor_links"] = flattenPnPDevicesReadItemsDeviceInfoNeighborLinks(&response.NeighborLinks)
		oi["onb_state"] = response.OnbState
		oi["pid"] = response.Pid
		oi["pnp_profile_list"] = flattenPnPDevicesReadItemsDeviceInfoPnpProfileList(&response.PnpProfileList)
		oi["populate_inventory"] = response.PopulateInventory
		oi["pre_workflow_cli_ouputs"] = flattenPnPDevicesReadItemsDeviceInfoPreWorkflowCliOuputs(&response.PreWorkflowCliOuputs)
		oi["project_id"] = response.ProjectID
		oi["project_name"] = response.ProjectName
		oi["reload_requested"] = response.ReloadRequested
		oi["serial_number"] = response.SerialNumber
		oi["site_id"] = response.SiteID
		oi["site_name"] = response.SiteName
		oi["smart_account_id"] = response.SmartAccountID
		oi["source"] = response.Source
		oi["stack"] = response.Stack
		oi["stack_info"] = flattenPnPDevicesReadItemsDeviceInfoStackInfo(&response.StackInfo)
		oi["state"] = response.State
		oi["sudi_required"] = response.SudiRequired
		oi["tags"] = response.Tags
		oi["user_mic_numbers"] = response.UserMicNumbers
		oi["user_sudi_serial_nos"] = response.UserSudiSerialNos
		oi["virtual_account_id"] = response.VirtualAccountID
		oi["workflow_id"] = response.WorkflowID
		oi["workflow_name"] = response.WorkflowName

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItemsRunSummaryListHistoryTaskInfoAddnDetails(response *[]dnac.GetPnpDeviceListResponseRunSummaryListHistoryTaskInfoAddnDetails) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})
			oi["key"] = element.Key
			oi["value"] = element.Value
			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItemsRunSummaryListHistoryTaskInfoWorkItemList(response *[]dnac.GetPnpDeviceListResponseRunSummaryListHistoryTaskInfoWorkItemList) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["command"] = element.Command
			oi["end_time"] = element.EndTime
			oi["output_str"] = element.OutputStr
			oi["start_time"] = element.StartTime
			oi["state"] = element.State
			oi["time_taken"] = element.TimeTaken

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItemsRunSummaryListHistoryTaskInfo(response *dnac.GetPnpDeviceListResponseRunSummaryListHistoryTaskInfo) []interface{} {
	ois := make([]interface{}, 1, 1)
	if response != nil {
		oi := make(map[string]interface{})

		oi["addn_details"] = flattenPnPDevicesReadItemsRunSummaryListHistoryTaskInfoAddnDetails(&response.AddnDetails)
		oi["name"] = response.Name
		oi["time_taken"] = response.TimeTaken
		oi["type"] = response.Type
		oi["work_item_list"] = flattenPnPDevicesReadItemsRunSummaryListHistoryTaskInfoWorkItemList(&response.WorkItemList)

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItemsRunSummaryList(response *[]dnac.GetPnpDeviceListResponseRunSummaryList) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, summary := range *response {
			oi := make(map[string]interface{})
			oi["details"] = summary.Details
			oi["error_flag"] = summary.ErrorFlag
			oi["history_task_info"] = flattenPnPDevicesReadItemsRunSummaryListHistoryTaskInfo(&summary.HistoryTaskInfo)
			oi["timestamp"] = summary.Timestamp

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItemsSystemResetWorkflowTasksWorkItemList(response *[]dnac.GetPnpDeviceListResponseSystemResetWorkflowTasksWorkItemList) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["command"] = element.Command
			oi["end_time"] = element.EndTime
			oi["output_str"] = element.OutputStr
			oi["start_time"] = element.StartTime
			oi["state"] = element.State
			oi["time_taken"] = element.TimeTaken

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItemsSystemResetWorkflowTasks(response *[]dnac.GetPnpDeviceListResponseSystemResetWorkflowTasks) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["curr_work_item_idx"] = element.CurrWorkItemIDx
			oi["end_time"] = element.EndTime
			oi["name"] = element.Name
			oi["start_time"] = element.StartTime
			oi["state"] = element.State
			oi["task_seq_no"] = element.TaskSeqNo
			oi["time_taken"] = element.TimeTaken
			oi["type"] = element.Type
			oi["work_item_list"] = flattenPnPDevicesReadItemsSystemResetWorkflowTasksWorkItemList(&element.WorkItemList)

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItemsSystemResetWorkflow(response *dnac.GetPnpDeviceListResponseSystemResetWorkflow) []interface{} {
	ois := make([]interface{}, 1, 1)
	if response != nil {
		oi := make(map[string]interface{})

		oi["id"] = response.TypeID
		oi["add_to_inventory"] = response.AddToInventory
		oi["added_on"] = response.AddedOn
		oi["config_id"] = response.ConfigID
		oi["curr_task_idx"] = response.CurrTaskIDx
		oi["description"] = response.Description
		oi["end_time"] = response.EndTime
		oi["exec_time"] = response.ExecTime
		oi["image_id"] = response.ImageID
		oi["instance_type"] = response.InstanceType
		oi["lastupdate_on"] = response.LastupdateOn
		oi["name"] = response.Name
		oi["start_time"] = response.StartTime
		oi["state"] = response.State
		oi["tasks"] = flattenPnPDevicesReadItemsSystemResetWorkflowTasks(&response.Tasks)
		oi["tenant_id"] = response.TenantID
		oi["type"] = response.Type
		oi["use_state"] = response.UseState
		oi["version"] = response.Version

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItemsSystemWorkflowTasksWorkItemList(response *[]dnac.GetPnpDeviceListResponseSystemWorkflowTasksWorkItemList) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["command"] = element.Command
			oi["end_time"] = element.EndTime
			oi["output_str"] = element.OutputStr
			oi["start_time"] = element.StartTime
			oi["state"] = element.State
			oi["time_taken"] = element.TimeTaken

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItemsSystemWorkflowTasks(response *[]dnac.GetPnpDeviceListResponseSystemWorkflowTasks) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["curr_work_item_idx"] = element.CurrWorkItemIDx
			oi["end_time"] = element.EndTime
			oi["name"] = element.Name
			oi["start_time"] = element.StartTime
			oi["state"] = element.State
			oi["task_seq_no"] = element.TaskSeqNo
			oi["time_taken"] = element.TimeTaken
			oi["type"] = element.Type
			oi["work_item_list"] = flattenPnPDevicesReadItemsSystemWorkflowTasksWorkItemList(&element.WorkItemList)

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItemsSystemWorkflow(response *dnac.GetPnpDeviceListResponseSystemWorkflow) []interface{} {
	ois := make([]interface{}, 1, 1)
	if response != nil {
		oi := make(map[string]interface{})

		oi["id"] = response.TypeID
		oi["add_to_inventory"] = response.AddToInventory
		oi["added_on"] = response.AddedOn
		oi["config_id"] = response.ConfigID
		oi["curr_task_idx"] = response.CurrTaskIDx
		oi["description"] = response.Description
		oi["end_time"] = response.EndTime
		oi["exec_time"] = response.ExecTime
		oi["image_id"] = response.ImageID
		oi["instance_type"] = response.InstanceType
		oi["lastupdate_on"] = response.LastupdateOn
		oi["name"] = response.Name
		oi["start_time"] = response.StartTime
		oi["state"] = response.State
		oi["tasks"] = flattenPnPDevicesReadItemsSystemWorkflowTasks(&response.Tasks)
		oi["tenant_id"] = response.TenantID
		oi["type"] = response.Type
		oi["use_state"] = response.UseState
		oi["version"] = response.Version

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItemsWorkflowTasksWorkItemList(response *[]dnac.GetPnpDeviceListResponseWorkflowTasksWorkItemList) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["command"] = element.Command
			oi["end_time"] = element.EndTime
			oi["output_str"] = element.OutputStr
			oi["start_time"] = element.StartTime
			oi["state"] = element.State
			oi["time_taken"] = element.TimeTaken

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItemsWorkflowTasks(response *[]dnac.GetPnpDeviceListResponseWorkflowTasks) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["curr_work_item_idx"] = element.CurrWorkItemIDx
			oi["end_time"] = element.EndTime
			oi["name"] = element.Name
			oi["start_time"] = element.StartTime
			oi["state"] = element.State
			oi["task_seq_no"] = element.TaskSeqNo
			oi["time_taken"] = element.TimeTaken
			oi["type"] = element.Type
			oi["work_item_list"] = flattenPnPDevicesReadItemsWorkflowTasksWorkItemList(&element.WorkItemList)

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItemsWorkflow(response *dnac.GetPnpDeviceListResponseWorkflow) []interface{} {
	ois := make([]interface{}, 1, 1)
	if response != nil {
		oi := make(map[string]interface{})

		oi["id"] = response.TypeID
		oi["add_to_inventory"] = response.AddToInventory
		oi["added_on"] = response.AddedOn
		oi["config_id"] = response.ConfigID
		oi["curr_task_idx"] = response.CurrTaskIDx
		oi["description"] = response.Description
		oi["end_time"] = response.EndTime
		oi["exec_time"] = response.ExecTime
		oi["image_id"] = response.ImageID
		oi["instance_type"] = response.InstanceType
		oi["lastupdate_on"] = response.LastupdateOn
		oi["name"] = response.Name
		oi["start_time"] = response.StartTime
		oi["state"] = response.State
		oi["tasks"] = flattenPnPDevicesReadItemsWorkflowTasks(&response.Tasks)
		oi["tenant_id"] = response.TenantID
		oi["type"] = response.Type
		oi["use_state"] = response.UseState
		oi["version"] = response.Version

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItemsWorkflowParametersConfigListConfigParameters(response *[]dnac.GetPnpDeviceListResponseWorkflowParametersConfigListConfigParameters) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})
			oi["key"] = element.Key
			oi["value"] = element.Value
			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItemsWorkflowParametersConfigList(response *[]dnac.GetPnpDeviceListResponseWorkflowParametersConfigList) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, summary := range *response {
			oi := make(map[string]interface{})

			oi["config_id"] = summary.ConfigID
			oi["config_parameters"] = flattenPnPDevicesReadItemsWorkflowParametersConfigListConfigParameters(&summary.ConfigParameters)

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItemsWorkflowParameters(response *dnac.GetPnpDeviceListResponseWorkflowParameters) []interface{} {
	ois := make([]interface{}, 1, 1)
	if response != nil {
		oi := make(map[string]interface{})

		oi["config_list"] = flattenPnPDevicesReadItemsWorkflowParametersConfigList(&response.ConfigList)
		oi["license_level"] = response.LicenseLevel
		oi["license_type"] = response.LicenseType
		oi["top_of_stack_serial_number"] = response.TopOfStackSerialNumber

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesReadItems(response *[]dnac.GetPnpDeviceListResponse) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))

		for i, deviceResponse := range *response {
			oi := make(map[string]interface{})

			oi["id"] = deviceResponse.TypeID

			oi["day_zero_config"] = flattenPnPDevicesReadItemsDayZeroConfig(&deviceResponse.DayZeroConfig)

			oi["day_zero_config_preview"] = deviceResponse.DayZeroConfigPreview

			oi["device_info"] = flattenPnPDevicesReadItemsDeviceInfo(&deviceResponse.DeviceInfo)

			oi["run_summary_list"] = flattenPnPDevicesReadItemsRunSummaryList(&deviceResponse.RunSummaryList)

			oi["system_reset_workflow"] = flattenPnPDevicesReadItemsSystemResetWorkflow(&deviceResponse.SystemResetWorkflow)

			oi["system_workflow"] = flattenPnPDevicesReadItemsSystemWorkflow(&deviceResponse.SystemWorkflow)

			oi["tenant_id"] = deviceResponse.TenantID
			oi["version"] = deviceResponse.Version

			oi["workflow"] = flattenPnPDevicesReadItemsWorkflow(&deviceResponse.Workflow)

			oi["workflow_parameters"] = flattenPnPDevicesReadItemsWorkflowParameters(&deviceResponse.WorkflowParameters)

			ois[i] = oi

		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesHistoryReadItemsHistoryTaskInfoAddnDetails(response *[]dnac.GetDeviceHistoryResponseResponseHistoryTaskInfoAddnDetails) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})
			oi["key"] = element.Key
			oi["value"] = element.Value
			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesHistoryReadItemsHistoryTaskInfoWorkItemList(response *[]dnac.GetDeviceHistoryResponseResponseHistoryTaskInfoWorkItemList) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["command"] = element.Command
			oi["end_time"] = element.EndTime
			oi["output_str"] = element.OutputStr
			oi["start_time"] = element.StartTime
			oi["state"] = element.State
			oi["time_taken"] = element.TimeTaken

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesHistoryReadItemsHistoryTaskInfo(response *dnac.GetDeviceHistoryResponseResponseHistoryTaskInfo) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["addn_details"] = flattenPnPDevicesHistoryReadItemsHistoryTaskInfoAddnDetails(&response.AddnDetails)
		oi["name"] = response.Name
		oi["time_taken"] = response.TimeTaken
		oi["type"] = response.Type
		oi["work_item_list"] = flattenPnPDevicesHistoryReadItemsHistoryTaskInfoWorkItemList(&response.WorkItemList)

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDevicesHistoryReadItems(response *[]dnac.GetDeviceHistoryResponseResponse) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, item := range *response {

			oi := make(map[string]interface{})
			oi["details"] = item.Details
			oi["error_flag"] = item.ErrorFlag
			oi["history_task_info"] = flattenPnPDevicesHistoryReadItemsHistoryTaskInfo(&item.HistoryTaskInfo)
			oi["timestamp"] = item.Timestamp

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceUnClaimReadItem(response *dnac.UnClaimDeviceResponse) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["json_array_response"] = response.JSONArrayResponse
		oi["json_response"] = response.JSONResponse
		oi["message"] = response.Message
		oi["status_code"] = response.StatusCode

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceClaimReadItem(response *dnac.ClaimDeviceResponse) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["json_array_response"] = response.JSONArrayResponse
		oi["json_response"] = response.JSONResponse
		oi["message"] = response.Message
		oi["status_code"] = response.StatusCode

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceClaimSiteReadItem(response *dnac.ClaimADeviceToASiteResponse) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["response"] = response.Response
		oi["version"] = response.Version

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceConfigPreviewReadItemResponse(response *dnac.PreviewConfigResponseResponse) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["complete"] = response.Complete
		oi["config"] = response.Config
		oi["error"] = response.Error
		oi["error_message"] = response.ErrorMessage
		oi["expired_time"] = response.ExpiredTime
		oi["rf_profile"] = response.RfProfile
		oi["sensor_profile"] = response.SensorProfile
		oi["site_id"] = response.SiteID
		oi["start_time"] = response.StartTime
		oi["task_id"] = response.TaskID

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceConfigPreviewReadItem(response *dnac.PreviewConfigResponse) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["response"] = flattenPnPDeviceConfigPreviewReadItemResponse(&response.Response)
		oi["version"] = response.Version

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceResetReadItem(response *dnac.ResetDeviceResponse) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["json_array_response"] = response.JSONArrayResponse
		oi["json_response"] = response.JSONResponse
		oi["message"] = response.Message
		oi["status_code"] = response.StatusCode

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceSyncVacctItemProfile(response *dnac.SyncVirtualAccountDevicesResponseProfile) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["address_fqdn"] = response.AddressFqdn
		oi["address_ip_v4"] = response.AddressIPV4
		oi["cert"] = response.Cert
		oi["make_default"] = response.MakeDefault
		oi["name"] = response.Name
		oi["port"] = response.Port
		oi["profile_id"] = response.ProfileID
		oi["proxy"] = response.Proxy

		ois[0] = oi
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceSyncVacctItemSyncResultSyncList(response *[]dnac.SyncVirtualAccountDevicesResponseSyncResultSyncList) []interface{} {
	if response != nil {
		ois := make([]interface{}, len(*response), len(*response))
		for i, element := range *response {
			oi := make(map[string]interface{})

			oi["device_sn_list"] = element.DeviceSnList
			oi["sync_type"] = element.SyncType

			ois[i] = oi
		}
		return ois
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceSyncVacctItemSyncResult(response *dnac.SyncVirtualAccountDevicesResponseSyncResult) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["sync_list"] = flattenPnPDeviceSyncVacctItemSyncResultSyncList(&response.SyncList)
		oi["sync_msg"] = response.SyncMsg

		ois[0] = oi
	}
	return make([]interface{}, 0)
}

func flattenPnPDeviceSyncVacctItem(response *dnac.SyncVirtualAccountDevicesResponse) []interface{} {
	if response != nil {
		ois := make([]interface{}, 1, 1)
		oi := make(map[string]interface{})

		oi["auto_sync_period"] = response.AutoSyncPeriod
		oi["cco_user"] = response.CcoUser
		oi["expiry"] = response.Expiry
		oi["last_sync"] = response.LastSync
		oi["profile"] = flattenPnPDeviceSyncVacctItemProfile(&response.Profile)
		oi["smart_account_id"] = response.SmartAccountID
		oi["sync_result"] = flattenPnPDeviceSyncVacctItemSyncResult(&response.SyncResult)
		oi["sync_result_str"] = response.SyncResultStr
		oi["sync_start_time"] = response.SyncStartTime
		oi["sync_status"] = response.SyncStatus
		oi["tenant_id"] = response.TenantID
		oi["token"] = response.Token
		oi["virtual_account_id"] = response.VirtualAccountID

		ois[0] = oi
		return ois
	}
	return make([]interface{}, 0)
}

///// end other pnp elements
