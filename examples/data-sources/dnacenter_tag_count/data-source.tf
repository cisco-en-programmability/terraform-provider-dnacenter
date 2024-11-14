
data "dnacenter_tag_count" "example" {
  provider       = dnacenter
  attribute_name = "string"
  name           = "string"
  name_space     = "string"
  size           = "string"
  system_tag     = "string"
}

output "dnacenter_tag_count_example" {
  value = data.dnacenter_tag_count.example.item
}
