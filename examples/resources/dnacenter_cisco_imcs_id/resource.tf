
resource "dnacenter_cisco_imcs_id" "example" {
  provider = dnacenter
  parameters {

    id         = "string"
    ip_address = "string"
    password   = "******"
    username   = "string"
  }
}

output "dnacenter_cisco_imcs_id_example" {
  value = dnacenter_cisco_imcs_id.example
}
