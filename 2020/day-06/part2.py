import sys


def calc(string):
    count = 0
    for line in string.split("\n\n"):
        lines = line.splitlines()
        all_counted = set(lines[0])
        for char in lines:
            all_counted &= set(char)
        count += len(all_counted)

    return count


with open(sys.argv[1]) as f:
    print(calc(f.read()))
