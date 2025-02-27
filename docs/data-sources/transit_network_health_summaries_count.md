---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_transit_network_health_summaries_count Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on SDA.
  Get a count of transit networks. Use available query parameters to get the count of a subset of transit networks.
  This data source provides the latest health data until the given endTime. If data is not ready for the provided
  endTime, the request will fail with error code 400 Bad Request, and the error message will indicate the recommended
  endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we
  are not a real time system. When endTime is not provided, the API returns the latest data.
  For detailed information about the usage of the API, please refer to the Open API specification document
  https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CECatCenter_Org-
  transitNetworkHealthSummaries-1.0.1-resolved.yaml
---

# dnacenter_transit_network_health_summaries_count (Data Source)

It performs read operation on SDA.

- Get a count of transit networks. Use available query parameters to get the count of a subset of transit networks.
This data source provides the latest health data until the given *endTime*. If data is not ready for the provided
endTime, the request will fail with error code *400 Bad Request*, and the error message will indicate the recommended
endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we
are not a real time system. When *endTime* is not provided, the API returns the latest data.
For detailed information about the usage of the API, please refer to the Open API specification document
https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
transitNetworkHealthSummaries-1.0.1-resolved.yaml

## Example Usage

```terraform
data "dnacenter_transit_network_health_summaries_count" "example" {
  provider    = dnacenter
  end_time    = 1609459200
  id          = "string"
  start_time  = 1609459200
  xca_lle_rid = "string"
}

output "dnacenter_transit_network_health_summaries_count_example" {
  value = data.dnacenter_transit_network_health_summaries_count.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `xca_lle_rid` (String) X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.

### Optional

- `end_time` (Number) endTime query parameter. End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
- `id` (String) id query parameter. The list of transit entity ids. (Ex "1551156a-bc97-3c63-aeda-8a6d3765b5b9") Examples: id=1551156a-bc97-3c63-aeda-8a6d3765b5b9 (single entity uuid requested) id=1551156a-bc97-3c63-aeda-8a6d3765b5b9&id=4aa20652-237c-4625-b2b4-fd7e82b6a81e (multiple entity uuids with '&' separator)
- `start_time` (Number) startTime query parameter. Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.

### Read-Only

- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `detail` (String)
- `error_code` (Number)
- `message` (String)
