import re
import sys

def test_id(id):
    id_split = [i.split(":") for i in id]
    id_dict = {k: v for k, v in id_split}
    if (
        1920 <= int(id_dict['byr']) <= 2002 and
        2010 <= int(id_dict['iyr']) <= 2020 and
        2020 <= int(id_dict['eyr']) <= 2030 and
        (m1 := re.match(r'^(\d+)(cm|in)$', id_dict['hgt'])) and
        (
            m1[2] == 'cm' and 150 <= int(m1[1]) <= 193 or
            m1[2] == 'in' and 59 <= int(m1[1]) <= 76
        ) and
        re.match('^#[a-f0-9]{6}$', id_dict['hcl']) and
        id_dict['ecl'] in set('amb blu brn gry grn hzl oth'.split()) and
        re.match('^[0-9]{9}$', id_dict['pid'])
    ):
        return True
    return False



def get_answers(arr):
    must_include = [
        "byr",
        "iyr",
        "eyr",
        "hgt",
        "hcl",
        "ecl",
        "pid",
    ]
    valid_passports = 0
    for a in arr:
        id = a.split()
        if len(id) == 8:
            if test_id(id):
                valid_passports += 1
        else:
            parts = []
            valid = True
            for p in id:
                parts.append(p.split(":")[0])
            for include in must_include:
                if include not in parts:
                    valid = False
            if valid:
                if test_id(id):
                    valid_passports += 1


    return valid_passports



with open(sys.argv[1]) as f:
    arr = f.read().split("\n\n")
    print(get_answers(arr))
