
resource "dnacenter_application_sets" "example" {
  provider = dnacenter

  parameters {
    payload {
      name = "string"
    }
  }
}

output "dnacenter_application_sets_example" {
  value = dnacenter_application_sets.example
}