
data "dnacevent_webhook_create" "example" {
  provider    = dnac
  description = "string"
  headers {

    default_value = "string"
    encrypt       = "false"
    name          = "string"
    value         = "string"
  }
  item {

    # api_status = ------
    error_message {

      # errors = [------]
    }
    # status_message = ------
  }
  method     = "string"
  name       = "string"
  trust_cert = "false"
  url        = "string"
  webhook_id = "string"
}