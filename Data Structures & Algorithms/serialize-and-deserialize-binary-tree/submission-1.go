/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


import (
	"slices"
)

func remove(s []*TreeNode, index int) []*TreeNode {
	return append(s[:index], s[index+1:]...)
}

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return ""
    }
    var ret string
    var q []*TreeNode
    q = append(q, root)
    for len(q) != 0 {
        cur := q[0] // front
        q = remove(q, 0)
        if cur.Val == -1001 {
            ret += ".L."
            continue
        } else if cur.Val == 1001 {
            ret += ".R."
            continue
        } else {
            ret += "."
            ret += strconv.Itoa(cur.Val)
            ret += "."
        }
        if cur.Left != nil {
            q = append(q, cur.Left)
        } else {
            q = append(q, &TreeNode{
                Val: -1001,
                Left: nil,
                Right: nil,
            })
        }
        if cur.Right != nil {
            q = append(q, cur.Right)
        } else {
            q = append(q, &TreeNode{
                Val: 1001,
                Left: nil,
                Right: nil,
            })
        }
    }
	q = nil
    return ret
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    if len(data) == 0 {
        return nil
    }
    split_values := strings.Split(data, ".")
    values := slices.DeleteFunc(split_values, func(ele string) bool {
        return (len(ele) == 0)
    })
    val, _ := strconv.Atoi(values[0])
    root := &TreeNode {
        Val: val,
        Left: nil,
        Right: nil,
    }
    var q []*TreeNode
    q = append(q, root)
    j := 1
    for len(q) != 0 {
        prev := q[0]
        q = remove(q, 0)
        if values[j] == "L" {
            prev.Left = nil
        } else {
            val, _ = strconv.Atoi(values[j])
            prev.Left = &TreeNode {
                Val : val,
                Left: nil,
                Right: nil,
            }
            q = append(q, prev.Left)
        }
        j++

        if values[j] == "R" {
            prev.Right = nil
        } else {
            val, _ = strconv.Atoi(values[j])
            prev.Right = &TreeNode {
                Val : val,
                Left: nil,
                Right: nil,
            }
            q = append(q, prev.Right)
        }
        j++
    }
	q = nil
    return root
}
