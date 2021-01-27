---
page_title: "dna_snmpv3_credential Resource - terraform-provider-dnacenter"
subcategory: "Discovery"
description: |-
  The dna_snmpv3_credential resource allows you to configure a Cisco DNA Center SNMPv3 credential.
---

# Resource dna_snmpv3_credential

The dna_snmpv3_credential resource allows you to configure a Cisco DNA Center SNMPv3 credential.

## Example Usage

```hcl
resource "dna_snmpv3_credential" "response" {
  provider = dnacenter
  item {
    snmp_mode = "NOAUTHNOPRIV"
    username  = "user3"
  }
}
```

## Argument Reference

- `item` - (Required) Item in a Cisco DNA Center SNMPv3 credential. See [Credential item](#credential-item) below for details.

### Credential item

Each SNMPv3 credential item contains the following arguments.

- `auth_password` - (Optional) The SNMPv3 credential's auth_password.
- `auth_type` - (Optional) The SNMPv3 credential's auth_type.
- `comments` - (Optional) The SNMPv3 credential's comments.
- `credential_type` - (Optional) The SNMPv3 credential's credential type. Available values are "GLOBAL" and "APP".
- `description` - (Optional) The SNMPv3 credential's description.
- `id` - (Optional) The SNMPv3 credential's id.
- `instance_tenant_id` - (Optional) The SNMPv3 credential's instance tenant id.
- `instance_uuid` - (Optional) The SNMPv3 credential's instance uuid.
- `privacy_password` - (Optional) - The SNMPv3 credentials' privacy password.
- `privacy_type` - (Optional) - The SNMPv3 credentials' privacy type. Available values are "DES" and "AES128".
- `snmp_mode` - (Required) - The SNMPv3 credentials' SNMP mode. Available values are "AUTHPRIV", "AUTHNOPRIV" and "NOAUTHNOPRIV".
- `username` - (Required) - The SNMPv3 credentials' username.

~> The value of `snmp_mode` indicates if `privacy_password` or `privacy_type` are needed. This validation is not done by the Tf resource.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `last_updated` - The SNMPv3 credential's updated time with format RFC850.
