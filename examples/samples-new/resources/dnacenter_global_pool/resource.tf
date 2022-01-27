
terraform {
  required_providers {
  dnacenter = {
    version = "0.0.3"
    source  = "hashicorp.com/edu/dnacenter"
    # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
  }
  }
}

resource "dnacenter_global_pool" "example" {
  provider = dnacenter
  parameters {

    #id = "string"
    settings {
      ippool {
        #ip_address_space = "string"
        dhcp_server_ips  = ["100.100.100.100"]
        dns_server_ips   = ["101.101.101.101"]
        gateway          = "13s.0.0.1"
        #id               = "string"
        ip_pool_cidr     = "14.0.0.0/8"
        ip_pool_name     = "13Network"
        type             = "generic"
      }
    }
  }
}

output "dnacenter_global_pool_example" {
  value = dnacenter_global_pool.example
}