
terraform {
  required_providers {
    dnacenter = {
      versions = ["0.2"]
      source   = "hashicorp.com/edu/dnacenter"
    }
  }
}

provider "dnacenter" {
}

resource "dna_network_service_provider_profile" "response1" {
  provider     = dnacenter
  profile_name = "Test1"
  model        = "6-class-model"
  wan_provider = "test1-provider"
}
output "dna_network_service_provider_profile_response1" {
  value = dna_network_service_provider_profile.response1
}


data "dna_network_global_ip_pool" "response" {
  provider = dnacenter
}
output "dna_network_global_ip_pool_response" {
  value = data.dna_network_global_ip_pool.response
}

data "dna_network_service_provider_profile" "response" {
  provider = dnacenter
}
output "dna_network_service_provider_profile_response" {
  value = data.dna_network_service_provider_profile.response
}


data "dna_network_device_credential" "response" {
  provider = dnacenter
  # site_id  = "b85fe0be-d971-4eb7-92a9-3498356ad87f"
}
output "dna_network_device_credential_response" {
  value = data.dna_network_device_credential.response
}


data "dna_network" "response" {
  provider = dnacenter
  # site_id  = "b85fe0be-d971-4eb7-92a9-3498356ad87f"
}
output "dna_network_response" {
  value = data.dna_network.response
}

resource "dna_network_global_ip_pool" "response1" {
  provider         = dnacenter
  type             = "Generic"
  gateway          = ""
  ip_address_space = "IPv4"
  item {
    id             = "22f70f75-5dae-4494-9965-d4b85e101898"
    ip_pool_name   = "dna-usa"
    dns_server_ips = ["34.245.38.218"]
    ip_pool_cidr   = "10.64.0.0/12"
  }
}
output "dna_network_global_ip_pool_response1" {
  value = dna_network_global_ip_pool.response1
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

resource "dna_network_credential_site_assignment" "response1" {
  provider = dnacenter
  site_id  = "b85fe0be-d971-4eb7-92a9-3498356ad87f"
  http_read {
    id = "babc42b9-0bdd-49ef-912e-66f533fb5d59"
  }
  cli {
    id = "f979d842-f6fd-456a-8137-2cb5113cd2e8"
  }
}
output "dna_network_credential_site_assignment_response1" {
  value = dna_network_credential_site_assignment.response1
}
