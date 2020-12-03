import sys


def get_values(arr):
    for i in arr:
        other = 2020 - i
        try:
            arr.index(other)
            return i * other
        except ValueError:
            pass


with open(sys.argv[1]) as f:
    print(get_values([int(l.strip()) for l in f if l.strip()]))
