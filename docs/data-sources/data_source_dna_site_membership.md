---
page_title: "dna_site_membership Data Source - terraform-provider-dnacenter"
subcategory: "Sites"
description: |-
  The dna_site_membership data source allows you to retrieve information about a particular Cisco DNA Center site membership
---

# Data Source dna_site_membership

The dna_site_membership data source allows you to retrieve information about a particular Cisco DNA Center site membership

## Example Usage

```hcl
data "dna_site_membership" "response" {
  provider = dnacenter
  site_id  = "a013dd15-69a3-423f-82dc-c6a10eba2cb7"
}
```

## Argument Reference

- `site_id` - (Required) The site_id param.
- `offset` - (Optional) The offset param.
- `limit` - (Optional) The limit param.
- `device_family` - (Optional) The device_family param.
- `serial_number` - (Optional) The serial_number param.

## Attributes Reference

The following attributes are exported.

- `response` - The item response. See [Membership Response](#membership-response) below for details.

### Membership Response

- `device` - The item's device. See [device](#device) below for details.
- `site` - The item's site. See [site](#site) below for details.

#### device

- `response` - The device's response.
- `site_id` - The device's site_id.
- `version` - The device's version.
- `message` - The device's message.

#### site

- `site` - The site's response. See [site-response](#site-response) below for details.

##### site response

- `parent_id` - The site response's parent id.
- `additional_info` - The site response's additional info.
- `group_type_list` - The site response's group type list.
- `group_hierarchy` - The site response's group hierarchy.
- `group_name_hierarchy` - The site response's group name hierarchy.
- `name` - The site response's name.
- `instance_tenant_id` - The site response's instance tenant id.
- `id` - The site response's id.
