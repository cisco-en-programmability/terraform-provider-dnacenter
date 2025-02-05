
data "dnacenter_wireless_controllers_mesh_ap_neighbours" "example" {
  provider             = dnacenter
  ap_name              = "string"
  ethernet_mac_address = "string"
  wlc_ip_address       = "string"
}

output "dnacenter_wireless_controllers_mesh_ap_neighbours_example" {
  value = data.dnacenter_wireless_controllers_mesh_ap_neighbours.example.item
}
