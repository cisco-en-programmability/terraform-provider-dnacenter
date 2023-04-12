---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_event_syslog_config Resource - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It manages create, read and update operations on Event Management.
  Update Syslog DestinationCreate Syslog Destination
---

# dnacenter_event_syslog_config (Resource)

It manages create, read and update operations on Event Management.

- Update Syslog Destination

- Create Syslog Destination

## Example Usage

```terraform
resource "dnacenter_event_syslog_config" "example" {
  provider = dnacenter

  parameters {

    config_id   = "string"
    description = "string"
    host        = "string"
    name        = "string"
    port        = "string"
    protocol    = "string"
  }
}

output "dnacenter_event_syslog_config_example" {
  value = dnacenter_event_syslog_config.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `parameters` (Block List) (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `last_updated` (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Optional:

- `config_id` (String) Required only for update syslog configuration
- `description` (String) Description
- `host` (String) Host
- `name` (String) Name
- `port` (String) Port
- `protocol` (String) Protocol


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `api_status` (String)
- `error_message` (List of Object) (see [below for nested schema](#nestedobjatt--item--error_message))
- `status_message` (List of Object) (see [below for nested schema](#nestedobjatt--item--status_message))

<a id="nestedobjatt--item--error_message"></a>
### Nested Schema for `item.error_message`

Read-Only:

- `errors` (List of String)


<a id="nestedobjatt--item--status_message"></a>
### Nested Schema for `item.status_message`

Read-Only:

- `config_id` (String)
- `description` (String)
- `host` (String)
- `name` (String)
- `port` (Number)
- `protocol` (String)
- `tenant_id` (String)
- `version` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import dnacenter_event_syslog_config.example "id:=string"
```