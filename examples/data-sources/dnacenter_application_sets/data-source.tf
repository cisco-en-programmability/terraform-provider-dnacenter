
data "dnacenter_application_sets" "example" {
  provider = dnacenter
  limit    = 1
  name     = "string"
  offset   = 1
}

output "dnacenter_application_sets_example" {
  value = data.dnacenter_application_sets.example.items
}
