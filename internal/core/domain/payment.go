package domain

type PaymentPayload struct {
	UserID   ID
	CourseID ID
	PaySum   int64
}
