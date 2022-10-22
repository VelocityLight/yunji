package pkg

type AWSResourceType string

const AWSEC2 = AWSResourceType("AmazonEC2")
const AWSEKS = AWSResourceType("AmazonEKS")
const AWSS3 = AWSResourceType("AmazonS3")
const AWSVPC = AWSResourceType("AmazonVPC")
const AWSELB = AWSResourceType("AWSELB")
const AWSRoute53 = AWSResourceType("AmazonRoute53")

var AWSResourceTypes = []AWSResourceType{AWSEC2, AWSEKS, AWSS3, AWSVPC, AWSRoute53, AWSELB}

type ResourceOperationType string

const Create = ResourceOperationType("CREATE")
const Delete = ResourceOperationType("DELETE")
const Update = ResourceOperationType("UPDATE")
const Read = ResourceOperationType("READ")

var Operations = []ResourceOperationType{Create, Update, Delete, Read}
