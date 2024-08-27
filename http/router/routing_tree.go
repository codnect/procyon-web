package router

import (
	"codnect.io/procyon-web/http"
	"strings"
)

type routingTree struct {
	children     *routingNode
	staticRoutes map[string]any
	routes       []string
}

func (n *routingTree) add(mapping Mapping, chain http.HandlerChain) {

	path := mapping.pattern

	if !strings.ContainsAny(path, ":*") {
		if n.staticRoutes == nil {
			n.staticRoutes = make(map[string]any, 0)
		}

		n.staticRoutes[path] = chain
		return
	}

	node := n.children
	index := 0
	processed := 0

	for {
	begin:

		if index == len(path) {
			if (node.typ == variableMapping || index-processed == len(node.path)) && node.chain != nil {
				panic("You have already registered the same path : " + string(path))
			}
			node.chain = chain
			return
		}

		char := path[index]

		if node.typ == variableMapping {

			if char == '/' {
				if char >= node.start && char < node.end {
					tempIndex := node.indices[char-node.start]

					if tempIndex != 0 {
						node = node.children[tempIndex]
						processed = index
						index++
						goto begin
					}
				}

				if len(node.path) == 0 {
					//chain.pathVariableNameMap[path[processed+1:index]] = len(chain.pathVariableNameMap)
					//chain.pathVariableIndexMap[len(chain.pathVariableIndexMap)] = path[processed+1 : index]

					node.handlePathSegment(path[index:], chain)
					break
				}

				if node.variableNode != nil {
					node = node.variableNode
					processed = index
					goto begin
				}

				//chain.pathVariableNameMap[path[processed+1:index]] = len(chain.pathVariableNameMap)
				//chain.pathVariableIndexMap[len(chain.pathVariableIndexMap)] = path[processed+1 : index]

				node.handlePathSegment(path[index:], chain)
				break
			}
		} else {
			if index == len(path) {
				tempIndex := index - processed
				splitNode := &routingNode{
					path:         node.path[tempIndex:],
					pathLen:      uint(len(node.path[tempIndex:])),
					chain:        node.chain,
					indices:      node.indices,
					start:        node.start,
					end:          node.end,
					index:        node.index,
					children:     node.children,
					variableNode: node.variableNode,
					wildcardNode: node.wildcardNode,
					hasVariable:  node.hasVariable,
					hasWildcard:  node.hasWildcard,
					typ:          node.typ,
					childNode:    node.childNode,
				}

				node.typ = segmentMapping
				node.path = node.path[:tempIndex]
				node.pathLen = uint(len(node.path[:tempIndex]))
				node.chain = nil
				node.variableNode = nil
				node.wildcardNode = nil
				node.hasWildcard = false
				node.hasVariable = false
				node.start = 0
				node.end = 0
				node.index = 0
				node.indices = nil
				node.children = nil
				node.childNode = nil

				node.chain = chain
				node.addRoutingNode(splitNode)
				break
			}

			if index-processed == len(node.path) {

				if char >= node.start && char < node.end {
					tempIndex := node.indices[char-node.start]

					if tempIndex != 0 {
						node = node.children[tempIndex]
						processed = index
						index++
						goto begin
					}
				}

				if len(node.path) == 0 {
					node.handlePathSegment(path[index:], chain)
					break
				}

				if node.variableNode != nil {
					node = node.variableNode
					processed = index
					goto begin
				}

				node.handlePathSegment(path[index:], chain)
				break
			}

			tempIndex := index - processed
			if path[index] != node.path[index-processed] {
				splitNode := &routingNode{
					path:         node.path[tempIndex:],
					pathLen:      uint(len(node.path[tempIndex:])),
					chain:        node.chain,
					indices:      node.indices,
					start:        node.start,
					end:          node.end,
					index:        node.index,
					children:     node.children,
					variableNode: node.variableNode,
					wildcardNode: node.wildcardNode,
					hasVariable:  node.hasVariable,
					hasWildcard:  node.hasWildcard,
					typ:          node.typ,
					childNode:    node.childNode,
				}

				node.typ = segmentMapping
				node.path = node.path[:tempIndex]
				node.pathLen = uint(len(node.path[:tempIndex]))
				node.chain = nil
				node.variableNode = nil
				node.wildcardNode = nil
				node.hasWildcard = false
				node.hasVariable = false
				node.start = 0
				node.end = 0
				node.index = 0
				node.indices = nil
				node.children = nil
				node.childNode = nil

				if len(path[index:]) == 0 {
					node.chain = chain
					node.addRoutingNode(splitNode)
					break
				}

				node.addRoutingNode(splitNode)
				node.handlePathSegment(path[index:], chain)
				break
			}
		}
		index++
	}
}

func (n *routingTree) match(ctx http.Context) http.HandlerChain {
	var (
		index     uint
		path      = ctx.Request().Path()
		processed uint

		lastWildcardMapping *routingNode
		//lastWildcard        uint
		existLastWildcard bool
		handlerChain      http.HandlerChain

		node       = n.children
		pathLength = uint(len(path))
		//pathVariables      = ctx.Value(http.PathVariablesAttribute).(*http.PathVariables)
		//pathVariablesIndex int
	)

search:
	for {
		if node == nil {
			return handlerChain
		}

		if index == pathLength {
			if index-processed == node.pathLen || node.path[node.pathLen-1] == 47 {
				handlerChain = node.chain
			}
			break
		}

		if index-processed == node.pathLen {
			if node.hasWildcard {
				lastWildcardMapping = node.wildcardNode
				existLastWildcard = true
				//lastWildcard = index
			}

			character := path[index]

			if character >= node.start && character < node.end {
				childIndex := node.indices[character-node.start]

				if childIndex != 0 {
					node = node.children[childIndex]
					processed = index
					index++
					continue search
				}
			}

			if node.hasVariable {
				node = node.variableNode
				processed = index
				index++

				for {
					if index == pathLength {
						//ctx.pathVariables.nameMap = node.handlerChain.pathVariableNameMap
						//pathVariableName := node.handlerChain.pathVariableIndexMap[pathVariablesIndex]

						//pathVariables.Put(pathVariableName, path[processed:index])
						return node.chain
					}

					if path[index] == 47 {
						//ctx.pathVariables.nameMap = node.handlerChain.pathVariableNameMap
						//pathVariableName := node.handlerChain.pathVariableIndexMap[pathVariablesIndex]

						//pathVariables.Put(pathVariableName, path[processed:index])

						node = node.childNode
						processed = index
						index++
						continue search
					}

					index++
				}
			}

			if node.hasWildcard {
				//ctx.pathVariables.nameMap = node.handlerChain.pathVariableNameMap
				//pathVariableName := node.handlerChain.pathVariableIndexMap[pathVariablesIndex]

				//pathVariables.Put(pathVariableName, path[index:])
				handlerChain = node.wildcardNode.chain
			}
			break
		}

		if path[index] != node.path[index-processed] {
			if existLastWildcard {
				//ctx.pathVariables.nameMap = node.handlerChain.pathVariableNameMap
				//pathVariableName := node.handlerChain.pathVariableIndexMap[pathVariablesIndex]

				//pathVariables.Put(pathVariableName, path[lastWildcard:])
				handlerChain = lastWildcardMapping.chain
			}
			break
		}

		index++

	}

	//ctx.pathVariables.nameMap = node.handlerChain.pathVariableNameMap
	return handlerChain
}
