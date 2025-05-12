provider "aws" {
  region = var.aws_region
}

resource "aws_s3_bucket" "static_website" {
  bucket = var.s3_bucket_name

  tags = {
    Environment = var.environment
  }

  website {
    index_document = var.website_index_document
    error_document = var.website_error_document
  }

  # Enable bucket ownership enforcement
  ownership_controls {
    rule {
      object_ownership = "BucketOwnerEnforced"
    }
  }

  # Block all ACLs
  public_access_block {
    block_public_acls       = true
    block_public_policy     = false
    ignore_public_acls      = true
    restrict_public_buckets = false
  }
}

resource "aws_s3_bucket_policy" "static_website" {
  bucket = aws_s3_bucket.static_website.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid       = "PublicReadGetObject"
        Effect    = "Allow"
        Principal = "*"
        Action    = "s3:GetObject"
        Resource  = "${aws_s3_bucket.static_website.arn}/*"
      }
    ]
  })
}

resource "aws_s3_bucket_object" "static_website_files" {
  for_each = fileset("./static-site", "**/*")

  bucket = aws_s3_bucket.static_website.bucket
  key    = each.value
  source = "./static-site/${each.value}"
}

