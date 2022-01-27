output "dna_application_set_response" {
  value       = dna_application_set.response.item
  description = "The dna_application_set resource's response"
}

output "dna_application_set_query" {
  value       = data.dna_application_set.query.items
  description = "The dna_application_set data source's response"
}
