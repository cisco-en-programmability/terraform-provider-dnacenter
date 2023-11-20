
terraform {
  required_providers {
    dnacenter = {
      version = "1.1.27-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_global_credential_snmpv2_read_community" "example" {
  provider = dnacenter
  parameters {
    description     = "Description 4"
    comments        = "New Comments"
    credential_type = "APP"
    read_community  = "Test4"
  }
}

output "dnacenter_global_credential_snmpv2_read_community_example" {
  value = dnacenter_global_credential_snmpv2_read_community.example
}
