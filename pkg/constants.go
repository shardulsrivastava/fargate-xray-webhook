package pkg

const (
	// Default token expiration in seconds if none is defined,
	// which is 24hrs as that is max for EKS
	DefaultTokenExpiration = int64(86400)
	// 1hr is min for kube-apiserver
	MinTokenExpiration = int64(3600)
)