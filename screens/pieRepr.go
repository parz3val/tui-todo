package screens

import (
	"github.com/gizak/termui/v3/widgets"
)

func BarProgress() *widgets.StackedBarChart {

	sbc := widgets.NewStackedBarChart()
	sbc.Title = "Student's Marks: X-Axis=Name, Y-Axis=Grade% (Math, English, Science, Computer Science)"
	sbc.Labels = []string{"Ken", "Rob", "Dennis", "Linus"}

	sbc.Data = make([][]float64, 4)
	sbc.Data[0] = []float64{90, 85, 90, 80}
	sbc.Data[1] = []float64{70, 85, 75, 60}
	sbc.Data[2] = []float64{75, 60, 80, 85}
	sbc.Data[3] = []float64{100, 100, 100, 100}
	sbc.SetRect(5, 5, 100, 30)
	sbc.BarWidth = 5
	return sbc
}
