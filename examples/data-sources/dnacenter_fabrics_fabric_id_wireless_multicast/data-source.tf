
data "dnacenter_fabrics_fabric_id_wireless_multicast" "example" {
  provider  = dnacenter
  fabric_id = "string"
}

output "dnacenter_fabrics_fabric_id_wireless_multicast_example" {
  value = data.dnacenter_fabrics_fabric_id_wireless_multicast.example.item
}
