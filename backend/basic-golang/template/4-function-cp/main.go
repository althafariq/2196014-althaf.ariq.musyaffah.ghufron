package main

import (
	"bytes"
	"html/template"
)

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan function pada template.
// Lengkapi function CalculateScore yang berfungsi untuk menjumlahkan total score dari users
// Lengkapi function ExecuteToByteBuffer dan textTemplate sehingga template dapat mencetak objek Leaderboard dengan ketentuan:
// Lakukan looping untuk setiap user
// Pada setiap loop, cetak "Nama: Score", contoh: "Roger: 1000"
// Pada bagian terakhir, cetak "Total Score: total", contoh: "Total Score: 50000"

type UserRank struct {
	Name  string
	Email string
	Rank  int
	Score int
}

type Leaderboard struct {
	Users []*UserRank
}

func CalculateScore(leaderboard Leaderboard) int {
	// TODO: answer here
	var sum int
	for _, user := range leaderboard.Users {
		sum += user.Score
	}
	return sum
}

func ExecuteToByteBuffer(leaderboard Leaderboard) ([]byte, error) {
	var textTemplate string
	// TODO: answer here
	funcMap := template.FuncMap{
		"CalculateScore": CalculateScore,
	}

	textTemplate = `{{ range $user := .Users }}{{ $user.Name }}: {{ $user.Score }}{{ end }}Total Score: {{ CalculateScore . }}`

	tmpl, err := template.New("leaderboard").Funcs(funcMap).Parse(textTemplate)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, leaderboard)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
