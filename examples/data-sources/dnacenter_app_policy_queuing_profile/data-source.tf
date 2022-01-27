
data "dnacenter_app_policy_queuing_profile" "example" {
  provider = dnacenter
  name     = "string"
}

output "dnacenter_app_policy_queuing_profile_example" {
  value = data.dnacenter_app_policy_queuing_profile.example.items
}
