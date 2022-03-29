
resource "dnacenter_authentication_import_certificate" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  cert_file_path = "string"
  file_name      = "string"
  list_of_users  = ["string"]
  pk_file_path   = "string"
  pk_password    = "******"
}