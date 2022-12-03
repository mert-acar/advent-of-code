# Solution submitted to: https://adventofcode.com/2020/day/11
import numpy as np

if __name__ == "__main__":
  with open("../data/day11.txt", "r") as f:
    data = list(
      map(
        lambda x: [int(i) for i in x],
        f.read().replace('L', '1').replace('.', '0').splitlines()
      )
    )
    data = np.array(data)
  
