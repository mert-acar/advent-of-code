# Solution submitted to: https://adventofcode.com/2020/day/6
def solve(data):
  c1, c2 = 0, 0 
  for line in data:
    c1 += len(set(line.replace('\n','')))
    c2 += len(set.intersection(*[set(i) for i in line.split('\n')]))
  return c1, c2
     
if __name__ == "__main__":
  with open("../data/day6.txt", "r") as f:
    data = f.read()[:-1].split('\n\n')

  res1, res2 = solve(data)
  print("Part I:", res1)
  print("Part II:", res2)
