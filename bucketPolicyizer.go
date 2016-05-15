package bucketPolicyizer

import "encoding/json"

var defaultVersion = "2012-10-17"

// GetObjectAction the s3 action for getting an object
var GetObjectAction = "s3:GetObject"

// Policy is the Bucket policy
type Policy struct {
	Version   string
	Statement []Statement
}

// Statement is a single permission
// the Principal element is sometimes an
// array and sometimes a string
type Statement struct {
	Sid       string
	Effect    string
	Principal interface{}
	Action    []string
	Resource  []string
}

// Principal is a list of ARNs
type Principal struct {
	AWS []string
}

// EmptyPolicy creates a valid empty policy
func EmptyPolicy() Policy {
	return Policy{
		Version: defaultVersion,
	}
}

// CompilePolicy renders the policy to JSON
func CompilePolicy(policy Policy) (string, error) {
	p, err := json.Marshal(policy)

	if err != nil {
		return "", err
	}

	return string(p), nil
}
