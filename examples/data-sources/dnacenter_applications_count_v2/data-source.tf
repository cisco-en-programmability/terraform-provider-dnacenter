
data "dnacenter_applications_count_v2" "example" {
  provider            = dnacenter
  scalable_group_type = "string"
}

output "dnacenter_applications_count_v2_example" {
  value = data.dnacenter_applications_count_v2.example.item
}
