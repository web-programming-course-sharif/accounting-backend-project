package otp

type Request struct {
	Receptor string `json:"receptor"`
	Token    string `json:"token"`
	Token2   string `json:"token2"`
	Token3   string `json:"token3"`
	Template string `json:"template"`
	Type     string `json:"type"`
}
