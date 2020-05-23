package gomunk

type GomunkConfiguration struct {
	PROVIDER   string
	AWS_REGION string
	AWS_BUCKET string
}

var (
	Config GomunkConfiguration
)
