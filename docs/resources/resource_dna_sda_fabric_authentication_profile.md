---
page_title: "dna_sda_fabric_authentication_profile Resource - terraform-provider-dnacenter"
subcategory: "SDA"
description: |-
  The dna_sda_fabric_authentication_profile resource allows you to configure a DNACenter SDA auth profile.
---

# Resource dna_sda_fabric_authentication_profile

The dna_sda_fabric_authentication_profile resource allows you to configure a DNACenter SDA auth profile.

## Example Usage

```hcl
resource "dna_sda_fabric_authentication_profile" "template_1" {
  provider   = dnacenter
  site_name_hierarchy = var.sda_auth_site_name_hierarchy
  authenticate_template_name = var.sda_auth_authenticate_template_name
  authenticate_template_id = var.sda_auth_authenticate_template_id
}
```

## Argument Reference

- `site_name_hierarchy` - (Required) The auth profile's site name hierarchy. If it's changed it forces the creation of a new resource.
- `authenticate_template_name` - (Required) The auth profile's authenticate template name.
- `authenticate_template_id` - (Required) The auth profile's authenticate template id.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `last_updated` - The auth profile's updated time with format RFC850.
