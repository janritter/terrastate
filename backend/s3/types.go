package s3

type stateConfig struct {
	Bucket        string
	DynamoDBTable string
	Key           string
	Region        string
	ACL           string
}
