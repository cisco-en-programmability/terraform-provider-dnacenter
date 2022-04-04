
resource "dnacenter_wireless_provision_device_update" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    payload{
    device_name = "string"
    dynamic_interfaces {

      interface_gateway          = "string"
      interface_ipaddress        = "string"
      interface_name             = "string"
      interface_netmask_in_cid_r = 1
      lag_or_port_number         = 1
      vlan_id                    = 1
    }
    managed_aplocations = ["string"]
  }
  }
}