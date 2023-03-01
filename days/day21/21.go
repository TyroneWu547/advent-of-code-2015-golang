package day21

import (
	utils "aoc/utils"
	"math"
	"strconv"
	"strings"
)

type character struct {
	health int
	damage int
	armor  int
}

type item struct {
	cost   int
	damage int
	armor  int
}

func getWeapons() []item {
	var dagger item = item{8, 4, 0}
	var shortsword item = item{10, 5, 0}
	var warhammer item = item{25, 6, 0}
	var longsword item = item{40, 7, 0}
	var greataxe item = item{74, 8, 0}

	return []item{dagger, shortsword, warhammer, longsword, greataxe}
}

func getArmors() []item {
	var none item = item{0, 0, 0}
	var leather item = item{13, 0, 1}
	var chainmail item = item{31, 0, 2}
	var splintmail item = item{53, 0, 3}
	var brandedmail item = item{75, 0, 4}
	var platemail item = item{102, 0, 5}

	return []item{none, leather, chainmail, splintmail, brandedmail, platemail}
}

func getRings() []item {
	var none item = item{0, 0, 0}
	var damage1 item = item{25, 1, 0}
	var damage2 item = item{50, 2, 0}
	var damage3 item = item{100, 3, 0}
	var defense1 item = item{20, 0, 1}
	var defense2 item = item{40, 0, 2}
	var defense3 item = item{80, 0, 3}

	return []item{none, none, damage1, damage2, damage3, defense1, defense2, defense3}
}

func max(left int, right int) int {
	if left < right {
		return right
	}
	return left
}

func willWin(me character, boss character) bool {
	var myDmg int = max(me.damage-boss.armor, 1)
	var bossDmg int = max(boss.damage-me.armor, 1)

	var roundsSurvive int = int(math.Ceil(float64(me.health) / float64(bossDmg)))
	var bossSurvive int = int(math.Ceil(float64(boss.health) / float64(myDmg)))

	return roundsSurvive >= bossSurvive
}

func partOne(input string) string {
	var tokens []string = strings.Fields(input)
	health, _ := strconv.Atoi(tokens[2])
	damage, _ := strconv.Atoi(tokens[4])
	armor, _ := strconv.Atoi(tokens[6])
	var boss character = character{
		health,
		damage,
		armor,
	}

	var weapons []item = getWeapons()
	var armors []item = getArmors()
	var rings []item = getRings()

	var minCost int = int(^uint32(0))
	for _, w := range weapons {
		for _, a := range armors {
			for i, r1 := range rings {
				for j, r2 := range rings {
					if i == j {
						continue
					}

					var me character = character{
						100,
						w.damage + r1.damage + r2.damage,
						a.armor + r1.armor + r2.armor,
					}

					var cost int = w.cost + a.cost + r1.cost + r2.cost
					if willWin(me, boss) && cost < minCost {
						minCost = cost
					}
				}
			}
		}
	}

	return strconv.Itoa(minCost)
}

func partTwo(input string) string {
	var tokens []string = strings.Fields(input)
	health, _ := strconv.Atoi(tokens[2])
	damage, _ := strconv.Atoi(tokens[4])
	armor, _ := strconv.Atoi(tokens[6])
	var boss character = character{
		health,
		damage,
		armor,
	}

	var weapons []item = getWeapons()
	var armors []item = getArmors()
	var rings []item = getRings()

	var maxCost int = 0
	for _, w := range weapons {
		for _, a := range armors {
			for i, r1 := range rings {
				for j, r2 := range rings {
					if i == j {
						continue
					}

					var me character = character{
						100,
						w.damage + r1.damage + r2.damage,
						a.armor + r1.armor + r2.armor,
					}

					var cost int = w.cost + a.cost + r1.cost + r2.cost
					if !willWin(me, boss) && cost > maxCost {
						maxCost = cost
					}
				}
			}
		}
	}

	return strconv.Itoa(maxCost)
}

func init() {
	var day utils.Day = utils.Day{
		DayNum:  21,
		PartOne: partOne,
		PartTwo: partTwo,
	}
	utils.RegisterDay("Day 21", &day)
}
