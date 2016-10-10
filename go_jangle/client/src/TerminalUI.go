package main

import (
	"fmt"
	"log"
	"strings"
	
	"github.com/jroimartin/gocui"
)

type Client struct {
	debug bool
}

var client Client

const delta = 1

var (
	views   = []string{}
	curView = -1
	idxView = 0
)

func main() {
	g := gocui.NewGui()
	if err := g.Init(); err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetLayout(layout)
	if err := initKeybindings(g); err != nil {
		log.Panicln(err)
	}
	if err := newView(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size();
	v, err := g.SetView("legend", maxX-25, 0, maxX-1, 8);
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err;
		}
		fmt.Fprintln(v, "KEYBINDINGS");
		fmt.Fprintln(v, "Space: New View");
		fmt.Fprintln(v, "Tab: Next View");
		fmt.Fprintln(v, "← ↑ → ↓: Move View");
		fmt.Fprintln(v, "Backspace: Delete View");
		fmt.Fprintln(v, "t: Set view on top");
		fmt.Fprintln(v, "^C: Exit");
	}
	v, err = g.SetView("title", 0, 0, maxX-1, maxY/8);
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err;
		}
		//fmt.Fprintln(v, "Jangle");
	}
	v.Title = "Jangle";

	v, err = g.SetView("column", 0, maxY/8, maxX-1, maxY-1);
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err;
		}
		//fmt.Fprintln(v, "Jangle");
	}
	
	v, err = g.SetView("message", maxX/4, maxY/8, maxX-1, 7*maxY/8);
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err;
		}
		fmt.Fprintln(v, "Micky: Test Message");
		fmt.Fprintln(v, "Demi: This is this is testing");
		fmt.Fprintln(v, "Micky: Test Message");
		fmt.Fprintln(v, "Micky: Test Message");
		fmt.Fprintln(v, "Micky: Test Message");
	}

	v, err = g.SetView("input", maxX/4, 7*maxY/8, maxX-1, maxY-1);
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err;
		}
		//fmt.Fprintln(v, "Jangle");
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
			v.Clear();
			return nil;
		}); err != nil {
		return err;
	}
	if err := g.SetKeybinding("", gocui.KeySpace, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return newView(g);
		}); err != nil {
		return err;
	}
	if err := g.SetKeybinding("", gocui.KeyBackspace2, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return delView(g);
		}); err != nil {
		return err;
	}
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return nextView(g, true);
		}); err != nil {
		return err;
	}
	if err := g.SetKeybinding("", gocui.KeyArrowLeft, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return moveView(g, v, -delta, 0);
		}); err != nil {
		return err;
	}
	if err := g.SetKeybinding("", gocui.KeyArrowRight, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return moveView(g, v, delta, 0);
		}); err != nil {
		return err;
	}
	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return moveView(g, v, 0, delta);
		}); err != nil {
		return err;
	}
	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return moveView(g, v, 0, -delta);
		}); err != nil {
		return err;
	}
	if err := g.SetKeybinding("", 't', gocui.ModNone, ontop); err != nil {
		return err;
	}
	return nil;
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit;
}

func newView(g *gocui.Gui) error {
	maxX, maxY := g.Size();
	name := fmt.Sprintf("v%v", idxView);
	v, err := g.SetView(name, maxX/2-5, maxY/2-5, maxX/2+5, maxY/2+5);
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err;
		}
		v.Wrap = true;
		fmt.Fprintln(v, strings.Repeat(name+" ", 30));
	}
	if err := g.SetCurrentView(name); err != nil {
		return err;
	}
	v.BgColor = gocui.ColorRed;

	if curView >= 0 {
		cv, err := g.View(views[curView]);
		if err != nil {
			return err;
		}
		cv.BgColor = g.BgColor;
	}

	views = append(views, name);
	curView = len(views) - 1;
	idxView += 1;
	return nil;
}

func delView(g *gocui.Gui) error {
	if len(views) <= 1 {
		return nil;
	}

	if err := g.DeleteView(views[curView]); err != nil {
		return err;
	}
	views = append(views[:curView], views[curView+1:]...);

	return nextView(g, false);
}

func nextView(g *gocui.Gui, disableCurrent bool) error {
	next := curView + 1;
	if next > len(views)-1 {
		next = 0;
	}

	nv, err := g.View(views[next]);
	if err != nil {
		return err;
	}
	if err := g.SetCurrentView(views[next]); err != nil {
		return err;
	}
	nv.BgColor = gocui.ColorRed;

	if disableCurrent && len(views) > 1 {
		cv, err := g.View(views[curView]);
		if err != nil {
			return err;
		}
		cv.BgColor = g.BgColor;
	}

	curView = next;
	return nil;
}

func moveView(g *gocui.Gui, v *gocui.View, dx, dy int) error {
	name := v.Name();
	x0, y0, x1, y1, err := g.ViewPosition(name);
	if err != nil {
		return err;
	}
	if _, err := g.SetView(name, x0+dx, y0+dy, x1+dx, y1+dy); err != nil {
		return err;
	}
	return nil;
}

func ontop(g *gocui.Gui, v *gocui.View) error {
	_, err := g.SetViewOnTop(views[curView]);
	return err;
}
