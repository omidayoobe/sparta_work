package notifier

import (
	"testing"
	)

const payload = `
{
    "eventArn": "arn:aws:health:eu-west-1::event/RDS/AWS_RDS_SECURITY_NOTIFICATION/AWS_RDS_SECURITY_NOTIFICATION_2f3a2c82-ca08-3265-8372-3c5435ae5s6f",
    "service": "RDS",
    "eventTypeCode": "AWS_RDS_SECURITY_NOTIFICATION",
    "eventTypeCategory": "accountNotification",
    "startTime": "Wed, 8 Jan 2020 21:00:00 GMT",
    "eventDescription": [
      {
        "language": "en_US",
        "latestDescription": "We previously sent a communication in early October to update your RDS SSL/TLS certificates by October 31, 2019. We have extended the dates and now request that you act before February 5, 2020 to avoid interruption of your applications that use Secure Sockets Layer (SSL) or Transport Layer Security (TLS) to connect to your RDS and Aurora database instances. Note that this new date is only 4 weeks before the actual Certificate Authority (CA) expiration on March 5, 2020. Because our own deployments, testing, and scanning to validate all RDS instances are ready for the expiry must take place during the final 4 weeks, the February 5th date cannot be further extended.  You are receiving this message because you have an Amazon RDS database instance(s) that requires action listed in the 'Affected resources' tab.  To protect your communications with RDS database instances, a CA generates time-bound certificates that are checked by your client applications that connect via SSL/TLS to authenticate RDS databases before exchanging information. AWS renews the CA and creates new root certificates every five years to ensure RDS customer connections are properly protected for years to come.  "
      }
    ],
    "affectedEntities": [
      {
        "entityValue": "nonprod-staging-rds-1",
        "tags": {}
      }
    ]
  }
`

func TestShouldRenderMessage(t *testing.T) {
	msg := `
*Service:* RDS
*Event Name:* 
AWS_RDS_SECURITY_NOTIFICATION

*Event Description:* 
We previously sent a communication in early October to update your RDS SSL/TLS certificates by October 31, 2019. We have extended the dates and now request that you act before February 5, 2020 to avoid interruption of your applications that use Secure Sockets Layer (SSL) or Transport Layer Security (TLS) to connect to your RDS and Aurora database instances. Note that this new date is only 4 weeks before the actual Certificate Authority (CA) expiration on March 5, 2020. Because our own deployments, testing, and scanning to validate all RDS instances are ready for the expiry must take place during the final 4 weeks, the February 5th date cannot be further extended.

*Affected Entities:* 
*nonprod-staging-rds-1*
`

}