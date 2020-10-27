from collections import defaultdict

tree = defaultdict(list)

with open('input.txt') as f:
    for raw in f:
        a, b = raw.strip().split(")")
        tree[a].append(b)


def countSelf(obj):
    if obj not in tree:
        return 0
    children = tree[obj]
    return len(children) + sum([countSelf(child) for child in children])


print("day6: ", sum(countSelf(obj) for obj in tree.keys()))
