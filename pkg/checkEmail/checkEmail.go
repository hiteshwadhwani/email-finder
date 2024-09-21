package checkEmail

import (
	"fmt"
	"net"
	"net/smtp"
	"os"
	"strings"
)

type validEmails []string

func New() *validEmails {
	return &validEmails{}
}

func (v *validEmails) Add(email string) {
	*v = append(*v, email)
}

func (v *validEmails) Check(email string) {
	domain := strings.Split(email, "@")[1]

	mx, err := net.LookupMX(domain)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error looking up mx record: %v\n", err)
		return
	}

	client, err := smtp.Dial(fmt.Sprintf("%s:%d", mx[0].Host, 25))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error connecting to smtp server: %v\n", err)
		return
	}

	defer client.Close()

	if err := client.Hello("localhost"); err != nil {
		fmt.Fprintf(os.Stderr, "error sending hello: %v\n", err)
		return
	}

	if err := client.Mail("xanew65641@cetnob.com"); err != nil {
		fmt.Fprintf(os.Stderr, "error sending mail from: %v\n", err)
		return
	}

	if err := client.Rcpt(email); err != nil {
		fmt.Fprintf(os.Stderr, "error sending mail to: %v\n", err)
		return
	}

	v.Add(email)
}
