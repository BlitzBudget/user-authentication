# user-authentication
User Authentication Scenario for BlitzBudget

## Deployment

Github workflows are configured to ensure that the deployment is automatic and is deployed to the AWS Cloud.

The github workflow is developed under the following folder structure `.github/workflow`

## AWS CloudFormation

The cloudformation template is used to deploy the AWS Infrastructure using IaaC (Infrastructure as a Code). The code is deployed using the `AWS CLI`, check the README.md available in `aws/cloudformation`.

The cloudformation template is available for these environments.
1. Production Environment

## Projects

The serverless functions developed for this project available in the `root` folder structure. 