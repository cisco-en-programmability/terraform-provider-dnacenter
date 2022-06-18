
resource "dnacenter_file_import" "example" {
  provider = dnacenter
  parameters {

    name_space = "string"
  }
}

output "dnacenter_file_import_example" {
  value = dnacenter_file_import.example
}