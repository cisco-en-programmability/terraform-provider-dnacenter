
resource "dnacenter_diagnostic_validation_workflows" "example" {
  provider = dnacenter

  parameters {

    description        = "string"
    id                 = "string"
    name               = "string"
    validation_set_ids = ["string"]
  }
}

output "dnacenter_diagnostic_validation_workflows_example" {
  value = dnacenter_diagnostic_validation_workflows.example
}