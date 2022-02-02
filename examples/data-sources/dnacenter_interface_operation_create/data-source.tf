
data "dnacinterface_operation_create" "example" {
  provider         = dnac
  deployement_mode = "string"
  interface_uuid   = "string"
  item {

    # task_id = ------
    # url = ------
  }
  operation = "string"
  payload   = ["string"]
}