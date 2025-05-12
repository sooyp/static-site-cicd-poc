# Define a variable for the environment name ('dev', 'prod')
variable "environment" {
  description = "The environment for the deployment"
  type        = string
  default     = "dev"
}

# Define a variable for the AWS region
variable "aws_region" {
  description = "The AWS region where resources will be provisioned"
  type        = string
}

# Define a variable for the S3 bucket name
variable "s3_bucket_name" {
  description = "The name of the S3 bucket to be created"
  type        = string
}

# Define a variable for the static website index document
variable "website_index_document" {
  description = "The index document for the static website"
  type        = string
  default     = "index.html"
}

# Define a variable for the static website error document
variable "website_error_document" {
  description = "The error document for the static website"
  type        = string
  default     = "error.html"
}
