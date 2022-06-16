
data "dnacenter_app_policy_default" "example" {
  provider = dnacenter
}

output "dnacenter_app_policy_default_example" {
  value = data.dnacenter_app_policy_default.example.items
}
