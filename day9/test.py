import sys
inst_set = [line.rstrip() for line in open(sys.argv[1])]
moves = {
    'U': (0, 1),
    'D': (0, -1),
    'L': (-1, 0),
    'R': (1, 0)
}

def simulate(n):

    s = (0,0)
    vst = [s]
    knt = [s for x in range(n)]

    grid = [[0 for i in range(50)] for _ in range(50)]

    for inst in inst_set:

        dir, step = moves[inst.split(' ')[0]], int(inst.split(' ')[1])
        
        for i in range(step):

            # head always moves, first
            m_x, m_y = dir[0], dir[1]
            knt[0] = (knt[0][0] + m_x, knt[0][1] + m_y)

            for i in range(1,len(knt)):

                x_h, y_h = knt[i-1][0], knt[i-1][1]
                x_t, y_t = knt[i][0], knt[i][1]
                gap_x = x_h - x_t
                gap_y = y_h - y_t

                if abs(gap_x) >= 2 or abs(gap_y) >= 2:

                    xp = 1 if gap_x > 0 else (0 if gap_x == 0 else -1)
                    yp = 1 if gap_y > 0 else (0 if gap_y == 0 else -1)
                    knt[i] = (knt[i][0] + xp , knt[i][1] + yp)

            last_pos = knt[-1]
            if last_pos not in vst:
                vst.append(last_pos)
                x, y = last_pos
                grid[y+20][x+20] = 1

    for row in grid[::-1]:
        for c in row:
            print('.' if c == 0 else '#', end="")
        print()

    return len(vst)

for i in range(2, 11):
    print(simulate(i)) # part 2
