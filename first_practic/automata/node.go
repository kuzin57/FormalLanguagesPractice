package automata

type node struct {
	next       map[byte][]*edge
	lenPrefix  int
	isTerminal bool
}

type edge struct {
	to         *node
	prefixRead int
	visited    bool
}

func newNode() *node {
	return &node{next: make(map[byte][]*edge)}
}

func (n *node) createBranch(word string, lenPrefix int) *node {
	curNode := n

	for len(word) > 0 {
		curNode.lenPrefix = lenPrefix
		nextNode, exists := curNode.next[word[0]]

		switch exists {
		case true:
			curNode = nextNode[0].to
		case false:
			curNode.next[word[0]] = make([]*edge, 1)
			curNode.next[word[0]][0] = &edge{to: newNode()}
			curNode = curNode.next[word[0]][0].to
		}

		lenPrefix++
		word = word[1:]
	}

	curNode.isTerminal = true
	return curNode
}

func (n *node) findMaxPrefix(word string, currentPosition int, prefixRead int) (*node, int, int) {
	var (
		maxPosition        int
		maxPrefix          = prefixRead
		nodeMaxPosition    *node
		maxPositionSon     int
		maxPrefixSon       int
		nodeMaxPositionSon *node
	)

	if n.isTerminal {
		prefixRead = currentPosition
		maxPrefix = currentPosition
	}

	if (len(n.next) == 0 || len(word) == 0) && n.isTerminal {
		return n, currentPosition, prefixRead
	}

	nextNodes := make(map[byte][]*edge)
	if len(word) > 0 {
		nextNodes[word[0]] = n.next[word[0]]
	}
	nextNodes[empty] = n.next[empty]

	if (len(word) == 0 && len(nextNodes[empty]) == 0) ||
		(len(word) > 0 && len(nextNodes[word[0]]) == 0 && len(nextNodes[empty]) == 0) {
		return n, currentPosition, prefixRead
	}

	for letter, nextEdges := range nextNodes {
		for _, nextEdge := range nextEdges {
			if !nextEdge.to.isTerminal && nextEdge.visited && nextEdge.prefixRead == prefixRead {
				continue
			}

			if !nextEdge.to.isTerminal {
				nextEdge.prefixRead = prefixRead
			} else {
				switch {
				case letter != empty:
					nextEdge.prefixRead = currentPosition + 1
				case letter == empty:
					nextEdge.prefixRead = currentPosition
				}
			}

			if nextEdge.to.isTerminal && nextEdge.visited && nextEdge.prefixRead == prefixRead {
				continue
			}

			nextEdge.visited = true
			if len(word) > 0 && letter == word[0] {
				nodeMaxPositionSon, maxPositionSon, maxPrefixSon = nextEdge.to.findMaxPrefix(word[1:], currentPosition+1, prefixRead)
			} else {
				nodeMaxPositionSon, maxPositionSon, maxPrefixSon = nextEdge.to.findMaxPrefix(word, currentPosition, prefixRead)
			}

			if maxPositionSon > maxPosition {
				maxPosition = maxPositionSon
				nodeMaxPosition = nodeMaxPositionSon
			}

			if maxPrefixSon > maxPrefix {
				maxPrefix = maxPrefixSon
			}
		}
	}

	return nodeMaxPosition, maxPosition, maxPrefix
}
