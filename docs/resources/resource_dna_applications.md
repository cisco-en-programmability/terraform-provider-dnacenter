---
page_title: "dna_applications Resource - terraform-provider-dnacenter"
subcategory: "Application Policy"
description: |-
  The dna_applications resource allows you to configure a Cisco DNA Center application.
---

# Resource dna_applications

The dna_applications resource allows you to configure a Cisco DNA Center application.

## Example Usage

```hcl
resource "dna_applications" "data" {
  provider = dnacenter
  name = ""
  items {
    application_set_id = ""
    application_network_applications {
      app_protocol = ""
      application_subtype = ""
      application_type = ""
      category_id = ""
      display_name = ""
      dscp = ""
      engine_id = ""
      help_string = ""
      ignore_conflict = ""
      long_description = ""
      name = ""
      popularity = 0
      rank = 0
      server_name = ""
      traffic_class = ""
      url = ""
    }
    application_network_identity {
      display_name = ""
      lower_port = 0
      ports = ""
      protocol = ""
      upper_port = 8080
    }
  }
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Required) The application's name. If it's changed it forces the creation of a new resource.
- `items` - (Required) The application's items. See [Application items](#application-items) below for details.

### Application items

Each application item contains `application_set_id`, `application_network_applications` and `application_network_identity`.

- `application_set_id` - (Required) The tag's instanceTenantId.
- `application_network_applications` - (Optional) List the application's network applications. See [Application Network Applications](#application-network-applications) below for details.
- `application_network_identity` - (Optional) List the tag's dynamic rules. See [Application Network Identity](#application-network-identity) below for details.

#### Application Network Applications

- `app_protocol` - (Required) The network application's app protocol.
- `application_subtype` - (Required) The network application's application subtype.
- `application_type` - (Required) The network application's application type.
- `category_id` - (Required) The network application's category id.
- `display_name` - (Required) The network application's display name.
- `dscp` - (Required) The network application's dscp.
- `engine_id` - (Required) The network application's engine id.
- `help_string` - (Required) The network application's help string.
- `ignore_conflict` - (Required) The network application's ignore conflict.
- `long_description` - (Required) The network application's long description.
- `name` - (Required) The network application's name.
- `popularity` - (Required) The network application's popularity.
- `rank` - (Required) The network application's rank.
- `server_name` - (Required) The network application's server name.
- `traffic_class` - (Required) The network application's traffic class.
- `url` - (Required) The network application's url.

#### Application Network Identity

- `display_name` - (Required) The network identity's display name.
- `lower_port` - (Required) The network identity's lower port.
- `ports` - (Required) The network identity's ports.
- `protocol` - (Required) The network identity's protocol.
- `upper_port` - (Required) The network identity's upper port.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `last_updated` - The site's updated time with format RFC850.
- `items` - The application's items. See [Application items](#application-items-1) below for details.

### Application items

Each application item contains `application_set_id`, `application_network_applications` and `application_network_identity`.

- `application_network_applications` - List the application's network applications. See [Application Network Applications](#application-network-applications-1) below for details.
- `application_network_identity` - List the tag's dynamic rules. See [Application Network Identity](#application-network-identity-1) below for details.

#### Application Network Applications

- `id` - The network applications's id.

#### Application Network Identity

- `id` - The network identity's id.
