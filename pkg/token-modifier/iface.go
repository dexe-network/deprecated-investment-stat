package token_modifier

import "dex-trades-parser/pkg/jwtoken"

type Modifier interface {
	Getter
	Setter
	Remover
}

type Getter interface {
	Get() (tokens jwtoken.Tokens)
}

type Setter interface {
	Set(tokens jwtoken.Tokens)
}

type Remover interface {
	Remove()
}
