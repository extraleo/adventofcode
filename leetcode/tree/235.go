package main 

func lowestCommonAncestor_235(root, p, q *TreeNode) *TreeNode {

    if root.Val > q.Val && root.Val > p.Val{
        return lowestCommonAncestor_235(root.Left, p, q)
    }
    if root.Val < q.Val && root.Val < p.Val{
        return lowestCommonAncestor_235(root.Right, p, q)
    }
    return root
}