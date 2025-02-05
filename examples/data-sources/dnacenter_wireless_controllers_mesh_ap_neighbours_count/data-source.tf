
data "dnacenter_wireless_controllers_mesh_ap_neighbours_count" "example" {
  provider = dnacenter
}

output "dnacenter_wireless_controllers_mesh_ap_neighbours_count_example" {
  value = data.dnacenter_wireless_controllers_mesh_ap_neighbours_count.example.item
}
