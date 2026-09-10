package main

import (
	"errors"
	"fmt"
	"strings"
)

const (
	headerFrom     = "From"
	headerTo       = "To"
	headerCC       = "CC"
	headerSubject  = "Subject"
	headerPriority = "Priority"
	previewTitle   = "=== Email Preview ==="
)

type EmailPreviewBuilder struct {
	sb      strings.Builder
	hasFrom bool
	hasTo   bool
	hasSubj bool
}

func NewEmailPreviewBuilder() *EmailPreviewBuilder {
	b := &EmailPreviewBuilder{}
	b.sb.WriteString(previewTitle + "\n")
	return b
}

func (b *EmailPreviewBuilder) SetFrom(address string) EmailBuilder {
	b.sb.WriteString(fmt.Sprintf("%s:    %s\n", headerFrom, address))
	b.hasFrom = address != ""
	return b
}

func (b *EmailPreviewBuilder) SetTo(addresses ...string) EmailBuilder {
	b.sb.WriteString(fmt.Sprintf("%s:      %s\n", headerTo, strings.Join(addresses, ", ")))
	b.hasTo = len(addresses) > 0
	return b
}

func (b *EmailPreviewBuilder) SetCC(addresses ...string) EmailBuilder {
	if len(addresses) > 0 {
		b.sb.WriteString(fmt.Sprintf("%s:      %s\n", headerCC, strings.Join(addresses, ", ")))
	}
	return b
}

func (b *EmailPreviewBuilder) SetSubject(subject string) EmailBuilder {
	b.sb.WriteString(fmt.Sprintf("%s: %s\n", headerSubject, subject))
	b.hasSubj = subject != ""
	return b
}

func (b *EmailPreviewBuilder) SetBody(body string) EmailBuilder {
	b.sb.WriteString("\n")
	b.sb.WriteString(body)
	b.sb.WriteString("\n")
	return b
}

func (b *EmailPreviewBuilder) SetPriority(priority Priority) EmailBuilder {
	b.sb.WriteString(fmt.Sprintf("%s: %s\n", headerPriority, priority))
	return b
}

func (b *EmailPreviewBuilder) AddAttachment(filename string) EmailBuilder {
	b.sb.WriteString("Attachment: ")
	b.sb.WriteString(filename)
	b.sb.WriteString("\n")
	return b
}

func (b *EmailPreviewBuilder) Build() (string, error) {
	if !b.hasFrom {
		return "", errors.New("email preview: поле From обязательно")
	}
	if !b.hasTo {
		return "", errors.New("email preview: нужен хотя бы один адрес в To")
	}
	if !b.hasSubj {
		return "", errors.New("email preview: поле Subject обязательно")
	}
	return b.sb.String(), nil
}
