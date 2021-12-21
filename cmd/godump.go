package main

import (
	"github.com/pterm/pterm"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	introScreen()

}

func introScreen() {
	ptermLogo, _ := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("Go", pterm.NewStyle(pterm.FgLightCyan)),
		pterm.NewLettersFromStringWithStyle("Dump", pterm.NewStyle(pterm.FgLightGreen))).
		Srender()

	pterm.DefaultCenter.Print(ptermLogo)

	pterm.Info.Println("This program is for analyzing with TCP dump.")
	pterm.Println()

	infaces, err := net.Interfaces()
	if err != nil {
		log.Printf("Error \t %v", err)
		os.Exit(1)
	}

	header := pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.BgDarkGray))
	pterm.DefaultCenter.Println(header.Sprint("AVAILABLE INTERFACES"))

	table := pterm.TableData{{"Index", "Name", "MAC"}}

	for i, v := range infaces {
		if i == 0 {
			continue
		}
		table = append(table, []string{strconv.Itoa(v.Index), v.Name, v.HardwareAddr.String()})
	}

	if err := pterm.DefaultTable.WithHasHeader().WithData(table).Render(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

}
