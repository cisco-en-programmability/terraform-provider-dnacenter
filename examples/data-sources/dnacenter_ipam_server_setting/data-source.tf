
data "dnacenter_ipam_server_setting" "example" {
  provider = dnacenter
}

output "dnacenter_ipam_server_setting_example" {
  value = data.dnacenter_ipam_server_setting.example.item
}
