
data "dnacsyslog_config_update" "example" {
  provider    = dnac
  config_id   = "string"
  description = "string"
  host        = "string"
  item {

    # api_status = ------
    error_message {

      # errors = [------]
    }
    # status_message = ------
  }
  name     = "string"
  port     = "string"
  protocol = "string"
}