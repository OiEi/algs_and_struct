package graph

import (
	"context"
	"log"
	"math/rand"
	"time"
)

var (
	baseGraph = map[string][]string{
		"MOW": {"LED", "PEZ"},
		"PEZ": {"LED", "MOW", "KZN"},
		"LED": {"MOW", "IST"},
		"IST": {"PUK", "KZN"},
	}

	baseGraph2 = map[string][]string{
		"MOW": {"LED", "KZN"},
		"PEZ": {"LED", "MOW", "KZN"},
		"LED": {"MOW", "IST"},
		"IST": {"PUK", "KZN"},
	}
)

type SearchMachine struct {
	graph    map[string][]string
	resultCh chan []string
}

func FindPathBetweenVert(ctx context.Context, start, end string, countOfVariants, maxVariantLen int) [][]string {
	//for i := 1; i < 100000; i++ {
	//	addVert(getRandRout(), &baseGraph)
	//}

	timeStart := time.Now()
	result := Dfs(ctx, baseGraph, start, end, countOfVariants, maxVariantLen)
	log.Printf("Длинна графа %v, время обхода %v ms", len(baseGraph), time.Since(timeStart).Milliseconds())

	return result
}

func Dfs(baseCtx context.Context, graph map[string][]string, startAirport, destAirport string, countOfVariants int, maxVariantLen int) [][]string {
	checkedAirports := make(map[string]struct{})
	path := make([]string, 0, 0)
	resultCh := make(chan []string)

	//сущность нужна чтоб хранить граф по которому обходим и канал в котором собираем результат
	joinMachine := SearchMachine{
		graph:    graph,
		resultCh: resultCh,
	}

	ctx, cancel := context.WithTimeout(baseCtx, 10*time.Minute)
	defer cancel()

	go func() {
		joinMachine.strongDFS(ctx, startAirport, destAirport, &checkedAirports, &path, maxVariantLen)
		close(resultCh)
	}()

	result := make([][]string, 0, 0)

	for t := range resultCh {
		if len(result) == countOfVariants-1 {
			cancel()
		}
		result = append(result, t)
	}

	return result
}

func (machine SearchMachine) strongDFS(ctx context.Context, start, end string, visited *map[string]struct{}, path *[]string, joinsLen int) (res []string) {
	//перестаём углубляться по условию снаружи
	select {
	case <-ctx.Done():
		return nil
	default:
	}

	if start == end {
		*path = append(*path, end)
		return *path
	}

	(*visited)[start] = struct{}{}
	*path = append(*path, start) //тут надо именно через указатель, иначе после схлопывания рекурсии тут будут левые данные

	if joinsLen > 0 && len(*path) == joinsLen {
		return nil
	}

	for _, neighbor := range machine.graph[start] {
		if _, found := (*visited)[neighbor]; !found {
			resalt := machine.strongDFS(ctx, neighbor, end, visited, path, joinsLen)
			if resalt != nil {
				copyRes := make([]string, 0, len(resalt))
				copyRes = append(copyRes, resalt...)
				machine.resultCh <- copyRes
			}
			//уотакая хуйня собачка, по другому пока не придумал как сделать
			copyPath := *path
			trimPath := copyPath[:len(copyPath)-1]
			*path = trimPath

			delete(*visited, neighbor)
		}
	}

	return nil
}

func dfs(graph map[string][]string, start, testedNode string, checkedAirports *map[string]struct{}) bool {
	if start == testedNode {
		return true
	}

	if _, found := (*checkedAirports)[start]; found {
		return false
	}

	(*checkedAirports)[start] = struct{}{}

	for _, neighbor := range graph[start] {
		if _, found := (*checkedAirports)[neighbor]; !found {
			if dfs(graph, neighbor, testedNode, checkedAirports) {
				log.Print(neighbor)
				return true
			}
		}
	}

	return false
}

func bfs(graph map[string][]string, start, dest string) bool {
	visited := make(map[string]struct{})
	visited[start] = struct{}{}

	queue := make([]string, 0, 0)
	queue = append(queue, start)
	for len(queue) != 0 {
		lastElemInQueue := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		for _, neighborVert := range graph[lastElemInQueue] {
			if _, find := visited[neighborVert]; !find {
				queue = append(queue, neighborVert)
				visited[neighborVert] = struct{}{}
				if neighborVert == dest {
					log.Printf("Ага")
					return true
				}
			}
		}
	}

	return false
}

func addVert(route []string, graph *map[string][]string) {
	if len(route) < 2 {
		return
	}

	for i, newAirport := range route {
		if i == len(route)-1 {
			break
		}

		nextAirport := route[i+1]

		if _, found := (*graph)[newAirport]; !found {
			(*graph)[newAirport] = []string{nextAirport}
		} else {

			var isExistAirport bool

			for _, vert := range (*graph)[newAirport] {
				if vert == nextAirport {
					isExistAirport = true
				}
			}

			if !isExistAirport {
				(*graph)[newAirport] = append((*graph)[newAirport], nextAirport)
			}

		}
	}

}

func getRandRout() []string {

	source := []string{"PEZ", "MOW", "LED", "SDU", "SVO", "DME", "KJA", "PUK", "FGT", "SDY", "BVS", "DDK", "BJS", "KGT", "FVT", "DLE"}

	rand.Seed(time.Now().Unix())

	result := []string{source[rand.Intn(len(source))]}

	min := 2
	max := 6
	n := rand.Intn(max-min+1) + min

	for i := 1; i < n; i++ {
		result = append(result, source[rand.Intn(len(source))])
	}

	return result
}
