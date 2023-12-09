# https://www.reddit.com/r/adventofcode/comments/18e5ytd/comment/kcm1c0w/

from math import prod

data = [[*map(int, s.split())] for s in open('input.txt')]

def P(x):
    n = len(x)
    Pj = lambda j: prod((n-k)/(j-k) for k in range(n) if k!=j)

    return sum(x[j] * Pj(j) for j in range(n))

for dir in 1, -1:
    print(sum(P(l[::dir]) for l in data))
