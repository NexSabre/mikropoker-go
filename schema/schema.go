package schema

type SessionCreate struct {
	Name string `json:"name"`
}

type SessionReveal struct {
	Reveal bool `json:"reveal"`
}

type Action struct {
	Action string `json:"action"`
}

type UserCreate struct {
	Username string `json:"username"`
}

type UserPoints struct {
	Username string  `json:"username"`
	Salle    float32 `json:"salle"`
}
