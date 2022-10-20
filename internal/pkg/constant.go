package pkg

type AWSResourceType string

const AWSEC2 = AWSResourceType("AWS_EC2")
const AWSEKS = AWSResourceType("AWS_EKS")
const AWSS3 = AWSResourceType("AWS_S3")
const AWSECR = AWSResourceType("AWS_ECR")
const AWSELB = AWSResourceType("AWS_ELB")

var AWSResourceTypes = []AWSResourceType{AWSEC2, AWSEKS, AWSS3, AWSECR, AWSELB}

type ResourceOperationType string

const Create = ResourceOperationType("CREATE")
const Delete = ResourceOperationType("DELETE")
const Update = ResourceOperationType("UPDATE")
const Read = ResourceOperationType("READ")

var Operations = []ResourceOperationType{Create, Update, Delete, Read}
