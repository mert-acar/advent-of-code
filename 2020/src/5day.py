# Solution submitted to: https://adventofcode.com/2020/day/5
def findRow(code):
  return int(code.replace("F","0").replace("B","1"), 2)

def findCol(code):
  return int(code.replace("L","0").replace("R","1"), 2)

def findIdx(code):
  row = findRow(code[:-3])
  col = findCol(code[-3:])
  return row * 8 + col

def solve(data):
  ids = []
  for line in data:
    ids.append(findIdx(line))
  # Part I
  max_id = max(ids)
  # Part II
  all_ids = set(range(10, max_id))
  my_id = list(all_ids.difference(set(ids)))[-1]
  return max_id, my_id

if __name__ == "__main__":
  with open("../data/day5.txt", "r") as f:
    data = list(map(lambda x: x.strip(), f.readlines()))
  
  res1, res2 = solve(data)
  print("Part I:", res1)
  print("Part II:", res2)
