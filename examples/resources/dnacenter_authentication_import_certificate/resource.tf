
resource "dnacenter_authentication_import_certificate" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    cert_file_path = "string"
    cert_file_name = "string"
    pk_file_name   = "string"
    list_of_users  = ["string"]
    pk_file_path   = "string"
    pk_password    = "******"
  }
}