
data "dnacenter_dna_command_runner_keywords" "example" {
  provider = dnacenter
}

output "dnacenter_dna_command_runner_keywords_example" {
  value = data.dnacenter_dna_command_runner_keywords.example.items
}
