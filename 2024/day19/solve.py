from functools import lru_cache


@lru_cache
def solve(design, towels):
	if design == '':
		return True

	matches = []
	for t in towels:
		if design.startswith(t):
			matches.append(t)

	for m in matches:
		if solve(design[len(m):], towels):
			return True

	return False


@lru_cache
def count(design, towels):
	if design == '':
		return 1

	matches = []
	for t in towels:
		if design.startswith(t):
			matches.append(t)

	return sum(count(design[len(m):], towels) for m in matches)


ans1 = ans2 = 0

with open("input.txt", "r") as f:
  lines = f.readlines()

towels = tuple(lines[0].split(', '))
designs = lines[2:]

ans1 = sum([solve(d, towels) for d in designs])
print(f"Part I: {ans1}")

ans2 = sum([count(d, towels) for d in designs])
print(f"Part II: {ans1}")
