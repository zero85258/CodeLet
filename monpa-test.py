import monpa
import sys

if __name__ == '__main__':
    # param=文章string
    sentence = sys.argv[1]
    result = monpa.pseg(sentence)

    for t in result:
       print(t)
