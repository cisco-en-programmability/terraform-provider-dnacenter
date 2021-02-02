
terraform {
  required_providers {
    dnacenter = {
      versions = ["0.0.3"]
      source   = "hashicorp.com/edu/dnacenter"
    }
  }
}

provider "dnacenter" {
}

data "dna_network" "response" {
  provider = dnacenter
  # site_id  = "b85fe0be-d971-4eb7-92a9-3498356ad87f"
}
output "dna_network_response" {
  value = data.dna_network.response
}

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
output "dna_network_response1" {
  value = dna_network.response1
}
