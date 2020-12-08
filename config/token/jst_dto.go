package token

// TokenDetails struct
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

// AccessDetails struct
type AccessDetails struct {
	AccessUUID string
	UserID     uint64
}
