
data "dnacenter_application_sets" "example" {
  provider = dnacenter
  limit    = "#"
  name     = "string"
  offset   = "#"
}

output "dnacenter_application_sets_example" {
  value = data.dnacenter_application_sets.example.items
}
