import sys


def get_answer(arr, x_offset, y_offset):
    tree_count = 0
    x_pos = 0
    y_pos = 0
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
    data = [l.strip() for l in f if l.strip()]
    slopes = [(1, 1), (1, 3), (1, 5), (1, 7), (2, 1)]
    trees = 1
    for slope in slopes:
        trees *= get_answer(data, slope[0], slope[1])
    print(trees)
