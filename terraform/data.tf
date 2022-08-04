data "aws_caller_identity" "current" {}

data "aws_s3_bucket" "lotusgo" {
  bucket = "lotusgogogo"
}

# output "account_id" {
#   value = data.aws_caller_identity.current.account_id
# }

# output "s3_bucket" {
#   value = data.aws_s3_bucket.lotusgo.bucket
# }

data "aws_route53_zone" "lotusgo" {
  name = "mldn.cc"  // 본 서비스에 적용할 루트 도메인.
}

data "aws_acm_certificate" "lotusgo" {
  domain      = "mldn.cc"
  types       = ["AMAZON_ISSUED"]
  most_recent = true
}