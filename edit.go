package main

import (
	"strconv"
	"strings"

	"github.com/rivo/tview"
)

func validateUDPPort(textToCheck string, lastChar rune) bool {
	val, err := strconv.ParseUint(textToCheck, 10, 32)
	if err != nil {
		return false
	}
	if val > 65534 || val < 1 {
		return false
	}
	return true
}

func validateHWAddr(textToCheck string, lastChar rune) bool {
	// 00:11:22:33:44:55
	if len(textToCheck) > 17 {
		return false
	}
	if !strings.ContainsRune("0123456789ABCDEF:", lastChar) {
		return false
	}
	return true
}

func editForm(onReturn func(), entry *Entry) *tview.Form {
	form := tview.NewForm().
		AddInputField("Name", entry.Name, 20, nil, func(value string) {
			entry.Name = value
		}).
		AddInputField("Hardware Addr", entry.HWAddr, 17, validateHWAddr, func(value string) {
			entry.HWAddr = value
		}).
		AddInputField("UDP Port", strconv.Itoa(int(entry.UDPPort)), 5, validateUDPPort, func(value string) {
			newPort, _ := strconv.Atoi(value)
			entry.UDPPort = uint(newPort)
		}).
		AddButton("Save", onReturn)
	form.SetBorder(true).SetTitle("Enter some data").SetTitleAlign(tview.AlignLeft)
	return form
}
