---
page_title: "dna_applications Data Source - terraform-provider-dnacenter"
subcategory: "Application Policy"
description: |-
  The application data source allows you to retrieve information about a particular DNACenter application.
---

# Data Source dna_applications

The application data source allows you to retrieve information about a particular DNACenter application.

## Example Usage

```hcl
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

The following attributes are exported.

- `items` - Items in a DNACenter app. See [Application items](#application-items) below for details.

### Application items

Each application item contains `application_set_id`, `application_network_applications` and `application_network_identity`.

- `application_set_id` - (Required) The tag's instanceTenantId.
- `application_network_applications` - (Optional) List of the application's network applications. See [Network Applications](#network-applications) below for details.
- `application_network_identity` - (Optional) List of the tag's dynamic rules. See [Network Identity](#network-identity) below for details.

#### Network Applications

- `id` - The network applications's id.
- `app_protocol` - The network application's app protocol.
- `application_subtype` - The network application's application subtype.
- `application_type` - The network application's application type.
- `category_id` - The network application's category id.
- `display_name` - The network application's display name.
- `dscp` - The network application's dscp.
- `engine_id` - The network application's engine id.
- `help_string` - The network application's help string.
- `ignore_conflict` - The network application's ignore conflict.
- `long_description` - The network application's long description.
- `name` - The network application's name.
- `popularity` - The network application's popularity.
- `rank` - The network application's rank.
- `server_name` - The network application's server name.
- `traffic_class` - The network application's traffic class.
- `url` - The network application's url.

#### Network Identity

- `id` - The network identity's id.
- `display_name` - The network identity's display name.
- `lower_port` - The network identity's lower port.
- `ports` - The network identity's ports.
- `protocol` - The network identity's protocol.
- `upper_port` - The network identity's upper port.
