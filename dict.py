#!/usr/bin/env python3

import sys


def checkWordList(word):
    flag = None
    for line in open('/home/dave/dict/wordlist'):
        if line.strip() == word:
            flag = True
    if flag:
        return word
    else:
        print('false {}'.format(line.strip()))
        print("Are you sure {:-^8} is a word?".format(word))
        sys.exit()


def searchWord(word):
    flag = None
    text = None
    for line in open('/home/dave/dict/dict'):
        if line.split(':')[0] == word:
            text = line
            flag = True
    if flag:
        for _ in text.split(':'):
            print(_)
    else:
        print('{:-^8} not in list'.format(word))
        print(' ')
        yes = {'y', 'yes'}
        if input('add {}. y/n: '.format(word)) in yes:
            addWord(word)


def addWord(word):
    list = ['英音', '美音', 'vt.','vi.', 'adj.', 'adv.', 'n.', 'other. ']
    print(''+word)
    content = str() 
    for _ in list:
        text = input(_ )
        if text:
            content += ':' + _ + text
    line = word + content
    fout = open('/home/dave/dict/dict', 'a', encoding='utf8')
    fout.write(line+'\n')
    fout.close()
    print('Add successfully!')


def main():
    for word in words:
        checkWordList(word)
        searchWord(word)

if __name__ == "__main__":
    if len(sys.argv) == 1:
        print('Usage:{} word1 word2 ...'.format(sys.argv[0]))
        sys.exit()
    else:
        words = sys.argv[1:]

main()
