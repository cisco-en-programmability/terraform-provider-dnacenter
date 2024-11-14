
data "dnacenter_diagnostic_validation_sets" "example" {
  provider = dnacenter
  view     = "string"
}

output "dnacenter_diagnostic_validation_sets_example" {
  value = data.dnacenter_diagnostic_validation_sets.example.items
}

data "dnacenter_diagnostic_validation_sets" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_diagnostic_validation_sets_example" {
  value = data.dnacenter_diagnostic_validation_sets.example.item
}
