
data "dnacenter_sda_virtual_network" "example" {
  provider             = dnacenter
  site_name_hierarchy  = "string"
  virtual_network_name = "string"
}

output "dnacenter_sda_virtual_network_example" {
  value = data.dnacenter_sda_virtual_network.example.item
}
