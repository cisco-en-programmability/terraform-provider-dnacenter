
data "dnacenter_applications" "example" {
  provider = dnacenter
  limit    = 1
  name     = "string"
  offset   = 1
}

output "dnacenter_applications_example" {
  value = data.dnacenter_applications.example.items
}
