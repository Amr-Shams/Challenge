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
    return 0
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
func rankIt(cnt map[rune]int, h *hand){
     if len(cnt)==5{
        h.score=HighCard
        return
    }
    if len(cnt)==4{
        h.score=OnePair
        return
    }
    if len(cnt)==3{
       for v,_:=range cnt{
        if cnt[v]==3{
            h.score=ThreeOfAKind
            return
        }
       }
        h.score=TwoPair
        return
    }
    if len(cnt)==2{
       for v,_:=range cnt{
       if cnt[v]==3{
           h.score=FullHouse
           return
       
        }
       }
        h.score=FourOfAKind
        return
    }
    h.score=FiveOfAKind   
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
