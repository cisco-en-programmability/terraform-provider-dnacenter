---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_device_configurations_export Resource - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs create operation on Configuration Archive.
  Export Device configurations to an encrypted zip file
---

# dnacenter_device_configurations_export (Resource)

It performs create operation on Configuration Archive.

- Export Device configurations to an encrypted zip file
~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in DNACenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "dnacenter_device_configurations_export" "example" {
  provider = dnacenter
  parameters {

    device_id = "string"
    password  = "******"
  }
}

output "dnacenter_device_configurations_export_example" {
  value = dnacenter_device_configurations_export.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `parameters` (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `last_updated` (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Optional:

- `device_id` (String) UUIDs of the devices for which configurations need to be exported.
- `password` (String, Sensitive) Password for the zip file to protect exported configurations. Must contain, at minimum 8 characters, one lowercase letter, one uppercase letter, one number, one special character(-=[];,./~!@#$%^&*()_+{}|:?). It may not contain white space or the characters <>.


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `task_id` (String)
- `url` (String)
