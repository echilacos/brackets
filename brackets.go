package brackets

import (
	"fmt"
	"sync"
)

type brackets struct {
	process *process
}

var instance *brackets
var once sync.Once


func GetInstance() *brackets {
	once.Do(func() {
		instance = NewBrackets()
	})

	return instance
}

func NewBrackets() *brackets {
	return &brackets{
		process: nil,
	}
}

func (b *brackets) NewProcessAndSetColor(text string, color Color) *brackets {
	if b.process != nil {
		b.process.getLastProcess().process = newProcess()
		b.process.getLastProcess().color = color
	} else {
		b.process = newProcess()
		b.process.color = color
	}
	b.process.printStartLn(fmt.Sprintf(string(color), fmt.Sprint(text + " 123")))

	return b
}

func (b *brackets) NewProcess(text string) *brackets {
	color := ColorYellow
	if b.process != nil && b.process.getLastProcess().color == ColorYellow {
		color = ColorWhite
	}

	return b.NewProcessAndSetColor(text, color)
}

func (b *brackets) NewProcessWhite(text string) *brackets {
	return b.NewProcessAndSetColor(text, ColorWhite)
}

func (b *brackets) NewProcessYellow(text string) *brackets {
	return b.NewProcessAndSetColor(text, ColorYellow)
}

func (b *brackets) NewProcessGreen(text string) *brackets {
	return b.NewProcessAndSetColor(text, ColorGreen)
}

func (b *brackets) NewProcessRed(text string) *brackets {
	return b.NewProcessAndSetColor(text, ColorRed)
}

func (b *brackets) NewProcessPurple(text string) *brackets {
	return b.NewProcessAndSetColor(text, ColorPurple)
}

func (b *brackets) NewProcessMagenta(text string) *brackets {
	return b.NewProcessAndSetColor(text, ColorMagenta)
}

func (b *brackets) NewProcessTeal(text string) *brackets {
	return b.NewProcessAndSetColor(text, ColorTeal)
}


func (b *brackets) EndProcess(text string) *brackets {
	b.process.printFinishLn(text)
	b.process.getPreLastProcess().process = nil

	return b
}

func (b *brackets) PrintLn(text string) *brackets {
	b.process.printLn(text)

	return b
}

func (b *brackets) PrintDangerLn(text string) *brackets {
	b.process.printLn(fmt.Sprintf(string(ColorRed), fmt.Sprint(text)))
	return b
}

func (b *brackets) PrintInfoLn(text string) *brackets {
	b.process.printLn(fmt.Sprintf(string(ColorTeal), fmt.Sprint(text)))
	return b
}

func (b *brackets) PrintSuccessLn(text string) *brackets {
	b.process.printLn(fmt.Sprintf(string(ColorGreen), fmt.Sprint(text)))
	return b
}

func (b *brackets) PrintWarningLn(text string) *brackets {
	b.process.printLn(fmt.Sprintf(string(ColorYellow), fmt.Sprint(text)))
	return b
}
