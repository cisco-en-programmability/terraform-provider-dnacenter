---
page_title: "dna_network Resource - terraform-provider-dnacenter"
subcategory: "Network Settings"
description: |-
  The dna_network resource allows you to configure a Cisco DNA Center network.
---

# Resource dna_network

The dna_network resource allows you to configure a Cisco DNA Center network.

## Example Usage

```hcl
resource "dna_network" "response1" {
  provider = dnacenter
  item {
    site_id = "a013dd15-69a3-423f-82dc-c6a10eba2cb7"
    # client_and_endpoint_aaa {
    #   ip_address    = ""
    #   network       = ""
    #   protocol      = ""
    #   servers       = "1,2"
    #   shared_secret = ""
    # }
    # dhcp_server = ["", "", ""]
    # dns_server {
    #   domain_name          = ""
    #   primary_ip_address   = ""
    #   secondary_ip_address = ""
    # }
    # message_of_theday {
    #   banner_message         = ""
    #   retain_existing_banner = true
    # }
    # netflowcollector {
    #   ip_address = ""
    #   # port = 0
    # }
    # network_aaa {
    #   ip_address    = ""
    #   network       = ""
    #   protocol      = ""
    #   servers       = "1,2"
    #   shared_secret = ""
    # }
    # ntp_server = []
    # snmp_server {
    #   configure_dnac_ip = true
    #   ip_addresses      = [""]
    # }
    # syslog_server {
    #   configure_dnac_ip = true
    #   ip_addresses      = [""]
    # }
    timezone = "UTC"
  }
}
```

## Argument Reference

- `item` - (Required) Item in a Cisco DNA Center Network. See [Network item](#network-item) below for details.

### Network item

- `site_id` - (Required) The network's site id. If it's changed it forces the creation of a new resource.
- `client_and_endpoint_aaa` - (Optional) The network's client and endpoint aaa. See [Client and endpoint aaa](#client-and-endpoint-aaa) below for details.
- `dhcp_server` - (Optional) The network's DHCP server.
- `dns_server` - (Optional) The network's DNS server. See [DNS server](#dns-server) below for details.
- `message_of_theday` - (Optional) The network's message_of_theday. See [Message of the day](#message-of-theday) below for details.
- `netflowcollector` - (Optional) The network's netflowcollector. See [Netflow collector](#netflow-collector) below for details.
- `network_aaa` - (Optional) The network's aaa. See [Network AAA](#network-aaa) below for details.
- `ntp_server` - (Optional) The network's NTP server.
- `snmp_server` - (Optional) The network's SNMP server. See [SNMP server](#snmp-server) below for details.
- `syslog_server` - (Optional) The network's Syslog server. See [Syslog server](#syslog-server) below for details.
- `timezone` - (Optional) The network's timezone.

### client and endpoint aaa

- `ip_address` - (Optional) The aaa's IP address.
- `network` - (Optional) The aaa's network.
- `protocol` - (Optional) The aaa's protocol.
- `servers` - (Optional) The aaa's servers.
- `shared_secret` - (Optional) The aaa's shared secret.

### dns server

- `domain_name` - (Optional) The DNS server's domain name.
- `primary_ip_address` - (Optional) The DNS server's primary IP address.
- `secondary_ip_address` - (Optional) The DNS server's secondary IP address.

### message of theday

- `banner_message` - (Optional) The banner_message.
- `retain_existing_banner` - (Optional) The retain existing banner flag.

### netflow collector

- `ip_address` - (Optional) The netflow collector's IP address.
- `port` - (Optional) The netflow collector's port.

### network aaa

- `ip_address` - (Optional) The network aaa's IP address.
- `network` - (Optional) The network aaa's network.
- `protocol` - (Optional) The network aaa's protocol.
- `servers` - (Optional) The network aaa's servers.
- `shared_secret` - (Optional) The network aaa's shared secret.

### snmp server

- `configure_dnac_ip` - (Optional) The server's configure DNAC IP.
- `ip_addresses` - (Optional) The server's IP addresses.

### syslog server

- `configure_dnac_ip` - (Optional) The server's configure DNAC IP.
- `ip_addresses` - (Optional) The server's IP addresses.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `last_updated` - The network's updated time with format RFC850.
