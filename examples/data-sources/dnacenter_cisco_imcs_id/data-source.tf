
data "dnacenter_cisco_imcs_id" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_cisco_imcs_id_example" {
  value = data.dnacenter_cisco_imcs_id.example.item
}
