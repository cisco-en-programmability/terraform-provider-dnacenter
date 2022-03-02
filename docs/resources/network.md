---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_network Resource - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It manages create, read and update operations on Network Settings.
  API to create a network for DHCP and DNS center server settings.API to update network for DHCP and DNS center server settings.
---

# dnacenter_network (Resource)

It manages create, read and update operations on Network Settings.

- API to create a network for DHCP and DNS center server settings.

- API to update network for DHCP and DNS center server settings.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **parameters** (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--parameters))

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **last_updated** (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Required:

- **site_id** (String) siteId path parameter. Site id to update the network settings which is associated with the site

Optional:

- **settings** (Block List, Max: 1) (see [below for nested schema](#nestedblock--parameters--settings))

<a id="nestedblock--parameters--settings"></a>
### Nested Schema for `parameters.settings`

Optional:

- **client_and_endpoint_aaa** (Block List, Max: 1) (see [below for nested schema](#nestedblock--parameters--settings--client_and_endpoint_aaa))
- **dhcp_server** (List of String) Dhcp serve Ip (eg: 1.1.1.1)
- **dns_server** (Block List, Max: 1) (see [below for nested schema](#nestedblock--parameters--settings--dns_server))
- **message_of_theday** (Block List, Max: 1) (see [below for nested schema](#nestedblock--parameters--settings--message_of_theday))
- **netflowcollector** (Block List, Max: 1) (see [below for nested schema](#nestedblock--parameters--settings--netflowcollector))
- **network_aaa** (Block List, Max: 1) (see [below for nested schema](#nestedblock--parameters--settings--network_aaa))
- **ntp_server** (List of String) IP address for NTP server (eg: 1.1.1.2)
- **snmp_server** (Block List, Max: 1) (see [below for nested schema](#nestedblock--parameters--settings--snmp_server))
- **syslog_server** (Block List, Max: 1) (see [below for nested schema](#nestedblock--parameters--settings--syslog_server))
- **timezone** (String) Input for time zone (eg: Africa/Abidjan)

<a id="nestedblock--parameters--settings--client_and_endpoint_aaa"></a>
### Nested Schema for `parameters.settings.client_and_endpoint_aaa`

Optional:

- **ip_address** (String) IP address for ISE serve (eg: 1.1.1.4). Mandatory for ISE servers.
- **network** (String) IP address for AAA or ISE server (eg: 2.2.2.1)
- **protocol** (String) Protocol for AAA or ISE serve (eg: RADIUS)
- **servers** (String) Server type AAA or ISE server (eg: AAA)
- **shared_secret** (String) Shared secret for ISE server. Supported only by ISE servers


<a id="nestedblock--parameters--settings--dns_server"></a>
### Nested Schema for `parameters.settings.dns_server`

Optional:

- **domain_name** (String) Domain name of DHCP (eg; cisco). It can only contain alphanumeric characters or hyphen.
- **primary_ip_address** (String) Primary ip address for DHCP (eg: 2.2.2.2). valid range : 1.0.0.0 - 223.255.255.255
- **secondary_ip_address** (String) Secondary ip address for DHCP (eg: 3.3.3.3). valid range : 1.0.0.0 - 223.255.255.255


<a id="nestedblock--parameters--settings--message_of_theday"></a>
### Nested Schema for `parameters.settings.message_of_theday`

Optional:

- **banner_message** (String) Massage for banner message (eg; Good day)
- **retain_existing_banner** (String) Retain existing banner message (eg: "true" or "false")


<a id="nestedblock--parameters--settings--netflowcollector"></a>
### Nested Schema for `parameters.settings.netflowcollector`

Optional:

- **ip_address** (String) IP address for netflow collector (eg: 3.3.3.1)
- **port** (Number) Port for netflow collector (eg; 443)


<a id="nestedblock--parameters--settings--network_aaa"></a>
### Nested Schema for `parameters.settings.network_aaa`

Optional:

- **ip_address** (String) IP address for AAA and ISE server (eg: 1.1.1.1). Mandatory for ISE servers and for AAA consider this as additional Ip.
- **network** (String) IP address for AAA or ISE server (eg: 2.2.2.2). For AAA server consider it as primary IP and For ISE consider as Network
- **protocol** (String) Protocol for AAA or ISE serve (eg: RADIUS)
- **servers** (String) Server type for AAA network (eg: AAA)
- **shared_secret** (String) Shared secret for ISE server. Supported only by ISE servers


<a id="nestedblock--parameters--settings--snmp_server"></a>
### Nested Schema for `parameters.settings.snmp_server`

Optional:

- **configure_dnac_ip** (String) Configuration dnac ip for snmp server (eg: true)
- **ip_addresses** (List of String) IP address for snmp server (eg: 4.4.4.1)


<a id="nestedblock--parameters--settings--syslog_server"></a>
### Nested Schema for `parameters.settings.syslog_server`

Optional:

- **configure_dnac_ip** (String) Configuration dnac ip for syslog server (eg: true)
- **ip_addresses** (List of String) IP address for syslog server (eg: 4.4.4.4)

