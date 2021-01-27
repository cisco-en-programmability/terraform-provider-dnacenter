---
page_title: "dna_site Resource - terraform-provider-dnacenter"
subcategory: "Sites"
description: |-
  The dna_site resource allows you to configure a Cisco DNA Center site.
---

# Resource dna_site

The dna_site resource allows you to configure a Cisco DNA Center site.

## Example Usage

### Area

```hcl
resource "dna_site" "area" {
  provider = dnacenter
  item {
    type = "area"
    name = "Peru"
    parent_name = "Global"
  }
}
```

### Building

```hcl
resource "dna_site" "building" {
  provider = dnacenter
  depends_on = [ dna_site.area ]
  item {
    type = "building"
    name = "Miraflores"
    parent_name = "Global/Peru"
    address = "Miraflores, Lima, Lima Province, Peru"
    latitude = -12.1209
    longitude = -77.0289
  }
}
```

### Floor

```hcl
resource "dna_site" "floor" {
  provider = dnacenter
  depends_on = [ dna_site.building ]
  item {
    type = "floor"
    name = "Floor 1"
    parent_name = "Global/Peru/Miraflores"
    rf_model = "Cubes And Walled Offices"
    height = 100.1
    length = 100.2
    width = 100.1
  }
}
```

## Argument Reference

- `item` - (Required) Item in a Cisco DNA Center site. See [Site item](#site-item) below for details.

### Site item

Each site item contains `type`, `name` and `parent_name`.

- `type` - (Required) The site's type, available values are **area**, **building** and **floor**. If it's changed it forces the creation of a new resource.
- `name` - (Required) The site's name. If it's changed it forces the creation of a new resource.
- `parent_name` - (Required) The site's parent name. If it's changed it forces the creation of a new resource.

Each site item of type **building** contains `address`, `latitude` and `longitude`.

- `address` - (Optional) The building's address.
- `latitude` - (Optional) The building's latitute.
- `longitude` - (Optional) The building's longitude.

Each site item of type **floor** contains `height`, `length`, `rf_model` and `width`.

- `height` - (Optional) The floor's height.
- `length` - (Optional) The floor's length.
- `rf_model` - (Optional) The floor's rf model.
- `width` - (Optional) The floor's width.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `last_updated` - The site's updated time with format RFC850.
- `item` - Item in a Cisco DNA Center site. See [Site item](#site-item-1) below for details.

### Site item

- `id` - The site's id.
