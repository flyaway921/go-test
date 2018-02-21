package main

import (
	//"fmt"
	"time"

	"github.com/nsf/termbox-go"
)

func main() {
	//try to draw something
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	//设置输入模式
	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	//获取当前窗口大小
	w, h := termbox.Size()
	fmt.Printf("width=%d height=%d", w, h)
	//绘制边框
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(0, 0, '┏', termbox.ColorRed, termbox.ColorBlack)
	termbox.SetCell(w-1, h-1, '┛', termbox.ColorRed, termbox.ColorBlack)
	termbox.SetCell(0, h-1, '┗', termbox.ColorRed, termbox.ColorBlack)
	termbox.SetCell(w-1, 0, '┓', termbox.ColorRed, termbox.ColorBlack)
	for x := 1; x < w-1; x++ {
		termbox.SetCell(x, 0, '━', termbox.ColorRed, termbox.ColorBlack)
		termbox.SetCell(x, h-1, '━', termbox.ColorRed, termbox.ColorBlack)
	}
	for y := 1; y < h-1; y++ {
		termbox.SetCell(w-1, y, '┃', termbox.ColorRed, termbox.ColorBlack)
		termbox.SetCell(0, y, '┃', termbox.ColorRed, termbox.ColorBlack)
	}
	//画蛇的初始位置
	var bodyUnit rune = '■'
	start := []int{w / 3, h * 2 / 3}
	end := []int{start[0] + 3, start[1]}
	for x := start[0]; x <= end[0]; x++ {
		termbox.SetCell(x, start[1], bodyUnit, termbox.ColorGreen, termbox.ColorBlack)
	}

	termbox.Flush()
	time.Sleep(5 * time.Second)
}
