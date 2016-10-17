package main

import (
	"fmt"
	"log"	

	"github.com/jroimartin/gocui"
)

func GUI_Init(){
	client.g = gocui.NewGui()
	if err := client.g.Init(); err != nil {
		log.Panicln(err)
	}
	client.g.SetLayout(layout)
	if err := initKeybindings(client.g); err != nil {
		log.Panicln(err)
	}
}

func GUI_Run(){
	defer client.g.Close();
	if err := client.g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size();
	v, err := g.SetView("title", 0, 0, maxX-1, maxY/8);
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err;
		}
	}
	v.Title = "Jangle";

	v, err = g.SetView("column", 0, maxY/8, maxX-1, maxY-1);
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err;
		}
	}
	
	v, err = g.SetView("message", maxX/4, maxY/8, maxX-1, 7*maxY/8);
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err;
		}
	}
	v.Autoscroll = true;
	v, err = g.SetView("input", maxX/4, 7*maxY/8, maxX-1, maxY-1);
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err;
		}
	}
	if err := g.SetCurrentView("input"); err != nil {
		return err
	}
	v.BgColor = gocui.ColorBlack;
	v.Editable = true;
	v.Wrap = true;
	

	return nil;
}

func initKeybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err;
	}
	if err := g.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			var p []byte;
			_, err :=v.Read(p);
			if(err == nil){
				write_data := make([]byte, len(v.ViewBuffer()) + 12)
				write_data[0] = 16;
				copy(write_data[1:4], Int_Converter(1)); 
				copy(write_data[5:8], Int_Converter(1)); 
				copy(write_data[9:12], Int_Converter(1)); 
				copy(write_data[13:], []byte(v.ViewBuffer()));
				if(client.debug){
					fmt.Println("OUT: ",write_data);
				}
				Write_To_Server(write_data);
				v.Clear();
			}
			return nil;
		}); err != nil {
			return err;
		}
	return nil;
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit;
}

func Append_Message(text []byte){
	i := 0;
	for (text[i]==10){
		i++;
	}
	client.g.Execute(func(g *gocui.Gui) error {
		v, err := g.View("message")
		if err != nil {
			// handle error
		}
		fmt.Fprintln(v, string(text[i:]))
		return nil
	})
}
