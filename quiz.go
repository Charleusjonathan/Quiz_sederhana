package main

import "fmt"

func main() {
    var jawaban string
    var ulang string

    for {
        var nilai int = 0 // reset nilai setiap quiz

        fmt.Println("===== PROGRAM QUIZ SEDERHANA =====")
        fmt.Println()

        // Soal 1
        fmt.Println("1. Ibukota Indonesia adalah?")
        fmt.Print("Jawaban: ")
        fmt.Scanln(&jawaban)

        if jawaban == "Jakarta" || jawaban == "jakarta" {
            fmt.Println("Benar!")
            nilai++
        } else {
            fmt.Println("Salah! Jawaban yang benar: Jakarta")
        }

        fmt.Println()

        // Soal 2
        fmt.Println("2. 5 + 7 = ?")
        fmt.Print("Jawaban: ")
        fmt.Scanln(&jawaban)

        if jawaban == "12" {
            fmt.Println("Benar!")
            nilai++
        } else {
            fmt.Println("Salah! Jawaban yang benar: 12")
        }

        fmt.Println()

        // Soal 3
        fmt.Println("3. Bahasa pemrograman yang digunakan dalam tugas ini adalah?")
        fmt.Print("Jawaban: ")
        fmt.Scanln(&jawaban)

        if jawaban == "Go" || jawaban == "Golang" || jawaban == "go" {
            fmt.Println("Benar!")
            nilai++
        } else {
            fmt.Println("Salah! Jawaban yang benar: Go / Golang")
        }

        fmt.Println()

        // Hasil
        fmt.Println("===== HASIL AKHIR =====")
        fmt.Printf("Skor kamu: %d dari 3\n", nilai)

        // Logika berhenti / ulang
        if nilai == 3 {
            fmt.Println("Keren! Semua jawaban benar!")
            break
        }

        fmt.Print("Skor belum sempurna. Ingin mengulang? (Y/N): ")
        fmt.Scanln(&ulang)

        if ulang != "Y" && ulang != "y" {
            fmt.Println("Terima kasih sudah mencoba")
            break
        }

        fmt.Println("\nMengulang quiz...")
    }
}