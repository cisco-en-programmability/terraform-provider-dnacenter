
data "dnacenter_cisco_imcs" "example" {
  provider = dnacenter
}

output "dnacenter_cisco_imcs_example" {
  value = data.dnacenter_cisco_imcs.example.items
}
