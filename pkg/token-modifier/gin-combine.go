package token_modifier

import "dex-trades-parser/pkg/jwtoken"

type GinCombineGetter struct {
	getters []Getter
}

func NewGinCombineGetter(getters ...Getter) GinCombineGetter {
	return GinCombineGetter{
		getters: getters,
	}
}

func (g GinCombineGetter) Get() (tokens jwtoken.Tokens) {
	for _, getter := range g.getters {
		t := getter.Get()
		if t.HasAccessToken() {
			tokens.AccessToken = t.AccessToken
			return
		}
	}

	return
}

type GinCombineSetter struct {
	setters []Setter
}

func NewGinCombineSetter(setters ...Setter) GinCombineSetter {
	return GinCombineSetter{
		setters: setters,
	}
}

func (s GinCombineSetter) Set(tokens jwtoken.Tokens) {
	for _, setter := range s.setters {
		setter.Set(tokens)
	}
}

type GinCombineRemover struct {
	removers []Remover
}

func NewGinCombineRemover(removers ...Remover) GinCombineRemover {
	return GinCombineRemover{
		removers: removers,
	}
}

func (s GinCombineRemover) Remove() {
	for _, remover := range s.removers {
		remover.Remove()
	}
}
