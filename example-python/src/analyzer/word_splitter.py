import jieba
import jieba.posseg as pseg


def split(sentence):
    jieba.enable_paddle()
    str = sentence
    words = pseg.cut(sentence=str, use_paddle=True)
    return words
