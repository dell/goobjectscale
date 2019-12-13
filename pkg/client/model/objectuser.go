package model

type BlobUser struct {
	UserID    string `json:"userid"`
	Namespace string `json:"namespace"`
}

// ObjectUserList contains an array of object users.
type ObjectUserList struct {
	BlobUser []BlobUser `json:"blobuser"`
}

// ObjectUserInfo contains information about an object user.
type ObjectUserInfo struct {
	Namespace string   `json:"namespace"`
	Name      string   `json:"name"`
	Locked    bool     `json:"locked"`
	Created   string   `json:"created"`
	Tags      []string `json:"tags"`
}

// ObjectUserSecret contains information about object user's secrets.
type ObjectUserSecret struct {
	SecretKey1          string `json:"secret_key_1"`
	KeyTimestamp1       string `json:"key_timestamp_1"`
	KeyExpiryTimestamp1 string `json:"key_expiry_timestamp_1"`
	SecretKey2          string `json:"secret_key_2"`
	KeyTimestamp2       string `json:"key_timestamp_2"`
	KeyExpiryTimestamp2 string `json:"key_expiry_timestamp_2"`
	Link                Link   `json:"link"`
}
