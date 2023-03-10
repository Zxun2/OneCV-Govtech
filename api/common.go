package api

// Response is the model for the response
type Response struct {
	Message			string `json:"message,omitempty"`
}

// SuspendStudentPayload - incoming request
type SuspendStudentPayload struct {
	Student 			string `json:"student"`
}

// SuspendStudentResponse - outgoing response
type SuspendStudentResponse struct {
	Response
}

// RegistersStudentsPayload - incoming request
type RegistersStudentsPayload struct {
	Teacher 			string   `json:"teacher"`
	Students 			[]string `json:"students"`
}

// RegistersStudentsResponse - outgoing response
type RegistersStudentsResponse struct {
	Response
}

// RetrieveCommonStudentsPayload - incoming request
type RetrieveCommonStudentsPayload struct {
	Students 			[]string `json:"students"`
}

// RetrieveCommonStudentsResponse - outgoing response
type RetrieveCommonStudentsResponse struct {
	Response
	Students 			[]string `json:"students"`
}

// ListStudentReceivingNotificationPayload - incoming request
type ListStudentReceivingNotificationPayload struct {
	Teacher 			string `json:"teacher"`
	Notification 	string `json:"notification"`
}

// ListStudentReceivingNotificationResponse - outgoing response
type ListStudentReceivingNotificationResponse struct {
	Response
	Recipients 		[]string `json:"recipients"`
}

func makeResponseErr(err error) Response {
	return Response{Message: err.Error()}
}