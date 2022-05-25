package auth

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
)

type LoginOptions struct {
	Token []byte
}

func NewCmdLogin() *cobra.Command {
	opts := &LoginOptions{}

	var withToken bool
	cmd := &cobra.Command{
		Use:   "login",
		Short: "",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if !withToken {
				return errors.New("only token authentication is supported for now")
			}
			// Read token (Tmp)
			//buf := &bytes.Buffer{}
			//reader := io.NopCloser(buf)
			//defer reader.Close()
			reader := bufio.NewReader(os.Stdin)
			token, err := io.ReadAll(reader)
			log.Println(token, err)
			if err != nil {
				return fmt.Errorf("failed to read token from standard input: %w", err)
			}

			opts.Token = token
			return runLogin(opts)
		},
	}

	cmd.Flags().BoolVar(&withToken, "with-token", true, "Read token from standard input")

	return cmd
}

func runLogin(opts *LoginOptions) error {
	// Validate login
	
	return nil
}

func validateAuthCredentials() {
}
