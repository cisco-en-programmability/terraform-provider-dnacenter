
resource "dnacenter_site_kpi_summaries_query_count" "example" {
  provider    = dnacenter
  xca_lle_rid = "string"
  parameters = [{

    end_time = 1
    filters = [{

      key      = "string"
      operator = "string"
      value    = "string"
    }]
    start_time = 1
  }]
}

output "dnacenter_site_kpi_summaries_query_count_example" {
  value = dnacenter_site_kpi_summaries_query_count.example
}
