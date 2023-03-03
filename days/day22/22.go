package day22

import (
	utils "aoc/utils"
	"strconv"
	"strings"
)

type player struct {
	hp       int
	mana     int
	shield   int
	recharge int
}

type boss struct {
	hp     int
	damage int
	poison int
}

type state struct {
	me   player
	boss boss
}

func bossTurn(m *player, b *boss) {
	if b.poison > 0 {
		b.hp -= 3
		b.poison--
	}
	if m.recharge > 0 {
		m.mana += 101
		m.recharge--
	}

	if b.hp <= 0 {
		return
	}

	if m.shield > 0 {
		m.hp -= b.damage - 7
		m.shield--
	} else {
		m.hp -= b.damage
	}
}

func magicMissle(m player, b boss) (player, boss) {
	m.mana -= 53
	b.hp -= 4

	if b.hp <= 0 {
		return m, b
	}

	bossTurn(&m, &b)
	return m, b
}

func drain(m player, b boss) (player, boss) {
	m.mana -= 73
	m.hp += 2
	b.hp -= 2

	if b.hp <= 0 {
		return m, b
	}

	bossTurn(&m, &b)
	return m, b
}

func shield(m player, b boss) (player, boss) {
	m.mana -= 113
	m.shield = 6

	bossTurn(&m, &b)
	return m, b
}

func poison(m player, b boss) (player, boss) {
	m.mana -= 173
	b.poison = 6

	bossTurn(&m, &b)
	return m, b
}

func recharge(m player, b boss) (player, boss) {
	m.mana -= 229
	m.recharge = 5

	bossTurn(&m, &b)
	return m, b
}

func dfs(m player, b boss, cache map[state]int, partTwo bool) int {
	if b.poison > 0 {
		b.hp -= 3
		b.poison--
	}
	if m.recharge > 0 {
		m.mana += 101
		m.recharge--
	}
	if m.shield > 0 {
		m.shield--
	}
	if partTwo {
		m.hp--
	}

	if b.hp <= 0 {
		return 0
	}

	var state state = state{m, b}
	val, ok := cache[state]
	if ok {
		return val
	}

	var minSpent int = int(^uint32(0))
	// magic missle
	if m.mana >= 53 {
		missleM, missleB := magicMissle(m, b)
		if missleM.hp > 0 {
			var spent int = dfs(missleM, missleB, cache, partTwo)
			if (spent + 53) < minSpent {
				minSpent = spent + 53
			}
		}
	}
	// drain
	if m.mana >= 73 {
		drainM, drainB := drain(m, b)
		if drainM.hp > 0 {
			var spent int = dfs(drainM, drainB, cache, partTwo)
			if (spent + 73) < minSpent {
				minSpent = spent + 73
			}
		}
	}
	// shield
	if m.shield == 0 && m.mana >= 113 {
		shieldM, shieldB := shield(m, b)
		if shieldM.hp > 0 {
			var spent int = dfs(shieldM, shieldB, cache, partTwo)
			if (spent + 113) < minSpent {
				minSpent = spent + 113
			}
		}
	}
	// poison
	if b.poison == 0 && m.mana >= 173 {
		poisonM, poisonB := poison(m, b)
		if poisonM.hp > 0 {
			var spent int = dfs(poisonM, poisonB, cache, partTwo)
			if (spent + 173) < minSpent {
				minSpent = spent + 173
			}
		}
	}
	// recharge
	if m.recharge == 0 && m.mana >= 229 {
		rechargeM, rechargeB := recharge(m, b)
		if rechargeM.hp > 0 {
			var spent int = dfs(rechargeM, rechargeB, cache, partTwo)
			if (spent + 229) < minSpent {
				minSpent = spent + 229
			}
		}
	}

	cache[state] = minSpent
	return minSpent
}

func partOne(input string) string {
	var tokens []string = strings.Fields(input)
	hp, _ := strconv.Atoi(tokens[2])
	dmg, _ := strconv.Atoi(tokens[4])
	var boss = boss{hp, dmg, 0}
	var me player = player{50, 500, 0, 0}

	var cache map[state]int = map[state]int{}
	return strconv.Itoa(dfs(me, boss, cache, false))
}

func partTwo(input string) string {
	var tokens []string = strings.Fields(input)
	hp, _ := strconv.Atoi(tokens[2])
	dmg, _ := strconv.Atoi(tokens[4])
	var boss = boss{hp, dmg, 0}
	var me player = player{50, 500, 0, 0}

	var cache map[state]int = map[state]int{}
	return strconv.Itoa(dfs(me, boss, cache, true))
}

func init() {
	var day utils.Day = utils.Day{
		DayNum:  22,
		PartOne: partOne,
		PartTwo: partTwo,
	}
	utils.RegisterDay("Day 22", &day)
}
