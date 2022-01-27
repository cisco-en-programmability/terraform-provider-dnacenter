
terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}

data "dna_network_global_ip_pool" "response" {
  provider = dnacenter
}
output "dna_network_global_ip_pool_response" {
  value = data.dna_network_global_ip_pool.response
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
