package dto

type UserInfo struct {
	Code          string `json:"code"`
	RawData       string `json:"raw_data"`
	IV            string `json:"iv"`
	EncryptedData string `json:"encrypted_data"`
	Signature     string `json:"signature"`
}
