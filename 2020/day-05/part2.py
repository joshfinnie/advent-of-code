import sys


def get_answers(arr):
    possible_seats = set(range(1024))
    for a in arr:
        a = a.replace("F", "0").replace("B", "1")
        a = a.replace("R", "1").replace("L", "0")
        seat = int(a, 2)
        possible_seats.remove(seat)

    for seat in possible_seats:
        if seat - 1 not in possible_seats and seat + 1 not in possible_seats:
            return seat


with open(sys.argv[1]) as f:
    arr = [l.strip() for l in f if l.strip()]
    print(get_answers(arr))
