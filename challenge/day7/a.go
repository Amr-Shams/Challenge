package day7 
import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/Amr-Shams/aoc2024/challenge"
    "github.com/Amr-Shams/aoc2024/util"
    "sort"
    "strings"
)


func aCommand() *cobra.Command {
    return &cobra.Command{
        Use:   "a",
        Short: "Run day7 challenge a",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
        },
    }
}
type hand struct {
    bid int 
    hand string 
    score int 
}
func (h *hand) parseHand(line string){
    h.hand= strings.Split(line," ")[0]
    h.bid=util.Atoi(strings.Split(line," ")[1])
}
func mapRune(r byte)int{ 
    //A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2 the order of the cards is A>K>Q>J>T>9>8>7>6>5>4>3>2
    switch r{
        case 'A':
            return 14
        case 'K':
            return 13
        case 'Q':
            return 12
        case 'J':
            return 11
        case 'T':
            return 10
        default:
            return int(r-'0')
    }
}
const (
    HighCard = iota 
    OnePair
    TwoPair
    ThreeOfAKind
    FullHouse
    FourOfAKind
    FiveOfAKind
)

func  compare(a,b hand, mapRune func(byte)int)int{
    if a.score>b.score{
        return 1
    }
    if a.score<b.score{
        return -1
    }
    // break tie with the highest card
    for i:=0;i<len(a.hand);i++{
        if mapRune(a.hand[i])>mapRune(b.hand[i]){
            return 1
        }
        if mapRune(a.hand[i])<mapRune(b.hand[i]){
            return -1
        }
    }
    return 0
}
func rankIt(cnt map[rune]int, h *hand) {
    // Create a list of values
    values := make([]int, 0, len(cnt))
    for _, v := range cnt {
        values = append(values, v)
    }
    sort.Ints(values)

    // Custom function to compare slices
    equals := func(a, b []int) bool {
        if len(a) != len(b) {
            return false
        }
        for i := range a {
            if a[i] != b[i] {
                return false
            }
        }
        return true
    }

    if equals(values, []int{5}) {
        h.score = FiveOfAKind
        return
    }
    if equals(values, []int{1, 4}) {
        h.score = FourOfAKind
        return
    }
    if equals(values, []int{2, 3}) {
        h.score = FullHouse
        return
    }
    if equals(values, []int{1, 1, 3}) {
        h.score = ThreeOfAKind
        return
    }
    if equals(values, []int{1, 2, 2}) {
        h.score = TwoPair
        return
    }
    if equals(values, []int{1, 1, 1, 2}) {
        h.score = OnePair
        return
    }
    h.score = HighCard
}

  
func (h *hand) scoreHand(){
    cnt:=map[rune]int{}
    for _,c:=range h.hand{
        cnt[c]++
    }
    rankIt(cnt,h)
}

func (h *hand) String()string{
    return fmt.Sprintf("hand: %s, bid: %d, score: %d",h.hand,h.bid,h.score)
}

 func partA(challenge *challenge.Input) int {
    hands:=[]hand{}
    for line:=range challenge.Lines(){
        h:=hand{}
        h.parseHand(line)
        h.scoreHand()
        hands=append(hands,h)
    }
    sort.Slice(hands,func(i,j int)bool{
        return compare(hands[j],hands[i],mapRune)>0
    })
    result:=0 
    for i:=0;i<len(hands);i++{
        result+=hands[i].bid * (i+1)
    }
    return result
}
