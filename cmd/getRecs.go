package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

var getRecsCmd = &cobra.Command{
	Use:   "getRecs",
	Short: "A way to get book recommendations given a folder of book files",
	Long: `Use --num / -n command to specify the number of recommendations desired, and --dir/-d to specify the target directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()
		dir, err := cmd.Flags().GetString("dir")
		if err != nil || dir == "" {
			fmt.Println("Error: please provide target directory via -d or --dir flags")
			return
		}
		numString, err := cmd.Flags().GetString("num")
		if err != nil || numString == "" {
			fmt.Println("Error: please provide number of books via -n or --num flags")
			return
		}
		num, err := strconv.Atoi(numString)
		if err != nil {
			fmt.Printf("Error: please provide valid number of books via -n or --num flags: %v", err)
			return
		}

		titleArray := listBooksInDir(dir)

		rand.Seed(time.Now().UTC().UnixNano())
		var chosenBooks []string

		for i := 0 ; i < num; i++ {
			newRand := rand.Intn(len(titleArray))
			chosenBooks = append(chosenBooks, titleArray[newRand])
		}

		fmt.Println("\nHere are your suggestions:")
		for i, book := range chosenBooks {
			fmt.Printf("\n %s: %s", strconv.Itoa(i + 1), book)
		}
		fmt.Println()
		elapsed := time.Since(start)
		fmt.Printf("Process took %v\n", elapsed)
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
			var directoryDivider string
			if runtime.GOOS == "windows" {
				directoryDivider = `\`
			} else {
				directoryDivider = "/"
			}
			titleArray = append(titleArray, listBooksInDir(dir + directoryDivider + f.Name())...)
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
