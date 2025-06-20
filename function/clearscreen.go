package function

import (
    "os"
    "os/exec"
    "runtime"
)

func ClearScreen() {
    // Menghapus layar terminal

    var cmd *exec.Cmd
    switch runtime.GOOS {
    case "windows":
        cmd = exec.Command("cmd", "/c", "cls")
    default:
        cmd = exec.Command("clear")
    }

    cmd.Stdout = os.Stdout
    cmd.Run()
}