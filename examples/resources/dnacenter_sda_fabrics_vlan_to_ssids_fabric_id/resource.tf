
resource "dnacenter_sda_fabrics_vlan_to_ssids_fabric_id" "example" {
  provider = dnacenter

  parameters {

    fabric_id = "string"
    ssid_details {

      name               = "string"
      security_group_tag = "string"
    }
    vlan_name = "string"
  }
}

output "dnacenter_sda_fabrics_vlan_to_ssids_fabric_id_example" {
  value = dnacenter_sda_fabrics_vlan_to_ssids_fabric_id.example
}