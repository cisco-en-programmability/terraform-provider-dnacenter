
terraform {
  required_providers {
    dnacenter = {
      version = "1.3.1-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_snmp_properties" "example" {
  provider = dnacenter
  parameters {

    #id                   = "string"
    #instance_tenant_id   = "string"
    #instance_uuid        = "string"
    #int_value            = 1
    system_property_name = "Timeout"
  }
}

output "dnacenter_snmp_properties_example" {
  value = dnacenter_snmp_properties.example
}
