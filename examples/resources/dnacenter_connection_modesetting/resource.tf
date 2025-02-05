
resource "dnacenter_connection_modesetting" "example" {
  provider = dnacenter

  parameters {

    connection_mode = "string"
    parameters {

      client_id          = "string"
      client_secret      = "string"
      on_premise_host    = "string"
      smart_account_name = "string"
    }
  }
}

output "dnacenter_connection_modesetting_example" {
  value = dnacenter_connection_modesetting.example
}
