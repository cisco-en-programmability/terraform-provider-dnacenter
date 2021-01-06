output "applications_amount" {
  value       = data.dna_applications_count.amount.response
  description = "The amount of applications"
}

output "applications_list" {
  value       = data.dna_applications.list.items
  description = "The list of applications"
}

output "applications_response" {
  value       = dna_applications.response.items
  description = "The result of the application resource"
}
