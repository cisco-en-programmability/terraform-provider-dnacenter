
data "dnacenter_sda_extranet_policies_count" "example" {
  provider = dnacenter
}

output "dnacenter_sda_extranet_policies_count_example" {
  value = data.dnacenter_sda_extranet_policies_count.example.item
}
