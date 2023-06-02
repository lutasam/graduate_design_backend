package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/lutasam/doctors/biz/common"
	"github.com/lutasam/doctors/biz/utils"
)

type TrieNode struct {
	Children map[rune]*TrieNode
	IsEnd    bool
}

func BuildTrie(words []string) *TrieNode {
	root := &TrieNode{
		Children: make(map[rune]*TrieNode),
		IsEnd:    false,
	}
	for _, word := range words {
		curNode := root
		for _, char := range word {
			if _, ok := curNode.Children[char]; !ok {
				curNode.Children[char] = &TrieNode{
					Children: make(map[rune]*TrieNode),
					IsEnd:    false,
				}
			}
			curNode = curNode.Children[char]
		}
		curNode.IsEnd = true
	}
	return root
}

var trieRoot *TrieNode

func init() {
	sensitiveWords := []string{"傻逼", "操", "死全家", "你妈", "fuck"}
	trieRoot = BuildTrie(sensitiveWords)
}

func SensitiveFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		var jsonData map[string]interface{}
		if err := c.ShouldBindBodyWith(&jsonData, binding.JSON); err != nil {
			utils.ResponseClientError(c, common.USERINPUTERROR)
			c.Abort()
			return
		}
		for _, value := range jsonData {
			if str, ok := value.(string); ok {
				curNode := trieRoot
				for _, char := range str {
					if _, ok := curNode.Children[char]; !ok {
						break
					}
					curNode = curNode.Children[char]
					if curNode.IsEnd {
						utils.ResponseClientError(c, common.SENSITIVEINPUT)
						c.Abort()
						return
					}
				}
			}
		}
		c.Next()
	}
}
