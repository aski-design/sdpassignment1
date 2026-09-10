package main

import "errors"

type EmailObjectBuilder struct {
	from        string
	to          []string
	cc          []string
	subject     string
	body        string
	priority    Priority
	attachments []string
}

func NewEmailObjectBuilder() *EmailObjectBuilder {
	return &EmailObjectBuilder{priority: PriorityNormal}
}

func (b *EmailObjectBuilder) SetFrom(address string) EmailBuilder {
	b.from = address
	return b
}

func (b *EmailObjectBuilder) SetTo(addresses ...string) EmailBuilder {
	b.to = addresses
	return b
}

func (b *EmailObjectBuilder) SetCC(addresses ...string) EmailBuilder {
	b.cc = addresses
	return b
}

func (b *EmailObjectBuilder) SetSubject(subject string) EmailBuilder {
	b.subject = subject
	return b
}

func (b *EmailObjectBuilder) SetBody(body string) EmailBuilder {
	b.body = body
	return b
}

func (b *EmailObjectBuilder) SetPriority(priority Priority) EmailBuilder {
	b.priority = priority
	return b
}

func (b *EmailObjectBuilder) AddAttachment(filename string) EmailBuilder {
	b.attachments = append(b.attachments, filename)
	return b
}

func (b *EmailObjectBuilder) Build() (*Email, error) {
	if b.from == "" {
		return nil, errors.New("email: поле From обязательно")
	}
	if len(b.to) == 0 {
		return nil, errors.New("email: нужен хотя бы один адрес в To")
	}
	if b.subject == "" {
		return nil, errors.New("email: поле Subject обязательно")
	}
	return newEmail(b.from, b.to, b.cc, b.subject, b.body, b.priority, b.attachments), nil
}
