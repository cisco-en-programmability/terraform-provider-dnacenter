
resource "dnacenter_cisco_imcs" "example" {
  provider = dnacenter

  parameters {

    ip_address = "string"
    node_id    = "string"
    password   = "******"
    username   = "string"
  }
}

output "dnacenter_cisco_imcs_example" {
  value = dnacenter_cisco_imcs.example
}
