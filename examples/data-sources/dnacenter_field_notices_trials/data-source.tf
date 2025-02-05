
data "dnacenter_field_notices_trials" "example" {
  provider = dnacenter
}

output "dnacenter_field_notices_trials_example" {
  value = data.dnacenter_field_notices_trials.example.item
}
