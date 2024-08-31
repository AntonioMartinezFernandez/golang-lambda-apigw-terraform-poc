resource "aws_dynamodb_table" "users_table" {
  name = "users"

  read_capacity  = 5
  write_capacity = 5
  stream_enabled = false

  hash_key = "Id"
  # range_key = "Name"

  attribute {
    name = "Id"
    type = "S"
  }

  # attribute {
  #   name = "Name"
  #   type = "S"
  # }
}
