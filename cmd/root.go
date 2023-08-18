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
	caseName     string
	evidenceList []string
)

var rootCmd = &cobra.Command{
	Use:   "init-case",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		currentPath, _ := os.Getwd()

		if caseName == "" {
			timestamp := time.Now().Format("20060102150405")
			caseName = fmt.Sprintf("%s_NEWCASE", timestamp)
		}

		if len(evidenceList) == 0 {
			evidenceList = []string{"EVIDENCE"}
		}

		createFolder(currentPath, caseName, evidenceList)
		fmt.Println("Successfully created case (" + caseName + ") folder with evidence (" + strings.Join(evidenceList, ", ") + ") and sub-folders.")
	},
}

func createFolder(basePath string, caseName string, evidenceList []string) error {
	evidenceStructure := []string{"Documents", "Extracts", "Pictures", "Videos", "Reports", "CaseFiles"}
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
	rootCmd.Flags().StringVarP(&caseName, "case", "c", "", "Name of the case :: F01-66")
	rootCmd.Flags().StringSliceVarP(&evidenceList, "evidence", "e", []string{}, "List of evidence :: EV01 EV02")
}
