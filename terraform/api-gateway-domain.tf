resource "aws_api_gateway_domain_name" "go-note-api" {
  domain_name              = "test22-api.mldn.cc"  // api 엔드포인트로 사용할 서브도메인
  regional_certificate_arn = data.aws_acm_certificate.lotusgo.arn

  endpoint_configuration {
    types = ["REGIONAL"]
  }
}

resource "aws_route53_record" "go-note-api" {
  name    = aws_api_gateway_domain_name.go-note-api.domain_name
  type    = "A"
  zone_id = data.aws_route53_zone.lotusgo.id

  alias {
    evaluate_target_health = true
    name                   = aws_api_gateway_domain_name.go-note-api.regional_domain_name
    zone_id                = aws_api_gateway_domain_name.go-note-api.regional_zone_id
  }
}

resource "aws_api_gateway_base_path_mapping" "go-note-api" {
  api_id      = aws_api_gateway_rest_api.go-note-api.id
  stage_name  = "v1"  // stage 이름을 v1로 정했었다.
  domain_name = aws_api_gateway_domain_name.go-note-api.domain_name
}