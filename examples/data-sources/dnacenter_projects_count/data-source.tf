
data "dnacenter_projects_count" "example" {
  provider = dnacenter
  name     = "string"
}

output "dnacenter_projects_count_example" {
  value = data.dnacenter_projects_count.example.item
}
