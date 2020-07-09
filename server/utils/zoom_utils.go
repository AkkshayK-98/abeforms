package utils

type ZoomMeeting struct {
	ClientEmail  string `json:"clientEmail,omitempty"`
	UserEmail    string `json:"userEmail,omitempty"`
	SelectedTime string `json:"selectedTime,omitempty"`
}

type LawyerSendTime struct {
	FirstSelectedTime  string `json:"firstSelectedTime,omitempty"`
	SecondSelectedTime string `json:"secondSelectedTime,omitempty"`
	ThirdSelectedTime  string `json:"thirdSelectedTime,omitempty"`
}
