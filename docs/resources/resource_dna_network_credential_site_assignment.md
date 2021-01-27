---
page_title: "dna_network_credential_site_assignment Resource - terraform-provider-dnacenter"
subcategory: "Network Settings"
description: |-
  The dna_network_credential_site_assignment resource allows you to assign DNAC credentials to site.
---

# Resource dna_network_credential_site_assignment

The dna_network_credential_site_assignment resource allows you to assign DNAC credentials to site.

## Example Usage

```hcl
resource "dna_network_credential_site_assignment" "response" {
  provider = dnacenter
  site_id  = "b85fe0be-d971-4eb7-92a9-3498356ad87f"
  http_read {
    id = "babc42b9-0bdd-49ef-912e-66f533fb5d59"
  }
  cli {
    id = "f979d842-f6fd-456a-8137-2cb5113cd2e8"
  }
}
```

## Argument Reference

- `item` - (Required) Item in a Cisco DNA Center assignment. See [Credential item](#credential-item) below for details.

- `site_id` - (Required) The site id.
- `cli` - (Optional) The CLI credential to assign. See [Credential item](#credential-item) below for details.
- `http_read` - (Optional) The HTTP read credential to assign. See [Credential item](#credential-item) below for details.
- `http_write` - (Optional) The HTTP write credential to assign. See [Credential item](#credential-item) below for details.
- `snmp_v2_read` - (Optional) The SNMPv2 Read credential to assign. See [Credential item](#credential-item) below for details.
- `snmp_v2_write` - (Optional) The SNMPv2 Write credential to assign. See [Credential item](#credential-item) below for details.
- `snmp_v3` - (Optional) The SNMPv3 credential to assign. See [Credential item](#credential-item) below for details.

~> There is no delete operation for this resource on DNAC. Once assigned it only can be replaced.

### Credential item

Each credential item contains the following arguments.

- `id` - (Required) The credential's id.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `last_updated` - The updated time with format RFC850.

- `cli` - The CLI credential assigned. See [CLI item](#cli-item) below for details.
- `http_read` - The HTTP read credential assigned. See [HTTP Read item](#http-item) below for details.
- `http_write` - The HTTP write credential assigned. See [HTTP Write item](#http--item) below for details.
- `snmp_v2_read` - The SNMPv2 Read credential assigned. See [SNMPv2 Read item](#snmp-v2-read-item) below for details.
- `snmp_v2_write` - The SNMPv2 Write credential assigned. See [SNMPv2 Write item](#snmp-v2-write-item) below for details.
- `snmp_v3` - The SNMPv3 credential assigned. See [SNMPv3 item](#snmp-v3-item) below for details.

### cli item

- `comments` - The CLI credential's comments.
- `credential_type` - The CLI credential's credential type.
- `description` - The CLI credential's description.
- `enable_password` - The CLI credential's enable password.
- `id` - The CLI credential's id.
- `instance_tenant_id` - The CLI credential's instance tenant id.
- `instance_uuid` - The CLI credential's instance uuid.
- `password` - The CLI credential's password.
- `username` - The CLI credential's username.

### http item

- `comments` - The HTTP Read credential's comments.
- `credential_type` - The HTTP Read credential's credential type.
- `description` - The HTTP Read credential's description.
- `id` - The HTTP Read credential's id.
- `instance_tenant_id` - The HTTP Read credential's instance tenant id.
- `instance_uuid` - The HTTP Read credential's instance uuid.
- `password` - The HTTP Read credential's password.
- `port` - The HTTP Read credential's port.
- `secure` - The HTTP Read credential's secure flag.
- `username` - The HTTP Read credential's username.

### snmp v2 read item

- `comments` - The SNMPv2 read community credential's comments.
- `credential_type` - The SNMPv2 read community credential's credential type.
- `description` - The SNMPv2 read community credential's description.
- `id` - The SNMPv2 read community credential's id.
- `instance_tenant_id` - The SNMPv2 read community credential's instance tenant id.
- `instance_uuid` - The SNMPv2 read community credential's instance uuid.
- `read_community` - The SNMPv2 read community credential's community.

### snmp v2 write item

- `comments` - The SNMPv2 write community credential's comments.
- `credential_type` - The SNMPv2 write community credential's credential type.
- `description` - The SNMPv2 write community credential's description.
- `id` - The SNMPv2 write community credential's id.
- `instance_tenant_id` - The SNMPv2 write community credential's instance tenant id.
- `instance_uuid` - The SNMPv2 write community credential's instance uuid.
- `write_community` - The SNMPv2 write community credential's community.

### snmp v3 item

- `auth_password` - The SNMPv3 credential's auth_password.
- `auth_type` - The SNMPv3 credential's auth_type.
- `comments` - The SNMPv3 credential's comments.
- `credential_type` - The SNMPv3 credential's credential type.
- `description` - The SNMPv3 credential's description.
- `id` - The SNMPv3 credential's id.
- `instance_tenant_id` - The SNMPv3 credential's instance tenant id.
- `instance_uuid` - The SNMPv3 credential's instance uuid.
- `privacy_password` - The SNMPv3 credentials' privacy password.
- `privacy_type` - The SNMPv3 credentials' privacy type.
- `snmp_mode` - The SNMPv3 credentials' SNMP mode.
- `username` - The SNMPv3 credentials' username.
