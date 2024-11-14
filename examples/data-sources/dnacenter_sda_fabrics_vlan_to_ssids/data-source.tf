
data "dnacenter_sda_fabrics_vlan_to_ssids" "example" {
  provider = dnacenter
  limit    = 1
  offset   = 1
}

output "dnacenter_sda_fabrics_vlan_to_ssids_example" {
  value = data.dnacenter_sda_fabrics_vlan_to_ssids.example.items
}
