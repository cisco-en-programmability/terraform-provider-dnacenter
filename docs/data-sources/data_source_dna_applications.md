---
page_title: "dna_applications Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  The application data source allows you to retrieve information about a particular DNACenter application.
---

# Data Source dna_applications

The application data source allows you to retrieve information about a particular DNACenter application.

## Example Usage

```terraform
data "dna_applications" "list" {
  provider = dnacenter
  offset = 0
  limit = 4
}
```

## Argument Reference

- `name` - (Optional) DNACenter application name.
- `offset` - (Optional) DNACenter application offset.
- `limit` - (Optional) DNACenter application limit.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `items` - (Computed) Items in a DNACenter tag. See [Application items](#application-items) below for details.

### Application items

Each application item contains `application_set_id`, `application_network_applications` and `application_network_identity`.

- `application_set_id` - (Required) The tag's instanceTenantId.
- `application_network_applications` - (List, Optional) The application's network applications. See [Network Applications](#network-applications) below for details.
- `application_network_identity` - (List, Optional) The tag's dynamic rules. See [Network Identity](#network-identity) below for details.

#### Network Applications

- `id` - (Computed) The network applications's id.
- `app_protocol` - (Computed) The network application's app protocol.
- `application_subtype` - (Computed) The network application's application subtype.
- `application_type` - (Computed) The network application's application type.
- `category_id` - (Computed) The network application's category id.
- `display_name` - (Computed) The network application's display name.
- `dscp` - (Computed) The network application's dscp.
- `engine_id` - (Computed) The network application's engine id.
- `help_string` - (Computed) The network application's help string.
- `ignore_conflict` - (Computed) The network application's ignore conflict.
- `long_description` - (Computed) The network application's long description.
- `name` - (Computed) The network application's name.
- `popularity` - (Computed) The network application's popularity.
- `rank` - (Computed) The network application's rank.
- `server_name` - (Computed) The network application's server name.
- `traffic_class` - (Computed) The network application's traffic class.
- `url` - (Computed) The network application's url.

#### Network Identity

- `id` - (Computed) The network identity's id.
- `display_name` - (Computed) The network identity's display name.
- `lower_port` - (Computed) The network identity's lower port.
- `ports` - (Computed) The network identity's ports.
- `protocol` - (Computed) The network identity's protocol.
- `upper_port` - (Computed) The network identity's upper port.
