package main

import "fmt"

func test_arrays() {
	var rgbImage [1080][1920][3]uint8
	line := rgbImage[2]      // 3-я строка в изображении
	pixel := rgbImage[2][3]  // 4-й пиксель в третьей строке изображения
	red := rgbImage[2][3][1] // значение синей компоненты (второй байт) 4-го пикселя в третьей строке изображения
	fmt.Println(line)
	fmt.Println(pixel)
	fmt.Println(red)
}
