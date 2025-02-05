
resource "dnacenter_qos_policy_setting" "example" {
  provider = dnacenter

  parameters {

    deploy_by_default_on_wired_devices = "false"
  }
}

output "dnacenter_qos_policy_setting_example" {
  value = dnacenter_qos_policy_setting.example
}
