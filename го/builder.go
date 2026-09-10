package main

type EmailBuilder interface {
	SetFrom(address string) EmailBuilder
	SetTo(addresses ...string) EmailBuilder
	SetCC(addresses ...string) EmailBuilder
	SetSubject(subject string) EmailBuilder
	SetBody(body string) EmailBuilder
	SetPriority(priority Priority) EmailBuilder
	AddAttachment(filename string) EmailBuilder
}
