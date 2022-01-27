
data "dnacenter_applications_count" "example" {
  provider = dnacenter
}

output "dnacenter_applications_count_example" {
  value = data.dnacenter_applications_count.example.item
}
