#!/usr/bin/env python3

import sys


DICT_DIR = '/home/dave/dict/dict'
WORD_LIST_DIR = '/usr/share/dict/american-english'


def checkWordList(word):
    """
    check wordlist if word in wordlist return Ture
    """
    flag = None
    with open(WORD_LIST_DIR) as f:
        for line in f:
            if line.strip().lower() == word:
                # convert the word in list to lowercase
                flag = True
    return flag


def searchWord(word):
    """
    search wordlist.
    if not found return None
    else return the found line
    """
    text = None
    with open(DICT_DIR) as f:
        for line in f:
            if line.split(':')[0] == word:
                text = line
    return text


def sortList(dir=DICT_DIR):
    """
    sort the dictlist default value is DICT_DIR
    and remove the space
    """
    with open(DICT_DIR) as f:
        lines = [line for line in f if line.strip()]  # remove space
        lines.sort()
    with open(DICT_DIR, 'w') as f:
        f.writelines(lines)


def addWord(word):
    """
    add the user input
    """
    list = ['英音', '美音', 'vt.', 'vi.', 'adj.', 'adv.', 'n.', 'other. ']
    print(''+word)
    content = str()
    for _ in list:
        text = input(_)
        if text:
            content += ':' + _ + text
    line = word + content
    with open(DICT_DIR, 'a', encoding='utf8') as f:
        f.write(line+'\n')
    sortList()  # call sortList function
    print('Add successfully!')


def countWords():
    """
    traversal
    """
    nu = 0
    for line in open('/home/dave/dict/dict'):
        nu += 1
    print()
    print('(^_^): already have {} words!'.format(nu))
    print()


def main():
    for word in words:
        word = word.lower()  # all input convert to lowercase
        """
        if not in wordlist sys.exit()
        """
        flag = checkWordList(word)
        if not flag:
            print()
            print("(~_~): Are you sure {:-^8} is a word?".format(word))
            print()
            sys.exit()

        """
        search wordlist
        """
        text = searchWord(word)
        if text:  # if find the word record print it.
            for _ in text.split(':'):
                print(_)
        else:     # else interactive ask user wheather add the word
            print('(#_#): {:-^8} not in my dictionary!'.format(word))
            print(' ')
            yes = {'y', 'yes'}
            if input('add {}. y/n: '.format(word)) in yes:
                addWord(word)


if __name__ == "__main__":
    if len(sys.argv) == 1:
        print('(*_*)Usage:{} [--options] [word1] [word2] ...'
              .format(sys.argv[0]))
        sys.exit()
    elif sys.argv[1] == '--count':
        countWords()
    else:
        words = sys.argv[1:]
        main()
