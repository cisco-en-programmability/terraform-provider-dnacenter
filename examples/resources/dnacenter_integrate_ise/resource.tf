
resource "dnacenter_integrate_ise" "example" {
  provider = dnacenter
  id       = "string"
  parameters {

    is_cert_accepted_by_user = "false"
  }
}

output "dnacenter_integrate_ise_example" {
  value = dnacenter_integrate_ise.example
}