
data "dnacenter_sda_virtual_network_ip_pool" "example" {
  provider             = dnacenter
  ip_pool_name         = "string"
  virtual_network_name = "string"
}

output "dnacenter_sda_virtual_network_ip_pool_example" {
  value = data.dnacenter_sda_virtual_network_ip_pool.example.item
}
