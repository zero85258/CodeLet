package main

import "fmt"

// import "net/http"
import "github.com/gin-gonic/gin"
import "sync"

// deployee
// apt-get update && apt-get install -y vim less && gofmt -w utw.go && go run utw.go

type Game struct {
	Ans   []int `json:"ans"`
	Input []int `json:"input"`
}

func (g *Game) getB() (b int) {
	return g.Ans.intersect(g.Input).count()
}

func (g *Game) getA() (a int) {
	b := 0
	for i := 0; i < 4; i++ {
		if g.Ans[i] == g.Input[i] {
			b++
		}
	}
	return b
}

func worker() {
	var wg sync.WaitGroup
	wg.Add(1)

	ch := make(chan int)
	go func() {
		for {
			foo, ok := <-ch
			if !ok {
				println("done")
				wg.Done()
				return
			}
			println(foo)
		}
  }()
  
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

  wg.Wait()
}

func main() {
	router := gin.Default()

	router.GET("/framyExam/simpleSum", func(c *gin.Context) {
		a := c.Param("a")
		b := c.Param("b")

		c.JSON(200, map[string]interface{}{
			"host":     "HOST_NAME",
			"function": "framyExam",
			"method":   "simpleSum",
			"key":      "a+b",
			"value":    a + b,
		})
	})

	router.GET("/star", func(c *gin.Context) {
		for i := 1; i <= 9; i++ {
			for j := 1; j <= i; j++ {
				fmt.Printf("*")
			}
			fmt.Printf("\n")
		}
	})

	// router.GET("/framyExam/simpleSum", func(c *gin.Context) {
	// 	// a := c.Param("a")
	// 	// b := c.Param("b")

	//   g = &Game{
	//     []int{9,5,2,7},
	//     []int{6,3,5,2},
	//   }

	// 	c.JSON(200, map[string]interface{}{
	// 		"host":"HOST_NAME",
	// 		"function":"framyExam",
	// 		"method":"simpleSum",
	// 		"key":"a+b",
	// 		"value":g.getA() + "A" + g() + "B"
	// 	})
	// })

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// curl --request POST   --url 'http://srv0api.placeyapp.com/test1.0/backstage/exm1' --header 'authorization: 0123456789#0#examId'
func parserApi(){
  req , _ := http.NewRequest("POST", "http://srv0api.placeyapp.com/test1.0/backstage/exm1")
  req.Header.Set("authorization", "0123456789#0#examId")
  client := &http.Client{}
  res _ := client.Do(req)

  sources := []interface{}res[p].map(func(e){
    return e["source"]
  })

  sources

  fmt.Printf("\n")
  // api()
}

// func api()(map[string]interface{}){
// 	return map[string]interface{}{
//   "p": [
//     {
//       "at": 1506183372853,
//       "att": {
//         "id": "!566MVa9ZLh",
//         "l": 1,
//         "p_ids": "!4RnnzS2JLg#!4eQ1Ss!!9H#!4uSz!nys9H",
//         "type": "1"
//       },
//       "count": [
//         4,
//         0,
//         2004
//       ],
//       "desc": "#zoolitude  #animal  #magie  #musique  #amour  #art  #amis  #touslesjours  #nature #sport ",
//       "id": "!566MVa9ZLh",
//       "is_favorite": false,
//       "is_prv": "0",
//       "p_stk": [
//         [
//           "!4RnnzS2JLg",
//           [
//             "!4K3VIBBSOq",
//             "minizoo.01",
//             "MiniZoo - AnimalRest",
//             0
//           ]
//         ],
//         [
//           "!4eQ1Ss!!9H",
//           [
//             "!4dVlmjRN9H",
//             "sports.01",
//             "sports.01",
//             0
//           ]
//         ],
//         [
//           "!4uSz!nys9H",
//           [
//             "!4dVlITQ79H",
//             "natureelements.02",
//             "natureelements.02",
//             0
//           ]
//         ]
//       ],
//       "plc": [
//         "ChIJQ2Dro1Ir0kgRmkXB5TQEim8",
//         "64.963051",
//         "-19.0208350000",
//         "Islande",
//         "Islande",
//         "",
//         "",
//         "",
//         "",
//         "",
//         "10"
//       ],
//       "private": [
//         0,
//         []
//       ],
//       "reply": [],
//       "source": [
//         "!566MVa9ZLh",
//         1506183372853,
//         [
//           "!03_EzYUgNy",
//           "songci",
//           "宋词",
//           0
//         ],
//         "0",
//         "0"
//       ]
//     },
//     {
//       "at": 1506120480176,
//       "att": {
//         "id": "!55sN1iGkLh",
//         "l": 1,
//         "p_ids": "!4i33EXVJ5g#!4Rno8D73Lg#!4vBG8HI75g",
//         "type": "1"
//       },
//       "count": [
//         7,
//         0,
//         1526
//       ],
//       "desc": "#zoolitude # #animal  #magie  #musique   #amour  #art  #amis  #touslesjours  #nature",
//       "id": "!55sN1iGkLh",
//       "is_favorite": false,
//       "is_prv": "0",
//       "p_stk": [
//         [
//           "!4i33EXVJ5g",
//           [
//             "!4K43MNQ78q",
//             "cooleffects.01",
//             "cooleffects.01",
//             0
//           ]
//         ],
//         [
//           "!4Rno8D73Lg",
//           [
//             "!4K3VIBBSOq",
//             "minizoo.01",
//             "MiniZoo - AnimalRest",
//             0
//           ]
//         ],
//         [
//           "!4vBG8HI75g",
//           [
//             "!4K43MNQ78q",
//             "cooleffects.01",
//             "cooleffects.01",
//             0
//           ]
//         ]
//       ],
//       "plc": [
//         "ChIJQ2Dro1Ir0kgRmkXB5TQEim8",
//         "64.963051",
//         "-19.0208350000",
//         "Islande",
//         "Islande",
//         "",
//         "",
//         "",
//         "",
//         "",
//         "10"
//       ],
//       "private": [
//         0,
//         []
//       ],
//       "reply": [],
//       "source": [
//         "!55sN1iGkLh",
//         1506120480176,
//         [
//           "!03_EzYUgNy",
//           "songci",
//           "宋词",
//           0
//         ],
//         "1",
//         "0"
//       ]
//     },
//     {
//       "at": 1506088740747,
//       "att": {
//         "id": "!55knTA0gLh",
//         "l": 1,
//         "p_ids": "other392a_ch011#!4t4P6oxN9H#!4Rno3KboLg",
//         "type": "1"
//       },
//       "count": [
//         6,
//         0,
//         883
//       ],
//       "desc": "#zoolitude  #animal  #magie  #musique  #amour  #art  #amis  #touslesjours  #nature",
//       "id": "!55knTA0gLh",
//       "is_favorite": false,
//       "is_prv": "0",
//       "p_stk": [
//         [
//           "other392a_ch011",
//           [
//             "!!pfS_ZJjOB",
//             "tooz",
//             "Tooz Inc.",
//             0
//           ]
//         ],
//         [
//           "!4t4P6oxN9H",
//           [
//             "!4f466Jvc9H",
//             "animal.01",
//             "animal.01",
//             0
//           ]
//         ],
//         [
//           "!4Rno3KboLg",
//           [
//             "!4K3VIBBSOq",
//             "minizoo.01",
//             "MiniZoo - AnimalRest",
//             0
//           ]
//         ]
//       ],
//       "plc": [
//         "ChIJQ2Dro1Ir0kgRmkXB5TQEim8",
//         "64.963051",
//         "-19.0208350000",
//         "Islande",
//         "Islande",
//         "",
//         "",
//         "",
//         "",
//         "",
//         "10"
//       ],
//       "private": [
//         0,
//         []
//       ],
//       "reply": [],
//       "source": [
//         "!55knTA0gLh",
//         1506088740747,
//         [
//           "!03_EzYUgNy",
//           "songci",
//           "宋词",
//           0
//         ],
//         "0",
//         "0"
//       ]
//     },
//     {
//       "at": 1505994103617,
//       "att": {
//         "id": "!55PEdY3!Lh",
//         "l": 1,
//         "p_ids": "!4Rno7Ec!Lg#!3zXaBmzg5h#!4yzgJYMF9H",
//         "type": "1"
//       },
//       "count": [
//         6,
//         0,
//         584
//       ],
//       "desc": "#zooltude #tiger #animal  #musique  #magie  #amour  #art  #amis  #touslesjours  #nature",
//       "id": "!55PEdY3!Lh",
//       "is_favorite": false,
//       "is_prv": "0",
//       "p_stk": [
//         [
//           "!4Rno7Ec!Lg",
//           [
//             "!4K3VIBBSOq",
//             "minizoo.01",
//             "MiniZoo - AnimalRest",
//             0
//           ]
//         ],
//         [
//           "!3zXaBmzg5h",
//           [
//             "!2bgeCMR075",
//             "repo.framy.co/framy-pkg/env",
//             "repo.framy.co/framy-pkg/env",
//             0
//           ]
//         ],
//         [
//           "!4yzgJYMF9H",
//           [
//             "!4dVkDZ_s9H",
//             "cosmic.01",
//             "cosmic.01",
//             0
//           ]
//         ]
//       ],
//       "plc": [
//         "ChIJQ2Dro1Ir0kgRmkXB5TQEim8",
//         "64.963051",
//         "-19.0208350000",
//         "Islande",
//         "Islande",
//         "",
//         "",
//         "",
//         "",
//         "",
//         "10"
//       ],
//       "private": [
//         0,
//         []
//       ],
//       "reply": [],
//       "source": [
//         "!55PEdY3!Lh",
//         1505994103617,
//         [
//           "!03_EzYUgNy",
//           "songci",
//           "宋词",
//           0
//         ],
//         "0",
//         "0"
//       ]
//     },
//     {
//       "at": 1505916162792,
//       "att": {
//         "id": "!556e0zasLh",
//         "l": 1,
//         "p_ids": "!4Rno2M5RLg#!4uSuOtZR5g#!2a1ZIBnF5h",
//         "type": "1"
//       },
//       "count": [
//         2,
//         0,
//         555
//       ],
//       "desc": "#zoolitude  #animal  #magie  #musique  #amour  #art  #touslesjours  #nature",
//       "id": "!556e0zasLh",
//       "is_favorite": false,
//       "is_prv": "0",
//       "p_stk": [
//         [
//           "!4Rno2M5RLg",
//           [
//             "!4K3VIBBSOq",
//             "minizoo.01",
//             "MiniZoo - AnimalRest",
//             0
//           ]
//         ],
//         [
//           "!4uSuOtZR5g",
//           [
//             "!4K42boXw8q",
//             "props.02",
//             "props.02",
//             0
//           ]
//         ],
//         [
//           "!2a1ZIBnF5h",
//           [
//             "!2Z3N_KfN7Q",
//             "plantplanet",
//             "PlantPlanet",
//             0
//           ]
//         ]
//       ],
//       "plc": [
//         "ChIJQ2Dro1Ir0kgRmkXB5TQEim8",
//         "64.963051",
//         "-19.0208350000",
//         "Islande",
//         "Islande",
//         "",
//         "",
//         "",
//         "",
//         "",
//         "10"
//       ],
//       "private": [
//         0,
//         []
//       ],
//       "reply": [],
//       "source": [
//         "!556e0zasLh",
//         1505916162792,
//         [
//           "!03_EzYUgNy",
//           "songci",
//           "宋词",
//           0
//         ],
//         "0",
//         "0"
//       ]
//     },
//     {
//       "at": 1505857992456,
//       "att": {
//         "id": "!54tml2VVLh",
//         "l": 1,
//         "p_ids": "!4Rno9BfsLg#!2a1ZIBn!5h#!4uSz5f0F9H",
//         "type": "1"
//       },
//       "count": [
//         7,
//         0,
//         1398
//       ],
//       "desc": "#animal  #magie  #musique  #amour  #art  #zoolitude #touslesjours  #nature",
//       "id": "!54tml2VVLh",
//       "is_favorite": false,
//       "is_prv": "0",
//       "p_stk": [
//         [
//           "!4Rno9BfsLg",
//           [
//             "!4K3VIBBSOq",
//             "minizoo.01",
//             "MiniZoo - AnimalRest",
//             0
//           ]
//         ],
//         [
//           "!2a1ZIBn!5h",
//           [
//             "!2Z3N_KfN7Q",
//             "plantplanet",
//             "PlantPlanet",
//             0
//           ]
//         ],
//         [
//           "!4uSz5f0F9H",
//           [
//             "!4dVlITQ79H",
//             "natureelements.02",
//             "natureelements.02",
//             0
//           ]
//         ]
//       ],
//       "plc": [
//         "ChIJQ2Dro1Ir0kgRmkXB5TQEim8",
//         "64.963051",
//         "-19.0208350000",
//         "Islande",
//         "Islande",
//         "",
//         "",
//         "",
//         "",
//         "",
//         "10"
//       ],
//       "private": [
//         0,
//         []
//       ],
//       "reply": [],
//       "source": [
//         "!54tml2VVLh",
//         1505857992456,
//         [
//           "!03_EzYUgNy",
//           "songci",
//           "宋词",
//           0
//         ],
//         "1",
//         "0"
//       ]
//     },
//     {
//       "at": 1505809851978,
//       "att": {
//         "id": "!54iJEJi7Lh",
//         "l": 1,
//         "p_ids": "!4dW3G4oR9H#!4Rno6G9RLg",
//         "type": "1"
//       },
//       "count": [
//         5,
//         0,
//         829
//       ],
//       "desc": "#zoolitude  #animal  #magie  #musique  #amour  #art  #amis  #nature",
//       "id": "!54iJEJi7Lh",
//       "is_favorite": false,
//       "is_prv": "0",
//       "p_stk": [
//         [
//           "!4dW3G4oR9H",
//           [
//             "!4dVkDZ_s9H",
//             "cosmic.01",
//             "cosmic.01",
//             0
//           ]
//         ],
//         [
//           "!4Rno6G9RLg",
//           [
//             "!4K3VIBBSOq",
//             "minizoo.01",
//             "MiniZoo - AnimalRest",
//             0
//           ]
//         ]
//       ],
//       "plc": [
//         "ChIJQ2Dro1Ir0kgRmkXB5TQEim8",
//         "64.963051",
//         "-19.0208350000",
//         "Islande",
//         "Islande",
//         "",
//         "",
//         "",
//         "",
//         "",
//         "10"
//       ],
//       "private": [
//         0,
//         []
//       ],
//       "reply": [],
//       "source": [
//         "!54iJEJi7Lh",
//         1505809851978,
//         [
//           "!03_EzYUgNy",
//           "songci",
//           "宋词",
//           0
//         ],
//         "0",
//         "0"
//       ]
//     },
//     {
//       "at": 1505746857674,
//       "att": {
//         "id": "!54UHzI6!Lh",
//         "l": 1,
//         "p_ids": "!3cmznP0N5h#!4dW3D9Ig9H#!4Rno4J4RLg",
//         "type": "1"
//       },
//       "count": [
//         29,
//         1,
//         5355
//       ],
//       "desc": "#zoolitude  #animal  #magie  #musique  #amour  #art  #spectacle  #touslesjours  #nature",
//       "id": "!54UHzI6!Lh",
//       "is_favorite": false,
//       "is_prv": "0",
//       "p_stk": [
//         [
//           "!3cmznP0N5h",
//           [
//             "!2bgeCMR075",
//             "repo.framy.co/framy-pkg/env",
//             "repo.framy.co/framy-pkg/env",
//             0
//           ]
//         ],
//         [
//           "!4dW3D9Ig9H",
//           [
//             "!4dVkDZ_s9H",
//             "cosmic.01",
//             "cosmic.01",
//             0
//           ]
//         ],
//         [
//           "!4Rno4J4RLg",
//           [
//             "!4K3VIBBSOq",
//             "minizoo.01",
//             "MiniZoo - AnimalRest",
//             0
//           ]
//         ]
//       ],
//       "plc": [
//         "ChIJQ2Dro1Ir0kgRmkXB5TQEim8",
//         "64.963051",
//         "-19.0208350000",
//         "Islande",
//         "Islande",
//         "",
//         "",
//         "",
//         "",
//         "",
//         "10"
//       ],
//       "private": [
//         0,
//         []
//       ],
//       "reply": [
//         [
//           [
//             "!54ghtyCZ!3",
//             "faizakhushi",
//             "faizakhushi",
//             0
//           ],
//           "nice",
//           1505804419460
//         ]
//       ],
//       "source": [
//         "!54UHzI6!Lh",
//         1505746857674,
//         [
//           "!03_EzYUgNy",
//           "songci",
//           "宋词",
//           0
//         ],
//         "1",
//         "0"
//       ]
//     },
//     {
//       "at": 1505663006343,
//       "att": {
//         "id": "!54AHo4L3Lh",
//         "l": 1,
//         "p_ids": "!4dW3E7mo9H#!4Rno!QXBLg",
//         "type": "1"
//       },
//       "count": [
//         4,
//         0,
//         450
//       ],
//       "desc": "#animal  #musique  #art  #touslesjours  #nature",
//       "id": "!54AHo4L3Lh",
//       "is_favorite": false,
//       "is_prv": "0",
//       "p_stk": [
//         [
//           "!4dW3E7mo9H",
//           [
//             "!4dVkDZ_s9H",
//             "cosmic.01",
//             "cosmic.01",
//             0
//           ]
//         ],
//         [
//           "!4Rno!QXBLg",
//           [
//             "!4K3VIBBSOq",
//             "minizoo.01",
//             "MiniZoo - AnimalRest",
//             0
//           ]
//         ]
//       ],
//       "plc": [
//         "ChIJQ2Dro1Ir0kgRmkXB5TQEim8",
//         "64.963051",
//         "-19.0208350000",
//         "Islande",
//         "Islande",
//         "",
//         "",
//         "",
//         "",
//         "",
//         "10"
//       ],
//       "private": [
//         0,
//         []
//       ],
//       "reply": [],
//       "source": [
//         "!54AHo4L3Lh",
//         1505663006343,
//         [
//           "!03_EzYUgNy",
//           "songci",
//           "宋词",
//           0
//         ],
//         "0",
//         "0"
//       ]
//     }
//   ]
// }
// }
