def solve():
    import sys
    input = sys.stdin.read
    from functools import lru_cache

    data = input().strip().split()
    N = int(data[0])
    cards = [(int(data[2*i+1]), int(data[2*i+2])) for i in range(N)]

    @lru_cache(None)
    def can_win(remaining_cards):
        rc = list(remaining_cards)
        for i in range(len(rc)):
            for j in range(i+1, len(rc)):
                a_i, b_i = rc[i]
                a_j, b_j = rc[j]
                if a_i == a_j or b_i == b_j:
                    next_state = tuple(rc[k] for k in range(len(rc)) if k != i and k != j)
                    if not can_win(next_state):
                        return True
        return False

    initial_state = tuple(cards)
    if can_win(initial_state):
        print("Takahashi")
    else:
        print("Aoki")

solve()


