
data "dnacenter_sda_fabrics_vlan_to_ssids_fabric_id_count" "example" {
  provider  = dnacenter
  fabric_id = "string"
}

output "dnacenter_sda_fabrics_vlan_to_ssids_fabric_id_count_example" {
  value = data.dnacenter_sda_fabrics_vlan_to_ssids_fabric_id_count.example.item
}
