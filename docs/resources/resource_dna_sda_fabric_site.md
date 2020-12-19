---
page_title: "dna_sda_fabric_site Resource - terraform-provider-dnacenter"
subcategory: "SDA"
description: |-
  The dna_sda_fabric_site resource allows you to configure a DNACenter SDA site.
---

# Resource dna_sda_fabric_site

The dna_sda_fabric_site resource allows you to configure a DNACenter SDA site.

## Example Usage

```hcl
resource "dna_sda_fabric_site" "response" {
  provider   = dnacenter
  site_name_hierarchy = site.site_name_hierarchy
  fabric_name = site.fabric_name
}
```

## Argument Reference

- `site_name_hierarchy` - (Required) The site's name hierarchy. If it's changed it forces the creation of a new resource.
- `fabric_name` - (Required) The site's fabric name. If it's changed it forces the creation of a new resource.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `last_updated` - The site's updated time with format RFC850.
