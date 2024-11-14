
data "dnacenter_sda_fabrics_vlan_to_ssids_fabric_id" "example" {
  provider  = dnacenter
  fabric_id = "string"
  limit     = 1
  offset    = 1
}

output "dnacenter_sda_fabrics_vlan_to_ssids_fabric_id_example" {
  value = data.dnacenter_sda_fabrics_vlan_to_ssids_fabric_id.example.items
}
