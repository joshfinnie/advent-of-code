import sys


def get_answer(arr):
    tree_count = 0
    x_pos = 0
    y_pos = 0
    x_offset = 1
    y_offset = 3
    width = len(arr[0])
    height = len(arr)

    while x_pos < height:
        test = arr[x_pos][y_pos]
        if test == "#":
            tree_count += 1
        x_pos = x_pos + x_offset
        y_pos = (y_pos + y_offset) % width

    return tree_count

with open(sys.argv[1]) as f:
    print(get_answer([l.strip() for l in f if l.strip()]))
