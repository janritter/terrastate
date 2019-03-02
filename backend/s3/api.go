package s3

type S3BackendAPI interface {
	GenerateStatefileForBackend(in interface{}) error
}

type S3Backend struct{}

func NewS3Backend() *S3Backend {
	return &S3Backend{}
}
