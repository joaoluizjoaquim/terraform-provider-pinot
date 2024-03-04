terraform {
  required_providers {
    pinot = {
      source = "hashicorp.com/edu/pinot"
    }
  }
}



provider "pinot" {
  controller_url     = "http://localhost:9000"
}

resource "pinot_schema" "block_schema" {
  schema_name = "ethereum_block_headers"
  schema = file("block_schema.json")
}
