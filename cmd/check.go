/*
Copyright © 2019 Joel Montes de Oca

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// Used for flags
var (
	simple bool
	ip     bool
)

// Argument used
var url string

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check URL",
	Short: "Checks to see if a website is up",
	Args:  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url = strings.Join(args, " ")
		fmt.Println("Check ran.")
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	checkCmd.Flags().BoolVarP(&ip, "ip", "i", false, "Returns the IP address of the URL entered")
	checkCmd.Flags().BoolVarP(&simple, "simple", "s", false, "Return a simple answer")
}
