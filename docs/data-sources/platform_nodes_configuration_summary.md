---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_platform_nodes_configuration_summary Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Platform Configuration.
  Provides details about the current Cisco DNA Center node configuration, such as API version, node name, NTP server,
  intracluster link, LACP mode, network static routes, DNS server, subnet mask, host IP, default gateway, and interface
  information.
---

# dnacenter_platform_nodes_configuration_summary (Data Source)

It performs read operation on Platform Configuration.

- Provides details about the current Cisco DNA Center node configuration, such as API version, node name, NTP server,
intracluster link, LACP mode, network static routes, DNS server, subnet mask, host IP, default gateway, and interface
information.

## Example Usage

```terraform
data "dnacenter_platform_nodes_configuration_summary" "example" {
  provider = dnacenter
}

output "dnacenter_platform_nodes_configuration_summary_example" {
  value = data.dnacenter_platform_nodes_configuration_summary.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `nodes` (List of Object) (see [below for nested schema](#nestedobjatt--item--nodes))

<a id="nestedobjatt--item--nodes"></a>
### Nested Schema for `item.nodes`

Read-Only:

- `id` (String)
- `name` (String)
- `network` (List of Object) (see [below for nested schema](#nestedobjatt--item--nodes--network))
- `ntp` (List of Object) (see [below for nested schema](#nestedobjatt--item--nodes--ntp))
- `platform` (List of Object) (see [below for nested schema](#nestedobjatt--item--nodes--platform))
- `proxy` (List of Object) (see [below for nested schema](#nestedobjatt--item--nodes--proxy))

<a id="nestedobjatt--item--nodes--network"></a>
### Nested Schema for `item.nodes.network`

Read-Only:

- `inet` (List of Object) (see [below for nested schema](#nestedobjatt--item--nodes--network--inet))
- `inet6` (List of Object) (see [below for nested schema](#nestedobjatt--item--nodes--network--inet6))
- `interface` (String)
- `intra_cluster_link` (String)
- `lacp_mode` (String)
- `lacp_supported` (String)
- `slave` (List of String)

<a id="nestedobjatt--item--nodes--network--inet"></a>
### Nested Schema for `item.nodes.network.slave`

Read-Only:

- `dns_servers` (List of String)
- `gateway` (String)
- `host_ip` (String)
- `netmask` (String)
- `routes` (List of String)


<a id="nestedobjatt--item--nodes--network--inet6"></a>
### Nested Schema for `item.nodes.network.slave`

Read-Only:

- `host_ip` (String)
- `netmask` (String)



<a id="nestedobjatt--item--nodes--ntp"></a>
### Nested Schema for `item.nodes.ntp`

Read-Only:

- `servers` (List of String)


<a id="nestedobjatt--item--nodes--platform"></a>
### Nested Schema for `item.nodes.platform`

Read-Only:

- `product` (String)
- `serial` (String)
- `vendor` (String)


<a id="nestedobjatt--item--nodes--proxy"></a>
### Nested Schema for `item.nodes.proxy`

Read-Only:

- `http_proxy` (String)
- `https_proxy` (String)
- `https_proxy_password` (String)
- `https_proxy_username` (String)
- `no_proxy` (List of String)
