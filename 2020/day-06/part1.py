import sys


def calc(string):
    count = 0
    for line in string.split("\n\n"):
        count += len(set(line) - {' ', '\n'})

    return count


with open(sys.argv[1]) as f:
    print(calc(f.read()))
