package types

var (
	WELCOME_EMAIL                    = "welcome"
	SUBMISSION_SUCCESS_EMAIL         = "user_reservation_submission"
	RESERVATION_APPROVED_EMAIL       = "reservation_approved"
	RESERVATION_ADMIN_EMAIL          = "admin_reservation"
)

type Email struct {
	// use this to get the email from firebase
	UserID   string
	FullName string
	Username string
	Type     string
	Data     EmailReservationData
}

type EmailReservationData struct {
	ReservationID 	string
	ProductID       string
	ProductName 	string
	StartDate       string
	EndDate         string
	MediaURL        string
}
