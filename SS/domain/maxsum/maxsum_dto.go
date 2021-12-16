// DTO onde declaramos nossas variáveis

package maxsum

type Numbers struct {
	List []int `json:"list"`
}

type Result struct {
	Sum       int   `json:"sum"`
	Positions []int `json:"positions"`
}
