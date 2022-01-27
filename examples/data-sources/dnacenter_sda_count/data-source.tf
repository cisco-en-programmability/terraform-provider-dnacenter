
data "dnacenter_sda_count" "example" {
  provider = dnacenter
}

output "dnacenter_sda_count_example" {
  value = data.dnacenter_sda_count.example.item
}
