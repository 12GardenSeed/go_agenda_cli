package cmd

import (
	"fmt"


	"github.com/HinanawiTenshi/Agenda/entity"
	"github.com/spf13/cobra"
)

// clearmeetingCmd represents the clearmeeting command
var clearmeetingCmd = &cobra.Command{
	Use:   "clearmeeting",
	Short: "Remove all the meetings you host.",
	Long:  `All the meetings that you created will be removed. Think twice.`,
	Run: func(cmd *cobra.Command, args []string) {
		curUser, _ := getCurUser()
		if curUser == "" {
			fmt.Println(argsError{permissionDeny: true}.Error())
			_errorLog.Println(argsError{permissionDeny: true}.Error())
			return
		}
		if len(args) != 0 {
			fmt.Println(argsError{invalidNArgs: true}.Error())
			_errorLog.Println(argsError{invalidNArgs: true}.Error())
			return
		}
		meetings := entity.GetMeetings()
		newMeetings := make([]entity.Meeting, 0)
		for _, meeting := range meetings {
			if !(meeting.Host == curUser) {
				newMeetings = append(newMeetings, meeting)
			}
		}
		entity.UpdateMeeting(newMeetings)
		fmt.Println("All meetings have been removed.")
		_infoLog.Printf("[" + curUser + "] All meetings have been removed\n")
	},
}

func init() {
	RootCmd.AddCommand(clearmeetingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clearmeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clearmeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
