---
page_title: "dna_application_set Resource - terraform-provider-dnacenter"
subcategory: "Application Policy"
description: |-
  The dna_application_set resource allows you to configure a Cisco DNA Center application set.
---

# Resource dna_application_set

The dna_application_set resource allows you to configure a Cisco DNA Center application set.

## Example Usage

```hcl
resource "dna_application_set" "response" {
  provider = dnacenter
  item {
    name = "test-set"
  }
}
```

## Argument Reference

The following arguments are supported:

- `item` - (Required) The application set's item. See [Application Set item](#application-set-item) below for details.

### Application Set item

Each application set item contains `name`.

- `name` - (Required) The application set's name. If it's changed it forces the creation of a new resource.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `last_updated` - The site's updated time with format RFC850.
- `id` - The application set's id.
- `identity_source` - The application set's identity source. See [Identity Source](#identity-source) below for details.

### Identity Source

Each application set's identity source contains `id` and `type`.

- `id` - The identity source's id.
- `type` - The identity source's type.
