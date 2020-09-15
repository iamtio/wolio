package main

import (
	"strconv"

	"github.com/linde12/gowol"
	"github.com/rivo/tview"
)

// App -
type App struct {
	mode        rune
	application *tview.Application
	config      ConfigStore
	entries     *[]Entry
}

// NewApp create instance
func NewApp() *App {
	return &App{
		mode:        's',
		application: tview.NewApplication(),
		config:      JSONConfigStore{},
	}
}

// DrawMenu -
func (a *App) DrawMenu() {
	entries, _ := a.config.loadEntries()
	a.entries = entries
	list := a.entriesView()

	list.AddItem("Add new", "", 'a', func() {
		*a.entries = append(*a.entries, Entry{Name: "New Entry", UDPPort: 9})
		a.config.storeEntries(a.entries)
		a.DrawMenu()
	})

	switch a.mode {
	case 's':
		list.AddItem("(SEND) Change mode", "", 'm', func() {
			a.mode = 'e'
			a.DrawMenu()
			return
		})
	case 'e':
		list.AddItem("(EDIT) Change mode", "", 'm', func() {
			a.mode = 'd'
			a.DrawMenu()
			return
		})
	case 'd':
		list.AddItem("(DELETE) Change mode", "", 'm', func() {
			a.mode = 's'
			a.DrawMenu()
			return
		})
	}

	list.AddItem("Quit", "", 'q', func() {
		a.application.Stop()
	})
	a.application.SetRoot(list, true).SetFocus(list)
}

func (a *App) entriesView() *tview.List {
	list := tview.NewList()
	for i, entry := range *a.entries {
		ic := i

		var r rune = '-'
		if i < 9 {
			r = rune('1' + i)
		}
		list.AddItem(entry.Name, entry.HWAddr, r, func() {
			switch a.mode {
			case 'e':
				f := editForm(func() {
					a.config.storeEntries(a.entries)
					a.DrawMenu()
				}, &(*a.entries)[ic])
				a.application.SetRoot(f, true).SetFocus(f)
			case 's':
				e := (*a.entries)[ic]
				packet, err := gowol.NewMagicPacket(e.HWAddr)
				if err != nil {
					// packet.SendPort("255.255.255.255", "9")
					panic(err)
				}
				packet.SendPort("255.255.255.255", strconv.Itoa(int(e.UDPPort)))

			case 'd':
				*a.entries = append((*a.entries)[:ic], (*a.entries)[ic+1:]...)
				a.config.storeEntries(a.entries)
				a.DrawMenu()
			}

		})
	}
	return list
}

// Run -
func (a *App) Run() {
	if err := a.application.Run(); err != nil {
		panic(err)
	}
}
