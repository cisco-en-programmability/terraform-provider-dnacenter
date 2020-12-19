---
page_title: "dna_snmpv2_write_community_credential Resource - terraform-provider-dnacenter"
subcategory: "Discovery"
description: |-
  The dna_snmpv2_write_community_credential resource allows you to configure a DNACenter SNMPv2 write community credential.
---

# Resource dna_snmpv2_write_community_credential

The dna_snmpv2_write_community_credential resource allows you to configure a DNACenter SNMPv2 write community credential.

## Example Usage

```hcl
resource "dna_snmpv2_write_community_credential" "response" {
  provider = dnacenter
  item {
    description     = "SNMP RO test 1"
    write_community = "ThisI5aP4s_sW0rd"
    credential_type = "APP"
    id              = "e566053f-d5cd-4a81-841e-cb91a712af20"
  }
}
```

## Argument Reference

- `item` - (Required) Item in a DNACenter SNMPv2 write community credential. See [Credential item](#credential-item) below for details.

### Credential item

Each SNMPv2 write community credential item contains the following arguments.

- `comments` - (Optional) The SNMPv2 write community credential's comments.
- `credential_type` - (Optional) The SNMPv2 write community credential's credential type. Available values are "GLOBAL" and "APP".
- `description` - (Required) The SNMPv2 write community credential's description.
- `id` - (Optional) The SNMPv2 write community credential's id.
- `instance_tenant_id` - (Optional) The SNMPv2 write community credential's instance tenant id.
- `instance_uuid` - (Optional) The SNMPv2 write community credential's instance uuid.
- `write_community` - (Required) The SNMPv2 write community credential's community.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `last_updated` - The SNMPv2 write community credential's updated time with format RFC850.
