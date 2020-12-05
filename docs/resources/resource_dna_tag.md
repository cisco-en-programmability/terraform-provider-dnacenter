---
page_title: "dna_tag Resource - terraform-provider-dnacenter"
subcategory: ""
description: |-
  The dna_tag resource allows you to configure a DNACenter tag.
---

# Resource dna_tag

The dna_tag resource allows you to configure a DNACenter tag.

## Example Usage

```terraform
resource "dna_tag" "data" {
  provider = dnacenter
  item {
    system_tag = false
    description = "New tag description for Terraform 012"
    name = "Tag012"
    instance_tenant_id = "15cdc6c45a8405f00c80c6ba3"
    dynamic_rules {
      member_type = "networkdevice"
      rules {
        operation = "ILIKE"
        name = "family"
        value = "%Switches and Hubs%"
      }
    }
  }
}
```

## Argument Reference

- `item` - (Required) Item in a DNACenter tag. See [Tag item](#tag-item) below for details.

### Tag item

Each tag item contains `system_tag`, `description`, `name`, `instance_tenant_id` and the `dynamic_rules` list.

- `system_tag` - (Required) The tag's systemTag flag.
- `description` - (Required) The tag's description.
- `name` - (Required) The tag's name.
- `instance_tenant_id` - (Required) The tag's instanceTenantId.
- `dynamic_rules` - (Optional) The tag's dynamic rules. See [Dynamic Rules](#dynamic-rules) below for details.

#### Dynamic Rules

- `member_type` - (Optional) The dynamic rule's member type.
- `rules` - (Optional) The dynamic rule's rules definition. See [Rules](#rules) below for details.

##### Rules

- `name` - (Optional) The rule's name.
- `operation` - (Optional) The rule's operation.
- `value` - (Optional) The rule's value.
- `values` - (Optional) The rule's values, a list of strings.
- `items` - (Optional) The rule's items, a list of strings.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `last_updated` - (Computed) The tag's updated time with format time.RFC850.
