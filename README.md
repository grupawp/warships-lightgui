# Warships-LightGUI

Warships-LightGUI provides an easy to use graphical user interface 
for the `Warships Online` game.

<img src="doc/warships.png" width=50%>

This package is suggested for beginners. If you want a more sophisticated
GUI, please check out the `Warships-GUI` package:
http://github.com/grupawp/warships-gui

## Installation

```
go get github.com/grupawp/warships-lightgui
```

## Quick Start

To initialize the *board* use the **New()** method. **NewConfig()** will 
load the default configuration.

```go
board := gui.New(gui.NewConfig())
board.Display()
```

To customize colours and characters used to indicate ships, misses, etc, 
use methods available on the `config` instance. 

```go
board := gui.New(
    gui.ConfigParams().
        HitChar('#').
        HitColor(color.FgRed).
        BorderColor(color.BgRed).
        RulerTextColor(color.BgYellow).
        NewConfig())
board.Display()
```

## Documentation
https://pkg.go.dev/github.com/grupawp/warships-lightgui
