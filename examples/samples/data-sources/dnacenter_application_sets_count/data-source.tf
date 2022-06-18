
data "dnacenter_application_sets_count" "example" {
  provider = dnacenter
}

output "dnacenter_application_sets_count_example" {
  value = data.dnacenter_application_sets_count.example.item
}
