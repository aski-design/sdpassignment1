package main

import "fmt"

type Priority int

const (
	PriorityLow Priority = iota
	PriorityNormal
	PriorityHigh
)

func (p Priority) String() string {
	switch p {
	case PriorityLow:
		return "Low"
	case PriorityHigh:
		return "High"
	default:
		return "Normal"
	}
}

type Email struct {
	from        string
	to          []string
	cc          []string
	subject     string
	body        string
	priority    Priority
	attachments []string
}

func newEmail(from string, to, cc []string, subject, body string, priority Priority, attachments []string) *Email {
	return &Email{
		from:        from,
		to:          to,
		cc:          cc,
		subject:     subject,
		body:        body,
		priority:    priority,
		attachments: attachments,
	}
}

func (e *Email) From() string          { return e.from }
func (e *Email) To() []string          { return e.to }
func (e *Email) CC() []string          { return e.cc }
func (e *Email) Subject() string       { return e.subject }
func (e *Email) Body() string          { return e.body }
func (e *Email) Priority() Priority    { return e.priority }
func (e *Email) Attachments() []string { return e.attachments }

func (e *Email) String() string {
	return fmt.Sprintf("Email{from=%s, to=%v, subject=%q, priority=%s}",
		e.from, e.to, e.subject, e.priority)
}
