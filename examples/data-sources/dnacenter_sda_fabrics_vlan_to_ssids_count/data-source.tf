
data "dnacenter_sda_fabrics_vlan_to_ssids_count" "example" {
  provider = dnacenter
}

output "dnacenter_sda_fabrics_vlan_to_ssids_count_example" {
  value = data.dnacenter_sda_fabrics_vlan_to_ssids_count.example.item
}
