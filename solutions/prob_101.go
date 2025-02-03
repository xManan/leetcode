package solutions

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func IsSymmetric(root *TreeNode) bool {
    return isMirror(root.Left, root.Right)
}

func isMirror(n1 *TreeNode, n2 *TreeNode) bool {
    if n1 == nil && n2 == nil {
        return true
    }
    if n1 == nil || n2 == nil {
        return false
    }
    return n1.Val == n2.Val && isMirror(n1.Left, n2.Right) && isMirror(n1.Right, n2.Left)
}
