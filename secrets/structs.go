package secrets

//TwitterCredentials stores all values needed to connect to api
type TwitterCredentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

//AmazonCredentials for the api
type AmazonCredentials struct {
	AccessKeyID     string
	SecretAccessKey string
}
