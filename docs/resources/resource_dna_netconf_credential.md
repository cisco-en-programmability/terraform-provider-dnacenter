---
page_title: "dna_netconf_credential Resource - terraform-provider-dnacenter"
subcategory: "Discovery"
description: |-
  The dna_netconf_credential resource allows you to configure a Cisco DNA Center Netconf credential.
---

# Resource dna_netconf_credential

The dna_netconf_credential resource allows you to configure a Cisco DNA Center Netconf credential.

## Example Usage

```hcl
resource "dna_netconf_credential" "response" {
  provider = dnacenter
  item {
    netconf_port = 23
    description  = "netconf 23"
  }
}
```

## Argument Reference

- `item` - (Required) Item in a Cisco DNA Center Netconf credential. See [Credential item](#credential-item) below for details.

### Credential item

Each Netconf credential item contains the following arguments.

- `comments` - (Optional) The Netconf credential's comments.
- `credential_type` - (Optional) The Netconf credential's credential type. Available values are "GLOBAL" and "APP".
- `description` - (Optional) The Netconf credential's description.
- `id` - (Optional) The Netconf credential's id.
- `instance_tenant_id` - (Optional) The Netconf credential's instance tenant id.
- `instance_uuid` - (Optional) The Netconf credential's instance uuid.
- `netconf_port` - (Required) The Netconf credential's port.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `last_updated` - The Netconf credential's updated time with format RFC850.
