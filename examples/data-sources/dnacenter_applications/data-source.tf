
data "dnacenter_applications" "example" {
  provider = dnacenter
  limit    = "#"
  name     = "string"
  offset   = "#"
}

output "dnacenter_applications_example" {
  value = data.dnacenter_applications.example.items
}
