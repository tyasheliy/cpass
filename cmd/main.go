package main

import (
	"context"
	"fmt"
	"github.com/tyasheliy/cpass/internal/passcl"
)

func main() {
	c := passcl.OsClient{}
	ctx := context.Background()

	err := c.InsertOtp(ctx, "test", "otpauth://totp/Example:alice@google.com?secret=JBSWY3DPEHPK3PXP&issuer=Example", passcl.InsertOtpOptions{Force: true})
	if err != nil {
		fmt.Println("insert otp error")
	}

	fmt.Println(c.ShowOtp(ctx, "test"))
}
