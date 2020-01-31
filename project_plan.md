# AWS Personal Health Dashboard monitoring

## Description

The purpose of this report is to provide an understanding on monitoring AWS Personal Health Dashboard using Golang, Terraform and Slack. When an event occurs in AWS Personal health dashboard, it is mainly due to a resource within the account which needs attention. This could be a certificate expiring or an EC2 instance going down. In a big organisation managing 20+ accounts, keeping track of each account becomes impracticable, hence you set a monitoring alerts to notify you when an event occurs.

## Diagram
![alt text](https://github.com/omidayoobe/sparta_work/blob/master/src/aws_diagram.png)

## Main 


## Goal
The goal of this project is to set up the infrastructure, where when an event occurs in AWS Personal Health Dashboard, the lambda should pick up the payload and send the main information to a slack channel. This will allow the team to take actions on time in order to make sure the infrastructure is efficient and resilient. 

## Resources and languages
- Terraform
- Golang
- AWS Lambda
- AWS CloudWatch Rules
- AWS Personal Health Dashboard
- Slack
 





