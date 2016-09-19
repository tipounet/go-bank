package model

// Goversion Image de la version de GO avec l'os et l'architecture
type Goversion struct {
	Version string
	OS      string
	Arch    string
}

// Version : Version detaillée de l'application et de go
type Version struct {
	Application string    `json:"application"`
	GoVersion   Goversion `json:"goversion"`
}
