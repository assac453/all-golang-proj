package main

import (
	"fmt"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func main() {
    // Инициализация GLFW
    if err := glfw.Init(); err != nil {
        panic(err)
    }
    defer glfw.Terminate()

    // Создание окна
    window, err := glfw.CreateWindow(800, 600, "Пример GLFW", nil, nil)
    if err != nil {
        panic(err)
    }
    defer window.Destroy()

    // Устанавливаем обработчик событий для нажатия кнопки мыши
    window.SetMouseButtonCallback(mouseButtonCallback)

    // Главный цикл
    for !window.ShouldClose() {
        // Здесь будет ваш код отрисовки

        // Проверка событий
        glfw.PollEvents()
    }
}

func mouseButtonCallback(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
    if button == glfw.MouseButtonLeft && action == glfw.Press {
        xpos, ypos := w.GetCursorPos()
		fmt.Println(xpos, ypos)

        // Здесь можно добавить логику для обработки нажатия левой кнопки мыши
        // Например, проверить, была ли нажата кнопка "Старт", используя координаты мыши (xpos, ypos)
    }
}


func drawButton(){
	gl.PushMatrix()
	
	gl.PopMatrix()
}