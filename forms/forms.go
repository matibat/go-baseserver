package forms

// Credentials form
type Credentials struct {
	User string
	Pass []byte
}

// SignUpData form
type SignUpData struct {
	User  string
	Pass  string
	Email string
}

