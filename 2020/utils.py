from time import time
from functools import wraps

def generate(day):
  string = '''# Solution submitted to: https://adventofcode.com/2020/day/{0}
if __name__ == "__main__":
  with open("../data/day{0}.txt", "r") as f:
  # with open("test.txt", "r") as f:
    data = f.read()
  print(data)'''.format(day)
  with open('{}day.py'.format(day), 'w') as f:
    f.write(string)  
    
def timing(f):
  @wraps(f)
  def wrap(*args, **kwargs):
    start = time()
    result = f(*args, **kwargs)
    end = time()
    print('func:%r took: %2.8f sec' %(f.__name__, end-start))
    return result
  return wrap

if __name__ == "__main__":
  for i in range(1,26):
    generate(i)
