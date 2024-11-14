
data "dnacenter_applications_v2" "example" {
  provider   = dnacenter
  attributes = "string"
  limit      = 1
  name       = "string"
  offset     = 1
}

output "dnacenter_applications_v2_example" {
  value = data.dnacenter_applications_v2.example.items
}
