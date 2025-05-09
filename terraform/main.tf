provider "aws" {
  region = var.aws_region
}

resource "aws_s3_bucket" "static_website" {
  bucket = var.s3_bucket_name
  acl    = "public-read"

  website {
    index_document = var.website_index_document
    # Specify the error document for the website
    error_document = var.website_error_document
  }
}

resource "aws_s3_bucket_object" "static_website_files" {
  for_each = fileset("./static-site", "**/*")

  bucket = aws_s3_bucket.static_website.bucket
  key    = each.value
  source = "./static-site/${each.value}"
  acl    = "public-read"
}
