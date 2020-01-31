# AWS Personal Health Dashboard monitoring

## Description

The purpose of this report is to provide an understanding on monitoring AWS Personal Health Dashboard using Golang, Terraform and Slack. When an event occurs in AWS Personal health dashboard, it is mainly due to a resource within the account that needs attention. This could be a certificate expiring or an EC2 instance going down. In a big organisation managing 20+ accounts, keeping track of each account becomes impracticable, hence you set monitoring alerts to notify you when an event occurs.

## Diagram
![alt text](https://github.com/omidayoobe/sparta_work/blob/master/src/aws_diagram.png)

## Main 
The diagram above shows the infrastructure of the project, which will be created. Whenever there is an event in personal Health Dashboard the logs of that event will be stored in CloudWatch logs. We will write CloudWatch Event Rules which will target the lambda. The rules will specify that anything related to aws.health, trigger the lambda and the lambda will grab the payload of the event and send the important parts of the payload to a slack channel.
<br>
The lambda which will be written in Golang will have two parts, the notifier.go which will grab the main information from slack and main.go which will create our lambda and integrate it with the notifier.go. We will be also testing our code using an example payload available on the AWS website. The lambda and the notifier will be both written in GoLang, that is because Go is a lightweight language, making it efficient for infrastructure.
<br>
We will integrate slack with the lambda by using the webhook provided by Slack and deploy the lambda and CloudWatch Event rules to AWS using Terraform.

## Overall goal
The goal of this project is to set up an infrastructure, whenever an event occurs in AWS Personal Health Dashboard, the lambda should be triggered, picking up the payload and sending the main information to a slack channel. This will allow the team to take action on time in order to make sure the infrastructure is efficient and resilient. 

## Resources and languages
- Terraform
- Golang
- AWS Lambda
- AWS CloudWatch Rules
- AWS Personal Health Dashboard
- Slack
 





