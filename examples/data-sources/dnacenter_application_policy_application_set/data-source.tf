
data "dnacenter_application_policy_application_set" "example" {
  provider   = dnacenter
  attributes = "string"
  limit      = 1
  name       = "string"
  offset     = 1
}

output "dnacenter_application_policy_application_set_example" {
  value = data.dnacenter_application_policy_application_set.example.items
}
