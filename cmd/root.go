package cmd

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/12yanogden/clipboard"
	"github.com/spf13/cobra"
)

// Base command
var rootCmd = &cobra.Command{
	Use:   "passgen",
	Short: "Generate a random password",
	Long: `Generate a random passord with alphanumeric and special characters.
	
	For example: passgen -l 12 --no-special`,

	Run: generatePassword,
}

func init() {
	rootCmd.Flags().IntP("length", "l", 20, "The length of the password to generate. Default: 20.")
	rootCmd.Flags().BoolP("no-alpha", "a", false, "Exclude alphabetical characters from the password generation.")
	rootCmd.Flags().BoolP("no-num", "n", false, "Exclude numberical characters from the password generation.")
	rootCmd.Flags().BoolP("no-special", "s", false, "Exclude special characters from the password generation.")
	rootCmd.Flags().StringP("specials", "p", "!@#$%^&*()_+{}[]|;:,.<>?-=", "The special characters used to generate the password. Default: !@#$%^&*()_+{}[]|;:,.<>?-=")
}

func generatePassword(cmd *cobra.Command, args []string) {
	length, _ := cmd.Flags().GetInt("length")
	noAlpha, _ := cmd.Flags().GetBool("no-alpha")
	noNum, _ := cmd.Flags().GetBool("no-num")
	noSpecial, _ := cmd.Flags().GetBool("no-special")
	specials, _ := cmd.Flags().GetString("specials")
	charset := ""
	password := make([]byte, length)

	if noAlpha && noNum && noSpecial {
		fmt.Println("passgen: must allow for at least one character type")
		os.Exit(1)
	}

	if !noAlpha {
		charset += "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	if !noNum {
		charset += "0123456789"
	}

	if !noSpecial {
		charset += specials
	}

	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	clipboard.PushBytes(password)
	fmt.Println("Copied to your clipboard")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}
