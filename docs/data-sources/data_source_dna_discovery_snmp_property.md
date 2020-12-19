---
page_title: "dna_discovery_snmp_property Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  The dna_discovery_snmp_property data source allows you to retrieve information about a particular DNACenter discovery SNMP properties.
---

# Data Source dna_discovery_snmp_property

The dna_discovery_snmp_property data source allows you to retrieve information about a particular DNACenter discovery SNMP properties.

## Example Usage

```hcl
data "dna_discovery_snmp_property" "response" {
  provider = dnacenter
}
```

## Attributes Reference

The following attributes are exported.

- `items` - Items in a DNACenter discovery SNMP property. See [Items](#items) below for details.

### Items

Each item contains the following attributes.

- `id` - The item's id.
- `instance_tenant_id` - The item's instance tenant id.
- `instance_uuid` - The item's instance uuid.
- `int_value` - The item's int value.
- `system_property_name` - The item's system property name.
