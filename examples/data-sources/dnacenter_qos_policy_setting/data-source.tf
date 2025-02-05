
data "dnacenter_qos_policy_setting" "example" {
  provider = dnacenter
}

output "dnacenter_qos_policy_setting_example" {
  value = data.dnacenter_qos_policy_setting.example.item
}
