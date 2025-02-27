
terraform {
  required_providers {
    dnacenter = {
      version = "1.3.1-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

resource "dnacenter_interface_update" "example" {
  provider = dnacenter
  parameters {

    //admin_status = "string"
    description    = "test"
    interface_uuid = "c6820b57-ecde-4b6d-98db-06ba10486809"
    vlan_id        = 2
    //voice_vlan_id = 1
  }
}

output "dnacenter_interface_update_example" {
  value = dnacenter_interface_update.example
}
