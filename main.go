package main

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var content = "User-agent: *\nDisallow: /\n"

type Handler struct {
	PathPattern  *regexp.Regexp
	MetaImport   string
	RedirectName string
	RedirectTo   string
}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	_, err := rw.Write([]byte(content))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	var cmdServe = &cobra.Command{
		Use:   "serve",
		Short: "Start serving robots.txt content",
		Long:  "Starts a HTTP server serving content of a robots.txt file disallowing all robots.",
		Run: func(cmd *cobra.Command, args []string) {
			h := &Handler{}
			fmt.Println("Listening on " + viper.GetString("listen") + "...")
			panic(http.ListenAndServe(viper.GetString("listen"), h))
		},
	}

	viper.SetEnvPrefix("ROBOTS_DISALLOW")
	viper.AutomaticEnv()

	cmdServe.Flags().String("listen", "0.0.0.0:8080", "address to listen on [$ROBOTS_DISALLOW_LISTEN]")
	viper.BindPFlag("listen", cmdServe.Flags().Lookup("listen"))

	var rootCmd = &cobra.Command{Use: "robots-disallow"}
	rootCmd.AddCommand(cmdServe)
	rootCmd.Execute()
}
