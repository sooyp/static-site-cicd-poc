# Static Website Deployment Pipeline to AWS S3 (POC)

This project is a hands-on Proof of Concept (POC) showcasing a complete CI/CD pipeline that provisions infrastructure and deploys a static website to an AWS S3 bucket using **Terraform** and **GitHub Actions**.

---

## 🚀 Project Overview

- **Infrastructure-as-Code**: AWS S3 bucket provisioned via Terraform
- **CI/CD Automation**: GitHub Actions for Terraform operations and site upload

---

## Golang and Terratest Setup

In this project, **Terratest** is used to validate the resources created by **Terraform**. Terratest is written in **Go**, so you need to have **Go** installed in the CI/CD pipeline github runner.

---

## 🗂️ Structure
```. 
├── .github/workflows/ # CI pipeline definition 
├── terraform/ # Terraform code to provision S3 
├── static-site/ # HTML files for the website
├── tests/ # terratest verification of s3
└── README.md # This file
```

---

## 📦 Requirements

- **AWS account** with IAM user configured for S3 and Terraform permissions
- **GitHub repository secrets set:**
  - `AWS_ACCESS_KEY_ID`
  - `AWS_SECRET_ACCESS_KEY`
  - `S3_BUCKET`

---

## 🔧 How It Works

### On Push to `develop` Branch:

1. **Terraform Init & Apply**
   - Initializes and applies S3 bucket configuration

2. **Upload Static Site**
   - Syncs HTML content in `static-site/` to the S3 bucket

---

## 🌐 Static Site Example

The `static-site/index.html` contains a placeholder welcome page. After deployment, this site will be accessible via the S3 website endpoint.

---
## GitHub Setup Needed
Create two environments in GitHub > Settings > Environments:

  - dev
  - prod

In each environment, add the following secrets:
  - AWS_ACCESS_KEY_ID
  - AWS_SECRET_ACCESS_KEY
  - S3_BUCKET (specific to each env)

Optionally add required reviewers for the prod environment to enforce approval gates.

---

## ✅ Status

This project is to provide a fully functional and used as a live demonstration of deploying cloud-native infrastructure with CI/CD.

---

## 🙋‍♂️ Author

Maintained by Lee Broom for demonstration purposes. Contributions and feedback are welcome!
