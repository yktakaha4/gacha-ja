package gacha

import (
	"errors"
)

type Player struct {
	tickets int // ガチャ券の枚数
	coin    int // コイン
}

func NewPlayer(tickets, coin int) *Player {
	return &Player{tickets: tickets, coin: coin}
}

// DrawableNum プレイヤーが行えるガチャの回数
func (p *Player) DrawableNum() int {
	// ガチャ券は1枚で1回、コインは10枚で1回ガチャが行える
	return p.tickets + p.coin/10
}

func (p *Player) draw(n int) error {

	if p.DrawableNum() < n {
		return errors.New("ガチャ券またはコインが不足しています")
	}

	// ガチャ券から優先的に使う
	if p.tickets > n {
		p.tickets -= n
		return nil
	}

	p.coin -= (n - p.tickets) * 10 // 1回あたり10枚消費する
	p.tickets = 0

	return nil
}
