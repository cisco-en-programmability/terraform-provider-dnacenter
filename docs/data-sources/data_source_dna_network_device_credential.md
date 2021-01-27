---
page_title: "dna_network_device_credential Data Source - terraform-provider-dnacenter"
subcategory: "Network Settings"
description: |-
  The dna_network_device_credential data source allows you to retrieve information about a particular Cisco DNA Center network device credentials.
---

# Data Source dna_network_device_credential

The dna_network_device_credential data source allows you to retrieve information about a particular Cisco DNA Center network device credentials.

## Example Usage

```hcl
data "dna_network_device_credential" "response" {
  provider = dnacenter
  # site_id  = "b85fe0be-d971-4eb7-92a9-3498356ad87f"
}
```

## Argument Reference

- `site_id` - (Optional) The site_id param.

## Attributes Reference

The following attributes are exported.

- `items` - Items in a Cisco DNA Center network device credential. See [Items](#items) below for details.

### Items

Each item contains the following attributes.

- `cli` - The CLI credentials. See [cli](#cli) below for details.
- `http_read` - The HTTP Read credentials. See [http_read](#http_read) below for details.
- `http_write` - The HTTP Write credentials. See [http_write](#http_write) below for details.
- `snmp_v2_read` - The SNMPv2 Read credentials. See [snmp_v2_read](#snmp_v2_read) below for details.
- `snmp_v2_write` - The SNMPv2 Write credentials. See [snmp_v2_write](#snmp_v2_write) below for details.
- `snmp_v3` - The SNMPv3 credentials. See [snmp_v3](#snmp_v3) below for details.

#### cli

Each credential contains the following attributes.

comments
credential_type
description
enable_password
id
instance_tenant_id
instance_uuid
password
username

#### http_read

Each credential contains the following attributes.

- `comments` - The credential's comments.
- `credential_type` - The credential's type.
- `description` - The credential's description.
- `id` - The credential's id.
- `instance_tenant_id` - The credential's instance tenant id.
- `instance_uuid` - The credential's instance uuid.
- `password` - The credential's password.
- `port` - The credential's port.
- `secure` - The credential's secure.
- `username` - The credential's username.

#### http_write

Each credential contains the following attributes.

comments

- `credential_type` - The credential's type.
- `description` - The credential's description.
- `id` - The credential's id.
- `instance_tenant_id` - The credential's instance tenant id.
- `instance_uuid` - The credential's instance uuid.
- `password` - The credential's password.
- `port` - The credential's port.
- `secure` - The credential's secure.
- `username` - The credential's username.

#### snmp_v2_read

Each credential contains the following attributes.

- `comments` - The credential's comments.
- `credential_type` - The credential's type.
- `description` - The credential's description.
- `id` - The credential's id.
- `instance_tenant_id` - The credential's instance tenant id.
- `instance_uuid` - The credential's instance uuid.
- `read_community` - The credential's read community.

#### snmp_v2_write

Each credential contains the following attributes.

- `comments` - The credential's comments.
- `credential_type` - The credential's type.
- `description` - The credential's description.
- `id` - The credential's id.
- `instance_tenant_id` - The credential's instance tenant id.
- `instance_uuid` - The credential's instance uuid.
- `write_community` - The credential's write community.

#### snmp_v3

- `auth_password` - The credential's auth password.
- `auth_type` - The credential's auth type.
- `comments` - The credential's comments.
- `credential_type` - The credential's type.
- `description` - The credential's description.
- `id` - The credential's id.
- `instance_tenant_id` - The credential's instance tenant id.
- `instance_uuid` - The credential's instance uuid.
- `privacy_password` - The credential's privacy password.
- `privacy_type` - The credential's privacy type.
- `snmp_mode` - The credential's snmp mode.
- `username` - The credential's username.
