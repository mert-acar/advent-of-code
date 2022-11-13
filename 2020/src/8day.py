# Solution submitted to: https://adventofcode.com/2020/day/8

def run(code, visited, acc=0, ptr=0):
  while ptr < len(code) and ptr not in visited:
    visited[ptr] = acc
    op, arg = code[ptr]
    arg = int(arg)
    if op == "acc":
      acc += arg
    elif op == "jmp":
      ptr += arg - 1
    ptr += 1
  return acc, ptr


def findSwap(code, visited):
  for k, j in enumerate(set(visited.keys())):
    op, arg = code[j]
    if op == "nop":
      i = j + int(arg) # Treat as if nop was jmp and run from this point on
      if i not in visited:
        acc, i = run(code, visited, visited[j], i)
    elif op == "jmp":
      i = j + 1 # Treat as if jmp was nop and run from this point on
      if i not in visited:
        acc, i = run(code, visited, visited[j], i)
    if i >= len(code): # Did it terminate ?
      return acc


if __name__ == "__main__":
  with open("../data/day8.txt", "r") as f:
    data = list(map(lambda x: x.strip().split(), f.readlines()))

  visited = {}
  res1, _ = run(data, visited)
  print("Part I:", res1)
  res2 = findSwap(data, visited)     
  print("Part II:", res2)
