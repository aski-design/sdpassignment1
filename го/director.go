package main

const (
	senderAdmin   = "admin@example.com"
	senderSupport = "support@example.com"
)

type EmailDirector struct{}

func (d *EmailDirector) MakeWelcomeEmail(b EmailBuilder, userEmail string) {
	b.SetFrom(senderAdmin).
		SetTo(userEmail).
		SetSubject("Welcome aboard!").
		SetBody("Thanks for signing up. We are happy to have you.").
		SetPriority(PriorityNormal)
}

func (d *EmailDirector) MakePasswordResetEmail(b EmailBuilder, userEmail string) {
	b.SetFrom(senderSupport).
		SetTo(userEmail).
		SetSubject("Reset your password").
		SetBody("Click the link below to reset your password. It expires in 30 minutes.").
		SetPriority(PriorityHigh).
		AddAttachment("security-tips.pdf")
}
