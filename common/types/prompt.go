package types

import "github.com/c-bata/go-prompt"

type Prompt struct {
	CollectionPrompt *prompt.Prompt
	TxQueryPrompt    *prompt.Prompt
	BlockQueryPrompt *prompt.Prompt
	HashPrompt       *prompt.Prompt
	BlockNumPrompt   *prompt.Prompt
}
