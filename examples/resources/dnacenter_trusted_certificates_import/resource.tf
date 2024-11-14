
resource "dnacenter_trusted_certificates_import" "example" {
  provider = dnacenter
}

output "dnacenter_trusted_certificates_import_example" {
  value = dnacenter_trusted_certificates_import.example
}