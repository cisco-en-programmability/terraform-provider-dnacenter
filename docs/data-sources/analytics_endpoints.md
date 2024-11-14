---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_analytics_endpoints Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on AI Endpoint Analytics.
  Query the endpoints, optionally using various filter and pagination criteria. 'GET /endpoints/count' API can be used
  to find out the total number of endpoints matching the filter criteria.Fetches details of the endpoint for the given unique identifier 'epId'.
---

# dnacenter_analytics_endpoints (Data Source)

It performs read operation on AI Endpoint Analytics.

- Query the endpoints, optionally using various filter and pagination criteria. 'GET /endpoints/count' API can be used
to find out the total number of endpoints matching the filter criteria.

- Fetches details of the endpoint for the given unique identifier 'epId'.

## Example Usage

```terraform
data "dnacenter_analytics_endpoints" "example" {
  provider = dnacenter
  ep_id    = "string"
  include  = "string"
}

output "dnacenter_analytics_endpoints_example" {
  value = data.dnacenter_analytics_endpoints.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `ai_spoofing_trust_level` (String) aiSpoofingTrustLevel query parameter. Trust level of the endpoint due to AI spoofing. Possible values are 'low', 'medium', 'high'.
- `anc_policy` (String) ancPolicy query parameter. ANC policy. Only exact match will be returned.
- `auth_method` (String) authMethod query parameter. Authentication method. Partial string is allowed.
- `changed_profile_trust_level` (String) changedProfileTrustLevel query parameter. Trust level of the endpoint due to changing profile labels. Possible values are 'low', 'medium', 'high'.
- `concurrent_mac_trust_level` (String) concurrentMacTrustLevel query parameter. Trust level of the endpoint due to concurrent MAC address. Possible values are 'low', 'medium', 'high'.
- `device_type` (String) deviceType query parameter. Type of device to search for. Partial string is allowed.
- `ep_id` (String) epId path parameter. Unique identifier for the endpoint.
- `hardware_manufacturer` (String) hardwareManufacturer query parameter. Hardware manufacturer to search for. Partial string is allowed.
- `hardware_model` (String) hardwareModel query parameter. Hardware model to search for. Partial string is allowed.
- `include` (String) include query parameter. The datasets that should be included in the response. By default, value of this parameter is blank, and the response will include only basic details of the endpoint. To include other datasets or dictionaries, send comma separated list of following values: 'ALL' Include all attributes. 'CDP', 'DHCP', etc. Include attributes from given dictionaries. To get full list of dictionaries, use corresponding GET API. 'ANC' Include ANC policy related details. 'TRUST' Include trust score details.
- `ip` (String) ip query parameter. IP address to search for. Partial string is allowed.
- `ip_blocklist_detected` (Boolean) ipBlocklistDetected query parameter. Flag to fetch endpoints hitting IP blocklist or not.
- `limit` (Number) limit query parameter. Maximum number of records to be fetched. If not provided, 50 records will be fetched by default. Maximum 1000 records can be fetched at a time. Use pagination if more records need to be fetched.
- `mac_address` (String) macAddress query parameter. MAC address to search for. Partial string is allowed.
- `mac_addresses` (List of String) macAddresses query parameter. List of MAC addresses to filter on. Only exact matches will be returned.
- `nat_trust_level` (String) natTrustLevel query parameter. Trust level of the endpoint due to NAT access. Possible values are 'low', 'medium', 'high'.
- `offset` (Number) offset query parameter. Record offset to start data fetch at. Offset starts at zero.
- `operating_system` (String) operatingSystem query parameter. Operating system to search for. Partial string is allowed.
- `order` (String) order query parameter. Order to be used for sorting. Possible values are 'asc', 'desc'.
- `posture_status` (String) postureStatus query parameter. Posture status.
- `profiling_status` (String) profilingStatus query parameter. Profiling status of the endpoint. Possible values are 'profiled', 'partialProfiled', 'notProfiled'.
- `random_mac` (Boolean) randomMac query parameter. Flag to fetch endpoints having randomized MAC or not.
- `registered` (Boolean) registered query parameter. Flag to fetch manually registered or non-registered endpoints.
- `sort_by` (String) sortBy query parameter. Name of the column to sort the results on. Please note that fetch might take more time if sorting is requested. Possible values are 'macAddress', 'ip'.
- `trust_score` (String) trustScore query parameter. Overall trust score of the endpoint. It can be provided either as a number value (e.g. 5), or as a range (e.g. 3-7). Provide value as '-' if you want to search for all endpoints where trust score is not assigned.
- `unauth_port_detected` (Boolean) unauthPortDetected query parameter. Flag to fetch endpoints exposing unauthorized ports or not.
- `weak_cred_detected` (Boolean) weakCredDetected query parameter. Flag to fetch endpoints having weak credentials or not.

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `anc_policy` (String)
- `attributes` (String)
- `device_type` (List of String)
- `duid` (String)
- `granular_anc_policy` (List of Object) (see [below for nested schema](#nestedobjatt--item--granular_anc_policy))
- `hardware_manufacturer` (List of String)
- `hardware_model` (List of String)
- `id` (String)
- `last_probe_collection_timestamp` (Number)
- `mac_address` (String)
- `operating_system` (List of String)
- `random_mac` (String)
- `registered` (String)
- `trust_data` (List of Object) (see [below for nested schema](#nestedobjatt--item--trust_data))

<a id="nestedobjatt--item--granular_anc_policy"></a>
### Nested Schema for `item.granular_anc_policy`

Read-Only:

- `name` (String)
- `nas_ip_address` (String)


<a id="nestedobjatt--item--trust_data"></a>
### Nested Schema for `item.trust_data`

Read-Only:

- `ai_spoofing_trust_level` (String)
- `auth_method` (String)
- `changed_profile_trust_level` (String)
- `concurrent_mac_trust_level` (String)
- `ip_blocklist_detected` (String)
- `nat_trust_level` (String)
- `posture_status` (String)
- `trust_score` (Number)
- `unauth_port_detected` (String)
- `weak_cred_detected` (String)