
resource "dnacenter_file_import" "example" {
  provider   = dnacenter
  file_name  = "string"
  file_path  = "string"
  name_space = "string"
  parameters {

  }
}

output "dnacenter_file_import_example" {
  value = dnacenter_file_import.example
}