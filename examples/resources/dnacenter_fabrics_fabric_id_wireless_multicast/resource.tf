
resource "dnacenter_fabrics_fabric_id_wireless_multicast" "example" {
  provider = dnacenter

  parameters {

    fabric_id         = "string"
    multicast_enabled = "false"
  }
}

output "dnacenter_fabrics_fabric_id_wireless_multicast_example" {
  value = dnacenter_fabrics_fabric_id_wireless_multicast.example
}
