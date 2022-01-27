
data "dnacenter_app_policy" "example" {
  provider     = dnacenter
  policy_scope = "string"
}

output "dnacenter_app_policy_example" {
  value = data.dnacenter_app_policy.example.items
}
