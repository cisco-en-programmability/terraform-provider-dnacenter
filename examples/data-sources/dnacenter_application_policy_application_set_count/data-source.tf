
data "dnacenter_application_policy_application_set_count" "example" {
  provider            = dnacenter
  scalable_group_type = "string"
}

output "dnacenter_application_policy_application_set_count_example" {
  value = data.dnacenter_application_policy_application_set_count.example.item
}
