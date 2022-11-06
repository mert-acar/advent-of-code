# Solution submitted to: https://adventofcode.com/2020/day/2
def solve(data):
  c1, c2 = 0, 0
  for line in data:
    lims = list(map(int, line[0].split('-')))
    # Part I
    if lims[0] <= line[-1].count(line[-2]) <= lims[1]:
      c1 += 1
    # Part II
    chrs = set([line[-1][lims[0] - 1], line[-1][lims[1] - 1]])
    if (line[-2] in chrs) and (len(chrs) != 1):
      c2 += 1
  return c1, c2


if __name__ == "__main__":
  with open("../data/day2.txt", "r") as f:
    data = list(map(lambda x: x.replace(':', '').split(), f.readlines()))

  ans1, ans2 = solve(data)
  print("Part I:", ans1)
  print("Part II:", ans2)
