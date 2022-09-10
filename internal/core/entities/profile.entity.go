package entities

const ProfileTableName = "profiles"

type Profile struct {
	Id        string `json:"Id,omitempty"`
	Name      string `json:"Name,omitempty"`
	Avatar    string `json:"Avatar,omitempty"`
	User      User   `json:"User,omitempty"`
	CreatedAt string `json:"CreatedAt,omitempty"`
	UpdatedAt string `json:"UpdatedAt,omitempty"`
}
