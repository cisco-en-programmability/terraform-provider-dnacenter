---
page_title: "dna_site_health Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  The dna_site_health data source allows you to retrieve information about a particular DNACenter site health.
---

# Data Source dna_site_health

The dna_site_health data source allows you to retrieve information about a particular DNACenter site health.

## Example Usage

```hcl
data "dna_site_health" "response" {
  provider = dnacenter
}
```

## Argument Reference

- `timestamp` - (Optional) The timestamp param.

## Attributes Reference

The following attributes are exported.

- `items` - The item response. See [Items](#items) below for details.

### Items

- `access_good_count` - The access good count.
- `access_total_count` - The access total count.
- `application_bytes_total_count` - The application bytes total count.
- `application_good_count` - The application good count.
- `application_health` - The application health.
- `application_total_count` - The application total count.
- `client_health_wired` - The client health wired.
- `client_health_wireless` - The client health wireless.
- `core_good_count` - The core good count.
- `core_total_count` - The core total count.
- `distribution_good_count` - The distribution good count.
- `distribution_total_count` - The distribution total count.
- `dnac_info` - The dnac info.
- `healthy_clients_percentage` - The healthy clients percentage.
- `healthy_network_device_percentage` - The healthy network device percentage.
- `latitude` - The latitude.
- `longitude` - The longitude.
- `network_health_access` - The network health access.
- `network_health_average` - The network health average.
- `network_health_core` - The network health core.
- `network_health_distribution` - The network health distribution.
- `network_health_others` - The network health others.
- `network_health_router` - The network health router.
- `network_health_wireless` - The network health wireless.
- `number_of_clients` - The number of clients.
- `number_of_network_device` - The number of network device.
- `number_of_wired_clients` - The number of wired clients.
- `number_of_wireless_clients` - The number of wireless clients.
- `overall_good_devices` - The overall good devices.
- `parent_site_id` - The parent site id.
- `parent_site_name` - The parent site name.
- `router_good_count` - The router good count.
- `router_total_count` - The router total count.
- `site_id` - The site id.
- `site_name` - The site name.
- `site_type` - The site type.
- `total_number_of_active_wireless_clients` - The total number of active wireless clients.
- `total_number_of_connected_wired_clients` - The total number of connected wired clients.
- `wired_good_clients` - The wired good clients.
- `wireless_device_good_count` - The wireless device good count.
- `wireless_device_total_count` - The wireless device total count.
- `wireless_good_clients` - The wireless good clients.
- `application_health_stats` - The application health stats. See [application_health_stats](#application_health_stats)

#### application_health_stats

- `app_total_count` - The app total count.
- `business_irrelevant_app_fair` - The business irrelevant app that are fair.
- `business_irrelevant_app_good` - The business irrelevant app that are good.
- `business_irrelevant_app_poor` - The business irrelevant app that are poor.
- `business_relevant_app_fair` - The business relevant app that are fair.
- `business_relevant_app_good` - The business relevant app that are good.
- `business_relevant_app_poor` - The business relevant app that are poor.
- `default_health_app_fair` - The default health app that are fair.
- `default_health_app_good` - The default health app that are good.
- `default_health_app_poor` - The default health app that are poor.
