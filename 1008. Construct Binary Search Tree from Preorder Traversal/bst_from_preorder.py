# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

# 4 cases
# 1. singleton [x]
# 2. [x, z1, z2, ..., zn], s.t. z > x for all n
# 3. [x, y1, y2, ..., ym], s.t. y < x for all m
# 4. [x, y1, y2, ..., ym, z1, z2, ..., zn], s.t. yi < x for all i and zj > x for all j
class Solution:
    def bstFromPreorder(self, preorder: List[int]) -> Optional[TreeNode]:
        def build(ls):
            root_val = ls[0]

            # ase 1
            if len(ls) == 1:
                return TreeNode(root_val)

            # case 2
            if ls[1] > root_val:
                return TreeNode(root_val, None, build(ls[1:]))

            i = 1
            while i < len(ls) and ls[i] < root_val:
                i += 1
            # case 3
            if i == len(ls):
                return TreeNode(root_val, build(ls[1:]), None)
            # case 4
            else:
                return TreeNode(root_val, build(ls[1:i]), build(ls[i:]))

        return build(preorder)
