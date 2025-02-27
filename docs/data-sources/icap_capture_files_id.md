---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_icap_capture_files_id Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Sensors.
  Retrieves details of a specific ICAP packet capture file. For detailed information about the usage of the API, please
  refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
  specs/blob/main/Assurance/CECatCenter_Org-icap-1.0.0-resolved.yaml
---

# dnacenter_icap_capture_files_id (Data Source)

It performs read operation on Sensors.

- Retrieves details of a specific ICAP packet capture file. For detailed information about the usage of the API, please
refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-icap-1.0.0-resolved.yaml

## Example Usage

```terraform
data "dnacenter_icap_capture_files_id" "example" {
  provider    = dnacenter
  id          = "string"
  xca_lle_rid = "string"
}

output "dnacenter_icap_capture_files_id_example" {
  value = data.dnacenter_icap_capture_files_id.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) id path parameter. The name of the packet capture file, as given by the GET /captureFiles API response.
- `xca_lle_rid` (String) X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.

### Read-Only

- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `ap_mac` (String)
- `client_mac` (String)
- `file_creation_timestamp` (Number)
- `file_name` (String)
- `file_size` (Number)
- `id` (String)
- `last_updated_timestamp` (Number)
- `type` (String)
