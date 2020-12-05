import sys


def get_answers(arr):
    maximum = 0
    for a in arr:
        a = a.replace("F", "0").replace("B", "1")
        a = a.replace("R", "1").replace("L", "0")
        maximum = max(maximum, int(a, 2))

    return maximum


with open(sys.argv[1]) as f:
    arr = [l.strip() for l in f if l.strip()]
    print(get_answers(arr))
