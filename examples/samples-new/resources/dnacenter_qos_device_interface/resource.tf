terraform {
  required_providers {
  dnacenter = {
    version = "0.0.3"
    source  = "hashicorp.com/edu/dnacenter"
    # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
  }
  }
}

resource "dnacenter_qos_device_interface" "example" {
  provider = dnacenter
  parameters {

    excluded_interfaces = ["a"]
    #id                  = "string"
    name                = "test"
    network_device_id   = "3eb928b8-2414-4121-ac35-1247e5d666a4"
    qos_device_interface_info {

      #dmvpn_remote_sites_bw = [1]
      #instance_id           = 1
      #interface_id          = "string"
      interface_name        = "a"
      #label                 = "string"
      role                  = "DMVPN_SPOKE"
      #upload_bw             = 1
    }
  }
}

output "dnacenter_qos_device_interface_example" {
  value = dnacenter_qos_device_interface.example
}