package brackets

import "fmt"

const StartProcessSymbol = "┌"
const StepProcessSymbol = "│"
const FinishProcessSymbol = "└"

type process struct {
	color   Color
	process *process
}

func newProcess() *process {
	return &process{
		color:   ColorYellow,
		process: nil,
	}
}


func (p *process) setProcessColor(color Color) {
	p.color = color
}

func (p process) getColorSymbol(symbol string) string {
	return fmt.Sprintf(string(p.color), fmt.Sprint(symbol))
}

func (p process) getPrefixLine(symbol string) string {
	if p.process != nil {
		return p.getColorSymbol(StepProcessSymbol) + p.process.getPrefixLine(symbol)
	} else {
		return p.getColorSymbol(symbol)
	}
}

func (p *process) printStartLn(text string) {
	fmt.Println(p.getPrefixLine(StartProcessSymbol) + " " + text)
}

func (p *process) printLn(text string) {
	fmt.Println(p.getPrefixLine(StepProcessSymbol) + " " + text)
}

func (p *process) printFinishLn(text string) {
	fmt.Println(p.getPrefixLine(FinishProcessSymbol) + " " + text)
}



func (p process) printProcessSymbol(symbol string) {

}

func (p *process) getPreLastProcess() *process {
	if p.process != nil && p.process.process != nil {
		return p.process.getPreLastProcess()
	}

	return p
}


func (p *process) getLastProcess() *process {
	if p.process != nil {
		return p.process.getLastProcess()
	}

	return p
}
