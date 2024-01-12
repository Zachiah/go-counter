package main

import (
    "log"
    "strconv"
    "os"
    "fmt"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    // "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func main() {
    content, err := os.ReadFile("./count.txt")

    var originalCount int
    if err == nil {
        parsedCount, err := strconv.Atoi(string(content))
        if err != nil {
            log.Fatal(err)
        }
        originalCount = parsedCount
    } else {
        originalCount = 0
    }


    var count = originalCount

    a := app.New()
    w := a.NewWindow("Counter")

    w.Resize(fyne.NewSize(200,200))

    var countButton *widget.Button
    countButton = widget.NewButton(fmt.Sprint(count), func() {
        count += 1
        text := fmt.Sprintf("%d", count);
        countButton.SetText(text)
        if err := os.WriteFile("./count.txt", []byte(fmt.Sprint(count)), 0770); err != nil {
            log.Fatal(err)
        }
    })
    w.SetContent(countButton)

    w.ShowAndRun()
}
