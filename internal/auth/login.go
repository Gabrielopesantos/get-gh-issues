package auth

import (
	"errors"
	"github.com/spf13/cobra"
	"net/http"
)

type LoginOptions struct {
	Token       string
	GitProtocol string
	Hostname    string
}

func NewCmdLogin() *cobra.Command {
	opts := &LoginOptions{
		GitProtocol: "https",
		Hostname:    "github.com",
	}

	var token string
	cmd := &cobra.Command{
		Use:   "login",
		Short: "",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if token == "" {
				return errors.New("only token authentication is supported for now")
			}

			opts.Token = token
			return runLogin(opts)
		},
	}

	cmd.Flags().StringVar(&token, "token", "", "GitHub personal access token")

	return cmd
}

func runLogin(opts *LoginOptions) error {
	httpClient := newHttpClient()
	// Check min scopes
	// Get current logged username
	// Write to files
	return nil
}

// Has to be put in a separate package
func newHttpClient() *http.Client {
	return http.DefaultClient
}
