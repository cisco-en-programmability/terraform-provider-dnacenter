---
page_title: "dna_discovery_snmp_property_add Data Source - terraform-provider-dnacenter"
subcategory: "Discovery"
description: |-
  The dna_discovery_snmp_property_add data source allows you to add/update information about a particular Cisco DNA Center discovery SNMP property.
---

# Data Source dna_discovery_snmp_property_add

The dna_discovery_snmp_property_add data source allows you to add/update information about a particular Cisco DNA Center discovery SNMP property.

## Example Usage

```hcl
data "dna_discovery_snmp_property_add" "response" {
  provider = dnacenter
  items {
    system_property_name = var.system_property_name
  }
}
```

## Argument Reference

- `items` - (Required) Items in a Cisco DNA Center discovery SNMP property. See [Items](#items) below for details.

### Items

Each item contains the following arguments.

- `id` - (Optional) The item's id.
- `instance_tenant_id` - (Optional) The item's instance tenant id.
- `instance_uuid` - (Optional) The item's instance uuid.
- `int_value` - (Optional) The item's int value.
- `system_property_name` - (Required) The item's system property name.

## Attributes Reference

In addition to all the attributes above, no other attributes are exported.
