package _00_399

import "fmt"

//332.重新安排行程 集合大法好 还是得用深度优先,用回溯

func findItinerary(tickets [][]string) []string {

    var current [][]string
    dfs(tickets, &current, []string{})
    if len(current) == 0 {
        return nil
    }
    return current[0]

}

func dfs(tickets [][]string, list *[][]string, current []string) {
    if len(tickets) == 0 {
        *list = append(*list, current)
        return
    }
    //fmt.Println(current,tickets)
    for i := 0; i < len(tickets); i++ {
        if len(current) == 0 {
            if tickets[i][0] == "JFK" {
                dfs(append(tickets[:i], tickets[i+1:]...), list, append(current, tickets[i]...))
                //current = append(current, tickets[i]...)
            }
        } else {
            fmt.Println(tickets[i],current[len(current)-1])
            if tickets[i][0] == current[len(current)-1] {
                //current = append(current, tickets[i][1])
                dfs(append(tickets[:i], tickets[i+1:]...), list, append(current, tickets[i][1]))
            }

        }
    }

}

func main() {
    f := [][]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}}
    //f := [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}
    fmt.Println(findItinerary(f))
}
