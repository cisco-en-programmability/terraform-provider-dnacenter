
data "dnacenter_app_policy_queuing_profile_count" "example" {
  provider = dnacenter
}

output "dnacenter_app_policy_queuing_profile_count_example" {
  value = data.dnacenter_app_policy_queuing_profile_count.example.item
}
