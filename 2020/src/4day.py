# Solution submitted to: https://adventofcode.com/2020/day/4
valid = set(map(lambda x: hex(x)[2:], list(range(16))))

def isValidColor(color):
  if color[0] != '#':
    return False
  if len(color[1:]) != 6:
    return False
  if len(set(color[1:]).union(valid)) != 16:
    return False
  return True
  
def solve(data):
  c1, c2 = 0, 0
  for record in data:
    r = {key: value for (key, value) in list(map(lambda x: x.split(':'), record))}
    keys = r.keys()
    if len(keys) == 8 or (len(keys) == 7 and 'cid' not in keys):
      c1 += 1
    try:
      if not (1920 <= int(r['byr']) <= 2002):
        continue
      if not (2010 <= int(r['iyr']) <= 2020):
        continue
      if not (2020 <= int(r['eyr']) <= 2030):
        continue
      if r['hgt'][-2:] == 'cm':
        if not (150 <= int(r['hgt'][:-2]) <= 193):
          continue
      elif r['hgt'][-2:] == 'in':
        if not (59 <= int(r['hgt'][:-2]) <= 76):
          continue
      else:
        continue
      if r['ecl'] not in ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth']:
        continue
      if not(len(r['pid']) == 9 and r['pid'].isnumeric()):
        continue
      if not (isValidColor(r['hcl'])):
        continue
      c2 += 1
    except KeyError:
      continue
  return c1, c2

if __name__ == "__main__":
  with open("../data/day4.txt", "r") as f:
  # with open("test.txt", "r") as f:
    data = f.read().split('\n\n')
    data = list(map(lambda x: x.split(), data))

  res1, res2 = solve(data)
  print("Part I:", res1)
  print("Part II:", res2)
