/*
laptop settings
Copyright (C) 2021  Ashok Pappu

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.

Gnomovision version 69, Copyright (C) year Ashok Pappu
Gnomovision comes with ABSOLUTELY NO WARRANTY; for details
type `show w'.  This is free software, and you are welcome
to redistribute it under certain conditions; type `show c'
for details.




*/

package main

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"log"
)

func main() {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)
	builder, err := gtk.BuilderNewFromFile("/home/apappu/GolandProjects/laptop/laptop.glade")
	if err != nil {
		log.Fatal("Unable to create builder:", err)
	}
	// you just place them in a map that names the signals, then feed the map to the builder
	var signals = map[string]interface{}{
		"on_savebtn_clicked":  onSaveClicked,
		"closebtn_clicked_cb": closetClicked,
	}

	builder.ConnectSignals(signals)

	obj, err := builder.GetObject("window")
	win := obj.(*gtk.Window)

	//// Create a new toplevel window, set its title, and connect it to the
	//// "destroy" signal to exit the GTK main loop when it is destroyed.
	//win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Simple Example")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	//
	//// Create a new label widget to show in the window.
	//l, err := gtk.LabelNew("Hello, gotk3!")
	//if err != nil {
	//	log.Fatal("Unable to create label:", err)
	//}
	//
	//// Add the label to the window.
	//win.Add(l)
	//
	//// Set the default window size.
	//win.SetDefaultSize(800, 600)
	//
	//// Recursively show all widgets contained in this window.

	//
	//// Begin executing the GTK main loop.  This blocks until
	//// gtk.MainQuit() is run.
	win.ShowAll()
	gtk.Main()
}

func onSaveClicked() {
	fmt.Println("save clicked")
}

func closetClicked() {
	fmt.Println(" close clicked")
	gtk.MainQuit()
}
