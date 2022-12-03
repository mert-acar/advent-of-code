import string

def part1(data):
  count = 0
  for item in data:
    f = set(item[:len(item) // 2])
    s = set(item[len(item) // 2:])
    inter = list(f.intersection(s))
    for ch in inter:
      count += priorities[ch]
  return count


def part2(data):
  count = 0
  step = 3
  for i in range(0, len(data) - step + 1, step):
    p1 = set(data[i])
    p2 = set(data[i + 1])
    p3 = set(data[i + 2])
    item = list(p1.intersection(p2).intersection(p3))
    count += priorities[item[0]]
  return count

if __name__ == "__main__":
  with open("../data/day3.txt", 'r') as f:
    data = f.read().splitlines()
  
  i = 1
  priorities = {}
  for alp in [string.ascii_lowercase, string.ascii_uppercase]:
    for char in alp:
      priorities[char] = i
      i += 1

  res = part1(data)
  print("Part I:", res)
  res = part2(data)
  print("Part II:", res)
