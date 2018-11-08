package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var getRecsCmd = &cobra.Command{
	Use:   "getRecs",
	Short: "A way to get book recommendations given a folder of book files",
	Long: `Use --num / -n command to specify the number of recommendations desired, and --dir/-d to specify the target directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, _ := cmd.Flags().GetString("dir")
		numString, _ := cmd.Flags().GetString("num")
		num, _ := strconv.Atoi(numString)

		files, err := ioutil.ReadDir(dir)
		if err != nil {
			log.Fatal(err)
		}

		var titleArray []string

		for _, f := range files {
			if f.IsDir() {
				continue
			}
			titleArray = append(titleArray, f.Name())
		}

		rand.Seed(time.Now().UTC().UnixNano())
		var chosenBooks []string

		for i := 0 ; i < num; i++ {
			newRand := rand.Intn(len(titleArray))
			chosenBooks = append(chosenBooks, titleArray[newRand])
		}

		fmt.Println("\nHere are your suggestions:")
		for i, book := range chosenBooks {
			fmt.Println("\n" + strconv.Itoa(i) + ": " + book)
		}
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(getRecsCmd)
	getRecsCmd.Flags().StringP("dir", "d", "", "set source directory")
	getRecsCmd.Flags().StringP("num", "n", "", "set number of recommendations")
}
