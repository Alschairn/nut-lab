import re

re_sentence_splitter = re.compile('([﹒﹔﹖﹗．；。！？]["’”」』]{0,2}|：(?=["‘“「『]{1,2}|$))')


def split(text):
    s = text
    sentence_list = []
    for i in re_sentence_splitter.split(s):
        if re_sentence_splitter.match(i) and sentence_list:
            sentence_list[-1] += i
        elif i:
            sentence_list.append(i)
    return sentence_list

