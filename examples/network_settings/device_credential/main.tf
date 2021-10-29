terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source   = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}

data "dna_network_device_credential" "response" {
  provider = dnacenter
  site_id  = "2397da83-4e12-4d04-9bd3-a57b2ad91652"
}

output "dna_network_device_credential_response" {
  value = data.dna_network_device_credential.response
}
