package main

import "fmt"

func main() {
	director := &EmailDirector{}
	objBuilder := NewEmailObjectBuilder()
	director.MakeWelcomeEmail(objBuilder, "new.user@example.com")
	welcomeEmail, err := objBuilder.Build()
	if err != nil {
		fmt.Println("build error:", err)
		return
	}
	fmt.Println(welcomeEmail)
	fmt.Println()

	previewBuilder := NewEmailPreviewBuilder()
	director.MakePasswordResetEmail(previewBuilder, "new.user@example.com")
	preview, err := previewBuilder.Build()
	if err != nil {
		fmt.Println("build error:", err)
		return
	}
	fmt.Println(preview)

	badBuilder := NewEmailObjectBuilder()
	badBuilder.SetSubject("Oops, no from/to")
	if _, err := badBuilder.Build(); err != nil {
		fmt.Println("validation works:", err)
	}
}
