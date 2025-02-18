resource "fmc_application_filter" "example" {
  name = "my_application_filter"
  applications = [
    {
      id   = "535"
      name = "1-800-Flowers"
    }
  ]
  filters = [
    {
      types = [
        {
          id = "WEBAPP"
        }
      ]
      risks = [
        {
          id = "VERY_LOW"
        }
      ]
      business_relevances = [
        {
          id = "LOW"
        }
      ]
      categories = [
        {
          id = "118"
        }
      ]
      tags = [
        {
          id = "24"
        }
      ]
    }
  ]
}
