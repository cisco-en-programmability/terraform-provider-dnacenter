---
page_title: "dna_site Data Source - terraform-provider-dnacenter"
subcategory: "Sites"
description: |-
  The dna_site data source allows you to retrieve information about a particular Cisco DNA Center site.
---

# Data Source dna_site

The dna_site data source allows you to retrieve information about a particular Cisco DNA Center site.

## Example Usage

```hcl
data "dna_site" "response" {
  provider = dnacenter
  offset   = "1"
  limit    = "3"
  name     = "Global/USA"
}
```

## Argument Reference

- `name` - (Optional) The name param.
- `site_id` - (Optional) The site_id param.
- `type` - (Optional) The type param.
- `offset` - (Optional) The offset param.
- `limit` - (Optional) The limit param.

## Attributes Reference

The following attributes are exported.

- `items` - The item response. See [Items](#items) below for details.

### Items

Each site item contains `id`, `type`, `name` and `parent_name`.

- `id` - The site's id.
- `type` - The site's type, available values are **area**, **building** and **floor**.
- `name` - The site's name.
- `parent_name` - The site's parent name.

Each site item of type **building** contains `address`, `latitude` and `longitude`.

- `address` - The building's address.
- `latitude` - The building's latitute.
- `longitude` - The building's longitude.

Each site item of type **floor** contains `height`, `length`, `rf_model` and `width`.

- `height` - The floor's height.
- `length` - The floor's length.
- `rf_model` - The floor's rf model.
- `width` - The floor's width.
