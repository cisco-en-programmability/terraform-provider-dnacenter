
resource "dnacenter_site_kpi_summaries_query" "example" {
  provider    = dnacenter
  xca_lle_rid = "string"
  parameters = [{

    attributes = ["string"]
    end_time   = 1
    filters = [{

      key      = "string"
      operator = "string"
      value    = "string"
    }]
    page = [{

      limit  = 1
      offset = 1
      sort_by = [{

        name  = "string"
        order = "string"
      }]
    }]
    start_time = 1
    views      = ["string"]
  }]
}

output "dnacenter_site_kpi_summaries_query_example" {
  value = dnacenter_site_kpi_summaries_query.example
}
