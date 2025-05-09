# static-site-cicd-poc
This project is a hands-on Proof of Concept (POC) showcasing a complete CI/CD pipeline that provisions infrastructure and deploys a static website to an AWS S3 bucket using Terraform, GitHub Actions, and Terratest.

## Project Overview

Infrastructure-as-Code: AWS S3 bucket provisioned via Terraform

CI/CD Automation: GitHub Actions for Terraform operations, site upload, and testing

Validation: Terratest ensures infrastructure is correctly deployed

---
.
├── .github/workflows/     # CI pipeline definition
├── terraform/             # Terraform code to provision S3
├── static-site/           # HTML files for the website
├── test/                  # Terratest Go tests
└── README.md              # This file
---
