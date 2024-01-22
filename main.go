package main

import (
  "fmt"

  "github.com/charmbracelet/lipgloss"
  "github.com/charmbracelet/lipgloss/table"

)

// https://www.ibm.com/docs/en/sdse/6.4.0?topic=configuration-ascii-characters-from-33-126
func main() {
  m, n := 12, 8
  chars := make([][]string, m)
  v := 33 
  max := 126
  for i := 0; i < m; i++ {
    chars[i] = make([]string, n)
    for j := 0; j < n; j++ {
      if v < max {
        chars[i][j] = fmt.Sprintf("%v: %c",v, v)
      }
      v++
    }
  }
   
  t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderRow(false).
		BorderColumn(false).
		Rows(chars...).
		StyleFunc(func(row, col int) lipgloss.Style {
			return lipgloss.NewStyle().Padding(0, 1)
		})


	fmt.Println(lipgloss.JoinVertical(lipgloss.Right, 
    lipgloss.JoinHorizontal(lipgloss.Center, 
    t.Render())) + "\n")

}
