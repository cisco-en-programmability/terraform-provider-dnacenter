
resource "dnacenter_security_advisories_trials" "example" {
  provider = dnacenter

}

output "dnacenter_security_advisories_trials_example" {
  value = dnacenter_security_advisories_trials.example
}
