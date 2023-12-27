/*
Copyright © 2023 jay.trairat@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	caseName     int
	year         int
	evidenceList []int
)

var rootCmd = &cobra.Command{
	Use:   "init-case",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		currentPath, _ := os.Getwd()

		if len(evidenceList) == 0 {
			evidenceList = []int{1}
		}
		formattedCaseID := fmt.Sprintf("F-%04d-%03d", year, caseName)

		evidenceListFormat := []string{}
		for _, evidenceNumber := range evidenceList {
			formattedEvidence := formattedCaseID + "-" + fmt.Sprintf("EV%03d", evidenceNumber)
			evidenceListFormat = append(evidenceListFormat, formattedEvidence)
		}

		createFolder(currentPath, formattedCaseID, evidenceListFormat)
		fmt.Println("Successfully created case (" + formattedCaseID + ") folder with evidence (" + strings.Join(evidenceListFormat, ", ") + ") and sub-folders.")
	},
}

func createFolder(basePath string, caseName string, evidenceList []string) error {
	evidenceStructure := []string{"Documents", "Images&Extractions", "Pictures", "Reports", "Request", "Videos"}
	casePath := filepath.Join(basePath, caseName)

	os.MkdirAll(casePath, 0755)
	for _, dir := range evidenceStructure {
		structureFolder := filepath.Join(casePath, dir)
		os.MkdirAll(structureFolder, 0755)
		for _, subDir := range evidenceList {
			os.MkdirAll(filepath.Join(structureFolder, subDir), 0755)
		}
	}
	return nil
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVarP(&caseName, "case", "c", 1, "Name of the case :: 1")
	rootCmd.Flags().IntVarP(&year, "year", "y", time.Now().Year(), "Year :: YYYY")
	rootCmd.Flags().IntSliceVarP(&evidenceList, "evidence", "e", []int{}, "List of evidence number :: 1 2")

	rootCmd.MarkFlagRequired("case")
	rootCmd.MarkFlagRequired("year")
}
