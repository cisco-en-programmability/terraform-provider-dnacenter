
data "dnacenter_fabrics_fabric_id_switch_wireless_setting" "example" {
  provider  = dnacenter
  fabric_id = "string"
}

output "dnacenter_fabrics_fabric_id_switch_wireless_setting_example" {
  value = data.dnacenter_fabrics_fabric_id_switch_wireless_setting.example.items
}
