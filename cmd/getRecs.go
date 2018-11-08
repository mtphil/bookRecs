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
		start := time.Now()
		dir, _ := cmd.Flags().GetString("dir")
		numString, _ := cmd.Flags().GetString("num")
		num, _ := strconv.Atoi(numString)

		titleArray := listBooksInDir(dir)

		rand.Seed(time.Now().UTC().UnixNano())
		var chosenBooks []string

		for i := 0 ; i < num; i++ {
			newRand := rand.Intn(len(titleArray))
			chosenBooks = append(chosenBooks, titleArray[newRand])
		}

		fmt.Println("\nHere are your suggestions:")
		for i, book := range chosenBooks {
			fmt.Println("\n" + strconv.Itoa(i + 1) + ": " + book)
		}
		fmt.Println()
		elapsed := time.Since(start)
		fmt.Printf("Process took %s\n", elapsed)
		fmt.Println()
	},
}

func listBooksInDir(dir string) []string {
	var titleArray []string

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			titleArray = append(titleArray, listBooksInDir(dir + "/" + f.Name())...)
		}
		titleArray = append(titleArray, f.Name())
	}
	return titleArray
}

func init() {
	rootCmd.AddCommand(getRecsCmd)
	getRecsCmd.Flags().StringP("dir", "d", "", "set source directory")
	getRecsCmd.Flags().StringP("num", "n", "", "set number of recommendations")
}
