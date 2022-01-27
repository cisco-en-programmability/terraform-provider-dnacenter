
data "dnacenter_sda_virtual_network_v2" "example" {
  provider             = dnacenter
  virtual_network_name = "string"
}

output "dnacenter_sda_virtual_network_v2_example" {
  value = data.dnacenter_sda_virtual_network_v2.example.item
}
