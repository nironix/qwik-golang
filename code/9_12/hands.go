package main

import "fmt"

type sall func(int, int, int) int
type are struct {
	len int
	bre int
}

type mon_sal struct {
	mon1 sal
}

type sal struct {
	basic  int
	hra    int
	ta     int
	salary sall
}

func main() {
	fmt.Println("---q1---")
	ar := are{4, 5}
	fmt.Println("area of rectangle is ", ar.len*ar.bre)

	fmt.Println("\n---q2---")
	/*mon := mon_sal{sal{14000, 600, 400,
		func(hra int, basic int, ta int) int {
			return (basic + hra + ta)
		}},

	/*sal{15000, 850, 900, func(hra int, basic int, ta int) int {
		return basic + hra + ta
	}}, sal{17000, 200, 150, func(hra int, basic int, ta int) int {
		return basic + hra + ta
	}}}
	fmt.Println("SALARY\n MONTH 1\nBASIC : ", mon.mon1.basic, "\t HRA : ", mon.mon1.hra, "\tTA : ", mon.mon1.ta, "\ttotal : ", mon.mon1.salary)
	//fmt.Println(" MONTH 2\nBASIC : ", mon.mon2.basic, "\t HRA : ", mon.mon2.hra, "\tTA : ", mon.mon2.ta, "\ttotal : ", mon.mon2.salary)
	//fmt.Println(" MONTH 3\nBASIC : ", mon.mon3.basic, "\t HRA : ", mon.mon3.hra, "\tTA : ", mon.mon3.ta, "\ttotal : ", mon.mon3.salary)
	*/
	tt := sal{basic: 1400, hra: 400, ta: 200, salary: func(basic int, hra int, ta int) int { return hra + basic + ta }}
	fmt.Println(tt.salary(tt.basic, tt.hra, tt.ta))
	fmt.Println(tt.salary)
	fmt.Println("\n---q4---")

	anim := map[string]string{
		"tiger":    "wild",
		"lion":     "wild",
		"dog":      "domestic",
		"cat":      "domestic",
		"bear":     "wild",
		"goat":     "domestic",
		"elephant": "wild",
		"cow":      "domestic",
		"chicken":  "domestic",
		"cheetah":  "wild",
	}
	count := map[string]int{
		"wild":     0,
		"domestic": 0,
	}
	//var t string
	_, ok := anim["cow"]
	fmt.Println(ok)

	for _, j := range anim {
		if j == "wild" {
			count["wild"] = count["wild"] + 1
		} else if j == "domestic" {
			count["domestic"] = count["domestic"] + 1
		}
	}
	for i, j := range count {
		fmt.Println(i, " : ", j)
	}

}
