
data "dnacenter_network_bugs_trials" "example" {
  provider = dnacenter
}

output "dnacenter_network_bugs_trials_example" {
  value = data.dnacenter_network_bugs_trials.example.item
}
