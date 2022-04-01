import xmind
import sys

if __name__ == '__main__':
    wb = xmind.load(sys.argv[1])  # Requires '.xmind' extension
    ws = wb.getPrimarySheet()
    rt = ws.getRootTopic()

    for topic in rt.getSubTopics():
       print(topic.getTitle())

    # 目標:parse成json

