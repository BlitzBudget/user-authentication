package config

import "os"

var ClientId = os.Getenv("CLIENT_ID")
var ClientSecret = os.Getenv("CLIENT_SECRET")

var CompanyIDUserAttribute = "custom:company_id"
var FirstNameUserAttribute = "custom:first_name"
var LastNameUserAttribute = "custom:last_name"
var MiddleNameUserAttribute = "custom:middle_name"
var OrganizationIDUserAttribute = "custom:organization_id"
var ProjectManagerUserAttribute = "custom:project_manager"
var TransformChannelUserAttribute = "custom:transform_channel"
