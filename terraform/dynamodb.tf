resource "aws_dynamodb_table" "lotusgo" {
  name           = "lotusgo"
  billing_mode   = "PAY_PER_REQUEST"

  hash_key       = "id"
  range_key      = "user"

  attribute {
    name = "user"
    type = "S"
  }

  attribute {
    name = "id"
    type = "S"
  }
}