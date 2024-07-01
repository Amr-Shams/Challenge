package day7 

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/Amr-Shams/aoc2024/challenge"
    "sort"
)


func bCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "b",
        Short: "Day 7, Problem B",
        Run: func(_ *cobra.Command, _ []string) {
            fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
        },
    }
    return cmd
}
func (h*hand) rankHand() {
    cnt:=map[rune]int{}
    var maxChar rune
    for _,c:=range h.hand{
        cnt[c]++
        if cnt[c]>cnt[maxChar]{
            maxChar=c
        }
    }
    cnt[maxChar]+=cnt['J']
    if cnt['J']==5{
        h.score=FiveOfAKind
        return
    }
    delete(cnt,'J')
    rankIt(cnt,h)

}
func mapByte(b byte) int {
    switch b {
        case 'A': return 14
        case 'K': return 13
        case 'Q': return 12
        case 'J': return 0
        case 'T': return 10
        default: return int(b-'0')
    }
}
func partB(challenge *challenge.Input) int {
    hands:=[]hand{}
    for line := range challenge.Lines() {
        h:=hand{}
        h.parseHand(line)
        h.rankHand()
        hands=append(hands,h)
    }
    sort.Slice(hands,func(i,j int) bool{
        return compare(hands[j],hands[i],mapByte) > 0
    })
    reuslt:=0
    for i:=0;i<len(hands);i++{
        reuslt+=hands[i].bid * (i+1)
    }
    return reuslt
}
